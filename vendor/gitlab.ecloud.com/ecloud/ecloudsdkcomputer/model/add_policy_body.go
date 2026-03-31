// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddPolicyBody struct {
    position.Body
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 定时任务类型（开机available,关机shutdown，重启restart）
	TaskTypeList *[]AddPolicyRequestTaskTypeList `json:"taskTypeList,omitempty"`
    // 所属资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
}

func (s AddPolicyBody) String() string {
	return utils.Beautify(s)
}

func (s AddPolicyBody) GoString() string {
	return s.String()
}

func (s AddPolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddPolicyBody) SetPolicyName(v string) *AddPolicyBody {
	s.PolicyName = &v
	return s
}

func (s *AddPolicyBody) SetPolicyDesc(v string) *AddPolicyBody {
	s.PolicyDesc = &v
	return s
}

func (s *AddPolicyBody) SetTaskTypeList(v []AddPolicyRequestTaskTypeList) *AddPolicyBody {
	s.TaskTypeList = &v
	return s
}

func (s *AddPolicyBody) SetPoolCode(v string) *AddPolicyBody {
	s.PoolCode = &v
	return s
}


type AddPolicyBodyBuilder struct {
	s *AddPolicyBody
}

func NewAddPolicyBodyBuilder() *AddPolicyBodyBuilder {
	s := &AddPolicyBody{}
	b := &AddPolicyBodyBuilder{s: s}
	return b
}

func (b *AddPolicyBodyBuilder) PolicyName(v string) *AddPolicyBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *AddPolicyBodyBuilder) PolicyDesc(v string) *AddPolicyBodyBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *AddPolicyBodyBuilder) TaskTypeList(v []AddPolicyRequestTaskTypeList) *AddPolicyBodyBuilder {
	b.s.TaskTypeList = &v
	return b
}

func (b *AddPolicyBodyBuilder) PoolCode(v string) *AddPolicyBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *AddPolicyBodyBuilder) Build() *AddPolicyBody {
	return b.s
}


