// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnifyInstanceQueryRequest struct {


	UnifyInstanceQueryBody *UnifyInstanceQueryBody `json:"unifyInstanceQueryBody,omitempty"`
}

func (s UnifyInstanceQueryRequest) String() string {
	return utils.Beautify(s)
}

func (s UnifyInstanceQueryRequest) GoString() string {
	return s.String()
}

func (s UnifyInstanceQueryRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnifyInstanceQueryRequest) SetUnifyInstanceQueryBody(v *UnifyInstanceQueryBody) *UnifyInstanceQueryRequest {
	s.UnifyInstanceQueryBody = v
	return s
}


type UnifyInstanceQueryRequestBuilder struct {
	s *UnifyInstanceQueryRequest
}

func NewUnifyInstanceQueryRequestBuilder() *UnifyInstanceQueryRequestBuilder {
	s := &UnifyInstanceQueryRequest{}
	b := &UnifyInstanceQueryRequestBuilder{s: s}
	return b
}

func (b *UnifyInstanceQueryRequestBuilder) UnifyInstanceQueryBody(v *UnifyInstanceQueryBody) *UnifyInstanceQueryRequestBuilder {
	b.s.UnifyInstanceQueryBody = v
	return b
}

func (b *UnifyInstanceQueryRequestBuilder) Build() *UnifyInstanceQueryRequest {
	return b.s
}


