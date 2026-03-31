// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBwInstanceListResponseOrders struct {

    // 是否升序
	Asc *bool `json:"asc,omitempty"`
    // 排序列
	Column *string `json:"column,omitempty"`
}

func (s GetBwInstanceListResponseOrders) String() string {
	return utils.Beautify(s)
}

func (s GetBwInstanceListResponseOrders) GoString() string {
	return s.String()
}

func (s GetBwInstanceListResponseOrders) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBwInstanceListResponseOrders) SetAsc(v bool) *GetBwInstanceListResponseOrders {
	s.Asc = &v
	return s
}

func (s *GetBwInstanceListResponseOrders) SetColumn(v string) *GetBwInstanceListResponseOrders {
	s.Column = &v
	return s
}


type GetBwInstanceListResponseOrdersBuilder struct {
	s *GetBwInstanceListResponseOrders
}

func NewGetBwInstanceListResponseOrdersBuilder() *GetBwInstanceListResponseOrdersBuilder {
	s := &GetBwInstanceListResponseOrders{}
	b := &GetBwInstanceListResponseOrdersBuilder{s: s}
	return b
}

func (b *GetBwInstanceListResponseOrdersBuilder) Asc(v bool) *GetBwInstanceListResponseOrdersBuilder {
	b.s.Asc = &v
	return b
}

func (b *GetBwInstanceListResponseOrdersBuilder) Column(v string) *GetBwInstanceListResponseOrdersBuilder {
	b.s.Column = &v
	return b
}

func (b *GetBwInstanceListResponseOrdersBuilder) Build() *GetBwInstanceListResponseOrders {
	return b.s
}


