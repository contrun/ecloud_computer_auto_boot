// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTaskRecordListByInstanceIdBody struct {
    position.Body
    // 任务类型 shutdown/restart/available
	TaskType *string `json:"taskType,omitempty"`
    // 资源实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 执行结果 成功：success  执行中：processing 等待：waiting 失败：fail
	Status *string `json:"status,omitempty"`
}

func (s GetTaskRecordListByInstanceIdBody) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListByInstanceIdBody) GoString() string {
	return s.String()
}

func (s GetTaskRecordListByInstanceIdBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListByInstanceIdBody) SetTaskType(v string) *GetTaskRecordListByInstanceIdBody {
	s.TaskType = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdBody) SetInstanceId(v string) *GetTaskRecordListByInstanceIdBody {
	s.InstanceId = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdBody) SetPageSize(v int32) *GetTaskRecordListByInstanceIdBody {
	s.PageSize = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdBody) SetCurrentPage(v int32) *GetTaskRecordListByInstanceIdBody {
	s.CurrentPage = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdBody) SetStatus(v string) *GetTaskRecordListByInstanceIdBody {
	s.Status = &v
	return s
}


type GetTaskRecordListByInstanceIdBodyBuilder struct {
	s *GetTaskRecordListByInstanceIdBody
}

func NewGetTaskRecordListByInstanceIdBodyBuilder() *GetTaskRecordListByInstanceIdBodyBuilder {
	s := &GetTaskRecordListByInstanceIdBody{}
	b := &GetTaskRecordListByInstanceIdBodyBuilder{s: s}
	return b
}

func (b *GetTaskRecordListByInstanceIdBodyBuilder) TaskType(v string) *GetTaskRecordListByInstanceIdBodyBuilder {
	b.s.TaskType = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdBodyBuilder) InstanceId(v string) *GetTaskRecordListByInstanceIdBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdBodyBuilder) PageSize(v int32) *GetTaskRecordListByInstanceIdBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdBodyBuilder) CurrentPage(v int32) *GetTaskRecordListByInstanceIdBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdBodyBuilder) Status(v string) *GetTaskRecordListByInstanceIdBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdBodyBuilder) Build() *GetTaskRecordListByInstanceIdBody {
	return b.s
}


