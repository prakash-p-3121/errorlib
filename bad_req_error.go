package errorlib

import (
	"github.com/gin-gonic/gin"
	restlib "github.com/prakash-p-3121/restlib"
)

type BadReqErrorImpl struct {
	AppErrorImpl
}

func NewBadReqError(desc string) AppError {
	return &BadReqErrorImpl{AppErrorImpl{
		errorDescription: desc,
	}}
}

func (err *BadReqErrorImpl) Error() string {
	return err.errorDescription
}

func (err *BadReqErrorImpl) SendRestResponse(ctx *gin.Context) {
	restlib.BadReqResponse(ctx, err.errorDescription)
}
