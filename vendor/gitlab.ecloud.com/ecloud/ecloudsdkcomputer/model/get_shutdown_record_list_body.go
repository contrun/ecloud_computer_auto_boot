// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetShutdownRecordListBody struct {
    position.Body
    // 定时任务主键id
	PolicyId *string `json:"policyId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s GetShutdownRecordListBody) String() string {
	return utils.Beautify(s)
}

func (s GetShutdownRecordListBody) GoString() string {
	return s.String()
}

func (s GetShutdownRecordListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShutdownRecordListBody) SetPolicyId(v string) *GetShutdownRecordListBody {
	s.PolicyId = &v
	return s
}

func (s *GetShutdownRecordListBody) SetPageSize(v int32) *GetShutdownRecordListBody {
	s.PageSize = &v
	return s
}

func (s *GetShutdownRecordListBody) SetCurrentPage(v int32) *GetShutdownRecordListBody {
	s.CurrentPage = &v
	return s
}


type GetShutdownRecordListBodyBuilder struct {
	s *GetShutdownRecordListBody
}

func NewGetShutdownRecordListBodyBuilder() *GetShutdownRecordListBodyBuilder {
	s := &GetShutdownRecordListBody{}
	b := &GetShutdownRecordListBodyBuilder{s: s}
	return b
}

func (b *GetShutdownRecordListBodyBuilder) PolicyId(v string) *GetShutdownRecordListBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetShutdownRecordListBodyBuilder) PageSize(v int32) *GetShutdownRecordListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetShutdownRecordListBodyBuilder) CurrentPage(v int32) *GetShutdownRecordListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetShutdownRecordListBodyBuilder) Build() *GetShutdownRecordListBody {
	return b.s
}


