// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PolicyDescResponseTaskDescInfoList struct {

    // 任务类型,定时任务类型（开机available,关机shutdown，重启restart）
	TaskType *string `json:"taskType,omitempty"`
    // 执行策略，eg:1:周日,2:周一,3:周二,4:周三,5:周四,6:周五,7:周六，如果选多个日期传字符串即可，用逗号隔开
	ExecutePolicy *string `json:"executePolicy,omitempty"`
    // 执行时间，时间自动默认为00:00
	ExecuteTime *string `json:"executeTime,omitempty"`
}

func (s PolicyDescResponseTaskDescInfoList) String() string {
	return utils.Beautify(s)
}

func (s PolicyDescResponseTaskDescInfoList) GoString() string {
	return s.String()
}

func (s PolicyDescResponseTaskDescInfoList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyDescResponseTaskDescInfoList) SetTaskType(v string) *PolicyDescResponseTaskDescInfoList {
	s.TaskType = &v
	return s
}

func (s *PolicyDescResponseTaskDescInfoList) SetExecutePolicy(v string) *PolicyDescResponseTaskDescInfoList {
	s.ExecutePolicy = &v
	return s
}

func (s *PolicyDescResponseTaskDescInfoList) SetExecuteTime(v string) *PolicyDescResponseTaskDescInfoList {
	s.ExecuteTime = &v
	return s
}


type PolicyDescResponseTaskDescInfoListBuilder struct {
	s *PolicyDescResponseTaskDescInfoList
}

func NewPolicyDescResponseTaskDescInfoListBuilder() *PolicyDescResponseTaskDescInfoListBuilder {
	s := &PolicyDescResponseTaskDescInfoList{}
	b := &PolicyDescResponseTaskDescInfoListBuilder{s: s}
	return b
}

func (b *PolicyDescResponseTaskDescInfoListBuilder) TaskType(v string) *PolicyDescResponseTaskDescInfoListBuilder {
	b.s.TaskType = &v
	return b
}

func (b *PolicyDescResponseTaskDescInfoListBuilder) ExecutePolicy(v string) *PolicyDescResponseTaskDescInfoListBuilder {
	b.s.ExecutePolicy = &v
	return b
}

func (b *PolicyDescResponseTaskDescInfoListBuilder) ExecuteTime(v string) *PolicyDescResponseTaskDescInfoListBuilder {
	b.s.ExecuteTime = &v
	return b
}

func (b *PolicyDescResponseTaskDescInfoListBuilder) Build() *PolicyDescResponseTaskDescInfoList {
	return b.s
}


