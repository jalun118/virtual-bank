package controllers

import (
	"context"
	"net/http"
	"time"
	"virtual-bank/internal/constants"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/models"
	"virtual-bank/internal/response"
	"virtual-bank/internal/usecase"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userUsecase    *usecase.UserUsecase
	sessionUsecase *usecase.SessionUsecase
	validator      *helpers.Validation
}

func NewAuthController(userUsecase *usecase.UserUsecase, sessionUsecase *usecase.SessionUsecase, validator *helpers.Validation) *AuthController {
	return &AuthController{
		userUsecase:    userUsecase,
		validator:      validator,
		sessionUsecase: sessionUsecase,
	}
}

type errRegisterUser struct {
	Username  int `json:"username"`
	Password  int `json:"password"`
	Email     int `json:"email"`
	FullName  int `json:"full_name"`
	BirthDate int `json:"birth_date"`
}

func (ctrl *AuthController) RegisterUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var req models.CreateUser

	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.UnsupportedContentType,
			Errors:  nil,
		})
		return
	}

	errRes := errRegisterUser{
		Username:  0,
		Password:  0,
		Email:     0,
		FullName:  0,
		BirthDate: 0,
	}

	if errorType := ctrl.validator.Struct(req); errorType != nil || len(errorType) > 0 {
		errRes.Email = helpers.FilterTagCode(errorType, "Email")
		errRes.BirthDate = helpers.FilterTagCode(errorType, "BirthDate")
		errRes.FullName = helpers.FilterTagCode(errorType, "FullName")
		errRes.Password = helpers.FilterTagCode(errorType, "Password")
		errRes.Username = helpers.FilterTagCode(errorType, "Username")

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.NotCompleteForm,
			Errors:  errRes,
		})
		return
	}

	httpCode := http.StatusCreated

	user, resNum := ctrl.userUsecase.Create(ctx, req)
	if resNum != constants.SuccessInsert {
		switch resNum {
		case constants.DuplicateRecord:
			errRes.Email = resNum
			httpCode = http.StatusBadRequest
		case constants.DatabaseError:
			httpCode = http.StatusInternalServerError
		case constants.InternalServerError:
			httpCode = http.StatusInternalServerError
		}

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    httpCode,
			Errors:  errRes,
			Message: constants.InputError,
		})
		return
	}

	response.NewResponseError(c, response.ResponseError{
		Data:    user,
		Code:    httpCode,
		Errors:  nil,
		Message: resNum,
	})
}

type errLoginUser struct {
	Password int `json:"password"`
	Email    int `json:"email"`
}

func (ctrl *AuthController) LoginUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var req models.LoginUser

	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.UnsupportedContentType,
			Errors:  nil,
		})
		return
	}

	errRes := errLoginUser{
		Password: 0,
		Email:    0,
	}

	if errorType := ctrl.validator.Struct(req); errorType != nil || len(errorType) > 0 {
		errRes.Email = helpers.FilterTagCode(errorType, "Email")
		errRes.Password = helpers.FilterTagCode(errorType, "Password")

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.NotCompleteForm,
			Errors:  errRes,
		})
		return
	}

	user, resNum := ctrl.userUsecase.GetByEmail(ctx, req.Email)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			errRes.Email = resNum
			response.NewResponseError(c, response.ResponseError{
				Data:    nil,
				Code:    http.StatusNotFound,
				Errors:  errRes,
				Message: resNum,
			})
			return
		}

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Errors:  errRes,
			Message: constants.NotCompleteForm,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		errRes.Password = constants.PasswordNotMatch
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Errors:  errRes,
			Message: constants.InputError,
		})
		return
	}

	session, resNum := ctrl.sessionUsecase.Create(ctx, models.PayloadSession{
		UserId: user.ID,
	})
	if resNum != constants.SuccessInsert {
		if resNum == constants.UserNotFound {
			response.NewResponseError(c, response.ResponseError{
				Data:    nil,
				Code:    http.StatusNotFound,
				Errors:  nil,
				Message: resNum,
			})
			return
		}

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Errors:  nil,
			Message: resNum,
		})
		return
	}

	refreshToken, err := helpers.CreateRefreshToken(helpers.PayloadRefreshToken{
		UserId:    user.ID,
		Email:     user.Email,
		ExpiredAt: time.Now().Add(3 * time.Minute),
	})
	if err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Errors:  nil,
			Message: constants.InternalServerError,
		})
		return
	}

	helpers.SetCookieAuth(c, helpers.CookieAuth{
		RefreshToken:   refreshToken,
		IdRefreshToken: session.ID,
	})

	response.NewResponseError(c, response.ResponseError{
		Data:    user,
		Code:    http.StatusOK,
		Errors:  nil,
		Message: resNum,
	})
}

func (ctrl *AuthController) SetRefreshToken(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cookieId := helpers.GetCookieId(c)

	session, resNum := ctrl.sessionUsecase.FindByIdAndNotExpired(ctx, cookieId)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponseUnauthorized(c)
			return
		}

		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.DatabaseError,
		})
		return
	}

	user, resNum := ctrl.userUsecase.GetById(ctx, session.UserId)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponseUnauthorized(c)
			return
		}

		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.DatabaseError,
		})
		return
	}

	token, err := helpers.CreateRefreshToken(helpers.PayloadRefreshToken{
		UserId:    session.UserId,
		Email:     user.Email,
		ExpiredAt: time.Now().Add(3 * time.Minute),
	})
	if err != nil {
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
		return
	}

	helpers.SetRefreshToken(c, token)

	response.NewResponse(c, response.Response{
		Data: gin.H{
			"token": token,
		},
		Code:    http.StatusOK,
		Message: constants.SuccessUpdate,
	})
}

func (ctrl *AuthController) LogOut(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cookieId := helpers.GetCookieId(c)
	if cookieId == "" {
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusOK,
			Message: constants.SuccessDelete,
		})
		return
	}

	reNumSession := ctrl.sessionUsecase.DeleteById(ctx, cookieId)
	if reNumSession == constants.DatabaseError {
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.DatabaseError,
		})
		return
	}

	helpers.DeleteCookieAuth(c)

	response.NewResponse(c, response.Response{
		Data:    nil,
		Code:    http.StatusOK,
		Message: constants.SuccessDelete,
	})
}
