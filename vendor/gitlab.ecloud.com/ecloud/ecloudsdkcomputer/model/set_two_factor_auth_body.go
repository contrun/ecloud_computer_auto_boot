// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetTwoFactorAuthBody struct {
    position.Body
    // 双因子登录设置，0开启 1关闭
	TwoFactorAuth *int32 `json:"twoFactorAuth,omitempty"`
}

func (s SetTwoFactorAuthBody) String() string {
	return utils.Beautify(s)
}

func (s SetTwoFactorAuthBody) GoString() string {
	return s.String()
}

func (s SetTwoFactorAuthBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetTwoFactorAuthBody) SetTwoFactorAuth(v int32) *SetTwoFactorAuthBody {
	s.TwoFactorAuth = &v
	return s
}


type SetTwoFactorAuthBodyBuilder struct {
	s *SetTwoFactorAuthBody
}

func NewSetTwoFactorAuthBodyBuilder() *SetTwoFactorAuthBodyBuilder {
	s := &SetTwoFactorAuthBody{}
	b := &SetTwoFactorAuthBodyBuilder{s: s}
	return b
}

func (b *SetTwoFactorAuthBodyBuilder) TwoFactorAuth(v int32) *SetTwoFactorAuthBodyBuilder {
	b.s.TwoFactorAuth = &v
	return b
}

func (b *SetTwoFactorAuthBodyBuilder) Build() *SetTwoFactorAuthBody {
	return b.s
}


