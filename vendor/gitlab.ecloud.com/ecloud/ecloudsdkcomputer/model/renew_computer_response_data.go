// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewComputerResponseData struct {

    // 订单信息
	OrderInfos *[]RenewComputerResponseOrderInfos `json:"orderInfos,omitempty"`
}

func (s RenewComputerResponseData) String() string {
	return utils.Beautify(s)
}

func (s RenewComputerResponseData) GoString() string {
	return s.String()
}

func (s RenewComputerResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewComputerResponseData) SetOrderInfos(v []RenewComputerResponseOrderInfos) *RenewComputerResponseData {
	s.OrderInfos = &v
	return s
}


type RenewComputerResponseDataBuilder struct {
	s *RenewComputerResponseData
}

func NewRenewComputerResponseDataBuilder() *RenewComputerResponseDataBuilder {
	s := &RenewComputerResponseData{}
	b := &RenewComputerResponseDataBuilder{s: s}
	return b
}

func (b *RenewComputerResponseDataBuilder) OrderInfos(v []RenewComputerResponseOrderInfos) *RenewComputerResponseDataBuilder {
	b.s.OrderInfos = &v
	return b
}

func (b *RenewComputerResponseDataBuilder) Build() *RenewComputerResponseData {
	return b.s
}


