// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeRenewRequest struct {


	UnsubscribeRenewBody *UnsubscribeRenewBody `json:"unsubscribeRenewBody,omitempty"`
}

func (s UnsubscribeRenewRequest) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeRenewRequest) GoString() string {
	return s.String()
}

func (s UnsubscribeRenewRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeRenewRequest) SetUnsubscribeRenewBody(v *UnsubscribeRenewBody) *UnsubscribeRenewRequest {
	s.UnsubscribeRenewBody = v
	return s
}


type UnsubscribeRenewRequestBuilder struct {
	s *UnsubscribeRenewRequest
}

func NewUnsubscribeRenewRequestBuilder() *UnsubscribeRenewRequestBuilder {
	s := &UnsubscribeRenewRequest{}
	b := &UnsubscribeRenewRequestBuilder{s: s}
	return b
}

func (b *UnsubscribeRenewRequestBuilder) UnsubscribeRenewBody(v *UnsubscribeRenewBody) *UnsubscribeRenewRequestBuilder {
	b.s.UnsubscribeRenewBody = v
	return b
}

func (b *UnsubscribeRenewRequestBuilder) Build() *UnsubscribeRenewRequest {
	return b.s
}


