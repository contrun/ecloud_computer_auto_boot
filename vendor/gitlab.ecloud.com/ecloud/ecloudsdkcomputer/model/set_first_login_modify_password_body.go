// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetFirstLoginModifyPasswordBody struct {
    position.Body
    // 首次登录强制修改密码开关 默认 0:关  1:开
	FirstLoginModifyPassword *int32 `json:"firstLoginModifyPassword,omitempty"`
}

func (s SetFirstLoginModifyPasswordBody) String() string {
	return utils.Beautify(s)
}

func (s SetFirstLoginModifyPasswordBody) GoString() string {
	return s.String()
}

func (s SetFirstLoginModifyPasswordBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetFirstLoginModifyPasswordBody) SetFirstLoginModifyPassword(v int32) *SetFirstLoginModifyPasswordBody {
	s.FirstLoginModifyPassword = &v
	return s
}


type SetFirstLoginModifyPasswordBodyBuilder struct {
	s *SetFirstLoginModifyPasswordBody
}

func NewSetFirstLoginModifyPasswordBodyBuilder() *SetFirstLoginModifyPasswordBodyBuilder {
	s := &SetFirstLoginModifyPasswordBody{}
	b := &SetFirstLoginModifyPasswordBodyBuilder{s: s}
	return b
}

func (b *SetFirstLoginModifyPasswordBodyBuilder) FirstLoginModifyPassword(v int32) *SetFirstLoginModifyPasswordBodyBuilder {
	b.s.FirstLoginModifyPassword = &v
	return b
}

func (b *SetFirstLoginModifyPasswordBodyBuilder) Build() *SetFirstLoginModifyPasswordBody {
	return b.s
}


