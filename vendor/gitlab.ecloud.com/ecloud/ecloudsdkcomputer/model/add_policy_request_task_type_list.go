// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddPolicyRequestTaskTypeList struct {

    // 任务类型,定时任务类型（开机available,关机shutdown，重启restart）
	TaskType *string `json:"taskType,omitempty"`
    // 执行策略，eg:1:周日,2:周一,3:周二,4:周三,5:周四,6:周五,7:周六，如果选多个日期传字符串即可，用逗号隔开
	ExecutePolicy *string `json:"executePolicy,omitempty"`
    // 执行时间，时间自动默认为00:00
	ExecuteTime *string `json:"executeTime,omitempty"`
}

func (s AddPolicyRequestTaskTypeList) String() string {
	return utils.Beautify(s)
}

func (s AddPolicyRequestTaskTypeList) GoString() string {
	return s.String()
}

func (s AddPolicyRequestTaskTypeList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddPolicyRequestTaskTypeList) SetTaskType(v string) *AddPolicyRequestTaskTypeList {
	s.TaskType = &v
	return s
}

func (s *AddPolicyRequestTaskTypeList) SetExecutePolicy(v string) *AddPolicyRequestTaskTypeList {
	s.ExecutePolicy = &v
	return s
}

func (s *AddPolicyRequestTaskTypeList) SetExecuteTime(v string) *AddPolicyRequestTaskTypeList {
	s.ExecuteTime = &v
	return s
}


type AddPolicyRequestTaskTypeListBuilder struct {
	s *AddPolicyRequestTaskTypeList
}

func NewAddPolicyRequestTaskTypeListBuilder() *AddPolicyRequestTaskTypeListBuilder {
	s := &AddPolicyRequestTaskTypeList{}
	b := &AddPolicyRequestTaskTypeListBuilder{s: s}
	return b
}

func (b *AddPolicyRequestTaskTypeListBuilder) TaskType(v string) *AddPolicyRequestTaskTypeListBuilder {
	b.s.TaskType = &v
	return b
}

func (b *AddPolicyRequestTaskTypeListBuilder) ExecutePolicy(v string) *AddPolicyRequestTaskTypeListBuilder {
	b.s.ExecutePolicy = &v
	return b
}

func (b *AddPolicyRequestTaskTypeListBuilder) ExecuteTime(v string) *AddPolicyRequestTaskTypeListBuilder {
	b.s.ExecuteTime = &v
	return b
}

func (b *AddPolicyRequestTaskTypeListBuilder) Build() *AddPolicyRequestTaskTypeList {
	return b.s
}


