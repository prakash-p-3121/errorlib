package errorlib

import (
	"github.com/gin-gonic/gin"
	restResponder "github.com/prakash-p-3121/rest-response-lib"
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
	restResponder.BadReqResponse(ctx, err.errorDescription)
}
