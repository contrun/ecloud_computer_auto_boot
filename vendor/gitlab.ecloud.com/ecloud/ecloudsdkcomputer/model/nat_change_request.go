// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type NatChangeRequest struct {


	NatChangeBody *NatChangeBody `json:"natChangeBody,omitempty"`
}

func (s NatChangeRequest) String() string {
	return utils.Beautify(s)
}

func (s NatChangeRequest) GoString() string {
	return s.String()
}

func (s NatChangeRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *NatChangeRequest) SetNatChangeBody(v *NatChangeBody) *NatChangeRequest {
	s.NatChangeBody = v
	return s
}


type NatChangeRequestBuilder struct {
	s *NatChangeRequest
}

func NewNatChangeRequestBuilder() *NatChangeRequestBuilder {
	s := &NatChangeRequest{}
	b := &NatChangeRequestBuilder{s: s}
	return b
}

func (b *NatChangeRequestBuilder) NatChangeBody(v *NatChangeBody) *NatChangeRequestBuilder {
	b.s.NatChangeBody = v
	return b
}

func (b *NatChangeRequestBuilder) Build() *NatChangeRequest {
	return b.s
}


