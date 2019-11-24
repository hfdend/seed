package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"seed/cli"
	"seed/errors"

	"github.com/gin-gonic/gin"
)

type mp map[string]interface{}

type Reply struct {
	Error   int         `json:"error"`
	Type    string      `json:"type,omitempty"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func NewReply(data interface{}) *Reply {
	var (
		err    error
		objErr *errors.Error
		ok     bool
	)
	r := new(Reply)
	if objErr, ok = data.(*errors.Error); ok {
		r.Error = 1
		r.Code = objErr.Code
		msg := objErr.Err.Error()
		r.Message = msg
		// 需要翻译的错误
		return r
	} else if err, ok = data.(error); ok {
		_, file, line, ok := runtime.Caller(2)

		if ok {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		cli.Logger.With("file", file, "line", line).Error(err)
		r.Error = 1
		r.Code = errors.System
		r.Message = "system error"
		return r
	}
	r.Error = 0
	if bts, ok := data.([]byte); ok {
		r.Data = json.RawMessage(bts)
	} else {
		r.Data = data
	}
	return r
}

func JSONStatus(c *gin.Context, statusCode int, v interface{}) {
	c.JSON(statusCode, NewReply(v))
}

func JSON(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, NewReply(v))
	c.Abort()
}

func JSONCode(c *gin.Context, v interface{}) {
	var code = http.StatusOK
	if _, ok := v.(error); ok {
		code = http.StatusInternalServerError
	}
	c.JSON(code, NewReply(v))
	c.Abort()
}

func Bind(c *gin.Context, v interface{}) error {
	if err := c.Bind(v); err != nil {
		JSONStatus(c, http.StatusBadRequest, errors.New(err.Error()))
		return err
	}
	return nil
}
