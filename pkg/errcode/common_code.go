package errcode

var (
	Success             = NewError(200, "OK")
	InvalidParams       = NewError(-1, "入参错误")
	UnAuthorization     = NewError(401, "未授权")
	NotFound            = NewError(404, "未找到")
	TooManyRequests     = NewError(429, "太多的请求")
	InternalServerError = NewError(500, "系统内部错误")
)
