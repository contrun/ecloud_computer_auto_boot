// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EditUserPolicyBody struct {
    position.Body
    // 锁定天 0：每天 1：周一 2：周.......7:周日
	LockMode []int32 `json:"lockMode,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 接入白名单设置
	IpRuleList *[]EditUserPolicyRequestIpRuleList `json:"ipRuleList,omitempty"`
    // 用户锁定时间
	LockTimes *[]EditUserPolicyRequestLockTimes `json:"lockTimes,omitempty"`
    // 用户锁定类型 0未锁定 1永久锁定 2按固定时间锁定
	LockType *int32 `json:"lockType,omitempty"`
    // 策略id
	UserPolicyId *string `json:"userPolicyId,omitempty"`
}

func (s EditUserPolicyBody) String() string {
	return utils.Beautify(s)
}

func (s EditUserPolicyBody) GoString() string {
	return s.String()
}

func (s EditUserPolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditUserPolicyBody) SetLockMode(v []int32) *EditUserPolicyBody {
	s.LockMode = v
	return s
}

func (s *EditUserPolicyBody) SetPolicyName(v string) *EditUserPolicyBody {
	s.PolicyName = &v
	return s
}

func (s *EditUserPolicyBody) SetPolicyDesc(v string) *EditUserPolicyBody {
	s.PolicyDesc = &v
	return s
}

func (s *EditUserPolicyBody) SetIpRuleList(v []EditUserPolicyRequestIpRuleList) *EditUserPolicyBody {
	s.IpRuleList = &v
	return s
}

func (s *EditUserPolicyBody) SetLockTimes(v []EditUserPolicyRequestLockTimes) *EditUserPolicyBody {
	s.LockTimes = &v
	return s
}

func (s *EditUserPolicyBody) SetLockType(v int32) *EditUserPolicyBody {
	s.LockType = &v
	return s
}

func (s *EditUserPolicyBody) SetUserPolicyId(v string) *EditUserPolicyBody {
	s.UserPolicyId = &v
	return s
}


type EditUserPolicyBodyBuilder struct {
	s *EditUserPolicyBody
}

func NewEditUserPolicyBodyBuilder() *EditUserPolicyBodyBuilder {
	s := &EditUserPolicyBody{}
	b := &EditUserPolicyBodyBuilder{s: s}
	return b
}

func (b *EditUserPolicyBodyBuilder) LockMode(v []int32) *EditUserPolicyBodyBuilder {
	b.s.LockMode = v
	return b
}

func (b *EditUserPolicyBodyBuilder) PolicyName(v string) *EditUserPolicyBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *EditUserPolicyBodyBuilder) PolicyDesc(v string) *EditUserPolicyBodyBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *EditUserPolicyBodyBuilder) IpRuleList(v []EditUserPolicyRequestIpRuleList) *EditUserPolicyBodyBuilder {
	b.s.IpRuleList = &v
	return b
}

func (b *EditUserPolicyBodyBuilder) LockTimes(v []EditUserPolicyRequestLockTimes) *EditUserPolicyBodyBuilder {
	b.s.LockTimes = &v
	return b
}

func (b *EditUserPolicyBodyBuilder) LockType(v int32) *EditUserPolicyBodyBuilder {
	b.s.LockType = &v
	return b
}

func (b *EditUserPolicyBodyBuilder) UserPolicyId(v string) *EditUserPolicyBodyBuilder {
	b.s.UserPolicyId = &v
	return b
}

func (b *EditUserPolicyBodyBuilder) Build() *EditUserPolicyBody {
	return b.s
}


