// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVpcBody struct {
    position.Body
    // vpc名称
	VpcName *string `json:"vpcName,omitempty"`
    // vpcid
	VpcUid *string `json:"vpcUid,omitempty"`
    // 当前页起始下标
	Start *int32 `json:"start,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // CMDB可用区名称
	PoolName *string `json:"poolName,omitempty"`
    // vpc状态
	Status *string `json:"status,omitempty"`
}

func (s GetVpcBody) String() string {
	return utils.Beautify(s)
}

func (s GetVpcBody) GoString() string {
	return s.String()
}

func (s GetVpcBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcBody) SetVpcName(v string) *GetVpcBody {
	s.VpcName = &v
	return s
}

func (s *GetVpcBody) SetVpcUid(v string) *GetVpcBody {
	s.VpcUid = &v
	return s
}

func (s *GetVpcBody) SetStart(v int32) *GetVpcBody {
	s.Start = &v
	return s
}

func (s *GetVpcBody) SetPageSize(v int32) *GetVpcBody {
	s.PageSize = &v
	return s
}

func (s *GetVpcBody) SetCurrentPage(v int32) *GetVpcBody {
	s.CurrentPage = &v
	return s
}

func (s *GetVpcBody) SetPoolName(v string) *GetVpcBody {
	s.PoolName = &v
	return s
}

func (s *GetVpcBody) SetStatus(v string) *GetVpcBody {
	s.Status = &v
	return s
}


type GetVpcBodyBuilder struct {
	s *GetVpcBody
}

func NewGetVpcBodyBuilder() *GetVpcBodyBuilder {
	s := &GetVpcBody{}
	b := &GetVpcBodyBuilder{s: s}
	return b
}

func (b *GetVpcBodyBuilder) VpcName(v string) *GetVpcBodyBuilder {
	b.s.VpcName = &v
	return b
}

func (b *GetVpcBodyBuilder) VpcUid(v string) *GetVpcBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *GetVpcBodyBuilder) Start(v int32) *GetVpcBodyBuilder {
	b.s.Start = &v
	return b
}

func (b *GetVpcBodyBuilder) PageSize(v int32) *GetVpcBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetVpcBodyBuilder) CurrentPage(v int32) *GetVpcBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetVpcBodyBuilder) PoolName(v string) *GetVpcBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetVpcBodyBuilder) Status(v string) *GetVpcBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetVpcBodyBuilder) Build() *GetVpcBody {
	return b.s
}


