// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserModifyPasswordResponseBody struct {

    // 用户修改密码开关 1:开启 0:关闭
	UserModifyPassword *int32 `json:"userModifyPassword,omitempty"`
}

func (s GetUserModifyPasswordResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetUserModifyPasswordResponseBody) GoString() string {
	return s.String()
}

func (s GetUserModifyPasswordResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserModifyPasswordResponseBody) SetUserModifyPassword(v int32) *GetUserModifyPasswordResponseBody {
	s.UserModifyPassword = &v
	return s
}


type GetUserModifyPasswordResponseBodyBuilder struct {
	s *GetUserModifyPasswordResponseBody
}

func NewGetUserModifyPasswordResponseBodyBuilder() *GetUserModifyPasswordResponseBodyBuilder {
	s := &GetUserModifyPasswordResponseBody{}
	b := &GetUserModifyPasswordResponseBodyBuilder{s: s}
	return b
}

func (b *GetUserModifyPasswordResponseBodyBuilder) UserModifyPassword(v int32) *GetUserModifyPasswordResponseBodyBuilder {
	b.s.UserModifyPassword = &v
	return b
}

func (b *GetUserModifyPasswordResponseBodyBuilder) Build() *GetUserModifyPasswordResponseBody {
	return b.s
}


