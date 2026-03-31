// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTwoFactorAuthResponseBody struct {

    // 双因子登录设置 0:开 1:关
	TwoFactorAuth *int32 `json:"twoFactorAuth,omitempty"`
}

func (s GetTwoFactorAuthResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetTwoFactorAuthResponseBody) GoString() string {
	return s.String()
}

func (s GetTwoFactorAuthResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTwoFactorAuthResponseBody) SetTwoFactorAuth(v int32) *GetTwoFactorAuthResponseBody {
	s.TwoFactorAuth = &v
	return s
}


type GetTwoFactorAuthResponseBodyBuilder struct {
	s *GetTwoFactorAuthResponseBody
}

func NewGetTwoFactorAuthResponseBodyBuilder() *GetTwoFactorAuthResponseBodyBuilder {
	s := &GetTwoFactorAuthResponseBody{}
	b := &GetTwoFactorAuthResponseBodyBuilder{s: s}
	return b
}

func (b *GetTwoFactorAuthResponseBodyBuilder) TwoFactorAuth(v int32) *GetTwoFactorAuthResponseBodyBuilder {
	b.s.TwoFactorAuth = &v
	return b
}

func (b *GetTwoFactorAuthResponseBodyBuilder) Build() *GetTwoFactorAuthResponseBody {
	return b.s
}


