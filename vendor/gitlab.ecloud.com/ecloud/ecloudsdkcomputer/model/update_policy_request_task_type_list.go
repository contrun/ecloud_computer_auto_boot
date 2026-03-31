// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdatePolicyRequestTaskTypeList struct {

    // 任务类型,定时任务类型（开机available,关机shutdown，重启restart）
	TaskType *string `json:"taskType,omitempty"`
    // 执行策略，eg:1:周日,2:周一,3:周二,4:周三,5:周四,6:周五,7:周六，如果选多个日期传字符串即可，用逗号隔开
	ExecutePolicy *string `json:"executePolicy,omitempty"`
    // 执行时间，时间自动默认为00:00
	ExecuteTime *string `json:"executeTime,omitempty"`
}

func (s UpdatePolicyRequestTaskTypeList) String() string {
	return utils.Beautify(s)
}

func (s UpdatePolicyRequestTaskTypeList) GoString() string {
	return s.String()
}

func (s UpdatePolicyRequestTaskTypeList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdatePolicyRequestTaskTypeList) SetTaskType(v string) *UpdatePolicyRequestTaskTypeList {
	s.TaskType = &v
	return s
}

func (s *UpdatePolicyRequestTaskTypeList) SetExecutePolicy(v string) *UpdatePolicyRequestTaskTypeList {
	s.ExecutePolicy = &v
	return s
}

func (s *UpdatePolicyRequestTaskTypeList) SetExecuteTime(v string) *UpdatePolicyRequestTaskTypeList {
	s.ExecuteTime = &v
	return s
}


type UpdatePolicyRequestTaskTypeListBuilder struct {
	s *UpdatePolicyRequestTaskTypeList
}

func NewUpdatePolicyRequestTaskTypeListBuilder() *UpdatePolicyRequestTaskTypeListBuilder {
	s := &UpdatePolicyRequestTaskTypeList{}
	b := &UpdatePolicyRequestTaskTypeListBuilder{s: s}
	return b
}

func (b *UpdatePolicyRequestTaskTypeListBuilder) TaskType(v string) *UpdatePolicyRequestTaskTypeListBuilder {
	b.s.TaskType = &v
	return b
}

func (b *UpdatePolicyRequestTaskTypeListBuilder) ExecutePolicy(v string) *UpdatePolicyRequestTaskTypeListBuilder {
	b.s.ExecutePolicy = &v
	return b
}

func (b *UpdatePolicyRequestTaskTypeListBuilder) ExecuteTime(v string) *UpdatePolicyRequestTaskTypeListBuilder {
	b.s.ExecuteTime = &v
	return b
}

func (b *UpdatePolicyRequestTaskTypeListBuilder) Build() *UpdatePolicyRequestTaskTypeList {
	return b.s
}


