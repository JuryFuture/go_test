package module

import "net/http"

type Response struct {
	// HTTP响应
	httpRes *http.Response

	// 响应的深度
	depth int
}

// 构造器
func NewResponse(httpRes *http.Response, depth int) *Response {
	return &Response{httpRes:httpRes, depth:depth}
}

// 获取HTTP响应
func (res *Response) HeepRes() *http.Response {
	return res.httpRes
}

// 获取响应深度
func (res *Response) Depth() int {
	return res.depth
}

// 用于判断响应是否有效
func (res *Response) Valid() bool {
	return res.httpRes != nil && res.httpRes.Body != nil
}
