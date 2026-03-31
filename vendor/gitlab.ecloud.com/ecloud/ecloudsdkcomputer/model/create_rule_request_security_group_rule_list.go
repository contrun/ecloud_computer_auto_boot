// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateRuleRequestSecurityGroupRuleList struct {

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

func (s CreateRuleRequestSecurityGroupRuleList) String() string {
	return utils.Beautify(s)
}

func (s CreateRuleRequestSecurityGroupRuleList) GoString() string {
	return s.String()
}

func (s CreateRuleRequestSecurityGroupRuleList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateRuleRequestSecurityGroupRuleList) SetRemoteType(v string) *CreateRuleRequestSecurityGroupRuleList {
	s.RemoteType = &v
	return s
}

func (s *CreateRuleRequestSecurityGroupRuleList) SetProtocol(v string) *CreateRuleRequestSecurityGroupRuleList {
	s.Protocol = &v
	return s
}

func (s *CreateRuleRequestSecurityGroupRuleList) SetMinPortRange(v string) *CreateRuleRequestSecurityGroupRuleList {
	s.MinPortRange = &v
	return s
}

func (s *CreateRuleRequestSecurityGroupRuleList) SetEtherType(v string) *CreateRuleRequestSecurityGroupRuleList {
	s.EtherType = &v
	return s
}

func (s *CreateRuleRequestSecurityGroupRuleList) SetRemark(v string) *CreateRuleRequestSecurityGroupRuleList {
	s.Remark = &v
	return s
}

func (s *CreateRuleRequestSecurityGroupRuleList) SetRemoteIpPrefix(v string) *CreateRuleRequestSecurityGroupRuleList {
	s.RemoteIpPrefix = &v
	return s
}

func (s *CreateRuleRequestSecurityGroupRuleList) SetDirection(v string) *CreateRuleRequestSecurityGroupRuleList {
	s.Direction = &v
	return s
}

func (s *CreateRuleRequestSecurityGroupRuleList) SetMaxPortRange(v string) *CreateRuleRequestSecurityGroupRuleList {
	s.MaxPortRange = &v
	return s
}


type CreateRuleRequestSecurityGroupRuleListBuilder struct {
	s *CreateRuleRequestSecurityGroupRuleList
}

func NewCreateRuleRequestSecurityGroupRuleListBuilder() *CreateRuleRequestSecurityGroupRuleListBuilder {
	s := &CreateRuleRequestSecurityGroupRuleList{}
	b := &CreateRuleRequestSecurityGroupRuleListBuilder{s: s}
	return b
}

func (b *CreateRuleRequestSecurityGroupRuleListBuilder) RemoteType(v string) *CreateRuleRequestSecurityGroupRuleListBuilder {
	b.s.RemoteType = &v
	return b
}

func (b *CreateRuleRequestSecurityGroupRuleListBuilder) Protocol(v string) *CreateRuleRequestSecurityGroupRuleListBuilder {
	b.s.Protocol = &v
	return b
}

func (b *CreateRuleRequestSecurityGroupRuleListBuilder) MinPortRange(v string) *CreateRuleRequestSecurityGroupRuleListBuilder {
	b.s.MinPortRange = &v
	return b
}

func (b *CreateRuleRequestSecurityGroupRuleListBuilder) EtherType(v string) *CreateRuleRequestSecurityGroupRuleListBuilder {
	b.s.EtherType = &v
	return b
}

func (b *CreateRuleRequestSecurityGroupRuleListBuilder) Remark(v string) *CreateRuleRequestSecurityGroupRuleListBuilder {
	b.s.Remark = &v
	return b
}

func (b *CreateRuleRequestSecurityGroupRuleListBuilder) RemoteIpPrefix(v string) *CreateRuleRequestSecurityGroupRuleListBuilder {
	b.s.RemoteIpPrefix = &v
	return b
}

func (b *CreateRuleRequestSecurityGroupRuleListBuilder) Direction(v string) *CreateRuleRequestSecurityGroupRuleListBuilder {
	b.s.Direction = &v
	return b
}

func (b *CreateRuleRequestSecurityGroupRuleListBuilder) MaxPortRange(v string) *CreateRuleRequestSecurityGroupRuleListBuilder {
	b.s.MaxPortRange = &v
	return b
}

func (b *CreateRuleRequestSecurityGroupRuleListBuilder) Build() *CreateRuleRequestSecurityGroupRuleList {
	return b.s
}


