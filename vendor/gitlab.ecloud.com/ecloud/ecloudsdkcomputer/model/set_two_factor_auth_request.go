// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetTwoFactorAuthRequest struct {


	SetTwoFactorAuthBody *SetTwoFactorAuthBody `json:"setTwoFactorAuthBody,omitempty"`
}

func (s SetTwoFactorAuthRequest) String() string {
	return utils.Beautify(s)
}

func (s SetTwoFactorAuthRequest) GoString() string {
	return s.String()
}

func (s SetTwoFactorAuthRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetTwoFactorAuthRequest) SetSetTwoFactorAuthBody(v *SetTwoFactorAuthBody) *SetTwoFactorAuthRequest {
	s.SetTwoFactorAuthBody = v
	return s
}


type SetTwoFactorAuthRequestBuilder struct {
	s *SetTwoFactorAuthRequest
}

func NewSetTwoFactorAuthRequestBuilder() *SetTwoFactorAuthRequestBuilder {
	s := &SetTwoFactorAuthRequest{}
	b := &SetTwoFactorAuthRequestBuilder{s: s}
	return b
}

func (b *SetTwoFactorAuthRequestBuilder) SetTwoFactorAuthBody(v *SetTwoFactorAuthBody) *SetTwoFactorAuthRequestBuilder {
	b.s.SetTwoFactorAuthBody = v
	return b
}

func (b *SetTwoFactorAuthRequestBuilder) Build() *SetTwoFactorAuthRequest {
	return b.s
}


