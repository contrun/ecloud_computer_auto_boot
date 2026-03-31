// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewNatResponseData struct {

    // 资源id
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单编号
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s RenewNatResponseData) String() string {
	return utils.Beautify(s)
}

func (s RenewNatResponseData) GoString() string {
	return s.String()
}

func (s RenewNatResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewNatResponseData) SetInstanceId(v string) *RenewNatResponseData {
	s.InstanceId = &v
	return s
}

func (s *RenewNatResponseData) SetOrderNo(v string) *RenewNatResponseData {
	s.OrderNo = &v
	return s
}


type RenewNatResponseDataBuilder struct {
	s *RenewNatResponseData
}

func NewRenewNatResponseDataBuilder() *RenewNatResponseDataBuilder {
	s := &RenewNatResponseData{}
	b := &RenewNatResponseDataBuilder{s: s}
	return b
}

func (b *RenewNatResponseDataBuilder) InstanceId(v string) *RenewNatResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *RenewNatResponseDataBuilder) OrderNo(v string) *RenewNatResponseDataBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *RenewNatResponseDataBuilder) Build() *RenewNatResponseData {
	return b.s
}


