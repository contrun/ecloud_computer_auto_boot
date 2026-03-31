// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OrderChangeResponseData struct {

    // instanceId
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单编号
	OrderId *string `json:"orderId,omitempty"`
}

func (s OrderChangeResponseData) String() string {
	return utils.Beautify(s)
}

func (s OrderChangeResponseData) GoString() string {
	return s.String()
}

func (s OrderChangeResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OrderChangeResponseData) SetInstanceId(v string) *OrderChangeResponseData {
	s.InstanceId = &v
	return s
}

func (s *OrderChangeResponseData) SetOrderId(v string) *OrderChangeResponseData {
	s.OrderId = &v
	return s
}


type OrderChangeResponseDataBuilder struct {
	s *OrderChangeResponseData
}

func NewOrderChangeResponseDataBuilder() *OrderChangeResponseDataBuilder {
	s := &OrderChangeResponseData{}
	b := &OrderChangeResponseDataBuilder{s: s}
	return b
}

func (b *OrderChangeResponseDataBuilder) InstanceId(v string) *OrderChangeResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *OrderChangeResponseDataBuilder) OrderId(v string) *OrderChangeResponseDataBuilder {
	b.s.OrderId = &v
	return b
}

func (b *OrderChangeResponseDataBuilder) Build() *OrderChangeResponseData {
	return b.s
}


