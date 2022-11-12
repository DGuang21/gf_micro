// Code generated by protoc-gen-go-ghttp. DO NOT EDIT.
// versions:
// 	protoc-gen-go-http v0.0.1
// 	protoc             v3.21.5
// source: echo/v1/echo.proto

package v1

import (
	context "context"

	gcode "github.com/gogf/gf/v2/errors/gcode"
	gerr "github.com/gogf/gf/v2/errors/gerror"
	g "github.com/gogf/gf/v2/frame/g"
)

var _ = g.Meta{}
var _ = gerr.Error{}
var _ = context.Background()
var notImplErrorCode = gcode.New(-1, "", nil)
var _ = notImplErrorCode

// UnimplementedEchoServer
type UnimplementedEchoServer struct {
	impl EchoImpl
}

// NewEchoApi is an entry that must be implemented.
func NewEchoApi(impl EchoImpl) UnimplementedEchoServer {
	return UnimplementedEchoServer{impl: impl}
}

// SayReq  SayReq is the request message for the Echo.Say method.
type SayReq struct {
	g.Meta     `path:"POST" method:"/v1/echo/say"`
	Content    string             `v:"required|max:64" d:"hello world" ` // 提交内容
	Nickname   string             // only comment,not rule
	Sex        string             // tail comment
	Data       *SayRes            `d:"nil" ` // 结构体调用
	MapData    map[string]*SayRes `d:"nil" ` // map 调用
	ArrayData  []*SayRes          `d:"nil" ` // 数组调用
	IntData    int32              // int
	Uint32Data uint32             // uint32
	Uint64Data int64              `d:"0" v:"required" json:"uint64_data" ` // uint64
}

// SayRes  SayRes is the response message for the Echo.Say method.
type SayRes struct {
	Content string
}

// Say  Echo returns the same message it receives.
func (Echo UnimplementedEchoServer) Say(ctx context.Context, req *SayReq) (*SayRes, error) {
	return nil, gerr.NewCode(notImplErrorCode, "Method Say not implemented.")
}

// EchoImpl is the server API for Echo service.
type EchoImpl interface {
	Say(ctx context.Context, req *SayReq) (*SayRes, error)
}
