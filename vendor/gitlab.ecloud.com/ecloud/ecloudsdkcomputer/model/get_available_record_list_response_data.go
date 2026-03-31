// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAvailableRecordListResponseData struct {

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

func (s GetAvailableRecordListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetAvailableRecordListResponseData) GoString() string {
	return s.String()
}

func (s GetAvailableRecordListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAvailableRecordListResponseData) SetRecordId(v string) *GetAvailableRecordListResponseData {
	s.RecordId = &v
	return s
}

func (s *GetAvailableRecordListResponseData) SetTaskType(v string) *GetAvailableRecordListResponseData {
	s.TaskType = &v
	return s
}

func (s *GetAvailableRecordListResponseData) SetFailNum(v int32) *GetAvailableRecordListResponseData {
	s.FailNum = &v
	return s
}

func (s *GetAvailableRecordListResponseData) SetPolicyId(v string) *GetAvailableRecordListResponseData {
	s.PolicyId = &v
	return s
}

func (s *GetAvailableRecordListResponseData) SetSuccessNum(v int32) *GetAvailableRecordListResponseData {
	s.SuccessNum = &v
	return s
}

func (s *GetAvailableRecordListResponseData) SetWaitNum(v int32) *GetAvailableRecordListResponseData {
	s.WaitNum = &v
	return s
}

func (s *GetAvailableRecordListResponseData) SetProcessNum(v int32) *GetAvailableRecordListResponseData {
	s.ProcessNum = &v
	return s
}

func (s *GetAvailableRecordListResponseData) SetTaskNum(v int32) *GetAvailableRecordListResponseData {
	s.TaskNum = &v
	return s
}

func (s *GetAvailableRecordListResponseData) SetExecuteTime(v string) *GetAvailableRecordListResponseData {
	s.ExecuteTime = &v
	return s
}

func (s *GetAvailableRecordListResponseData) SetStatus(v string) *GetAvailableRecordListResponseData {
	s.Status = &v
	return s
}


type GetAvailableRecordListResponseDataBuilder struct {
	s *GetAvailableRecordListResponseData
}

func NewGetAvailableRecordListResponseDataBuilder() *GetAvailableRecordListResponseDataBuilder {
	s := &GetAvailableRecordListResponseData{}
	b := &GetAvailableRecordListResponseDataBuilder{s: s}
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) RecordId(v string) *GetAvailableRecordListResponseDataBuilder {
	b.s.RecordId = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) TaskType(v string) *GetAvailableRecordListResponseDataBuilder {
	b.s.TaskType = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) FailNum(v int32) *GetAvailableRecordListResponseDataBuilder {
	b.s.FailNum = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) PolicyId(v string) *GetAvailableRecordListResponseDataBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) SuccessNum(v int32) *GetAvailableRecordListResponseDataBuilder {
	b.s.SuccessNum = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) WaitNum(v int32) *GetAvailableRecordListResponseDataBuilder {
	b.s.WaitNum = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) ProcessNum(v int32) *GetAvailableRecordListResponseDataBuilder {
	b.s.ProcessNum = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) TaskNum(v int32) *GetAvailableRecordListResponseDataBuilder {
	b.s.TaskNum = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) ExecuteTime(v string) *GetAvailableRecordListResponseDataBuilder {
	b.s.ExecuteTime = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) Status(v string) *GetAvailableRecordListResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *GetAvailableRecordListResponseDataBuilder) Build() *GetAvailableRecordListResponseData {
	return b.s
}


