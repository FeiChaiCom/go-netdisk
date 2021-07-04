package render

import (
	"github.com/gin-gonic/gin"
	cfg "go-netdisk/config"
	"go-netdisk/utils"
	"log"
	"net/http"
)

// ErrCode
type ErrCode int

const (
	Success ErrCode = iota
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
	Code    ErrCode     `json:"code"`
	Message string      `json:"msg"`
}

func JSONResponse(c *gin.Context, r *Response) {
	// Log response data for debug
	if cfg.DebugOn {
		log.Printf("---------------------\n"+
			"FullPath:\t%s\n",
			c.FullPath(),
		)

		log.Println(utils.PrettyJson(r))
	}

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

func OkOnly(c *gin.Context) {
	JSONResponse(c, &Response{
		Code:    Success,
		Result:  true,
		Data:    nil,
		Message: "success",
	})
}

func OkWithMsg(c *gin.Context, data interface{}, message string) {
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

func FailWithCode(c *gin.Context, message string, code ErrCode) {
	JSONResponse(c, &Response{
		Code:    code,
		Result:  false,
		Data:    nil,
		Message: message,
	})
}

func Error(c *gin.Context, err error) {
	JSONResponse(c, &Response{
		Code:    Failure,
		Result:  false,
		Data:    nil,
		Message: err.Error(),
	})
}
