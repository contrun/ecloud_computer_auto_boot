// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewComputerResponseOrderInfos struct {

    // 资源id
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单编号
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s RenewComputerResponseOrderInfos) String() string {
	return utils.Beautify(s)
}

func (s RenewComputerResponseOrderInfos) GoString() string {
	return s.String()
}

func (s RenewComputerResponseOrderInfos) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewComputerResponseOrderInfos) SetInstanceId(v string) *RenewComputerResponseOrderInfos {
	s.InstanceId = &v
	return s
}

func (s *RenewComputerResponseOrderInfos) SetOrderNo(v string) *RenewComputerResponseOrderInfos {
	s.OrderNo = &v
	return s
}


type RenewComputerResponseOrderInfosBuilder struct {
	s *RenewComputerResponseOrderInfos
}

func NewRenewComputerResponseOrderInfosBuilder() *RenewComputerResponseOrderInfosBuilder {
	s := &RenewComputerResponseOrderInfos{}
	b := &RenewComputerResponseOrderInfosBuilder{s: s}
	return b
}

func (b *RenewComputerResponseOrderInfosBuilder) InstanceId(v string) *RenewComputerResponseOrderInfosBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *RenewComputerResponseOrderInfosBuilder) OrderNo(v string) *RenewComputerResponseOrderInfosBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *RenewComputerResponseOrderInfosBuilder) Build() *RenewComputerResponseOrderInfos {
	return b.s
}


