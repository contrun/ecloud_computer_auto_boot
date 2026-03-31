// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitNatRequest struct {


	SubmitNatBody *SubmitNatBody `json:"submitNatBody,omitempty"`
}

func (s SubmitNatRequest) String() string {
	return utils.Beautify(s)
}

func (s SubmitNatRequest) GoString() string {
	return s.String()
}

func (s SubmitNatRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitNatRequest) SetSubmitNatBody(v *SubmitNatBody) *SubmitNatRequest {
	s.SubmitNatBody = v
	return s
}


type SubmitNatRequestBuilder struct {
	s *SubmitNatRequest
}

func NewSubmitNatRequestBuilder() *SubmitNatRequestBuilder {
	s := &SubmitNatRequest{}
	b := &SubmitNatRequestBuilder{s: s}
	return b
}

func (b *SubmitNatRequestBuilder) SubmitNatBody(v *SubmitNatBody) *SubmitNatRequestBuilder {
	b.s.SubmitNatBody = v
	return b
}

func (b *SubmitNatRequestBuilder) Build() *SubmitNatRequest {
	return b.s
}


