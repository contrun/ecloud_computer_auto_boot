// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CopySecurityGroupRequestSecurityGroupRuleList struct {

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

func (s CopySecurityGroupRequestSecurityGroupRuleList) String() string {
	return utils.Beautify(s)
}

func (s CopySecurityGroupRequestSecurityGroupRuleList) GoString() string {
	return s.String()
}

func (s CopySecurityGroupRequestSecurityGroupRuleList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CopySecurityGroupRequestSecurityGroupRuleList) SetRemoteType(v string) *CopySecurityGroupRequestSecurityGroupRuleList {
	s.RemoteType = &v
	return s
}

func (s *CopySecurityGroupRequestSecurityGroupRuleList) SetProtocol(v string) *CopySecurityGroupRequestSecurityGroupRuleList {
	s.Protocol = &v
	return s
}

func (s *CopySecurityGroupRequestSecurityGroupRuleList) SetMinPortRange(v string) *CopySecurityGroupRequestSecurityGroupRuleList {
	s.MinPortRange = &v
	return s
}

func (s *CopySecurityGroupRequestSecurityGroupRuleList) SetEtherType(v string) *CopySecurityGroupRequestSecurityGroupRuleList {
	s.EtherType = &v
	return s
}

func (s *CopySecurityGroupRequestSecurityGroupRuleList) SetRemark(v string) *CopySecurityGroupRequestSecurityGroupRuleList {
	s.Remark = &v
	return s
}

func (s *CopySecurityGroupRequestSecurityGroupRuleList) SetRemoteIpPrefix(v string) *CopySecurityGroupRequestSecurityGroupRuleList {
	s.RemoteIpPrefix = &v
	return s
}

func (s *CopySecurityGroupRequestSecurityGroupRuleList) SetDirection(v string) *CopySecurityGroupRequestSecurityGroupRuleList {
	s.Direction = &v
	return s
}

func (s *CopySecurityGroupRequestSecurityGroupRuleList) SetMaxPortRange(v string) *CopySecurityGroupRequestSecurityGroupRuleList {
	s.MaxPortRange = &v
	return s
}


type CopySecurityGroupRequestSecurityGroupRuleListBuilder struct {
	s *CopySecurityGroupRequestSecurityGroupRuleList
}

func NewCopySecurityGroupRequestSecurityGroupRuleListBuilder() *CopySecurityGroupRequestSecurityGroupRuleListBuilder {
	s := &CopySecurityGroupRequestSecurityGroupRuleList{}
	b := &CopySecurityGroupRequestSecurityGroupRuleListBuilder{s: s}
	return b
}

func (b *CopySecurityGroupRequestSecurityGroupRuleListBuilder) RemoteType(v string) *CopySecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.RemoteType = &v
	return b
}

func (b *CopySecurityGroupRequestSecurityGroupRuleListBuilder) Protocol(v string) *CopySecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.Protocol = &v
	return b
}

func (b *CopySecurityGroupRequestSecurityGroupRuleListBuilder) MinPortRange(v string) *CopySecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.MinPortRange = &v
	return b
}

func (b *CopySecurityGroupRequestSecurityGroupRuleListBuilder) EtherType(v string) *CopySecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.EtherType = &v
	return b
}

func (b *CopySecurityGroupRequestSecurityGroupRuleListBuilder) Remark(v string) *CopySecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.Remark = &v
	return b
}

func (b *CopySecurityGroupRequestSecurityGroupRuleListBuilder) RemoteIpPrefix(v string) *CopySecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.RemoteIpPrefix = &v
	return b
}

func (b *CopySecurityGroupRequestSecurityGroupRuleListBuilder) Direction(v string) *CopySecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.Direction = &v
	return b
}

func (b *CopySecurityGroupRequestSecurityGroupRuleListBuilder) MaxPortRange(v string) *CopySecurityGroupRequestSecurityGroupRuleListBuilder {
	b.s.MaxPortRange = &v
	return b
}

func (b *CopySecurityGroupRequestSecurityGroupRuleListBuilder) Build() *CopySecurityGroupRequestSecurityGroupRuleList {
	return b.s
}


