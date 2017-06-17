package module

import "net/http"

type Request struct {
	// HTTP请求
	httpReq *http.Request

	// 请求深度
	depth int
}

// 构造器
func NewRequest(httpReq * http.Request, depth int) *Request {
	return &Request{httpReq:httpReq, depth:depth}
}

// 获取Http请求
func (req *Request) HttpReq() (*http.Request) {
	return req.httpReq
}

// 获取请求的深度
func (req *Request) Depth() (int) {
	return req.depth
}
