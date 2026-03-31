// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetRestartRecordListResponseData struct {

    // 策略执行任务记录主键id
	RecordId *string `json:"recordId,omitempty"`
    // 任务类型 开机:available 关机:shutdown 重启:restart
	TaskType *string `json:"taskType,omitempty"`
    // 失败数量
	FailNum *int32 `json:"failNum,omitempty"`
    // 策略主键id
	PolicyId *string `json:"policyId,omitempty"`
    // 成功数量
	SuccessNum *int32 `json:"successNum,omitempty"`
    // 等待数量
	WaitNum *int32 `json:"waitNum,omitempty"`
    // 执行中数量
	ProcessNum *int32 `json:"processNum,omitempty"`
    // 任务数量
	TaskNum *int32 `json:"taskNum,omitempty"`
    // 执行时间
	ExecuteTime *string `json:"executeTime,omitempty"`
    // 执行结果 waiting:待执行 processing:执行中 success:成功 fail:失败
	Status *string `json:"status,omitempty"`
}

func (s GetRestartRecordListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetRestartRecordListResponseData) GoString() string {
	return s.String()
}

func (s GetRestartRecordListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetRestartRecordListResponseData) SetRecordId(v string) *GetRestartRecordListResponseData {
	s.RecordId = &v
	return s
}

func (s *GetRestartRecordListResponseData) SetTaskType(v string) *GetRestartRecordListResponseData {
	s.TaskType = &v
	return s
}

func (s *GetRestartRecordListResponseData) SetFailNum(v int32) *GetRestartRecordListResponseData {
	s.FailNum = &v
	return s
}

func (s *GetRestartRecordListResponseData) SetPolicyId(v string) *GetRestartRecordListResponseData {
	s.PolicyId = &v
	return s
}

func (s *GetRestartRecordListResponseData) SetSuccessNum(v int32) *GetRestartRecordListResponseData {
	s.SuccessNum = &v
	return s
}

func (s *GetRestartRecordListResponseData) SetWaitNum(v int32) *GetRestartRecordListResponseData {
	s.WaitNum = &v
	return s
}

func (s *GetRestartRecordListResponseData) SetProcessNum(v int32) *GetRestartRecordListResponseData {
	s.ProcessNum = &v
	return s
}

func (s *GetRestartRecordListResponseData) SetTaskNum(v int32) *GetRestartRecordListResponseData {
	s.TaskNum = &v
	return s
}

func (s *GetRestartRecordListResponseData) SetExecuteTime(v string) *GetRestartRecordListResponseData {
	s.ExecuteTime = &v
	return s
}

func (s *GetRestartRecordListResponseData) SetStatus(v string) *GetRestartRecordListResponseData {
	s.Status = &v
	return s
}


type GetRestartRecordListResponseDataBuilder struct {
	s *GetRestartRecordListResponseData
}

func NewGetRestartRecordListResponseDataBuilder() *GetRestartRecordListResponseDataBuilder {
	s := &GetRestartRecordListResponseData{}
	b := &GetRestartRecordListResponseDataBuilder{s: s}
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) RecordId(v string) *GetRestartRecordListResponseDataBuilder {
	b.s.RecordId = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) TaskType(v string) *GetRestartRecordListResponseDataBuilder {
	b.s.TaskType = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) FailNum(v int32) *GetRestartRecordListResponseDataBuilder {
	b.s.FailNum = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) PolicyId(v string) *GetRestartRecordListResponseDataBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) SuccessNum(v int32) *GetRestartRecordListResponseDataBuilder {
	b.s.SuccessNum = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) WaitNum(v int32) *GetRestartRecordListResponseDataBuilder {
	b.s.WaitNum = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) ProcessNum(v int32) *GetRestartRecordListResponseDataBuilder {
	b.s.ProcessNum = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) TaskNum(v int32) *GetRestartRecordListResponseDataBuilder {
	b.s.TaskNum = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) ExecuteTime(v string) *GetRestartRecordListResponseDataBuilder {
	b.s.ExecuteTime = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) Status(v string) *GetRestartRecordListResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *GetRestartRecordListResponseDataBuilder) Build() *GetRestartRecordListResponseData {
	return b.s
}


