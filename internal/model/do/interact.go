// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Interact is the golang structure of table gf_interact for DAO operations like Where/Data.
type Interact struct {
	g.Meta     `orm:"table:gf_interact, do:true"`
	Id         interface{} // 自增ID
	Type       interface{} // 操作类型。0:赞，1:踩。
	UserId     interface{} // 操作用户
	TargetId   interface{} // 对应内容ID，该内容可能是content, reply
	TargetType interface{} // 内容模型: content, reply, 具体由程序定义
	Count      interface{} // 操作数据值
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}