// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByInstanceIdResponseTaskDescInfoList struct {

    // 任务类型,定时任务类型（开机available,关机shutdown，重启restart）
	TaskType *string `json:"taskType,omitempty"`
    // 执行策略，eg:1:周一,2:周二,3:周三,4:周四,5:周五,6:周六,7:周日，如果选多个日期传字符串即可，用逗号隔开
	ExecutePolicy *string `json:"executePolicy,omitempty"`
    // 执行时间，时间自动默认为00:00
	ExecuteTime *string `json:"executeTime,omitempty"`
}

func (s GetPolicyInfoByInstanceIdResponseTaskDescInfoList) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByInstanceIdResponseTaskDescInfoList) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByInstanceIdResponseTaskDescInfoList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByInstanceIdResponseTaskDescInfoList) SetTaskType(v string) *GetPolicyInfoByInstanceIdResponseTaskDescInfoList {
	s.TaskType = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseTaskDescInfoList) SetExecutePolicy(v string) *GetPolicyInfoByInstanceIdResponseTaskDescInfoList {
	s.ExecutePolicy = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseTaskDescInfoList) SetExecuteTime(v string) *GetPolicyInfoByInstanceIdResponseTaskDescInfoList {
	s.ExecuteTime = &v
	return s
}


type GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder struct {
	s *GetPolicyInfoByInstanceIdResponseTaskDescInfoList
}

func NewGetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder() *GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder {
	s := &GetPolicyInfoByInstanceIdResponseTaskDescInfoList{}
	b := &GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder) TaskType(v string) *GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder {
	b.s.TaskType = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder) ExecutePolicy(v string) *GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder {
	b.s.ExecutePolicy = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder) ExecuteTime(v string) *GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder {
	b.s.ExecuteTime = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseTaskDescInfoListBuilder) Build() *GetPolicyInfoByInstanceIdResponseTaskDescInfoList {
	return b.s
}


