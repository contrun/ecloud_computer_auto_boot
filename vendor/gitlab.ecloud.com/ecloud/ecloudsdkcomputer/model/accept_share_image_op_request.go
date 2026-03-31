// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AcceptShareImageOpRequest struct {


	AcceptShareImageOpBody *AcceptShareImageOpBody `json:"acceptShareImageOpBody,omitempty"`
}

func (s AcceptShareImageOpRequest) String() string {
	return utils.Beautify(s)
}

func (s AcceptShareImageOpRequest) GoString() string {
	return s.String()
}

func (s AcceptShareImageOpRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AcceptShareImageOpRequest) SetAcceptShareImageOpBody(v *AcceptShareImageOpBody) *AcceptShareImageOpRequest {
	s.AcceptShareImageOpBody = v
	return s
}


type AcceptShareImageOpRequestBuilder struct {
	s *AcceptShareImageOpRequest
}

func NewAcceptShareImageOpRequestBuilder() *AcceptShareImageOpRequestBuilder {
	s := &AcceptShareImageOpRequest{}
	b := &AcceptShareImageOpRequestBuilder{s: s}
	return b
}

func (b *AcceptShareImageOpRequestBuilder) AcceptShareImageOpBody(v *AcceptShareImageOpBody) *AcceptShareImageOpRequestBuilder {
	b.s.AcceptShareImageOpBody = v
	return b
}

func (b *AcceptShareImageOpRequestBuilder) Build() *AcceptShareImageOpRequest {
	return b.s
}


