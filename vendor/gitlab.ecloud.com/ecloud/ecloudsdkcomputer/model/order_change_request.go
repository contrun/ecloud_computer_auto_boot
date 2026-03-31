// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OrderChangeRequest struct {


	OrderChangeBody *OrderChangeBody `json:"orderChangeBody,omitempty"`
}

func (s OrderChangeRequest) String() string {
	return utils.Beautify(s)
}

func (s OrderChangeRequest) GoString() string {
	return s.String()
}

func (s OrderChangeRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OrderChangeRequest) SetOrderChangeBody(v *OrderChangeBody) *OrderChangeRequest {
	s.OrderChangeBody = v
	return s
}


type OrderChangeRequestBuilder struct {
	s *OrderChangeRequest
}

func NewOrderChangeRequestBuilder() *OrderChangeRequestBuilder {
	s := &OrderChangeRequest{}
	b := &OrderChangeRequestBuilder{s: s}
	return b
}

func (b *OrderChangeRequestBuilder) OrderChangeBody(v *OrderChangeBody) *OrderChangeRequestBuilder {
	b.s.OrderChangeBody = v
	return b
}

func (b *OrderChangeRequestBuilder) Build() *OrderChangeRequest {
	return b.s
}


