// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeNatRequest struct {


	UnsubscribeNatBody *UnsubscribeNatBody `json:"unsubscribeNatBody,omitempty"`
}

func (s UnsubscribeNatRequest) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeNatRequest) GoString() string {
	return s.String()
}

func (s UnsubscribeNatRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeNatRequest) SetUnsubscribeNatBody(v *UnsubscribeNatBody) *UnsubscribeNatRequest {
	s.UnsubscribeNatBody = v
	return s
}


type UnsubscribeNatRequestBuilder struct {
	s *UnsubscribeNatRequest
}

func NewUnsubscribeNatRequestBuilder() *UnsubscribeNatRequestBuilder {
	s := &UnsubscribeNatRequest{}
	b := &UnsubscribeNatRequestBuilder{s: s}
	return b
}

func (b *UnsubscribeNatRequestBuilder) UnsubscribeNatBody(v *UnsubscribeNatBody) *UnsubscribeNatRequestBuilder {
	b.s.UnsubscribeNatBody = v
	return b
}

func (b *UnsubscribeNatRequestBuilder) Build() *UnsubscribeNatRequest {
	return b.s
}


