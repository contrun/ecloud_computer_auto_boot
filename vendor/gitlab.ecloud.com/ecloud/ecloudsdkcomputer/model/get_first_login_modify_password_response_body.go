// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetFirstLoginModifyPasswordResponseBody struct {

    // 首次登录强制修改密码开关 默认 0:关  1:开
	FirstLoginModifyPassword *int32 `json:"firstLoginModifyPassword,omitempty"`
}

func (s GetFirstLoginModifyPasswordResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetFirstLoginModifyPasswordResponseBody) GoString() string {
	return s.String()
}

func (s GetFirstLoginModifyPasswordResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetFirstLoginModifyPasswordResponseBody) SetFirstLoginModifyPassword(v int32) *GetFirstLoginModifyPasswordResponseBody {
	s.FirstLoginModifyPassword = &v
	return s
}


type GetFirstLoginModifyPasswordResponseBodyBuilder struct {
	s *GetFirstLoginModifyPasswordResponseBody
}

func NewGetFirstLoginModifyPasswordResponseBodyBuilder() *GetFirstLoginModifyPasswordResponseBodyBuilder {
	s := &GetFirstLoginModifyPasswordResponseBody{}
	b := &GetFirstLoginModifyPasswordResponseBodyBuilder{s: s}
	return b
}

func (b *GetFirstLoginModifyPasswordResponseBodyBuilder) FirstLoginModifyPassword(v int32) *GetFirstLoginModifyPasswordResponseBodyBuilder {
	b.s.FirstLoginModifyPassword = &v
	return b
}

func (b *GetFirstLoginModifyPasswordResponseBodyBuilder) Build() *GetFirstLoginModifyPasswordResponseBody {
	return b.s
}


