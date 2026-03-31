// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ShareImageOperationRequest struct {


	ShareImageOperationBody *ShareImageOperationBody `json:"shareImageOperationBody,omitempty"`
}

func (s ShareImageOperationRequest) String() string {
	return utils.Beautify(s)
}

func (s ShareImageOperationRequest) GoString() string {
	return s.String()
}

func (s ShareImageOperationRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ShareImageOperationRequest) SetShareImageOperationBody(v *ShareImageOperationBody) *ShareImageOperationRequest {
	s.ShareImageOperationBody = v
	return s
}


type ShareImageOperationRequestBuilder struct {
	s *ShareImageOperationRequest
}

func NewShareImageOperationRequestBuilder() *ShareImageOperationRequestBuilder {
	s := &ShareImageOperationRequest{}
	b := &ShareImageOperationRequestBuilder{s: s}
	return b
}

func (b *ShareImageOperationRequestBuilder) ShareImageOperationBody(v *ShareImageOperationBody) *ShareImageOperationRequestBuilder {
	b.s.ShareImageOperationBody = v
	return b
}

func (b *ShareImageOperationRequestBuilder) Build() *ShareImageOperationRequest {
	return b.s
}


