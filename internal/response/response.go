package response

import (
	"net/http"
	"virtual-bank/internal/constants"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    any
	Code    int
	Message int
}

func NewResponse(c *gin.Context, res Response) {
	c.JSON(res.Code, gin.H{
		"data":    res.Data,
		"code":    res.Code,
		"message": res.Message,
	})
}

type ResponseError struct {
	Data    any
	Code    int
	Errors  any
	Message int
}

func NewResponseError(c *gin.Context, res ResponseError) {
	c.JSON(res.Code, gin.H{
		"data":    res.Data,
		"code":    res.Code,
		"errors":  res.Errors,
		"message": res.Message,
	})
}

type MetaData struct {
	CurrentPage int
	TotalPage   int
	TotalData   int
}

type ResponseMeta struct {
	Data    any
	Code    int
	Message int
	Meta    MetaData
}

func NewResponseMeta(c *gin.Context, res ResponseMeta) {
	c.JSON(res.Code, gin.H{
		"data":    res.Data,
		"code":    res.Code,
		"message": res.Message,
		"meta": gin.H{
			"current_page": res.Meta.CurrentPage,
			"total_page":   res.Meta.TotalPage,
			"total_data":   res.Meta.TotalData,
		},
	})
}

type ResponseMetaError struct {
	Data    any
	Code    int
	Message int
	Meta    MetaData
	Errors  any
}

func NewResponseMetaError(c *gin.Context, res ResponseMetaError) {
	c.JSON(res.Code, gin.H{
		"data":    res.Data,
		"code":    res.Code,
		"message": res.Message,
		"errors":  res.Errors,
		"meta": gin.H{
			"current_page": res.Meta.CurrentPage,
			"total_page":   res.Meta.TotalPage,
			"total_data":   res.Meta.TotalData,
		},
	})
}

func NewResponseUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"data":    nil,
		"code":    http.StatusUnauthorized,
		"message": constants.AccessUnauthorized,
	})
}
func NewResponseInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    nil,
		"code":    http.StatusInternalServerError,
		"message": constants.InternalServerError,
	})
}
