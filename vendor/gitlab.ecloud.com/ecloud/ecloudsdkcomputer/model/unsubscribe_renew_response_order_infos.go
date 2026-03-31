// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeRenewResponseOrderInfos struct {

    // 实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // mopT
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s UnsubscribeRenewResponseOrderInfos) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeRenewResponseOrderInfos) GoString() string {
	return s.String()
}

func (s UnsubscribeRenewResponseOrderInfos) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeRenewResponseOrderInfos) SetInstanceId(v string) *UnsubscribeRenewResponseOrderInfos {
	s.InstanceId = &v
	return s
}

func (s *UnsubscribeRenewResponseOrderInfos) SetOrderNo(v string) *UnsubscribeRenewResponseOrderInfos {
	s.OrderNo = &v
	return s
}


type UnsubscribeRenewResponseOrderInfosBuilder struct {
	s *UnsubscribeRenewResponseOrderInfos
}

func NewUnsubscribeRenewResponseOrderInfosBuilder() *UnsubscribeRenewResponseOrderInfosBuilder {
	s := &UnsubscribeRenewResponseOrderInfos{}
	b := &UnsubscribeRenewResponseOrderInfosBuilder{s: s}
	return b
}

func (b *UnsubscribeRenewResponseOrderInfosBuilder) InstanceId(v string) *UnsubscribeRenewResponseOrderInfosBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *UnsubscribeRenewResponseOrderInfosBuilder) OrderNo(v string) *UnsubscribeRenewResponseOrderInfosBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *UnsubscribeRenewResponseOrderInfosBuilder) Build() *UnsubscribeRenewResponseOrderInfos {
	return b.s
}


