// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewBwRequest struct {


	RenewBwBody *RenewBwBody `json:"renewBwBody,omitempty"`
}

func (s RenewBwRequest) String() string {
	return utils.Beautify(s)
}

func (s RenewBwRequest) GoString() string {
	return s.String()
}

func (s RenewBwRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewBwRequest) SetRenewBwBody(v *RenewBwBody) *RenewBwRequest {
	s.RenewBwBody = v
	return s
}


type RenewBwRequestBuilder struct {
	s *RenewBwRequest
}

func NewRenewBwRequestBuilder() *RenewBwRequestBuilder {
	s := &RenewBwRequest{}
	b := &RenewBwRequestBuilder{s: s}
	return b
}

func (b *RenewBwRequestBuilder) RenewBwBody(v *RenewBwBody) *RenewBwRequestBuilder {
	b.s.RenewBwBody = v
	return b
}

func (b *RenewBwRequestBuilder) Build() *RenewBwRequest {
	return b.s
}


