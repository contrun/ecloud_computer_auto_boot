// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateRuleBody struct {
    position.Body
    // 协议类型，TCP/UPD/ICMP/GRE/VRRP/IGMP/ANY
	Protocol *string `json:"protocol,omitempty"`
    // 下行端口（协议为ANY和ICMP时填null，表示所有端口，TCP、UDP必填，与上行端口值相同时表示指定端口号）
	MinPortRange *string `json:"minPortRange,omitempty"`
    // IP类型,IPv4 或 IPv6
	EtherType *string `json:"etherType,omitempty"`
    // 备注
	Remark *string `json:"remark,omitempty"`
    // 规则UID
	RuleUid *string `json:"ruleUid,omitempty"`
    // 授权对象
	RemoteIpPrefix *string `json:"remoteIpPrefix,omitempty"`
    // 安全组UID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
    // 规则方向,ingress:入,egress:出
	Direction *string `json:"direction,omitempty"`
    // 上行端口（协议为ANY和ICMP时填null，表示所有端口，TCP、UDP必填，与下行端口值相同时表示指定端口号）
	MaxPortRange *string `json:"maxPortRange,omitempty"`
}

func (s UpdateRuleBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateRuleBody) GoString() string {
	return s.String()
}

func (s UpdateRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateRuleBody) SetProtocol(v string) *UpdateRuleBody {
	s.Protocol = &v
	return s
}

func (s *UpdateRuleBody) SetMinPortRange(v string) *UpdateRuleBody {
	s.MinPortRange = &v
	return s
}

func (s *UpdateRuleBody) SetEtherType(v string) *UpdateRuleBody {
	s.EtherType = &v
	return s
}

func (s *UpdateRuleBody) SetRemark(v string) *UpdateRuleBody {
	s.Remark = &v
	return s
}

func (s *UpdateRuleBody) SetRuleUid(v string) *UpdateRuleBody {
	s.RuleUid = &v
	return s
}

func (s *UpdateRuleBody) SetRemoteIpPrefix(v string) *UpdateRuleBody {
	s.RemoteIpPrefix = &v
	return s
}

func (s *UpdateRuleBody) SetSecurityGroupUid(v string) *UpdateRuleBody {
	s.SecurityGroupUid = &v
	return s
}

func (s *UpdateRuleBody) SetDirection(v string) *UpdateRuleBody {
	s.Direction = &v
	return s
}

func (s *UpdateRuleBody) SetMaxPortRange(v string) *UpdateRuleBody {
	s.MaxPortRange = &v
	return s
}


type UpdateRuleBodyBuilder struct {
	s *UpdateRuleBody
}

func NewUpdateRuleBodyBuilder() *UpdateRuleBodyBuilder {
	s := &UpdateRuleBody{}
	b := &UpdateRuleBodyBuilder{s: s}
	return b
}

func (b *UpdateRuleBodyBuilder) Protocol(v string) *UpdateRuleBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *UpdateRuleBodyBuilder) MinPortRange(v string) *UpdateRuleBodyBuilder {
	b.s.MinPortRange = &v
	return b
}

func (b *UpdateRuleBodyBuilder) EtherType(v string) *UpdateRuleBodyBuilder {
	b.s.EtherType = &v
	return b
}

func (b *UpdateRuleBodyBuilder) Remark(v string) *UpdateRuleBodyBuilder {
	b.s.Remark = &v
	return b
}

func (b *UpdateRuleBodyBuilder) RuleUid(v string) *UpdateRuleBodyBuilder {
	b.s.RuleUid = &v
	return b
}

func (b *UpdateRuleBodyBuilder) RemoteIpPrefix(v string) *UpdateRuleBodyBuilder {
	b.s.RemoteIpPrefix = &v
	return b
}

func (b *UpdateRuleBodyBuilder) SecurityGroupUid(v string) *UpdateRuleBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *UpdateRuleBodyBuilder) Direction(v string) *UpdateRuleBodyBuilder {
	b.s.Direction = &v
	return b
}

func (b *UpdateRuleBodyBuilder) MaxPortRange(v string) *UpdateRuleBodyBuilder {
	b.s.MaxPortRange = &v
	return b
}

func (b *UpdateRuleBodyBuilder) Build() *UpdateRuleBody {
	return b.s
}


