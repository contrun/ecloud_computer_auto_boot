// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SecurityGroupManagementListResponseData struct {

    // 公司编码, 例如:H3C
	CompanyCode *string `json:"companyCode,omitempty"`
    // 安全组名
	SecurityGroupName *string `json:"securityGroupName,omitempty"`
    // 公司编码, 例如:新华三
	CompanyName *string `json:"companyName,omitempty"`
    // 请求来源
	RequestSource *string `json:"requestSource,omitempty"`
    // 安全组状态
	SecurityGroupStatus *string `json:"securityGroupStatus,omitempty"`
    // 可用区名称, 例如:苏州资源池
	RegionAzName *string `json:"regionAzName,omitempty"`
    // 安全组描述
	SecurityGroupDescription *string `json:"securityGroupDescription,omitempty"`
    // 安全组状态类型(有状态 stateful、无状态 stateless)
	SecurityGroupStateType *string `json:"securityGroupStateType,omitempty"`
    // 租户表主键ID
	TenantId *string `json:"tenantId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 租户名称
	MopUserName *string `json:"mopUserName,omitempty"`
    // 资源池编码, 例如:CIDC-RP-30
	PoolCode *string `json:"poolCode,omitempty"`
    // 企业ID
	EnterpriseID *string `json:"enterpriseID,omitempty"`
    // 租户ID
	MopUserId *string `json:"mopUserId,omitempty"`
    // 可用区编码, 例如:CIDC-RP-30-AZ1
	Region *string `json:"region,omitempty"`
    // 安全组UID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
    // dc, 例如:H3C-region-0001
	Dc *string `json:"dc,omitempty"`
    // 资源池名称, 例如:华东-苏州
	PoolName *string `json:"poolName,omitempty"`
    // 安全组绑定机器数
	SgBindMachineCount *string `json:"sgBindMachineCount,omitempty"`
}

func (s SecurityGroupManagementListResponseData) String() string {
	return utils.Beautify(s)
}

func (s SecurityGroupManagementListResponseData) GoString() string {
	return s.String()
}

func (s SecurityGroupManagementListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SecurityGroupManagementListResponseData) SetCompanyCode(v string) *SecurityGroupManagementListResponseData {
	s.CompanyCode = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetSecurityGroupName(v string) *SecurityGroupManagementListResponseData {
	s.SecurityGroupName = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetCompanyName(v string) *SecurityGroupManagementListResponseData {
	s.CompanyName = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetRequestSource(v string) *SecurityGroupManagementListResponseData {
	s.RequestSource = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetSecurityGroupStatus(v string) *SecurityGroupManagementListResponseData {
	s.SecurityGroupStatus = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetRegionAzName(v string) *SecurityGroupManagementListResponseData {
	s.RegionAzName = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetSecurityGroupDescription(v string) *SecurityGroupManagementListResponseData {
	s.SecurityGroupDescription = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetSecurityGroupStateType(v string) *SecurityGroupManagementListResponseData {
	s.SecurityGroupStateType = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetTenantId(v string) *SecurityGroupManagementListResponseData {
	s.TenantId = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetCreatedTime(v string) *SecurityGroupManagementListResponseData {
	s.CreatedTime = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetMopUserName(v string) *SecurityGroupManagementListResponseData {
	s.MopUserName = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetPoolCode(v string) *SecurityGroupManagementListResponseData {
	s.PoolCode = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetEnterpriseID(v string) *SecurityGroupManagementListResponseData {
	s.EnterpriseID = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetMopUserId(v string) *SecurityGroupManagementListResponseData {
	s.MopUserId = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetRegion(v string) *SecurityGroupManagementListResponseData {
	s.Region = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetSecurityGroupUid(v string) *SecurityGroupManagementListResponseData {
	s.SecurityGroupUid = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetDc(v string) *SecurityGroupManagementListResponseData {
	s.Dc = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetPoolName(v string) *SecurityGroupManagementListResponseData {
	s.PoolName = &v
	return s
}

func (s *SecurityGroupManagementListResponseData) SetSgBindMachineCount(v string) *SecurityGroupManagementListResponseData {
	s.SgBindMachineCount = &v
	return s
}


type SecurityGroupManagementListResponseDataBuilder struct {
	s *SecurityGroupManagementListResponseData
}

func NewSecurityGroupManagementListResponseDataBuilder() *SecurityGroupManagementListResponseDataBuilder {
	s := &SecurityGroupManagementListResponseData{}
	b := &SecurityGroupManagementListResponseDataBuilder{s: s}
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) CompanyCode(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) SecurityGroupName(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.SecurityGroupName = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) CompanyName(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) RequestSource(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.RequestSource = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) SecurityGroupStatus(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.SecurityGroupStatus = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) RegionAzName(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.RegionAzName = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) SecurityGroupDescription(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.SecurityGroupDescription = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) SecurityGroupStateType(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.SecurityGroupStateType = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) TenantId(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.TenantId = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) CreatedTime(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) MopUserName(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) PoolCode(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) EnterpriseID(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.EnterpriseID = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) MopUserId(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) Region(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) SecurityGroupUid(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) Dc(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.Dc = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) PoolName(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) SgBindMachineCount(v string) *SecurityGroupManagementListResponseDataBuilder {
	b.s.SgBindMachineCount = &v
	return b
}

func (b *SecurityGroupManagementListResponseDataBuilder) Build() *SecurityGroupManagementListResponseData {
	return b.s
}


