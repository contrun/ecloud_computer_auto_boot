// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnifyInstanceQueryResponseOrders struct {

    // 是否升序
	Asc *bool `json:"asc,omitempty"`
    // 排序列
	Column *string `json:"column,omitempty"`
}

func (s UnifyInstanceQueryResponseOrders) String() string {
	return utils.Beautify(s)
}

func (s UnifyInstanceQueryResponseOrders) GoString() string {
	return s.String()
}

func (s UnifyInstanceQueryResponseOrders) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnifyInstanceQueryResponseOrders) SetAsc(v bool) *UnifyInstanceQueryResponseOrders {
	s.Asc = &v
	return s
}

func (s *UnifyInstanceQueryResponseOrders) SetColumn(v string) *UnifyInstanceQueryResponseOrders {
	s.Column = &v
	return s
}


type UnifyInstanceQueryResponseOrdersBuilder struct {
	s *UnifyInstanceQueryResponseOrders
}

func NewUnifyInstanceQueryResponseOrdersBuilder() *UnifyInstanceQueryResponseOrdersBuilder {
	s := &UnifyInstanceQueryResponseOrders{}
	b := &UnifyInstanceQueryResponseOrdersBuilder{s: s}
	return b
}

func (b *UnifyInstanceQueryResponseOrdersBuilder) Asc(v bool) *UnifyInstanceQueryResponseOrdersBuilder {
	b.s.Asc = &v
	return b
}

func (b *UnifyInstanceQueryResponseOrdersBuilder) Column(v string) *UnifyInstanceQueryResponseOrdersBuilder {
	b.s.Column = &v
	return b
}

func (b *UnifyInstanceQueryResponseOrdersBuilder) Build() *UnifyInstanceQueryResponseOrders {
	return b.s
}


