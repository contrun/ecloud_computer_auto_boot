// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartAutoRenewNatRequest struct {


	StartAutoRenewNatBody *StartAutoRenewNatBody `json:"startAutoRenewNatBody,omitempty"`
}

func (s StartAutoRenewNatRequest) String() string {
	return utils.Beautify(s)
}

func (s StartAutoRenewNatRequest) GoString() string {
	return s.String()
}

func (s StartAutoRenewNatRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartAutoRenewNatRequest) SetStartAutoRenewNatBody(v *StartAutoRenewNatBody) *StartAutoRenewNatRequest {
	s.StartAutoRenewNatBody = v
	return s
}


type StartAutoRenewNatRequestBuilder struct {
	s *StartAutoRenewNatRequest
}

func NewStartAutoRenewNatRequestBuilder() *StartAutoRenewNatRequestBuilder {
	s := &StartAutoRenewNatRequest{}
	b := &StartAutoRenewNatRequestBuilder{s: s}
	return b
}

func (b *StartAutoRenewNatRequestBuilder) StartAutoRenewNatBody(v *StartAutoRenewNatBody) *StartAutoRenewNatRequestBuilder {
	b.s.StartAutoRenewNatBody = v
	return b
}

func (b *StartAutoRenewNatRequestBuilder) Build() *StartAutoRenewNatRequest {
	return b.s
}


