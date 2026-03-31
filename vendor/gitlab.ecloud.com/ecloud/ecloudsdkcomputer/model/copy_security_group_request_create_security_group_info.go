// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CopySecurityGroupRequestCreateSecurityGroupInfo struct {

    // 安全组名称
	SecurityGroupName *string `json:"securityGroupName,omitempty"`

	SecurityGroupRuleList *[]CopySecurityGroupRequestSecurityGroupRuleList `json:"securityGroupRuleList,omitempty"`
    // 安全组描述
	Description *string `json:"description,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 可用区编码, 例如:CIDC-RP-30-AZ1
	Region *string `json:"region,omitempty"`
}

func (s CopySecurityGroupRequestCreateSecurityGroupInfo) String() string {
	return utils.Beautify(s)
}

func (s CopySecurityGroupRequestCreateSecurityGroupInfo) GoString() string {
	return s.String()
}

func (s CopySecurityGroupRequestCreateSecurityGroupInfo) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CopySecurityGroupRequestCreateSecurityGroupInfo) SetSecurityGroupName(v string) *CopySecurityGroupRequestCreateSecurityGroupInfo {
	s.SecurityGroupName = &v
	return s
}

func (s *CopySecurityGroupRequestCreateSecurityGroupInfo) SetSecurityGroupRuleList(v []CopySecurityGroupRequestSecurityGroupRuleList) *CopySecurityGroupRequestCreateSecurityGroupInfo {
	s.SecurityGroupRuleList = &v
	return s
}

func (s *CopySecurityGroupRequestCreateSecurityGroupInfo) SetDescription(v string) *CopySecurityGroupRequestCreateSecurityGroupInfo {
	s.Description = &v
	return s
}

func (s *CopySecurityGroupRequestCreateSecurityGroupInfo) SetPoolCode(v string) *CopySecurityGroupRequestCreateSecurityGroupInfo {
	s.PoolCode = &v
	return s
}

func (s *CopySecurityGroupRequestCreateSecurityGroupInfo) SetRegion(v string) *CopySecurityGroupRequestCreateSecurityGroupInfo {
	s.Region = &v
	return s
}


type CopySecurityGroupRequestCreateSecurityGroupInfoBuilder struct {
	s *CopySecurityGroupRequestCreateSecurityGroupInfo
}

func NewCopySecurityGroupRequestCreateSecurityGroupInfoBuilder() *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder {
	s := &CopySecurityGroupRequestCreateSecurityGroupInfo{}
	b := &CopySecurityGroupRequestCreateSecurityGroupInfoBuilder{s: s}
	return b
}

func (b *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder) SecurityGroupName(v string) *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder {
	b.s.SecurityGroupName = &v
	return b
}

func (b *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder) SecurityGroupRuleList(v []CopySecurityGroupRequestSecurityGroupRuleList) *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder {
	b.s.SecurityGroupRuleList = &v
	return b
}

func (b *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder) Description(v string) *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder {
	b.s.Description = &v
	return b
}

func (b *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder) PoolCode(v string) *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder) Region(v string) *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder {
	b.s.Region = &v
	return b
}

func (b *CopySecurityGroupRequestCreateSecurityGroupInfoBuilder) Build() *CopySecurityGroupRequestCreateSecurityGroupInfo {
	return b.s
}


