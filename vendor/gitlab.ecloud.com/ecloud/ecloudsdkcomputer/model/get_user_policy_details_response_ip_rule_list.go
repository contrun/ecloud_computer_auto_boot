// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserPolicyDetailsResponseIpRuleList struct {

    // 备注
	Remark *string `json:"remark,omitempty"`
    // ip地址
	IpAddr *string `json:"ipAddr,omitempty"`
    // 用户设置ip类型  1 ip 2 ip段
	IpType *int32 `json:"ipType,omitempty"`
}

func (s GetUserPolicyDetailsResponseIpRuleList) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyDetailsResponseIpRuleList) GoString() string {
	return s.String()
}

func (s GetUserPolicyDetailsResponseIpRuleList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyDetailsResponseIpRuleList) SetRemark(v string) *GetUserPolicyDetailsResponseIpRuleList {
	s.Remark = &v
	return s
}

func (s *GetUserPolicyDetailsResponseIpRuleList) SetIpAddr(v string) *GetUserPolicyDetailsResponseIpRuleList {
	s.IpAddr = &v
	return s
}

func (s *GetUserPolicyDetailsResponseIpRuleList) SetIpType(v int32) *GetUserPolicyDetailsResponseIpRuleList {
	s.IpType = &v
	return s
}


type GetUserPolicyDetailsResponseIpRuleListBuilder struct {
	s *GetUserPolicyDetailsResponseIpRuleList
}

func NewGetUserPolicyDetailsResponseIpRuleListBuilder() *GetUserPolicyDetailsResponseIpRuleListBuilder {
	s := &GetUserPolicyDetailsResponseIpRuleList{}
	b := &GetUserPolicyDetailsResponseIpRuleListBuilder{s: s}
	return b
}

func (b *GetUserPolicyDetailsResponseIpRuleListBuilder) Remark(v string) *GetUserPolicyDetailsResponseIpRuleListBuilder {
	b.s.Remark = &v
	return b
}

func (b *GetUserPolicyDetailsResponseIpRuleListBuilder) IpAddr(v string) *GetUserPolicyDetailsResponseIpRuleListBuilder {
	b.s.IpAddr = &v
	return b
}

func (b *GetUserPolicyDetailsResponseIpRuleListBuilder) IpType(v int32) *GetUserPolicyDetailsResponseIpRuleListBuilder {
	b.s.IpType = &v
	return b
}

func (b *GetUserPolicyDetailsResponseIpRuleListBuilder) Build() *GetUserPolicyDetailsResponseIpRuleList {
	return b.s
}


