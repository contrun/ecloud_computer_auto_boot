// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetInstanceListCemResponseData struct {

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
	Records *[]GetInstanceListCemResponseRecords `json:"records,omitempty"`
    // 是否进行 count 查询，设置false后不会返回total
	SearchCount *bool `json:"searchCount,omitempty"`
    // 排序字段信息
	Orders *[]GetInstanceListCemResponseOrders `json:"orders,omitempty"`
}

func (s GetInstanceListCemResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetInstanceListCemResponseData) GoString() string {
	return s.String()
}

func (s GetInstanceListCemResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetInstanceListCemResponseData) SetCurrent(v int64) *GetInstanceListCemResponseData {
	s.Current = &v
	return s
}

func (s *GetInstanceListCemResponseData) SetTotal(v int64) *GetInstanceListCemResponseData {
	s.Total = &v
	return s
}

func (s *GetInstanceListCemResponseData) SetHitCount(v bool) *GetInstanceListCemResponseData {
	s.HitCount = &v
	return s
}

func (s *GetInstanceListCemResponseData) SetPages(v int64) *GetInstanceListCemResponseData {
	s.Pages = &v
	return s
}

func (s *GetInstanceListCemResponseData) SetSize(v int64) *GetInstanceListCemResponseData {
	s.Size = &v
	return s
}

func (s *GetInstanceListCemResponseData) SetRecords(v []GetInstanceListCemResponseRecords) *GetInstanceListCemResponseData {
	s.Records = &v
	return s
}

func (s *GetInstanceListCemResponseData) SetSearchCount(v bool) *GetInstanceListCemResponseData {
	s.SearchCount = &v
	return s
}

func (s *GetInstanceListCemResponseData) SetOrders(v []GetInstanceListCemResponseOrders) *GetInstanceListCemResponseData {
	s.Orders = &v
	return s
}


type GetInstanceListCemResponseDataBuilder struct {
	s *GetInstanceListCemResponseData
}

func NewGetInstanceListCemResponseDataBuilder() *GetInstanceListCemResponseDataBuilder {
	s := &GetInstanceListCemResponseData{}
	b := &GetInstanceListCemResponseDataBuilder{s: s}
	return b
}

func (b *GetInstanceListCemResponseDataBuilder) Current(v int64) *GetInstanceListCemResponseDataBuilder {
	b.s.Current = &v
	return b
}

func (b *GetInstanceListCemResponseDataBuilder) Total(v int64) *GetInstanceListCemResponseDataBuilder {
	b.s.Total = &v
	return b
}

func (b *GetInstanceListCemResponseDataBuilder) HitCount(v bool) *GetInstanceListCemResponseDataBuilder {
	b.s.HitCount = &v
	return b
}

func (b *GetInstanceListCemResponseDataBuilder) Pages(v int64) *GetInstanceListCemResponseDataBuilder {
	b.s.Pages = &v
	return b
}

func (b *GetInstanceListCemResponseDataBuilder) Size(v int64) *GetInstanceListCemResponseDataBuilder {
	b.s.Size = &v
	return b
}

func (b *GetInstanceListCemResponseDataBuilder) Records(v []GetInstanceListCemResponseRecords) *GetInstanceListCemResponseDataBuilder {
	b.s.Records = &v
	return b
}

func (b *GetInstanceListCemResponseDataBuilder) SearchCount(v bool) *GetInstanceListCemResponseDataBuilder {
	b.s.SearchCount = &v
	return b
}

func (b *GetInstanceListCemResponseDataBuilder) Orders(v []GetInstanceListCemResponseOrders) *GetInstanceListCemResponseDataBuilder {
	b.s.Orders = &v
	return b
}

func (b *GetInstanceListCemResponseDataBuilder) Build() *GetInstanceListCemResponseData {
	return b.s
}


