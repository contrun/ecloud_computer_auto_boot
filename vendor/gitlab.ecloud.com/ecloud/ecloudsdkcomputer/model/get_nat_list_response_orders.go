// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNatListResponseOrders struct {

    // 是否升序
	Asc *bool `json:"asc,omitempty"`
    // 排序列
	Column *string `json:"column,omitempty"`
}

func (s GetNatListResponseOrders) String() string {
	return utils.Beautify(s)
}

func (s GetNatListResponseOrders) GoString() string {
	return s.String()
}

func (s GetNatListResponseOrders) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNatListResponseOrders) SetAsc(v bool) *GetNatListResponseOrders {
	s.Asc = &v
	return s
}

func (s *GetNatListResponseOrders) SetColumn(v string) *GetNatListResponseOrders {
	s.Column = &v
	return s
}


type GetNatListResponseOrdersBuilder struct {
	s *GetNatListResponseOrders
}

func NewGetNatListResponseOrdersBuilder() *GetNatListResponseOrdersBuilder {
	s := &GetNatListResponseOrders{}
	b := &GetNatListResponseOrdersBuilder{s: s}
	return b
}

func (b *GetNatListResponseOrdersBuilder) Asc(v bool) *GetNatListResponseOrdersBuilder {
	b.s.Asc = &v
	return b
}

func (b *GetNatListResponseOrdersBuilder) Column(v string) *GetNatListResponseOrdersBuilder {
	b.s.Column = &v
	return b
}

func (b *GetNatListResponseOrdersBuilder) Build() *GetNatListResponseOrders {
	return b.s
}


