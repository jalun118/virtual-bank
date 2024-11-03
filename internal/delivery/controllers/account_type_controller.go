package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"virtual-bank/internal/constants"
	"virtual-bank/internal/response"
	"virtual-bank/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AccountTypeController struct {
	accountTypeUsecase *usecase.AccountTypeUsecase
}

func NewAccountTypeController(accountTypeUsecase *usecase.AccountTypeUsecase) *AccountTypeController {
	return &AccountTypeController{
		accountTypeUsecase: accountTypeUsecase,
	}
}

func (ctrl *AccountTypeController) GetAllType(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	acountTypes, resNum := ctrl.accountTypeUsecase.GetAll(ctx)
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordFound {
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
		Data:    acountTypes,
		Code:    http.StatusOK,
		Message: resNum,
	})
}

func (ctrl *AccountTypeController) GetOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusNotFound,
			Message: constants.RecordNotFound,
		})
		return
	}

	if Id < 1 {
		response.NewResponse(c, response.Response{
			Data:    nil,
			Code:    http.StatusNotFound,
			Message: constants.RecordNotFound,
		})
		return
	}

	acountType, resNum := ctrl.accountTypeUsecase.GetById(ctx, uint(Id))
	if resNum != constants.SuccessRead {
		if resNum == constants.RecordFound {
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
		Data:    acountType,
		Code:    http.StatusOK,
		Message: resNum,
	})
}
