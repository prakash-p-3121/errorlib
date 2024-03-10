package errorlib

import "github.com/gin-gonic/gin"

/*
	AppError instances will not be created at all.
	Just the sub-classes of AppError like BadReqError, InternalServerError,
	NotFoudError, ConflictError etc.. will be instantiated.
*/
type AppError interface {
	Error() string
	SendRestResponse(c *gin.Context)
}
