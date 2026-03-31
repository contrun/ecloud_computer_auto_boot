// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddUserPolicyRequestIpRuleList struct {

    // 备注
	Remark *string `json:"remark,omitempty"`
    // ip地址
	IpAddr *string `json:"ipAddr,omitempty"`
    // 用户设置ip类型  1 ip 2 ip段
	IpType *int32 `json:"ipType,omitempty"`
}

func (s AddUserPolicyRequestIpRuleList) String() string {
	return utils.Beautify(s)
}

func (s AddUserPolicyRequestIpRuleList) GoString() string {
	return s.String()
}

func (s AddUserPolicyRequestIpRuleList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddUserPolicyRequestIpRuleList) SetRemark(v string) *AddUserPolicyRequestIpRuleList {
	s.Remark = &v
	return s
}

func (s *AddUserPolicyRequestIpRuleList) SetIpAddr(v string) *AddUserPolicyRequestIpRuleList {
	s.IpAddr = &v
	return s
}

func (s *AddUserPolicyRequestIpRuleList) SetIpType(v int32) *AddUserPolicyRequestIpRuleList {
	s.IpType = &v
	return s
}


type AddUserPolicyRequestIpRuleListBuilder struct {
	s *AddUserPolicyRequestIpRuleList
}

func NewAddUserPolicyRequestIpRuleListBuilder() *AddUserPolicyRequestIpRuleListBuilder {
	s := &AddUserPolicyRequestIpRuleList{}
	b := &AddUserPolicyRequestIpRuleListBuilder{s: s}
	return b
}

func (b *AddUserPolicyRequestIpRuleListBuilder) Remark(v string) *AddUserPolicyRequestIpRuleListBuilder {
	b.s.Remark = &v
	return b
}

func (b *AddUserPolicyRequestIpRuleListBuilder) IpAddr(v string) *AddUserPolicyRequestIpRuleListBuilder {
	b.s.IpAddr = &v
	return b
}

func (b *AddUserPolicyRequestIpRuleListBuilder) IpType(v int32) *AddUserPolicyRequestIpRuleListBuilder {
	b.s.IpType = &v
	return b
}

func (b *AddUserPolicyRequestIpRuleListBuilder) Build() *AddUserPolicyRequestIpRuleList {
	return b.s
}


