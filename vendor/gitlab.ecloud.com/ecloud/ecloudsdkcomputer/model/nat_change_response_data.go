// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type NatChangeResponseData struct {

    // 资源实例ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单号：MOP-O-XXX
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s NatChangeResponseData) String() string {
	return utils.Beautify(s)
}

func (s NatChangeResponseData) GoString() string {
	return s.String()
}

func (s NatChangeResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *NatChangeResponseData) SetInstanceId(v string) *NatChangeResponseData {
	s.InstanceId = &v
	return s
}

func (s *NatChangeResponseData) SetOrderNo(v string) *NatChangeResponseData {
	s.OrderNo = &v
	return s
}


type NatChangeResponseDataBuilder struct {
	s *NatChangeResponseData
}

func NewNatChangeResponseDataBuilder() *NatChangeResponseDataBuilder {
	s := &NatChangeResponseData{}
	b := &NatChangeResponseDataBuilder{s: s}
	return b
}

func (b *NatChangeResponseDataBuilder) InstanceId(v string) *NatChangeResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *NatChangeResponseDataBuilder) OrderNo(v string) *NatChangeResponseDataBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *NatChangeResponseDataBuilder) Build() *NatChangeResponseData {
	return b.s
}


