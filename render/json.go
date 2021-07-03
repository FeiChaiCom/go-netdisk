package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success = iota
	Failure
	ValidateError
	NotFoundError
	CreateError
	DeleteError
	UpdateError
)

// Define the response format
type Response struct {
	Result  bool        `json:"result"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func JSONResponse(c *gin.Context, r *Response) {
	c.JSON(http.StatusOK, r)
}

func Ok(c *gin.Context, data interface{}) {
	JSONResponse(c, &Response{
		Code:    Success,
		Result:  true,
		Data:    data,
		Message: "success",
	})
}

func OkWithMessage(c *gin.Context, data interface{}, message string) {
	JSONResponse(c, &Response{
		Code:    Success,
		Result:  true,
		Data:    data,
		Message: message,
	})
}

func Fail(c *gin.Context, message string) {
	JSONResponse(c, &Response{
		Code:    Failure,
		Result:  false,
		Data:    nil,
		Message: message,
	})
}

func FailWithCode(c *gin.Context, message string, code int) {
	JSONResponse(c, &Response{
		Code:    code,
		Result:  false,
		Data:    nil,
		Message: message,
	})
}

func FailWithError(c *gin.Context, err error) {
	JSONResponse(c, &Response{
		Code:    Failure,
		Result:  false,
		Data:    nil,
		Message: err.Error(),
	})
}
