// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTaskRecordListBody struct {
    position.Body
    // 定时任务执行记录主键id
	RecordId *string `json:"recordId,omitempty"`
    // 任务类型 关机：shutdown 重启：restart 开机：available
	TaskType *string `json:"taskType,omitempty"`
    // 定时任务执行策略主键id
	PolicyId *string `json:"policyId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 执行结果 成功：success 失败：fail 执行中：processing 待执行：waiting
	Status *string `json:"status,omitempty"`
}

func (s GetTaskRecordListBody) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListBody) GoString() string {
	return s.String()
}

func (s GetTaskRecordListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListBody) SetRecordId(v string) *GetTaskRecordListBody {
	s.RecordId = &v
	return s
}

func (s *GetTaskRecordListBody) SetTaskType(v string) *GetTaskRecordListBody {
	s.TaskType = &v
	return s
}

func (s *GetTaskRecordListBody) SetPolicyId(v string) *GetTaskRecordListBody {
	s.PolicyId = &v
	return s
}

func (s *GetTaskRecordListBody) SetPageSize(v int32) *GetTaskRecordListBody {
	s.PageSize = &v
	return s
}

func (s *GetTaskRecordListBody) SetCurrentPage(v int32) *GetTaskRecordListBody {
	s.CurrentPage = &v
	return s
}

func (s *GetTaskRecordListBody) SetStatus(v string) *GetTaskRecordListBody {
	s.Status = &v
	return s
}


type GetTaskRecordListBodyBuilder struct {
	s *GetTaskRecordListBody
}

func NewGetTaskRecordListBodyBuilder() *GetTaskRecordListBodyBuilder {
	s := &GetTaskRecordListBody{}
	b := &GetTaskRecordListBodyBuilder{s: s}
	return b
}

func (b *GetTaskRecordListBodyBuilder) RecordId(v string) *GetTaskRecordListBodyBuilder {
	b.s.RecordId = &v
	return b
}

func (b *GetTaskRecordListBodyBuilder) TaskType(v string) *GetTaskRecordListBodyBuilder {
	b.s.TaskType = &v
	return b
}

func (b *GetTaskRecordListBodyBuilder) PolicyId(v string) *GetTaskRecordListBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetTaskRecordListBodyBuilder) PageSize(v int32) *GetTaskRecordListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetTaskRecordListBodyBuilder) CurrentPage(v int32) *GetTaskRecordListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetTaskRecordListBodyBuilder) Status(v string) *GetTaskRecordListBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetTaskRecordListBodyBuilder) Build() *GetTaskRecordListBody {
	return b.s
}


