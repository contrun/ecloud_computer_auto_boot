// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSecurityGroupRuleInfoResponseRuleLists struct {


	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 下行端口（协议为ANY和ICMP时填null，表示所有端口，TCP、UDP必填，与下行端口值相同时表示指定端口号）
	MinPortRange *string `json:"minPortRange,omitempty"`
    // 备注
	Remark *string `json:"remark,omitempty"`
    // cidr,remoteType为cidr时必填
	RemoteIpPrefix *string `json:"remoteIpPrefix,omitempty"`
    // 填cidr或者securityGroup，cidr表示地址端，securityGroup表示远端安全组
	RemoteType *string `json:"remoteType,omitempty"`
    // 协议类型（TCP、UDP、ICMP、ANY）
	Protocol *string `json:"protocol,omitempty"`

	IsDeleted *string `json:"isDeleted,omitempty"`
    // 远端安全组id，remoteType为securityGroup时必填
	RemoteSecurityGroupUid *string `json:"remoteSecurityGroupUid,omitempty"`
    // 规则类型(默认-1、其他-0)
	RuleType *string `json:"ruleType,omitempty"`
    // IP类型,IPv4或IPv6
	EtherType *string `json:"etherType,omitempty"`

	CreatedTime *string `json:"createdTime,omitempty"`
    // 安全组规则uid
	RuleUid *string `json:"ruleUid,omitempty"`

	Id *int64 `json:"id,omitempty"`
    // 安全组uid
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
    // 出入方向（ingress:入，egress:出）
	Direction *string `json:"direction,omitempty"`
    // 上行端口（协议为ANY和ICMP时填null，表示所有端口，TCP、UDP必填，与下行端口值相同时表示指定端口号）
	MaxPortRange *string `json:"maxPortRange,omitempty"`
    // 安全组规则状态(creating创建中,active启用中,createError创建失败,deleting删除中,deleted已删除,deleteError删除失败,alter变更中,alterError变更失败)
	Status *string `json:"status,omitempty"`
}

func (s GetSecurityGroupRuleInfoResponseRuleLists) String() string {
	return utils.Beautify(s)
}

func (s GetSecurityGroupRuleInfoResponseRuleLists) GoString() string {
	return s.String()
}

func (s GetSecurityGroupRuleInfoResponseRuleLists) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetModifiedTime(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.ModifiedTime = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetMinPortRange(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.MinPortRange = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetRemark(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.Remark = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetRemoteIpPrefix(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.RemoteIpPrefix = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetRemoteType(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.RemoteType = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetProtocol(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.Protocol = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetIsDeleted(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.IsDeleted = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetRemoteSecurityGroupUid(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.RemoteSecurityGroupUid = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetRuleType(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.RuleType = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetEtherType(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.EtherType = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetCreatedTime(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.CreatedTime = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetRuleUid(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.RuleUid = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetId(v int64) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.Id = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetSecurityGroupUid(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.SecurityGroupUid = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetDirection(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.Direction = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetMaxPortRange(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.MaxPortRange = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseRuleLists) SetStatus(v string) *GetSecurityGroupRuleInfoResponseRuleLists {
	s.Status = &v
	return s
}


type GetSecurityGroupRuleInfoResponseRuleListsBuilder struct {
	s *GetSecurityGroupRuleInfoResponseRuleLists
}

func NewGetSecurityGroupRuleInfoResponseRuleListsBuilder() *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	s := &GetSecurityGroupRuleInfoResponseRuleLists{}
	b := &GetSecurityGroupRuleInfoResponseRuleListsBuilder{s: s}
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) ModifiedTime(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) MinPortRange(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.MinPortRange = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) Remark(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.Remark = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) RemoteIpPrefix(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.RemoteIpPrefix = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) RemoteType(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.RemoteType = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) Protocol(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.Protocol = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) IsDeleted(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.IsDeleted = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) RemoteSecurityGroupUid(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.RemoteSecurityGroupUid = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) RuleType(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.RuleType = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) EtherType(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.EtherType = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) CreatedTime(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) RuleUid(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.RuleUid = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) Id(v int64) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.Id = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) SecurityGroupUid(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) Direction(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.Direction = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) MaxPortRange(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.MaxPortRange = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) Status(v string) *GetSecurityGroupRuleInfoResponseRuleListsBuilder {
	b.s.Status = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseRuleListsBuilder) Build() *GetSecurityGroupRuleInfoResponseRuleLists {
	return b.s
}


