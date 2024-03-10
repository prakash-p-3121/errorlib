package errorlib

import (
	"github.com/gin-gonic/gin"
	restResponder "github.com/prakash-p-3121/rest-response-lib"
)

type AppErrorImpl struct {
	errorDescription string
}

func (err *AppErrorImpl) Error() string {
	return err.errorDescription
}

func (err *AppErrorImpl) SendRestResponse(ctx *gin.Context) {
	restResponder.InternalServerErrorResponse(ctx, err.errorDescription)
}
