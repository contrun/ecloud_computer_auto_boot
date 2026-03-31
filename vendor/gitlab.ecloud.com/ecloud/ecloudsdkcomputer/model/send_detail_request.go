// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SendDetailRequest struct {


	SendDetailBody *SendDetailBody `json:"sendDetailBody,omitempty"`
}

func (s SendDetailRequest) String() string {
	return utils.Beautify(s)
}

func (s SendDetailRequest) GoString() string {
	return s.String()
}

func (s SendDetailRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SendDetailRequest) SetSendDetailBody(v *SendDetailBody) *SendDetailRequest {
	s.SendDetailBody = v
	return s
}


type SendDetailRequestBuilder struct {
	s *SendDetailRequest
}

func NewSendDetailRequestBuilder() *SendDetailRequestBuilder {
	s := &SendDetailRequest{}
	b := &SendDetailRequestBuilder{s: s}
	return b
}

func (b *SendDetailRequestBuilder) SendDetailBody(v *SendDetailBody) *SendDetailRequestBuilder {
	b.s.SendDetailBody = v
	return b
}

func (b *SendDetailRequestBuilder) Build() *SendDetailRequest {
	return b.s
}


