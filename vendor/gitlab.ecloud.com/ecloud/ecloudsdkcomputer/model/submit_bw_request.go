// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitBwRequest struct {


	SubmitBwBody *SubmitBwBody `json:"submitBwBody,omitempty"`
}

func (s SubmitBwRequest) String() string {
	return utils.Beautify(s)
}

func (s SubmitBwRequest) GoString() string {
	return s.String()
}

func (s SubmitBwRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitBwRequest) SetSubmitBwBody(v *SubmitBwBody) *SubmitBwRequest {
	s.SubmitBwBody = v
	return s
}


type SubmitBwRequestBuilder struct {
	s *SubmitBwRequest
}

func NewSubmitBwRequestBuilder() *SubmitBwRequestBuilder {
	s := &SubmitBwRequest{}
	b := &SubmitBwRequestBuilder{s: s}
	return b
}

func (b *SubmitBwRequestBuilder) SubmitBwBody(v *SubmitBwBody) *SubmitBwRequestBuilder {
	b.s.SubmitBwBody = v
	return b
}

func (b *SubmitBwRequestBuilder) Build() *SubmitBwRequest {
	return b.s
}


