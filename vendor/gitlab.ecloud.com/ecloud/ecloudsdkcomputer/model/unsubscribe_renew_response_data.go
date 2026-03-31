// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeRenewResponseData struct {

    // 订单列表
	OrderInfos *[]UnsubscribeRenewResponseOrderInfos `json:"orderInfos,omitempty"`
}

func (s UnsubscribeRenewResponseData) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeRenewResponseData) GoString() string {
	return s.String()
}

func (s UnsubscribeRenewResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeRenewResponseData) SetOrderInfos(v []UnsubscribeRenewResponseOrderInfos) *UnsubscribeRenewResponseData {
	s.OrderInfos = &v
	return s
}


type UnsubscribeRenewResponseDataBuilder struct {
	s *UnsubscribeRenewResponseData
}

func NewUnsubscribeRenewResponseDataBuilder() *UnsubscribeRenewResponseDataBuilder {
	s := &UnsubscribeRenewResponseData{}
	b := &UnsubscribeRenewResponseDataBuilder{s: s}
	return b
}

func (b *UnsubscribeRenewResponseDataBuilder) OrderInfos(v []UnsubscribeRenewResponseOrderInfos) *UnsubscribeRenewResponseDataBuilder {
	b.s.OrderInfos = &v
	return b
}

func (b *UnsubscribeRenewResponseDataBuilder) Build() *UnsubscribeRenewResponseData {
	return b.s
}


