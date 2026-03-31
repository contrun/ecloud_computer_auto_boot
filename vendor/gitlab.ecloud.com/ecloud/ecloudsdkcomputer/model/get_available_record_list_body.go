// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAvailableRecordListBody struct {
    position.Body
    // 定时任务主键id
	PolicyId *string `json:"policyId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s GetAvailableRecordListBody) String() string {
	return utils.Beautify(s)
}

func (s GetAvailableRecordListBody) GoString() string {
	return s.String()
}

func (s GetAvailableRecordListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAvailableRecordListBody) SetPolicyId(v string) *GetAvailableRecordListBody {
	s.PolicyId = &v
	return s
}

func (s *GetAvailableRecordListBody) SetPageSize(v int32) *GetAvailableRecordListBody {
	s.PageSize = &v
	return s
}

func (s *GetAvailableRecordListBody) SetCurrentPage(v int32) *GetAvailableRecordListBody {
	s.CurrentPage = &v
	return s
}


type GetAvailableRecordListBodyBuilder struct {
	s *GetAvailableRecordListBody
}

func NewGetAvailableRecordListBodyBuilder() *GetAvailableRecordListBodyBuilder {
	s := &GetAvailableRecordListBody{}
	b := &GetAvailableRecordListBodyBuilder{s: s}
	return b
}

func (b *GetAvailableRecordListBodyBuilder) PolicyId(v string) *GetAvailableRecordListBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetAvailableRecordListBodyBuilder) PageSize(v int32) *GetAvailableRecordListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetAvailableRecordListBodyBuilder) CurrentPage(v int32) *GetAvailableRecordListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetAvailableRecordListBodyBuilder) Build() *GetAvailableRecordListBody {
	return b.s
}


