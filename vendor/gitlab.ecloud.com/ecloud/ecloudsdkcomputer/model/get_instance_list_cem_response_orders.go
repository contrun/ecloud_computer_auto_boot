// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetInstanceListCemResponseOrders struct {

    // 是否升序
	Asc *bool `json:"asc,omitempty"`
    // 排序列
	Column *string `json:"column,omitempty"`
}

func (s GetInstanceListCemResponseOrders) String() string {
	return utils.Beautify(s)
}

func (s GetInstanceListCemResponseOrders) GoString() string {
	return s.String()
}

func (s GetInstanceListCemResponseOrders) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetInstanceListCemResponseOrders) SetAsc(v bool) *GetInstanceListCemResponseOrders {
	s.Asc = &v
	return s
}

func (s *GetInstanceListCemResponseOrders) SetColumn(v string) *GetInstanceListCemResponseOrders {
	s.Column = &v
	return s
}


type GetInstanceListCemResponseOrdersBuilder struct {
	s *GetInstanceListCemResponseOrders
}

func NewGetInstanceListCemResponseOrdersBuilder() *GetInstanceListCemResponseOrdersBuilder {
	s := &GetInstanceListCemResponseOrders{}
	b := &GetInstanceListCemResponseOrdersBuilder{s: s}
	return b
}

func (b *GetInstanceListCemResponseOrdersBuilder) Asc(v bool) *GetInstanceListCemResponseOrdersBuilder {
	b.s.Asc = &v
	return b
}

func (b *GetInstanceListCemResponseOrdersBuilder) Column(v string) *GetInstanceListCemResponseOrdersBuilder {
	b.s.Column = &v
	return b
}

func (b *GetInstanceListCemResponseOrdersBuilder) Build() *GetInstanceListCemResponseOrders {
	return b.s
}


