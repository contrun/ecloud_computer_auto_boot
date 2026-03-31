// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNetworkBody struct {
    position.Body
    // vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
    // 当前页起始下标
	Start *int32 `json:"start,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s GetNetworkBody) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkBody) GoString() string {
	return s.String()
}

func (s GetNetworkBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkBody) SetVpcUid(v string) *GetNetworkBody {
	s.VpcUid = &v
	return s
}

func (s *GetNetworkBody) SetStart(v int32) *GetNetworkBody {
	s.Start = &v
	return s
}

func (s *GetNetworkBody) SetPageSize(v int32) *GetNetworkBody {
	s.PageSize = &v
	return s
}

func (s *GetNetworkBody) SetCurrentPage(v int32) *GetNetworkBody {
	s.CurrentPage = &v
	return s
}


type GetNetworkBodyBuilder struct {
	s *GetNetworkBody
}

func NewGetNetworkBodyBuilder() *GetNetworkBodyBuilder {
	s := &GetNetworkBody{}
	b := &GetNetworkBodyBuilder{s: s}
	return b
}

func (b *GetNetworkBodyBuilder) VpcUid(v string) *GetNetworkBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *GetNetworkBodyBuilder) Start(v int32) *GetNetworkBodyBuilder {
	b.s.Start = &v
	return b
}

func (b *GetNetworkBodyBuilder) PageSize(v int32) *GetNetworkBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetNetworkBodyBuilder) CurrentPage(v int32) *GetNetworkBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetNetworkBodyBuilder) Build() *GetNetworkBody {
	return b.s
}


