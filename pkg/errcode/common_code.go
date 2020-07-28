package errcode

var (
	Success         = NewError(200, "OK")
	TooManyRequests = NewError(429, "太多的请求")
)
