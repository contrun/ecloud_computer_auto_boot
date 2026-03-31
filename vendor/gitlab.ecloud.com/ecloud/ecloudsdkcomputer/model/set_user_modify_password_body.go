// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetUserModifyPasswordBody struct {
    position.Body
    // 用户修改密码开关 1:开启 0:关闭
	UserModifyPassword *int32 `json:"userModifyPassword,omitempty"`
}

func (s SetUserModifyPasswordBody) String() string {
	return utils.Beautify(s)
}

func (s SetUserModifyPasswordBody) GoString() string {
	return s.String()
}

func (s SetUserModifyPasswordBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetUserModifyPasswordBody) SetUserModifyPassword(v int32) *SetUserModifyPasswordBody {
	s.UserModifyPassword = &v
	return s
}


type SetUserModifyPasswordBodyBuilder struct {
	s *SetUserModifyPasswordBody
}

func NewSetUserModifyPasswordBodyBuilder() *SetUserModifyPasswordBodyBuilder {
	s := &SetUserModifyPasswordBody{}
	b := &SetUserModifyPasswordBodyBuilder{s: s}
	return b
}

func (b *SetUserModifyPasswordBodyBuilder) UserModifyPassword(v int32) *SetUserModifyPasswordBodyBuilder {
	b.s.UserModifyPassword = &v
	return b
}

func (b *SetUserModifyPasswordBodyBuilder) Build() *SetUserModifyPasswordBody {
	return b.s
}


