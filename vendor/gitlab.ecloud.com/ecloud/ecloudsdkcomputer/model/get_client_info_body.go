// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetClientInfoBody struct {
    position.Body
    // 密码
	Password *string `json:"password,omitempty"`
    // 登录账号
	Username *string `json:"username,omitempty"`
}

func (s GetClientInfoBody) String() string {
	return utils.Beautify(s)
}

func (s GetClientInfoBody) GoString() string {
	return s.String()
}

func (s GetClientInfoBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetClientInfoBody) SetPassword(v string) *GetClientInfoBody {
	s.Password = &v
	return s
}

func (s *GetClientInfoBody) SetUsername(v string) *GetClientInfoBody {
	s.Username = &v
	return s
}


type GetClientInfoBodyBuilder struct {
	s *GetClientInfoBody
}

func NewGetClientInfoBodyBuilder() *GetClientInfoBodyBuilder {
	s := &GetClientInfoBody{}
	b := &GetClientInfoBodyBuilder{s: s}
	return b
}

func (b *GetClientInfoBodyBuilder) Password(v string) *GetClientInfoBodyBuilder {
	b.s.Password = &v
	return b
}

func (b *GetClientInfoBodyBuilder) Username(v string) *GetClientInfoBodyBuilder {
	b.s.Username = &v
	return b
}

func (b *GetClientInfoBodyBuilder) Build() *GetClientInfoBody {
	return b.s
}


