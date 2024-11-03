package controllers

import (
	"context"
	"net/http"
	"sync"
	"time"
	"virtual-bank/internal/constants"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/models"
	"virtual-bank/internal/response"
	"virtual-bank/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountUsecase *usecase.AccountUsecase
	validator      *helpers.Validation
	mu             sync.Mutex
}

func NewAccountController(accountUsecase *usecase.AccountUsecase, validator *helpers.Validation) *AccountController {
	return &AccountController{
		validator:      validator,
		accountUsecase: accountUsecase,
	}
}

type errCreateAccount struct {
	AccountTypeId int `json:"account_type_id"`
	Balance       int `json:"balance"`
	Pin           int `json:"pin"`
	Currency      int `json:"currency"`
	MotherBirth   int `json:"mother_birth"`
}

func (ctrl *AccountController) CreateAccount(c *gin.Context) {
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

	var req models.CreateAccount

	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.UnsupportedContentType,
			Errors:  nil,
		})
		return
	}

	req.UserId = payloadToken.UserId

	errRes := errCreateAccount{
		AccountTypeId: 0,
		Balance:       0,
		Pin:           0,
		Currency:      0,
		MotherBirth:   0,
	}

	if errorType := ctrl.validator.Struct(req); errorType != nil || len(errorType) > 0 {
		errRes.Pin = helpers.FilterTagCode(errorType, "Pin")
		errRes.Balance = helpers.FilterTagCode(errorType, "Balance")
		errRes.Currency = helpers.FilterTagCode(errorType, "Currency")
		errRes.MotherBirth = helpers.FilterTagCode(errorType, "MotherBirth")
		errRes.AccountTypeId = helpers.FilterTagCode(errorType, "AccountTypeId")

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.NotCompleteForm,
			Errors:  errRes,
		})
		return
	}

	account, resNum := ctrl.accountUsecase.Create(ctx, req)
	if resNum != constants.SuccessInsert {
		httpCode := http.StatusBadRequest
		switch resNum {
		case constants.UserNotFound:
			httpCode = http.StatusUnauthorized
		case constants.AccountTypeNotFound:
			errRes.AccountTypeId = resNum
			httpCode = http.StatusNotFound
		case constants.DatabaseError:
			httpCode = http.StatusInternalServerError
		case constants.InternalServerError:
			httpCode = http.StatusInternalServerError
		}

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    httpCode,
			Message: constants.InputError,
			Errors:  nil,
		})
		return
	}
	response.NewResponseError(c, response.ResponseError{
		Data:    account,
		Code:    http.StatusCreated,
		Message: constants.SuccessInsert,
		Errors:  nil,
	})
}

type errTopUpAccount struct {
	AccountId int `json:"account_id"`
	Amount    int `json:"amount"`
}

func (ctrl *AccountController) TopAccount(c *gin.Context) {
	ctrl.mu.Lock()
	defer ctrl.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var req models.TopUpAccount

	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.UnsupportedContentType,
			Errors:  nil,
		})
		return
	}

	errRes := errTopUpAccount{
		AccountId: 0,
		Amount:    0,
	}

	if errorType := ctrl.validator.Struct(req); errorType != nil || len(errorType) > 0 {
		errRes.Amount = helpers.FilterTagCode(errorType, "Amount")
		errRes.AccountId = helpers.FilterTagCode(errorType, "AccountId")

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.NotCompleteForm,
			Errors:  errRes,
		})
		return
	}

	account, resNum := ctrl.accountUsecase.TopUpAccount(ctx, req)
	if resNum != constants.SuccessUpdate {
		if resNum == constants.RecordNotFound {
			errRes.AccountId = resNum
			response.NewResponseError(c, response.ResponseError{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: constants.InputError,
				Errors:  errRes,
			})
			return
		}

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InputError,
			Errors:  nil,
		})
		return
	}
	response.NewResponseError(c, response.ResponseError{
		Data:    account,
		Code:    http.StatusOK,
		Message: constants.SuccessUpdate,
		Errors:  nil,
	})
}

func (ctrl *AccountController) GetMyAccount(c *gin.Context) {
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

	account, resNum := ctrl.accountUsecase.GetByUserId(ctx, payloadToken.UserId)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponse(c, response.Response{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: resNum,
			})
			return
		}
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: resNum,
		})
		return
	}

	response.NewResponse(c, response.Response{
		Data:    account,
		Code:    http.StatusOK,
		Message: resNum,
	})
}

func (ctrl *AccountController) MatchPinAccount(c *gin.Context) {
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

	var req models.MatchPinAccount

	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.UnsupportedContentType,
		})
		return
	}

	if errorType := ctrl.validator.Struct(req); errorType != nil || len(errorType) > 0 {
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: helpers.FilterTagCode(errorType, "Pin"),
		})
		return
	}

	account, resNum := ctrl.accountUsecase.MatchPinByUserId(ctx, uint(payloadToken.UserId), req.Pin)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponse(c, response.Response{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: resNum,
			})
			return
		}
		if resNum == constants.PinNotMatch {
			response.NewResponse(c, response.Response{
				Data:    nil,
				Code:    http.StatusBadRequest,
				Message: resNum,
			})
			return
		}
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: resNum,
		})
		return
	}

	response.NewResponse(c, response.Response{
		Data:    account,
		Code:    http.StatusOK,
		Message: resNum,
	})
}

func (ctrl *AccountController) GetAccount(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Id := c.Param("id")

	account, resNum := ctrl.accountUsecase.GetById(ctx, Id)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponse(c, response.Response{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: resNum,
			})
			return
		}
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: resNum,
		})
		return
	}

	response.NewResponse(c, response.Response{
		Data:    account,
		Code:    http.StatusOK,
		Message: resNum,
	})
}
