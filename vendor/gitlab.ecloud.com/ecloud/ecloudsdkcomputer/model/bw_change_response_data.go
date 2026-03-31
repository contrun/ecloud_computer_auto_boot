// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BwChangeResponseData struct {

    // 资源实例ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单号：MOP-O-XXX
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s BwChangeResponseData) String() string {
	return utils.Beautify(s)
}

func (s BwChangeResponseData) GoString() string {
	return s.String()
}

func (s BwChangeResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BwChangeResponseData) SetInstanceId(v string) *BwChangeResponseData {
	s.InstanceId = &v
	return s
}

func (s *BwChangeResponseData) SetOrderNo(v string) *BwChangeResponseData {
	s.OrderNo = &v
	return s
}


type BwChangeResponseDataBuilder struct {
	s *BwChangeResponseData
}

func NewBwChangeResponseDataBuilder() *BwChangeResponseDataBuilder {
	s := &BwChangeResponseData{}
	b := &BwChangeResponseDataBuilder{s: s}
	return b
}

func (b *BwChangeResponseDataBuilder) InstanceId(v string) *BwChangeResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *BwChangeResponseDataBuilder) OrderNo(v string) *BwChangeResponseDataBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *BwChangeResponseDataBuilder) Build() *BwChangeResponseData {
	return b.s
}


