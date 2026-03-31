// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTaskRecordListByInstanceIdResponseData struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 策略id
	PolicyUid *string `json:"policyUid,omitempty"`
    // 任务类型
	TaskType *string `json:"taskType,omitempty"`
    // 资源实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // 桌面id
	MachineId *string `json:"machineId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 执行时间
	ExecuteTime *string `json:"executeTime,omitempty"`
    // 执行结果
	Status *string `json:"status,omitempty"`
}

func (s GetTaskRecordListByInstanceIdResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListByInstanceIdResponseData) GoString() string {
	return s.String()
}

func (s GetTaskRecordListByInstanceIdResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListByInstanceIdResponseData) SetModifiedTime(v string) *GetTaskRecordListByInstanceIdResponseData {
	s.ModifiedTime = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponseData) SetPolicyUid(v string) *GetTaskRecordListByInstanceIdResponseData {
	s.PolicyUid = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponseData) SetTaskType(v string) *GetTaskRecordListByInstanceIdResponseData {
	s.TaskType = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponseData) SetInstanceId(v string) *GetTaskRecordListByInstanceIdResponseData {
	s.InstanceId = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponseData) SetMachineId(v string) *GetTaskRecordListByInstanceIdResponseData {
	s.MachineId = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponseData) SetCreatedTime(v string) *GetTaskRecordListByInstanceIdResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponseData) SetExecuteTime(v string) *GetTaskRecordListByInstanceIdResponseData {
	s.ExecuteTime = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponseData) SetStatus(v string) *GetTaskRecordListByInstanceIdResponseData {
	s.Status = &v
	return s
}


type GetTaskRecordListByInstanceIdResponseDataBuilder struct {
	s *GetTaskRecordListByInstanceIdResponseData
}

func NewGetTaskRecordListByInstanceIdResponseDataBuilder() *GetTaskRecordListByInstanceIdResponseDataBuilder {
	s := &GetTaskRecordListByInstanceIdResponseData{}
	b := &GetTaskRecordListByInstanceIdResponseDataBuilder{s: s}
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseDataBuilder) ModifiedTime(v string) *GetTaskRecordListByInstanceIdResponseDataBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseDataBuilder) PolicyUid(v string) *GetTaskRecordListByInstanceIdResponseDataBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseDataBuilder) TaskType(v string) *GetTaskRecordListByInstanceIdResponseDataBuilder {
	b.s.TaskType = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseDataBuilder) InstanceId(v string) *GetTaskRecordListByInstanceIdResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseDataBuilder) MachineId(v string) *GetTaskRecordListByInstanceIdResponseDataBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseDataBuilder) CreatedTime(v string) *GetTaskRecordListByInstanceIdResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseDataBuilder) ExecuteTime(v string) *GetTaskRecordListByInstanceIdResponseDataBuilder {
	b.s.ExecuteTime = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseDataBuilder) Status(v string) *GetTaskRecordListByInstanceIdResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseDataBuilder) Build() *GetTaskRecordListByInstanceIdResponseData {
	return b.s
}


