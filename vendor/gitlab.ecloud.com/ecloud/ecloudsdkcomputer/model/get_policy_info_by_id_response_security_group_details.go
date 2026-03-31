// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdResponseSecurityGroupDetails struct {

    // 厂商编码
	CompanyCode *string `json:"companyCode,omitempty"`
    // 安全组名称
	SecurityGroupName *string `json:"securityGroupName,omitempty"`
    // 是否含有绑定中机器,含有则无法动态删减. true: 安全组含有绑定中机器, false: 安全组不含有绑定中机器
	HaveBindingMachine *string `json:"haveBindingMachine,omitempty"`
    // 状态类型（有状态和无状态）
	StateType *string `json:"stateType,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 安全组Uid
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
    // 安全组状态，creating创建中，create创建完成待生效，effective已生效，createError创建失败
	Status *string `json:"status,omitempty"`
}

func (s GetPolicyInfoByIdResponseSecurityGroupDetails) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseSecurityGroupDetails) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseSecurityGroupDetails) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetails) SetCompanyCode(v string) *GetPolicyInfoByIdResponseSecurityGroupDetails {
	s.CompanyCode = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetails) SetSecurityGroupName(v string) *GetPolicyInfoByIdResponseSecurityGroupDetails {
	s.SecurityGroupName = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetails) SetHaveBindingMachine(v string) *GetPolicyInfoByIdResponseSecurityGroupDetails {
	s.HaveBindingMachine = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetails) SetStateType(v string) *GetPolicyInfoByIdResponseSecurityGroupDetails {
	s.StateType = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetails) SetDescription(v string) *GetPolicyInfoByIdResponseSecurityGroupDetails {
	s.Description = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetails) SetPoolCode(v string) *GetPolicyInfoByIdResponseSecurityGroupDetails {
	s.PoolCode = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetails) SetRegion(v string) *GetPolicyInfoByIdResponseSecurityGroupDetails {
	s.Region = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetails) SetSecurityGroupUid(v string) *GetPolicyInfoByIdResponseSecurityGroupDetails {
	s.SecurityGroupUid = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetails) SetStatus(v string) *GetPolicyInfoByIdResponseSecurityGroupDetails {
	s.Status = &v
	return s
}


type GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder struct {
	s *GetPolicyInfoByIdResponseSecurityGroupDetails
}

func NewGetPolicyInfoByIdResponseSecurityGroupDetailsBuilder() *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	s := &GetPolicyInfoByIdResponseSecurityGroupDetails{}
	b := &GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) CompanyCode(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) SecurityGroupName(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	b.s.SecurityGroupName = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) HaveBindingMachine(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	b.s.HaveBindingMachine = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) StateType(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	b.s.StateType = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) Description(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	b.s.Description = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) PoolCode(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) Region(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	b.s.Region = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) SecurityGroupUid(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) Status(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder {
	b.s.Status = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailsBuilder) Build() *GetPolicyInfoByIdResponseSecurityGroupDetails {
	return b.s
}


