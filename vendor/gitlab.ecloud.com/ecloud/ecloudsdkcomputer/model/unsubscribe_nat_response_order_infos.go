// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeNatResponseOrderInfos struct {

    // 实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // mopT
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s UnsubscribeNatResponseOrderInfos) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeNatResponseOrderInfos) GoString() string {
	return s.String()
}

func (s UnsubscribeNatResponseOrderInfos) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeNatResponseOrderInfos) SetInstanceId(v string) *UnsubscribeNatResponseOrderInfos {
	s.InstanceId = &v
	return s
}

func (s *UnsubscribeNatResponseOrderInfos) SetOrderNo(v string) *UnsubscribeNatResponseOrderInfos {
	s.OrderNo = &v
	return s
}


type UnsubscribeNatResponseOrderInfosBuilder struct {
	s *UnsubscribeNatResponseOrderInfos
}

func NewUnsubscribeNatResponseOrderInfosBuilder() *UnsubscribeNatResponseOrderInfosBuilder {
	s := &UnsubscribeNatResponseOrderInfos{}
	b := &UnsubscribeNatResponseOrderInfosBuilder{s: s}
	return b
}

func (b *UnsubscribeNatResponseOrderInfosBuilder) InstanceId(v string) *UnsubscribeNatResponseOrderInfosBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *UnsubscribeNatResponseOrderInfosBuilder) OrderNo(v string) *UnsubscribeNatResponseOrderInfosBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *UnsubscribeNatResponseOrderInfosBuilder) Build() *UnsubscribeNatResponseOrderInfos {
	return b.s
}


