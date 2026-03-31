// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelInstanceListResponseData struct {

    // 订单列表
	OrderInfos *[]CancelInstanceListResponseOrderInfos `json:"orderInfos,omitempty"`
}

func (s CancelInstanceListResponseData) String() string {
	return utils.Beautify(s)
}

func (s CancelInstanceListResponseData) GoString() string {
	return s.String()
}

func (s CancelInstanceListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelInstanceListResponseData) SetOrderInfos(v []CancelInstanceListResponseOrderInfos) *CancelInstanceListResponseData {
	s.OrderInfos = &v
	return s
}


type CancelInstanceListResponseDataBuilder struct {
	s *CancelInstanceListResponseData
}

func NewCancelInstanceListResponseDataBuilder() *CancelInstanceListResponseDataBuilder {
	s := &CancelInstanceListResponseData{}
	b := &CancelInstanceListResponseDataBuilder{s: s}
	return b
}

func (b *CancelInstanceListResponseDataBuilder) OrderInfos(v []CancelInstanceListResponseOrderInfos) *CancelInstanceListResponseDataBuilder {
	b.s.OrderInfos = &v
	return b
}

func (b *CancelInstanceListResponseDataBuilder) Build() *CancelInstanceListResponseData {
	return b.s
}


