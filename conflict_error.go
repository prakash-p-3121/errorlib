package errorlib

import (
	"github.com/gin-gonic/gin"
	restlib "github.com/prakash-p-3121/restlib"
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
	restlib.ConflictResponse(ctx, err.errorDescription)
}
