// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ReallocateRequest struct {


	ReallocateBody *ReallocateBody `json:"reallocateBody,omitempty"`
}

func (s ReallocateRequest) String() string {
	return utils.Beautify(s)
}

func (s ReallocateRequest) GoString() string {
	return s.String()
}

func (s ReallocateRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ReallocateRequest) SetReallocateBody(v *ReallocateBody) *ReallocateRequest {
	s.ReallocateBody = v
	return s
}


type ReallocateRequestBuilder struct {
	s *ReallocateRequest
}

func NewReallocateRequestBuilder() *ReallocateRequestBuilder {
	s := &ReallocateRequest{}
	b := &ReallocateRequestBuilder{s: s}
	return b
}

func (b *ReallocateRequestBuilder) ReallocateBody(v *ReallocateBody) *ReallocateRequestBuilder {
	b.s.ReallocateBody = v
	return b
}

func (b *ReallocateRequestBuilder) Build() *ReallocateRequest {
	return b.s
}


