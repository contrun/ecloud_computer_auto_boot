// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserPolicyDetailsResponseLockRuleList struct {

    // 锁定天 0：每天 1：周一 2：周.......7:周日
	LockMode *int32 `json:"lockMode,omitempty"`
    // 开始时间
	BeginTime *string `json:"beginTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
}

func (s GetUserPolicyDetailsResponseLockRuleList) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyDetailsResponseLockRuleList) GoString() string {
	return s.String()
}

func (s GetUserPolicyDetailsResponseLockRuleList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyDetailsResponseLockRuleList) SetLockMode(v int32) *GetUserPolicyDetailsResponseLockRuleList {
	s.LockMode = &v
	return s
}

func (s *GetUserPolicyDetailsResponseLockRuleList) SetBeginTime(v string) *GetUserPolicyDetailsResponseLockRuleList {
	s.BeginTime = &v
	return s
}

func (s *GetUserPolicyDetailsResponseLockRuleList) SetEndTime(v string) *GetUserPolicyDetailsResponseLockRuleList {
	s.EndTime = &v
	return s
}


type GetUserPolicyDetailsResponseLockRuleListBuilder struct {
	s *GetUserPolicyDetailsResponseLockRuleList
}

func NewGetUserPolicyDetailsResponseLockRuleListBuilder() *GetUserPolicyDetailsResponseLockRuleListBuilder {
	s := &GetUserPolicyDetailsResponseLockRuleList{}
	b := &GetUserPolicyDetailsResponseLockRuleListBuilder{s: s}
	return b
}

func (b *GetUserPolicyDetailsResponseLockRuleListBuilder) LockMode(v int32) *GetUserPolicyDetailsResponseLockRuleListBuilder {
	b.s.LockMode = &v
	return b
}

func (b *GetUserPolicyDetailsResponseLockRuleListBuilder) BeginTime(v string) *GetUserPolicyDetailsResponseLockRuleListBuilder {
	b.s.BeginTime = &v
	return b
}

func (b *GetUserPolicyDetailsResponseLockRuleListBuilder) EndTime(v string) *GetUserPolicyDetailsResponseLockRuleListBuilder {
	b.s.EndTime = &v
	return b
}

func (b *GetUserPolicyDetailsResponseLockRuleListBuilder) Build() *GetUserPolicyDetailsResponseLockRuleList {
	return b.s
}


