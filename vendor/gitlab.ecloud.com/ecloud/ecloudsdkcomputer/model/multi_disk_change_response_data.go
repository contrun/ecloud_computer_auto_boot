// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MultiDiskChangeResponseData struct {

    // instanceId
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单编号
	OrderId *string `json:"orderId,omitempty"`
}

func (s MultiDiskChangeResponseData) String() string {
	return utils.Beautify(s)
}

func (s MultiDiskChangeResponseData) GoString() string {
	return s.String()
}

func (s MultiDiskChangeResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MultiDiskChangeResponseData) SetInstanceId(v string) *MultiDiskChangeResponseData {
	s.InstanceId = &v
	return s
}

func (s *MultiDiskChangeResponseData) SetOrderId(v string) *MultiDiskChangeResponseData {
	s.OrderId = &v
	return s
}


type MultiDiskChangeResponseDataBuilder struct {
	s *MultiDiskChangeResponseData
}

func NewMultiDiskChangeResponseDataBuilder() *MultiDiskChangeResponseDataBuilder {
	s := &MultiDiskChangeResponseData{}
	b := &MultiDiskChangeResponseDataBuilder{s: s}
	return b
}

func (b *MultiDiskChangeResponseDataBuilder) InstanceId(v string) *MultiDiskChangeResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *MultiDiskChangeResponseDataBuilder) OrderId(v string) *MultiDiskChangeResponseDataBuilder {
	b.s.OrderId = &v
	return b
}

func (b *MultiDiskChangeResponseDataBuilder) Build() *MultiDiskChangeResponseData {
	return b.s
}


