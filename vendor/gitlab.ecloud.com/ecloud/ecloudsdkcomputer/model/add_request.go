// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddRequest struct {


	AddBody *AddBody `json:"addBody,omitempty"`
}

func (s AddRequest) String() string {
	return utils.Beautify(s)
}

func (s AddRequest) GoString() string {
	return s.String()
}

func (s AddRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddRequest) SetAddBody(v *AddBody) *AddRequest {
	s.AddBody = v
	return s
}


type AddRequestBuilder struct {
	s *AddRequest
}

func NewAddRequestBuilder() *AddRequestBuilder {
	s := &AddRequest{}
	b := &AddRequestBuilder{s: s}
	return b
}

func (b *AddRequestBuilder) AddBody(v *AddBody) *AddRequestBuilder {
	b.s.AddBody = v
	return b
}

func (b *AddRequestBuilder) Build() *AddRequest {
	return b.s
}


