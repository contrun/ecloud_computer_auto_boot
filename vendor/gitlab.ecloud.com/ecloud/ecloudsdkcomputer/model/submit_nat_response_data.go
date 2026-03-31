// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitNatResponseData struct {

    // 资源实例ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单编号
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s SubmitNatResponseData) String() string {
	return utils.Beautify(s)
}

func (s SubmitNatResponseData) GoString() string {
	return s.String()
}

func (s SubmitNatResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitNatResponseData) SetInstanceId(v string) *SubmitNatResponseData {
	s.InstanceId = &v
	return s
}

func (s *SubmitNatResponseData) SetOrderNo(v string) *SubmitNatResponseData {
	s.OrderNo = &v
	return s
}


type SubmitNatResponseDataBuilder struct {
	s *SubmitNatResponseData
}

func NewSubmitNatResponseDataBuilder() *SubmitNatResponseDataBuilder {
	s := &SubmitNatResponseData{}
	b := &SubmitNatResponseDataBuilder{s: s}
	return b
}

func (b *SubmitNatResponseDataBuilder) InstanceId(v string) *SubmitNatResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *SubmitNatResponseDataBuilder) OrderNo(v string) *SubmitNatResponseDataBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *SubmitNatResponseDataBuilder) Build() *SubmitNatResponseData {
	return b.s
}


