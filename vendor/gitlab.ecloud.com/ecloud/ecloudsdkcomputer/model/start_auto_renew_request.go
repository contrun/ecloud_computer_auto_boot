// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartAutoRenewRequest struct {


	StartAutoRenewBody *StartAutoRenewBody `json:"startAutoRenewBody,omitempty"`
}

func (s StartAutoRenewRequest) String() string {
	return utils.Beautify(s)
}

func (s StartAutoRenewRequest) GoString() string {
	return s.String()
}

func (s StartAutoRenewRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartAutoRenewRequest) SetStartAutoRenewBody(v *StartAutoRenewBody) *StartAutoRenewRequest {
	s.StartAutoRenewBody = v
	return s
}


type StartAutoRenewRequestBuilder struct {
	s *StartAutoRenewRequest
}

func NewStartAutoRenewRequestBuilder() *StartAutoRenewRequestBuilder {
	s := &StartAutoRenewRequest{}
	b := &StartAutoRenewRequestBuilder{s: s}
	return b
}

func (b *StartAutoRenewRequestBuilder) StartAutoRenewBody(v *StartAutoRenewBody) *StartAutoRenewRequestBuilder {
	b.s.StartAutoRenewBody = v
	return b
}

func (b *StartAutoRenewRequestBuilder) Build() *StartAutoRenewRequest {
	return b.s
}


