// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnifyInstanceQueryResponseData struct {

    // 当前页
	Current *int64 `json:"current,omitempty"`
    // 符合查询条件的总记录数量
	Total *int64 `json:"total,omitempty"`
    // 是否命中count缓存
	HitCount *bool `json:"hitCount,omitempty"`
    // 根据&nbsp;total&nbsp;和&nbsp;size&nbsp;计算得出的总页数
	Pages *int64 `json:"pages,omitempty"`
    // 分页大小
	Size *int64 `json:"size,omitempty"`
    // 当前页的数据记录
	Records *[]UnifyInstanceQueryResponseRecords `json:"records,omitempty"`
    // 是否进行 count 查询，设置false后不会返回total
	SearchCount *bool `json:"searchCount,omitempty"`
    // 排序字段信息
	Orders *[]UnifyInstanceQueryResponseOrders `json:"orders,omitempty"`
}

func (s UnifyInstanceQueryResponseData) String() string {
	return utils.Beautify(s)
}

func (s UnifyInstanceQueryResponseData) GoString() string {
	return s.String()
}

func (s UnifyInstanceQueryResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnifyInstanceQueryResponseData) SetCurrent(v int64) *UnifyInstanceQueryResponseData {
	s.Current = &v
	return s
}

func (s *UnifyInstanceQueryResponseData) SetTotal(v int64) *UnifyInstanceQueryResponseData {
	s.Total = &v
	return s
}

func (s *UnifyInstanceQueryResponseData) SetHitCount(v bool) *UnifyInstanceQueryResponseData {
	s.HitCount = &v
	return s
}

func (s *UnifyInstanceQueryResponseData) SetPages(v int64) *UnifyInstanceQueryResponseData {
	s.Pages = &v
	return s
}

func (s *UnifyInstanceQueryResponseData) SetSize(v int64) *UnifyInstanceQueryResponseData {
	s.Size = &v
	return s
}

func (s *UnifyInstanceQueryResponseData) SetRecords(v []UnifyInstanceQueryResponseRecords) *UnifyInstanceQueryResponseData {
	s.Records = &v
	return s
}

func (s *UnifyInstanceQueryResponseData) SetSearchCount(v bool) *UnifyInstanceQueryResponseData {
	s.SearchCount = &v
	return s
}

func (s *UnifyInstanceQueryResponseData) SetOrders(v []UnifyInstanceQueryResponseOrders) *UnifyInstanceQueryResponseData {
	s.Orders = &v
	return s
}


type UnifyInstanceQueryResponseDataBuilder struct {
	s *UnifyInstanceQueryResponseData
}

func NewUnifyInstanceQueryResponseDataBuilder() *UnifyInstanceQueryResponseDataBuilder {
	s := &UnifyInstanceQueryResponseData{}
	b := &UnifyInstanceQueryResponseDataBuilder{s: s}
	return b
}

func (b *UnifyInstanceQueryResponseDataBuilder) Current(v int64) *UnifyInstanceQueryResponseDataBuilder {
	b.s.Current = &v
	return b
}

func (b *UnifyInstanceQueryResponseDataBuilder) Total(v int64) *UnifyInstanceQueryResponseDataBuilder {
	b.s.Total = &v
	return b
}

func (b *UnifyInstanceQueryResponseDataBuilder) HitCount(v bool) *UnifyInstanceQueryResponseDataBuilder {
	b.s.HitCount = &v
	return b
}

func (b *UnifyInstanceQueryResponseDataBuilder) Pages(v int64) *UnifyInstanceQueryResponseDataBuilder {
	b.s.Pages = &v
	return b
}

func (b *UnifyInstanceQueryResponseDataBuilder) Size(v int64) *UnifyInstanceQueryResponseDataBuilder {
	b.s.Size = &v
	return b
}

func (b *UnifyInstanceQueryResponseDataBuilder) Records(v []UnifyInstanceQueryResponseRecords) *UnifyInstanceQueryResponseDataBuilder {
	b.s.Records = &v
	return b
}

func (b *UnifyInstanceQueryResponseDataBuilder) SearchCount(v bool) *UnifyInstanceQueryResponseDataBuilder {
	b.s.SearchCount = &v
	return b
}

func (b *UnifyInstanceQueryResponseDataBuilder) Orders(v []UnifyInstanceQueryResponseOrders) *UnifyInstanceQueryResponseDataBuilder {
	b.s.Orders = &v
	return b
}

func (b *UnifyInstanceQueryResponseDataBuilder) Build() *UnifyInstanceQueryResponseData {
	return b.s
}


