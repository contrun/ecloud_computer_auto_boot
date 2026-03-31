// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EditUserPolicyRequestIpRuleList struct {

    // 备注
	Remark *string `json:"remark,omitempty"`
    // ip地址
	IpAddr *string `json:"ipAddr,omitempty"`
    // 用户设置ip类型  1 ip 2 ip段
	IpType *int32 `json:"ipType,omitempty"`
}

func (s EditUserPolicyRequestIpRuleList) String() string {
	return utils.Beautify(s)
}

func (s EditUserPolicyRequestIpRuleList) GoString() string {
	return s.String()
}

func (s EditUserPolicyRequestIpRuleList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditUserPolicyRequestIpRuleList) SetRemark(v string) *EditUserPolicyRequestIpRuleList {
	s.Remark = &v
	return s
}

func (s *EditUserPolicyRequestIpRuleList) SetIpAddr(v string) *EditUserPolicyRequestIpRuleList {
	s.IpAddr = &v
	return s
}

func (s *EditUserPolicyRequestIpRuleList) SetIpType(v int32) *EditUserPolicyRequestIpRuleList {
	s.IpType = &v
	return s
}


type EditUserPolicyRequestIpRuleListBuilder struct {
	s *EditUserPolicyRequestIpRuleList
}

func NewEditUserPolicyRequestIpRuleListBuilder() *EditUserPolicyRequestIpRuleListBuilder {
	s := &EditUserPolicyRequestIpRuleList{}
	b := &EditUserPolicyRequestIpRuleListBuilder{s: s}
	return b
}

func (b *EditUserPolicyRequestIpRuleListBuilder) Remark(v string) *EditUserPolicyRequestIpRuleListBuilder {
	b.s.Remark = &v
	return b
}

func (b *EditUserPolicyRequestIpRuleListBuilder) IpAddr(v string) *EditUserPolicyRequestIpRuleListBuilder {
	b.s.IpAddr = &v
	return b
}

func (b *EditUserPolicyRequestIpRuleListBuilder) IpType(v int32) *EditUserPolicyRequestIpRuleListBuilder {
	b.s.IpType = &v
	return b
}

func (b *EditUserPolicyRequestIpRuleListBuilder) Build() *EditUserPolicyRequestIpRuleList {
	return b.s
}


