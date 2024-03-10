package errorlib

import (
	"github.com/gin-gonic/gin"
	restResponder "github.com/prakash-p-3121/rest-response-lib"
)

type ConflictErrorImpl struct {
	AppErrorImpl
}

func NewConflictError(desc string) AppError {
	return &ConflictErrorImpl{AppErrorImpl{
		errorDescription: desc,
	}}
}

func (err *ConflictErrorImpl) Error() string {
	return err.errorDescription
}

func (err *ConflictErrorImpl) SendRestResponse(ctx *gin.Context) {
	restResponder.ConlictResponse(ctx, err.errorDescription)
}
