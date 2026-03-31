// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNatListResponseData struct {

    // 当前页数
	Current *int64 `json:"current,omitempty"`
    // 总数
	Total *int64 `json:"total,omitempty"`
    // 是否命中count缓存
	HitCount *bool `json:"hitCount,omitempty"`
    // 总页数
	Pages *int64 `json:"pages,omitempty"`
    // 页大小
	Size *int64 `json:"size,omitempty"`
    // 实例数据列表
	Records *[]GetNatListResponseRecords `json:"records,omitempty"`
    // 是否进行count查询
	SearchCount *bool `json:"searchCount,omitempty"`
    // 排序字段信息
	Orders *[]GetNatListResponseOrders `json:"orders,omitempty"`
}

func (s GetNatListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetNatListResponseData) GoString() string {
	return s.String()
}

func (s GetNatListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNatListResponseData) SetCurrent(v int64) *GetNatListResponseData {
	s.Current = &v
	return s
}

func (s *GetNatListResponseData) SetTotal(v int64) *GetNatListResponseData {
	s.Total = &v
	return s
}

func (s *GetNatListResponseData) SetHitCount(v bool) *GetNatListResponseData {
	s.HitCount = &v
	return s
}

func (s *GetNatListResponseData) SetPages(v int64) *GetNatListResponseData {
	s.Pages = &v
	return s
}

func (s *GetNatListResponseData) SetSize(v int64) *GetNatListResponseData {
	s.Size = &v
	return s
}

func (s *GetNatListResponseData) SetRecords(v []GetNatListResponseRecords) *GetNatListResponseData {
	s.Records = &v
	return s
}

func (s *GetNatListResponseData) SetSearchCount(v bool) *GetNatListResponseData {
	s.SearchCount = &v
	return s
}

func (s *GetNatListResponseData) SetOrders(v []GetNatListResponseOrders) *GetNatListResponseData {
	s.Orders = &v
	return s
}


type GetNatListResponseDataBuilder struct {
	s *GetNatListResponseData
}

func NewGetNatListResponseDataBuilder() *GetNatListResponseDataBuilder {
	s := &GetNatListResponseData{}
	b := &GetNatListResponseDataBuilder{s: s}
	return b
}

func (b *GetNatListResponseDataBuilder) Current(v int64) *GetNatListResponseDataBuilder {
	b.s.Current = &v
	return b
}

func (b *GetNatListResponseDataBuilder) Total(v int64) *GetNatListResponseDataBuilder {
	b.s.Total = &v
	return b
}

func (b *GetNatListResponseDataBuilder) HitCount(v bool) *GetNatListResponseDataBuilder {
	b.s.HitCount = &v
	return b
}

func (b *GetNatListResponseDataBuilder) Pages(v int64) *GetNatListResponseDataBuilder {
	b.s.Pages = &v
	return b
}

func (b *GetNatListResponseDataBuilder) Size(v int64) *GetNatListResponseDataBuilder {
	b.s.Size = &v
	return b
}

func (b *GetNatListResponseDataBuilder) Records(v []GetNatListResponseRecords) *GetNatListResponseDataBuilder {
	b.s.Records = &v
	return b
}

func (b *GetNatListResponseDataBuilder) SearchCount(v bool) *GetNatListResponseDataBuilder {
	b.s.SearchCount = &v
	return b
}

func (b *GetNatListResponseDataBuilder) Orders(v []GetNatListResponseOrders) *GetNatListResponseDataBuilder {
	b.s.Orders = &v
	return b
}

func (b *GetNatListResponseDataBuilder) Build() *GetNatListResponseData {
	return b.s
}


