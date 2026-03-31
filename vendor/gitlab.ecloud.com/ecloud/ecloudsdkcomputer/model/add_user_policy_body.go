// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddUserPolicyBody struct {
    position.Body
    // 锁定天 0：每天 1：周一 2：周.......7:周日
	LockMode []int32 `json:"lockMode,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 企业门户组织id
	CcempOrganizeId *string `json:"ccempOrganizeId,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 接入白名单设置
	IpRuleList *[]AddUserPolicyRequestIpRuleList `json:"ipRuleList,omitempty"`
    // 用户锁定时间
	LockTimes *[]AddUserPolicyRequestLockTimes `json:"lockTimes,omitempty"`
    // 用户锁定类型 0未锁定 1永久锁定 2按固定时间锁定
	LockType *int32 `json:"lockType,omitempty"`
}

func (s AddUserPolicyBody) String() string {
	return utils.Beautify(s)
}

func (s AddUserPolicyBody) GoString() string {
	return s.String()
}

func (s AddUserPolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddUserPolicyBody) SetLockMode(v []int32) *AddUserPolicyBody {
	s.LockMode = v
	return s
}

func (s *AddUserPolicyBody) SetPolicyName(v string) *AddUserPolicyBody {
	s.PolicyName = &v
	return s
}

func (s *AddUserPolicyBody) SetCcempOrganizeId(v string) *AddUserPolicyBody {
	s.CcempOrganizeId = &v
	return s
}

func (s *AddUserPolicyBody) SetPolicyDesc(v string) *AddUserPolicyBody {
	s.PolicyDesc = &v
	return s
}

func (s *AddUserPolicyBody) SetIpRuleList(v []AddUserPolicyRequestIpRuleList) *AddUserPolicyBody {
	s.IpRuleList = &v
	return s
}

func (s *AddUserPolicyBody) SetLockTimes(v []AddUserPolicyRequestLockTimes) *AddUserPolicyBody {
	s.LockTimes = &v
	return s
}

func (s *AddUserPolicyBody) SetLockType(v int32) *AddUserPolicyBody {
	s.LockType = &v
	return s
}


type AddUserPolicyBodyBuilder struct {
	s *AddUserPolicyBody
}

func NewAddUserPolicyBodyBuilder() *AddUserPolicyBodyBuilder {
	s := &AddUserPolicyBody{}
	b := &AddUserPolicyBodyBuilder{s: s}
	return b
}

func (b *AddUserPolicyBodyBuilder) LockMode(v []int32) *AddUserPolicyBodyBuilder {
	b.s.LockMode = v
	return b
}

func (b *AddUserPolicyBodyBuilder) PolicyName(v string) *AddUserPolicyBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *AddUserPolicyBodyBuilder) CcempOrganizeId(v string) *AddUserPolicyBodyBuilder {
	b.s.CcempOrganizeId = &v
	return b
}

func (b *AddUserPolicyBodyBuilder) PolicyDesc(v string) *AddUserPolicyBodyBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *AddUserPolicyBodyBuilder) IpRuleList(v []AddUserPolicyRequestIpRuleList) *AddUserPolicyBodyBuilder {
	b.s.IpRuleList = &v
	return b
}

func (b *AddUserPolicyBodyBuilder) LockTimes(v []AddUserPolicyRequestLockTimes) *AddUserPolicyBodyBuilder {
	b.s.LockTimes = &v
	return b
}

func (b *AddUserPolicyBodyBuilder) LockType(v int32) *AddUserPolicyBodyBuilder {
	b.s.LockType = &v
	return b
}

func (b *AddUserPolicyBodyBuilder) Build() *AddUserPolicyBody {
	return b.s
}


