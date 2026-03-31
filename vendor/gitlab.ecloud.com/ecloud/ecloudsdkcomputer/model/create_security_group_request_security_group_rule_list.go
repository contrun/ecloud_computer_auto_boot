// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSecurityGroupRequestSecurityGroupRuleList struct {

    // 填 cidr 或者 securityGroup, 当前写死为 cidr
	RemoteType *string `json:"remoteType,omitempty"`
    // 协议类型，TCP/UPD/ICMP/GRE/VRRP/IGMP/ANY
	Protocol *string `json:"protocol,omitempty"`
    // 下行端口（协议为ANY和ICMP时填null，表示所有端口，TCP、UDP必填，与上行端口值相同时表示指定端口号）
	MinPortRange *string `json:"minPortRange,omitempty"`
    // IP类型,IPv4 或 IPv6
	EtherType *string `json:"etherType,omitempty"`
    // 备注
	Remark *string `json:"remark,omitempty"`
    // 授权对象
	RemoteIpPrefix *string `json:"remoteIpPrefix,omitempty"`
    // 规则方向,ingress:入,egress:出
	Direction *string `json:"direction,omitempty"`
    // 上行端口（协议为ANY和ICMP时填null，表示所有端口，TCP、UDP必填，与下行端口值相同时表示指定端口号）
	MaxPortRange *string `json:"maxPortRange,omitempty"`
}

func (s CreateSecurityGroupRequestSecurityGroupRuleList) String() string {
	return utils.Beautify(s)
}

func (s CreateSecurityGroupRequestSecurityGroupRuleList) GoString() string {
	return s.String()
}

func (s CreateSecurityGroupRequestSecurityGroupRuleList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSecurityGroupRequestSecurityGroupRuleList) SetRemoteType(v string) *CreateSecurityGroupRequestSecurityGroupRuleList {
	s.RemoteType = &v
	return s
}

func (s *CreateSecurityGroupRequestSecurityGroupRuleList) SetProtocol(v string) *CreateSecurityGroupRequestSecurityGroupRuleList {
	s.Protocol = &v
	return s
}

func (s *CreateSecurityGroupRequestSecurityGroupRuleList) SetMinPortRange(v string) *CreateSecurityGroupRequestSecurityGroupRuleList {
	s.MinPortRange = &v
	return s
}

func (s *CreateSecurityGroupRequestSecurityGroupRuleList) SetEtherType(v string) *CreateSecurityGroupRequestSecurityGroupRuleList {
	s.EtherType = &v
	return s
}

func (s *CreateSecurityGroupRequestSecurityGroupRuleList) SetRemark(v string) *CreateSecurityGroupRequestSecurityGroupRuleList {
	s.Remark = &v
	return s
}

func (s *CreateSecurityGroupRequestSecurityGroupRuleList) SetRemoteIpPrefix(v string) *CreateSecurityGroupRequestSecurityGroupRuleList {
	s.RemoteIpPrefix = &v
	return s
}

func (s *CreateSecurityGroupRequestSecurityGroupRuleList) SetDirection(v string) *CreateSecurityGroupRequestSecurityGroupRuleList {
	s.Direction = &v
	return s
}

func (s *CreateSecurityGroupRequestSecurityGroupRuleList) SetMaxPortRange(v string) *CreateSecurityGroupRequestSecurityGroupRuleList {
	s.MaxPortRange = &v
	return s
}


type CreateSecurityGroupRequestSecurityGroupRuleListBuilder struct {
	s *CreateSecurityGroupRequestSecurityGroupRuleList
}

func NewCreateSecurityGroupRequestSecurityGroupRuleListBuilder() *CreateSecurityGroupRequestSecurityGroupRuleListBuilder {
	s := &CreateSecurityGroupRequestSecurityGroupRuleList{}
	b := &CreateSecurityGroupRequestSecurityGroupRuleListBuilder{s: s}
	return b
}

func (b *CreateSecurityGroupRequestSecurityGroupRuleListBuilder) RemoteType(v string) *CreateSecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.RemoteType = &v
	return b
}

func (b *CreateSecurityGroupRequestSecurityGroupRuleListBuilder) Protocol(v string) *CreateSecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.Protocol = &v
	return b
}

func (b *CreateSecurityGroupRequestSecurityGroupRuleListBuilder) MinPortRange(v string) *CreateSecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.MinPortRange = &v
	return b
}

func (b *CreateSecurityGroupRequestSecurityGroupRuleListBuilder) EtherType(v string) *CreateSecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.EtherType = &v
	return b
}

func (b *CreateSecurityGroupRequestSecurityGroupRuleListBuilder) Remark(v string) *CreateSecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.Remark = &v
	return b
}

func (b *CreateSecurityGroupRequestSecurityGroupRuleListBuilder) RemoteIpPrefix(v string) *CreateSecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.RemoteIpPrefix = &v
	return b
}

func (b *CreateSecurityGroupRequestSecurityGroupRuleListBuilder) Direction(v string) *CreateSecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.Direction = &v
	return b
}

func (b *CreateSecurityGroupRequestSecurityGroupRuleListBuilder) MaxPortRange(v string) *CreateSecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.MaxPortRange = &v
	return b
}

func (b *CreateSecurityGroupRequestSecurityGroupRuleListBuilder) Build() *CreateSecurityGroupRequestSecurityGroupRuleList {
	return b.s
}


