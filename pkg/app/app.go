package app

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) ToResponse()      {}
func (r *Response) ToListResponse()  {}
func (r *Response) ToErrorResponse() {}
