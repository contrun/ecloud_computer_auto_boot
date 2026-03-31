// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBwInstanceListResponseData struct {

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
	Records *[]GetBwInstanceListResponseRecords `json:"records,omitempty"`
    // 是否进行count查询
	SearchCount *bool `json:"searchCount,omitempty"`
    // 排序字段信息
	Orders *[]GetBwInstanceListResponseOrders `json:"orders,omitempty"`
}

func (s GetBwInstanceListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetBwInstanceListResponseData) GoString() string {
	return s.String()
}

func (s GetBwInstanceListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBwInstanceListResponseData) SetCurrent(v int64) *GetBwInstanceListResponseData {
	s.Current = &v
	return s
}

func (s *GetBwInstanceListResponseData) SetTotal(v int64) *GetBwInstanceListResponseData {
	s.Total = &v
	return s
}

func (s *GetBwInstanceListResponseData) SetHitCount(v bool) *GetBwInstanceListResponseData {
	s.HitCount = &v
	return s
}

func (s *GetBwInstanceListResponseData) SetPages(v int64) *GetBwInstanceListResponseData {
	s.Pages = &v
	return s
}

func (s *GetBwInstanceListResponseData) SetSize(v int64) *GetBwInstanceListResponseData {
	s.Size = &v
	return s
}

func (s *GetBwInstanceListResponseData) SetRecords(v []GetBwInstanceListResponseRecords) *GetBwInstanceListResponseData {
	s.Records = &v
	return s
}

func (s *GetBwInstanceListResponseData) SetSearchCount(v bool) *GetBwInstanceListResponseData {
	s.SearchCount = &v
	return s
}

func (s *GetBwInstanceListResponseData) SetOrders(v []GetBwInstanceListResponseOrders) *GetBwInstanceListResponseData {
	s.Orders = &v
	return s
}


type GetBwInstanceListResponseDataBuilder struct {
	s *GetBwInstanceListResponseData
}

func NewGetBwInstanceListResponseDataBuilder() *GetBwInstanceListResponseDataBuilder {
	s := &GetBwInstanceListResponseData{}
	b := &GetBwInstanceListResponseDataBuilder{s: s}
	return b
}

func (b *GetBwInstanceListResponseDataBuilder) Current(v int64) *GetBwInstanceListResponseDataBuilder {
	b.s.Current = &v
	return b
}

func (b *GetBwInstanceListResponseDataBuilder) Total(v int64) *GetBwInstanceListResponseDataBuilder {
	b.s.Total = &v
	return b
}

func (b *GetBwInstanceListResponseDataBuilder) HitCount(v bool) *GetBwInstanceListResponseDataBuilder {
	b.s.HitCount = &v
	return b
}

func (b *GetBwInstanceListResponseDataBuilder) Pages(v int64) *GetBwInstanceListResponseDataBuilder {
	b.s.Pages = &v
	return b
}

func (b *GetBwInstanceListResponseDataBuilder) Size(v int64) *GetBwInstanceListResponseDataBuilder {
	b.s.Size = &v
	return b
}

func (b *GetBwInstanceListResponseDataBuilder) Records(v []GetBwInstanceListResponseRecords) *GetBwInstanceListResponseDataBuilder {
	b.s.Records = &v
	return b
}

func (b *GetBwInstanceListResponseDataBuilder) SearchCount(v bool) *GetBwInstanceListResponseDataBuilder {
	b.s.SearchCount = &v
	return b
}

func (b *GetBwInstanceListResponseDataBuilder) Orders(v []GetBwInstanceListResponseOrders) *GetBwInstanceListResponseDataBuilder {
	b.s.Orders = &v
	return b
}

func (b *GetBwInstanceListResponseDataBuilder) Build() *GetBwInstanceListResponseData {
	return b.s
}


