package controllers

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"virtual-bank/internal/constants"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/models"
	"virtual-bank/internal/repository"
	"virtual-bank/internal/response"
	"virtual-bank/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionUsecase *usecase.TransactionUsecase
	accountUsecase     *usecase.AccountUsecase
	validator          *helpers.Validation
	mu                 sync.Mutex
}

func NewTransactionController(transactionUsecase *usecase.TransactionUsecase, accountUsecase *usecase.AccountUsecase, validator *helpers.Validation) *TransactionController {
	return &TransactionController{
		validator:          validator,
		transactionUsecase: transactionUsecase,
		accountUsecase:     accountUsecase,
	}
}

type errCreateTransaction struct {
	DestAccountId   int `json:"dest_account_id"`
	TransactionType int `json:"transaction_type"`
	Amount          int `json:"amount"`
	Currency        int `json:"currency"`
	Description     int `json:"description"`
}

func (ctrl *TransactionController) CreateTransaction(c *gin.Context) {
	ctrl.mu.Lock()
	defer ctrl.mu.Unlock()

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

	accountSender, resNum := ctrl.accountUsecase.GetByUserId(ctx, payloadToken.UserId)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponseError(c, response.ResponseError{
				Data:    nil,
				Code:    http.StatusUnauthorized,
				Message: constants.RecordFound,
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

	var req models.CreateTransaction

	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.UnsupportedContentType,
			Errors:  nil,
		})
		return
	}

	req.AccountId = accountSender.ID

	errRes := errCreateTransaction{
		DestAccountId:   0,
		TransactionType: 0,
		Amount:          0,
		Currency:        0,
		Description:     0,
	}

	if errorType := ctrl.validator.Struct(req); errorType != nil || len(errorType) > 0 {
		errRes.Amount = helpers.FilterTagCode(errorType, "Amount")
		errRes.Currency = helpers.FilterTagCode(errorType, "Currency")
		errRes.Description = helpers.FilterTagCode(errorType, "Description")
		errRes.DestAccountId = helpers.FilterTagCode(errorType, "DestAccountId")
		errRes.TransactionType = helpers.FilterTagCode(errorType, "TransactionType")

		response.NewResponseError(c, response.ResponseError{
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: constants.InputError,
			Errors:  errRes,
		})
		return
	}

	transaction, resNum := ctrl.transactionUsecase.Create(ctx, req)
	if resNum != constants.SuccessInsert {
		httpCode := http.StatusBadRequest
		switch resNum {
		case constants.AccountNotFound:
			httpCode = http.StatusNotFound
		case constants.DatabaseError:
			httpCode = http.StatusInternalServerError
		case constants.NotEnoughBalance:
			httpCode = http.StatusBadRequest
			errRes.Amount = resNum
		case constants.DestinationAccountNotFound:
			httpCode = http.StatusNotFound
			errRes.DestAccountId = resNum
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
		Data:    transaction,
		Code:    http.StatusOK,
		Errors:  nil,
		Message: constants.SuccessInsert,
	})
}

func (ctrl *TransactionController) GetOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Id := c.Param("id")

	transaction, resNum := ctrl.transactionUsecase.GetOne(ctx, Id)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
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
		Data:    transaction,
		Code:    http.StatusOK,
		Message: constants.SuccessRead,
	})
}

func (ctrl *TransactionController) GetAllTransaction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := helpers.GetCookieRefreshToken(c)
	payloadToken, err := helpers.DecodeRefreshToken(token)
	if err != nil {
		response.NewResponseMeta(c, response.ResponseMeta{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
			Meta: response.MetaData{
				CurrentPage: 0,
				TotalPage:   0,
				TotalData:   0,
			},
		})
		return
	}

	myAccount, resNum := ctrl.accountUsecase.GetByUserId(ctx, payloadToken.UserId)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponseMeta(c, response.ResponseMeta{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: constants.RecordFound,
				Meta: response.MetaData{
					CurrentPage: 0,
					TotalPage:   0,
					TotalData:   0,
				},
			})
			return
		}
		response.NewResponseMeta(c, response.ResponseMeta{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
			Meta: response.MetaData{
				CurrentPage: 0,
				TotalPage:   0,
				TotalData:   0,
			},
		})
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if limit > 100 {
		limit = 100
	}
	if err != nil {
		limit = 100
	}

	sort := strings.ToLower(c.Query("sort"))
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	orderBy := c.Query("order")
	if helpers.ValidationNotAllowChar(orderBy) {
		orderBy = "transaction_date"
	}
	if orderBy == "" {
		orderBy = "transaction_date"
	}

	transaction, meta, resNum := ctrl.transactionUsecase.GetAllPagination(ctx, myAccount.ID, repository.Pagination{
		Limit:   limit,
		Page:    page,
		Sort:    sort,
		OrderBy: orderBy,
	})
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponseMeta(c, response.ResponseMeta{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: constants.RecordNotFound,
				Meta: response.MetaData{
					CurrentPage: page,
					TotalPage:   0,
					TotalData:   0,
				},
			})
			return
		}
		response.NewResponseMeta(c, response.ResponseMeta{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
			Meta: response.MetaData{
				CurrentPage: page,
				TotalPage:   0,
				TotalData:   0,
			},
		})
		return
	}

	response.NewResponseMeta(c, response.ResponseMeta{
		Data:    transaction,
		Code:    http.StatusOK,
		Message: constants.SuccessRead,
		Meta: response.MetaData{
			CurrentPage: page,
			TotalPage:   meta.TotalPages,
			TotalData:   len(transaction),
		},
	})
}

func (ctrl *TransactionController) GetAllSubmition(c *gin.Context) {
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

	myAccount, resNum := ctrl.accountUsecase.GetByUserId(ctx, payloadToken.UserId)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponse(c, response.Response{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: constants.RecordFound,
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

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if limit > 100 {
		limit = 100
	}
	if err != nil {
		limit = 100
	}

	sort := strings.ToLower(c.Query("sort"))
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	orderBy := c.Query("order")
	if helpers.ValidationNotAllowChar(orderBy) {
		orderBy = "transaction_date"
	}
	if orderBy == "" {
		orderBy = "transaction_date"
	}

	transaction, meta, resNum := ctrl.transactionUsecase.GetAllPaginationSubmissions(ctx, myAccount.ID, repository.Pagination{
		Limit:   limit,
		Page:    page,
		Sort:    sort,
		OrderBy: orderBy,
	})
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordNotFound {
			response.NewResponseMeta(c, response.ResponseMeta{
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: constants.RecordNotFound,
				Meta: response.MetaData{
					CurrentPage: page,
					TotalPage:   0,
					TotalData:   0,
				},
			})
			return
		}
		response.NewResponseMeta(c, response.ResponseMeta{
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: constants.InternalServerError,
			Meta: response.MetaData{
				CurrentPage: page,
				TotalPage:   0,
				TotalData:   0,
			},
		})
		return
	}

	response.NewResponseMeta(c, response.ResponseMeta{
		Data:    transaction,
		Code:    http.StatusOK,
		Message: constants.SuccessRead,
		Meta: response.MetaData{
			CurrentPage: page,
			TotalPage:   meta.TotalPages,
			TotalData:   len(transaction),
		},
	})
}
