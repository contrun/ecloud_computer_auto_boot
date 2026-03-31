// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdResponseSecurityGroupDetail struct {

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

func (s GetPolicyInfoByIdResponseSecurityGroupDetail) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseSecurityGroupDetail) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseSecurityGroupDetail) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetail) SetCompanyCode(v string) *GetPolicyInfoByIdResponseSecurityGroupDetail {
	s.CompanyCode = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetail) SetSecurityGroupName(v string) *GetPolicyInfoByIdResponseSecurityGroupDetail {
	s.SecurityGroupName = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetail) SetHaveBindingMachine(v string) *GetPolicyInfoByIdResponseSecurityGroupDetail {
	s.HaveBindingMachine = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetail) SetStateType(v string) *GetPolicyInfoByIdResponseSecurityGroupDetail {
	s.StateType = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetail) SetDescription(v string) *GetPolicyInfoByIdResponseSecurityGroupDetail {
	s.Description = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetail) SetPoolCode(v string) *GetPolicyInfoByIdResponseSecurityGroupDetail {
	s.PoolCode = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetail) SetRegion(v string) *GetPolicyInfoByIdResponseSecurityGroupDetail {
	s.Region = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetail) SetSecurityGroupUid(v string) *GetPolicyInfoByIdResponseSecurityGroupDetail {
	s.SecurityGroupUid = &v
	return s
}

func (s *GetPolicyInfoByIdResponseSecurityGroupDetail) SetStatus(v string) *GetPolicyInfoByIdResponseSecurityGroupDetail {
	s.Status = &v
	return s
}


type GetPolicyInfoByIdResponseSecurityGroupDetailBuilder struct {
	s *GetPolicyInfoByIdResponseSecurityGroupDetail
}

func NewGetPolicyInfoByIdResponseSecurityGroupDetailBuilder() *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	s := &GetPolicyInfoByIdResponseSecurityGroupDetail{}
	b := &GetPolicyInfoByIdResponseSecurityGroupDetailBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) CompanyCode(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) SecurityGroupName(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	b.s.SecurityGroupName = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) HaveBindingMachine(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	b.s.HaveBindingMachine = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) StateType(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	b.s.StateType = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) Description(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	b.s.Description = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) PoolCode(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) Region(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	b.s.Region = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) SecurityGroupUid(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) Status(v string) *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder {
	b.s.Status = &v
	return b
}

func (b *GetPolicyInfoByIdResponseSecurityGroupDetailBuilder) Build() *GetPolicyInfoByIdResponseSecurityGroupDetail {
	return b.s
}


