// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewBwResponseData struct {

    // 资源id
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单编号
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s RenewBwResponseData) String() string {
	return utils.Beautify(s)
}

func (s RenewBwResponseData) GoString() string {
	return s.String()
}

func (s RenewBwResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewBwResponseData) SetInstanceId(v string) *RenewBwResponseData {
	s.InstanceId = &v
	return s
}

func (s *RenewBwResponseData) SetOrderNo(v string) *RenewBwResponseData {
	s.OrderNo = &v
	return s
}


type RenewBwResponseDataBuilder struct {
	s *RenewBwResponseData
}

func NewRenewBwResponseDataBuilder() *RenewBwResponseDataBuilder {
	s := &RenewBwResponseData{}
	b := &RenewBwResponseDataBuilder{s: s}
	return b
}

func (b *RenewBwResponseDataBuilder) InstanceId(v string) *RenewBwResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *RenewBwResponseDataBuilder) OrderNo(v string) *RenewBwResponseDataBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *RenewBwResponseDataBuilder) Build() *RenewBwResponseData {
	return b.s
}


