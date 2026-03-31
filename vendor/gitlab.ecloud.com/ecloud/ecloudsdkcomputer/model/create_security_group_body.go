// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSecurityGroupBody struct {
    position.Body
    // 安全组名称
	SecurityGroupName *string `json:"securityGroupName,omitempty"`

	SecurityGroupRuleList *[]CreateSecurityGroupRequestSecurityGroupRuleList `json:"securityGroupRuleList,omitempty"`
    // 安全组描述
	Description *string `json:"description,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 可用区编码, 例如:CIDC-RP-30-AZ1
	Region *string `json:"region,omitempty"`
}

func (s CreateSecurityGroupBody) String() string {
	return utils.Beautify(s)
}

func (s CreateSecurityGroupBody) GoString() string {
	return s.String()
}

func (s CreateSecurityGroupBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSecurityGroupBody) SetSecurityGroupName(v string) *CreateSecurityGroupBody {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateSecurityGroupBody) SetSecurityGroupRuleList(v []CreateSecurityGroupRequestSecurityGroupRuleList) *CreateSecurityGroupBody {
	s.SecurityGroupRuleList = &v
	return s
}

func (s *CreateSecurityGroupBody) SetDescription(v string) *CreateSecurityGroupBody {
	s.Description = &v
	return s
}

func (s *CreateSecurityGroupBody) SetPoolCode(v string) *CreateSecurityGroupBody {
	s.PoolCode = &v
	return s
}

func (s *CreateSecurityGroupBody) SetRegion(v string) *CreateSecurityGroupBody {
	s.Region = &v
	return s
}


type CreateSecurityGroupBodyBuilder struct {
	s *CreateSecurityGroupBody
}

func NewCreateSecurityGroupBodyBuilder() *CreateSecurityGroupBodyBuilder {
	s := &CreateSecurityGroupBody{}
	b := &CreateSecurityGroupBodyBuilder{s: s}
	return b
}

func (b *CreateSecurityGroupBodyBuilder) SecurityGroupName(v string) *CreateSecurityGroupBodyBuilder {
	b.s.SecurityGroupName = &v
	return b
}

func (b *CreateSecurityGroupBodyBuilder) SecurityGroupRuleList(v []CreateSecurityGroupRequestSecurityGroupRuleList) *CreateSecurityGroupBodyBuilder {
	b.s.SecurityGroupRuleList = &v
	return b
}

func (b *CreateSecurityGroupBodyBuilder) Description(v string) *CreateSecurityGroupBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateSecurityGroupBodyBuilder) PoolCode(v string) *CreateSecurityGroupBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *CreateSecurityGroupBodyBuilder) Region(v string) *CreateSecurityGroupBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *CreateSecurityGroupBodyBuilder) Build() *CreateSecurityGroupBody {
	return b.s
}


