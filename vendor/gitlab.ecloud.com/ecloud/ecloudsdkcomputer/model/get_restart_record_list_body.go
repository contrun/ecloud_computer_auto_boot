// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetRestartRecordListBody struct {
    position.Body
    // 定时任务主键id
	PolicyId *string `json:"policyId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s GetRestartRecordListBody) String() string {
	return utils.Beautify(s)
}

func (s GetRestartRecordListBody) GoString() string {
	return s.String()
}

func (s GetRestartRecordListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetRestartRecordListBody) SetPolicyId(v string) *GetRestartRecordListBody {
	s.PolicyId = &v
	return s
}

func (s *GetRestartRecordListBody) SetPageSize(v int32) *GetRestartRecordListBody {
	s.PageSize = &v
	return s
}

func (s *GetRestartRecordListBody) SetCurrentPage(v int32) *GetRestartRecordListBody {
	s.CurrentPage = &v
	return s
}


type GetRestartRecordListBodyBuilder struct {
	s *GetRestartRecordListBody
}

func NewGetRestartRecordListBodyBuilder() *GetRestartRecordListBodyBuilder {
	s := &GetRestartRecordListBody{}
	b := &GetRestartRecordListBodyBuilder{s: s}
	return b
}

func (b *GetRestartRecordListBodyBuilder) PolicyId(v string) *GetRestartRecordListBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetRestartRecordListBodyBuilder) PageSize(v int32) *GetRestartRecordListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetRestartRecordListBodyBuilder) CurrentPage(v int32) *GetRestartRecordListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetRestartRecordListBodyBuilder) Build() *GetRestartRecordListBody {
	return b.s
}


