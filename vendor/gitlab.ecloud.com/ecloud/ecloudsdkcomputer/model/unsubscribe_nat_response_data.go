// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeNatResponseData struct {

    // 订单列表
	OrderInfos *[]UnsubscribeNatResponseOrderInfos `json:"orderInfos,omitempty"`
}

func (s UnsubscribeNatResponseData) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeNatResponseData) GoString() string {
	return s.String()
}

func (s UnsubscribeNatResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeNatResponseData) SetOrderInfos(v []UnsubscribeNatResponseOrderInfos) *UnsubscribeNatResponseData {
	s.OrderInfos = &v
	return s
}


type UnsubscribeNatResponseDataBuilder struct {
	s *UnsubscribeNatResponseData
}

func NewUnsubscribeNatResponseDataBuilder() *UnsubscribeNatResponseDataBuilder {
	s := &UnsubscribeNatResponseData{}
	b := &UnsubscribeNatResponseDataBuilder{s: s}
	return b
}

func (b *UnsubscribeNatResponseDataBuilder) OrderInfos(v []UnsubscribeNatResponseOrderInfos) *UnsubscribeNatResponseDataBuilder {
	b.s.OrderInfos = &v
	return b
}

func (b *UnsubscribeNatResponseDataBuilder) Build() *UnsubscribeNatResponseData {
	return b.s
}


