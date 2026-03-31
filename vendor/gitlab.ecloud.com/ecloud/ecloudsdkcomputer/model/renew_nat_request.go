// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewNatRequest struct {


	RenewNatBody *RenewNatBody `json:"renewNatBody,omitempty"`
}

func (s RenewNatRequest) String() string {
	return utils.Beautify(s)
}

func (s RenewNatRequest) GoString() string {
	return s.String()
}

func (s RenewNatRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewNatRequest) SetRenewNatBody(v *RenewNatBody) *RenewNatRequest {
	s.RenewNatBody = v
	return s
}


type RenewNatRequestBuilder struct {
	s *RenewNatRequest
}

func NewRenewNatRequestBuilder() *RenewNatRequestBuilder {
	s := &RenewNatRequest{}
	b := &RenewNatRequestBuilder{s: s}
	return b
}

func (b *RenewNatRequestBuilder) RenewNatBody(v *RenewNatBody) *RenewNatRequestBuilder {
	b.s.RenewNatBody = v
	return b
}

func (b *RenewNatRequestBuilder) Build() *RenewNatRequest {
	return b.s
}


