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
)

type UserController struct {
	userUsecase *usecase.UserUsecase
	validator   *helpers.Validation
}

func NewUserController(userUsecase *usecase.UserUsecase, validator *helpers.Validation) *UserController {
	return &UserController{
		userUsecase: userUsecase,
		validator:   validator,
	}
}

func (ctrl *UserController) GetMe(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := helpers.GetCookieRefreshToken(c)
	payloadToken, err := helpers.DecodeRefreshToken(token)
	if err != nil {
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
		return
	}

	user, resNum := ctrl.userUsecase.GetById(ctx, payloadToken.UserId)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordFound {
			response.NewResponse(c, response.Response{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: constants.RecordNotFound,
			})
			return
		}
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
		})
		return
	}

	response.NewResponse(c, response.Response{
		Data:    user,
		Code:    http.StatusOK,
		Message: constants.RecordFound,
	})
}

type errUpdateUser struct {
	Username  int `json:"username"`
	FullName  int `json:"full_name"`
	BirthDate int `json:"birth_date"`
}

func (ctrl *UserController) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := helpers.GetCookieRefreshToken(c)
	payloadToken, err := helpers.DecodeRefreshToken(token)
	if err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
			Errors:  nil,
		})
		return
	}

	var req models.UpdateUser

	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.UnsupportedContentType,
			Errors:  nil,
		})
		return
	}

	errRes := errUpdateUser{
		Username:  0,
		FullName:  0,
		BirthDate: 0,
	}

	if errorType := ctrl.validator.Struct(req); errorType != nil || len(errorType) > 0 {
		errRes.BirthDate = helpers.FilterTagCode(errorType, "BirthDate")
		errRes.FullName = helpers.FilterTagCode(errorType, "FullName")
		errRes.Username = helpers.FilterTagCode(errorType, "Username")

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.NotCompleteForm,
			Errors:  errRes,
		})
		return
	}

	user, resNum := ctrl.userUsecase.Update(ctx, payloadToken.UserId, req)
	if resNum != constants.SuccessUpdate {
		if resNum == constants.RecordFound {
			response.NewResponseError(c, response.ResponseError{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: constants.RecordNotFound,
				Errors:  nil,
			})
			return
		}
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
			Errors:  nil,
		})
		return
	}

	response.NewResponseError(c, response.ResponseError{
		Data:    user,
		Code:    http.StatusOK,
		Message: constants.RecordFound,
		Errors:  nil,
	})
}

type errChangePasswordUser struct {
	OldPassword int `json:"old_password"`
	NewPassword int `json:"new_password"`
}

func (ctrl *UserController) ChangePassword(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := helpers.GetCookieRefreshToken(c)
	payloadToken, err := helpers.DecodeRefreshToken(token)
	if err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
			Errors:  nil,
		})
		return
	}

	var req models.ChangePassword

	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.UnsupportedContentType,
			Errors:  nil,
		})
		return
	}

	errRes := errChangePasswordUser{
		OldPassword: 0,
		NewPassword: 0,
	}

	if errorType := ctrl.validator.Struct(req); errorType != nil || len(errorType) > 0 {
		errRes.OldPassword = helpers.FilterTagCode(errorType, "OldPassword")
		errRes.NewPassword = helpers.FilterTagCode(errorType, "NewPassword")

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.NotCompleteForm,
			Errors:  errRes,
		})
		return
	}

	user, resNum := ctrl.userUsecase.ChangePassword(ctx, payloadToken.UserId, req)
	if resNum != constants.SuccessUpdate {
		if resNum == constants.RecordFound {
			response.NewResponseError(c, response.ResponseError{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: constants.RecordNotFound,
				Errors:  nil,
			})
			return
		}
		if resNum == constants.PasswordNotMatch {
			response.NewResponseError(c, response.ResponseError{
				Data:    nil,
				Code:    http.StatusBadRequest,
				Message: resNum,
				Errors:  nil,
			})
			return
		}
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: resNum,
			Errors:  nil,
		})
		return
	}

	response.NewResponseError(c, response.ResponseError{
		Data:    user,
		Code:    http.StatusOK,
		Message: constants.RecordFound,
		Errors:  nil,
	})
}
