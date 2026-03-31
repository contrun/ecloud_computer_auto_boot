// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdatePolicyBody struct {
    position.Body
    // 定时任务策略uuid
	PolicyUid *string `json:"policyUid,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 任务类型
	TaskTypeList *[]UpdatePolicyRequestTaskTypeList `json:"taskTypeList,omitempty"`
}

func (s UpdatePolicyBody) String() string {
	return utils.Beautify(s)
}

func (s UpdatePolicyBody) GoString() string {
	return s.String()
}

func (s UpdatePolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdatePolicyBody) SetPolicyUid(v string) *UpdatePolicyBody {
	s.PolicyUid = &v
	return s
}

func (s *UpdatePolicyBody) SetPolicyName(v string) *UpdatePolicyBody {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyBody) SetPolicyDesc(v string) *UpdatePolicyBody {
	s.PolicyDesc = &v
	return s
}

func (s *UpdatePolicyBody) SetTaskTypeList(v []UpdatePolicyRequestTaskTypeList) *UpdatePolicyBody {
	s.TaskTypeList = &v
	return s
}


type UpdatePolicyBodyBuilder struct {
	s *UpdatePolicyBody
}

func NewUpdatePolicyBodyBuilder() *UpdatePolicyBodyBuilder {
	s := &UpdatePolicyBody{}
	b := &UpdatePolicyBodyBuilder{s: s}
	return b
}

func (b *UpdatePolicyBodyBuilder) PolicyUid(v string) *UpdatePolicyBodyBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *UpdatePolicyBodyBuilder) PolicyName(v string) *UpdatePolicyBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *UpdatePolicyBodyBuilder) PolicyDesc(v string) *UpdatePolicyBodyBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *UpdatePolicyBodyBuilder) TaskTypeList(v []UpdatePolicyRequestTaskTypeList) *UpdatePolicyBodyBuilder {
	b.s.TaskTypeList = &v
	return b
}

func (b *UpdatePolicyBodyBuilder) Build() *UpdatePolicyBody {
	return b.s
}


