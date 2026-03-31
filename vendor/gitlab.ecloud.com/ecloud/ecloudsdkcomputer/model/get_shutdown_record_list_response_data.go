// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetShutdownRecordListResponseData struct {

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

func (s GetShutdownRecordListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetShutdownRecordListResponseData) GoString() string {
	return s.String()
}

func (s GetShutdownRecordListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShutdownRecordListResponseData) SetRecordId(v string) *GetShutdownRecordListResponseData {
	s.RecordId = &v
	return s
}

func (s *GetShutdownRecordListResponseData) SetTaskType(v string) *GetShutdownRecordListResponseData {
	s.TaskType = &v
	return s
}

func (s *GetShutdownRecordListResponseData) SetFailNum(v int32) *GetShutdownRecordListResponseData {
	s.FailNum = &v
	return s
}

func (s *GetShutdownRecordListResponseData) SetPolicyId(v string) *GetShutdownRecordListResponseData {
	s.PolicyId = &v
	return s
}

func (s *GetShutdownRecordListResponseData) SetSuccessNum(v int32) *GetShutdownRecordListResponseData {
	s.SuccessNum = &v
	return s
}

func (s *GetShutdownRecordListResponseData) SetWaitNum(v int32) *GetShutdownRecordListResponseData {
	s.WaitNum = &v
	return s
}

func (s *GetShutdownRecordListResponseData) SetProcessNum(v int32) *GetShutdownRecordListResponseData {
	s.ProcessNum = &v
	return s
}

func (s *GetShutdownRecordListResponseData) SetTaskNum(v int32) *GetShutdownRecordListResponseData {
	s.TaskNum = &v
	return s
}

func (s *GetShutdownRecordListResponseData) SetExecuteTime(v string) *GetShutdownRecordListResponseData {
	s.ExecuteTime = &v
	return s
}

func (s *GetShutdownRecordListResponseData) SetStatus(v string) *GetShutdownRecordListResponseData {
	s.Status = &v
	return s
}


type GetShutdownRecordListResponseDataBuilder struct {
	s *GetShutdownRecordListResponseData
}

func NewGetShutdownRecordListResponseDataBuilder() *GetShutdownRecordListResponseDataBuilder {
	s := &GetShutdownRecordListResponseData{}
	b := &GetShutdownRecordListResponseDataBuilder{s: s}
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) RecordId(v string) *GetShutdownRecordListResponseDataBuilder {
	b.s.RecordId = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) TaskType(v string) *GetShutdownRecordListResponseDataBuilder {
	b.s.TaskType = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) FailNum(v int32) *GetShutdownRecordListResponseDataBuilder {
	b.s.FailNum = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) PolicyId(v string) *GetShutdownRecordListResponseDataBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) SuccessNum(v int32) *GetShutdownRecordListResponseDataBuilder {
	b.s.SuccessNum = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) WaitNum(v int32) *GetShutdownRecordListResponseDataBuilder {
	b.s.WaitNum = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) ProcessNum(v int32) *GetShutdownRecordListResponseDataBuilder {
	b.s.ProcessNum = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) TaskNum(v int32) *GetShutdownRecordListResponseDataBuilder {
	b.s.TaskNum = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) ExecuteTime(v string) *GetShutdownRecordListResponseDataBuilder {
	b.s.ExecuteTime = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) Status(v string) *GetShutdownRecordListResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *GetShutdownRecordListResponseDataBuilder) Build() *GetShutdownRecordListResponseData {
	return b.s
}


