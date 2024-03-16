package errorlib

import (
	"github.com/gin-gonic/gin"
	restlib "github.com/prakash-p-3121/restlib"
)

type AppErrorImpl struct {
	errorDescription string
}

func (err *AppErrorImpl) Error() string {
	return err.errorDescription
}

func (err *AppErrorImpl) SendRestResponse(ctx *gin.Context) {
	restlib.InternalServerErrorResponse(ctx, RestError{Reason: err.errorDescription})
}
