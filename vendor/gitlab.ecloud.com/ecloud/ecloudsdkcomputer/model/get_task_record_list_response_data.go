// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTaskRecordListResponseData struct {

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
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 桌面名称
	MachineName *string `json:"machineName,omitempty"`
    // 执行时间
	ExecuteTime *string `json:"executeTime,omitempty"`
    // 执行结果
	Status *string `json:"status,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s GetTaskRecordListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListResponseData) GoString() string {
	return s.String()
}

func (s GetTaskRecordListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListResponseData) SetModifiedTime(v string) *GetTaskRecordListResponseData {
	s.ModifiedTime = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetPolicyUid(v string) *GetTaskRecordListResponseData {
	s.PolicyUid = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetTaskType(v string) *GetTaskRecordListResponseData {
	s.TaskType = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetInstanceId(v string) *GetTaskRecordListResponseData {
	s.InstanceId = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetMachineId(v string) *GetTaskRecordListResponseData {
	s.MachineId = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetCreatedTime(v string) *GetTaskRecordListResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetPoolCode(v string) *GetTaskRecordListResponseData {
	s.PoolCode = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetMachineName(v string) *GetTaskRecordListResponseData {
	s.MachineName = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetExecuteTime(v string) *GetTaskRecordListResponseData {
	s.ExecuteTime = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetStatus(v string) *GetTaskRecordListResponseData {
	s.Status = &v
	return s
}

func (s *GetTaskRecordListResponseData) SetPoolName(v string) *GetTaskRecordListResponseData {
	s.PoolName = &v
	return s
}


type GetTaskRecordListResponseDataBuilder struct {
	s *GetTaskRecordListResponseData
}

func NewGetTaskRecordListResponseDataBuilder() *GetTaskRecordListResponseDataBuilder {
	s := &GetTaskRecordListResponseData{}
	b := &GetTaskRecordListResponseDataBuilder{s: s}
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) ModifiedTime(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) PolicyUid(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) TaskType(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.TaskType = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) InstanceId(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) MachineId(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) CreatedTime(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) PoolCode(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) MachineName(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) ExecuteTime(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.ExecuteTime = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) Status(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) PoolName(v string) *GetTaskRecordListResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetTaskRecordListResponseDataBuilder) Build() *GetTaskRecordListResponseData {
	return b.s
}


