// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVpcResponseData struct {

    // 厂商
	CompanyCode *string `json:"companyCode,omitempty"`

	EcloudRegionCode *string `json:"ecloudRegionCode,omitempty"`
    // vpc名称
	VpcName *string `json:"vpcName,omitempty"`
    // 企业id
	CompanyTag *string `json:"companyTag,omitempty"`
    // 厂商名称
	CompanyName *string `json:"companyName,omitempty"`
    // vpc描述
	Description *string `json:"description,omitempty"`
    // 关联nat网关名称
	NatGatewayName *string `json:"natGatewayName,omitempty"`
    // vpc来源
	RequestSource *string `json:"requestSource,omitempty"`
    // NAT网关uid
	NatUid *string `json:"natUid,omitempty"`
    // vpc规格
	Specs *string `json:"specs,omitempty"`

	EcloudPoolName *string `json:"ecloudPoolName,omitempty"`
    // vpcid
	VpcUid *string `json:"vpcUid,omitempty"`
    // 租户名称
	TenantName *string `json:"tenantName,omitempty"`
    // 子网个数
	SubnetSize *int32 `json:"subnetSize,omitempty"`

	EcloudRegionName *string `json:"ecloudRegionName,omitempty"`

	EcloudPoolCode *string `json:"ecloudPoolCode,omitempty"`
    // 租户主键id
	TenantId *string `json:"tenantId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // CMDB可用区
	PoolCode *string `json:"poolCode,omitempty"`
    // 租户id
	MopUserId *string `json:"mopUserId,omitempty"`
    // CMDB可用区名称
	PoolName *string `json:"poolName,omitempty"`
    // vpc状态
	Status *string `json:"status,omitempty"`
}

func (s GetVpcResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetVpcResponseData) GoString() string {
	return s.String()
}

func (s GetVpcResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcResponseData) SetCompanyCode(v string) *GetVpcResponseData {
	s.CompanyCode = &v
	return s
}

func (s *GetVpcResponseData) SetEcloudRegionCode(v string) *GetVpcResponseData {
	s.EcloudRegionCode = &v
	return s
}

func (s *GetVpcResponseData) SetVpcName(v string) *GetVpcResponseData {
	s.VpcName = &v
	return s
}

func (s *GetVpcResponseData) SetCompanyTag(v string) *GetVpcResponseData {
	s.CompanyTag = &v
	return s
}

func (s *GetVpcResponseData) SetCompanyName(v string) *GetVpcResponseData {
	s.CompanyName = &v
	return s
}

func (s *GetVpcResponseData) SetDescription(v string) *GetVpcResponseData {
	s.Description = &v
	return s
}

func (s *GetVpcResponseData) SetNatGatewayName(v string) *GetVpcResponseData {
	s.NatGatewayName = &v
	return s
}

func (s *GetVpcResponseData) SetRequestSource(v string) *GetVpcResponseData {
	s.RequestSource = &v
	return s
}

func (s *GetVpcResponseData) SetNatUid(v string) *GetVpcResponseData {
	s.NatUid = &v
	return s
}

func (s *GetVpcResponseData) SetSpecs(v string) *GetVpcResponseData {
	s.Specs = &v
	return s
}

func (s *GetVpcResponseData) SetEcloudPoolName(v string) *GetVpcResponseData {
	s.EcloudPoolName = &v
	return s
}

func (s *GetVpcResponseData) SetVpcUid(v string) *GetVpcResponseData {
	s.VpcUid = &v
	return s
}

func (s *GetVpcResponseData) SetTenantName(v string) *GetVpcResponseData {
	s.TenantName = &v
	return s
}

func (s *GetVpcResponseData) SetSubnetSize(v int32) *GetVpcResponseData {
	s.SubnetSize = &v
	return s
}

func (s *GetVpcResponseData) SetEcloudRegionName(v string) *GetVpcResponseData {
	s.EcloudRegionName = &v
	return s
}

func (s *GetVpcResponseData) SetEcloudPoolCode(v string) *GetVpcResponseData {
	s.EcloudPoolCode = &v
	return s
}

func (s *GetVpcResponseData) SetTenantId(v string) *GetVpcResponseData {
	s.TenantId = &v
	return s
}

func (s *GetVpcResponseData) SetCreatedTime(v string) *GetVpcResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetVpcResponseData) SetPoolCode(v string) *GetVpcResponseData {
	s.PoolCode = &v
	return s
}

func (s *GetVpcResponseData) SetMopUserId(v string) *GetVpcResponseData {
	s.MopUserId = &v
	return s
}

func (s *GetVpcResponseData) SetPoolName(v string) *GetVpcResponseData {
	s.PoolName = &v
	return s
}

func (s *GetVpcResponseData) SetStatus(v string) *GetVpcResponseData {
	s.Status = &v
	return s
}


type GetVpcResponseDataBuilder struct {
	s *GetVpcResponseData
}

func NewGetVpcResponseDataBuilder() *GetVpcResponseDataBuilder {
	s := &GetVpcResponseData{}
	b := &GetVpcResponseDataBuilder{s: s}
	return b
}

func (b *GetVpcResponseDataBuilder) CompanyCode(v string) *GetVpcResponseDataBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *GetVpcResponseDataBuilder) EcloudRegionCode(v string) *GetVpcResponseDataBuilder {
	b.s.EcloudRegionCode = &v
	return b
}

func (b *GetVpcResponseDataBuilder) VpcName(v string) *GetVpcResponseDataBuilder {
	b.s.VpcName = &v
	return b
}

func (b *GetVpcResponseDataBuilder) CompanyTag(v string) *GetVpcResponseDataBuilder {
	b.s.CompanyTag = &v
	return b
}

func (b *GetVpcResponseDataBuilder) CompanyName(v string) *GetVpcResponseDataBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *GetVpcResponseDataBuilder) Description(v string) *GetVpcResponseDataBuilder {
	b.s.Description = &v
	return b
}

func (b *GetVpcResponseDataBuilder) NatGatewayName(v string) *GetVpcResponseDataBuilder {
	b.s.NatGatewayName = &v
	return b
}

func (b *GetVpcResponseDataBuilder) RequestSource(v string) *GetVpcResponseDataBuilder {
	b.s.RequestSource = &v
	return b
}

func (b *GetVpcResponseDataBuilder) NatUid(v string) *GetVpcResponseDataBuilder {
	b.s.NatUid = &v
	return b
}

func (b *GetVpcResponseDataBuilder) Specs(v string) *GetVpcResponseDataBuilder {
	b.s.Specs = &v
	return b
}

func (b *GetVpcResponseDataBuilder) EcloudPoolName(v string) *GetVpcResponseDataBuilder {
	b.s.EcloudPoolName = &v
	return b
}

func (b *GetVpcResponseDataBuilder) VpcUid(v string) *GetVpcResponseDataBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *GetVpcResponseDataBuilder) TenantName(v string) *GetVpcResponseDataBuilder {
	b.s.TenantName = &v
	return b
}

func (b *GetVpcResponseDataBuilder) SubnetSize(v int32) *GetVpcResponseDataBuilder {
	b.s.SubnetSize = &v
	return b
}

func (b *GetVpcResponseDataBuilder) EcloudRegionName(v string) *GetVpcResponseDataBuilder {
	b.s.EcloudRegionName = &v
	return b
}

func (b *GetVpcResponseDataBuilder) EcloudPoolCode(v string) *GetVpcResponseDataBuilder {
	b.s.EcloudPoolCode = &v
	return b
}

func (b *GetVpcResponseDataBuilder) TenantId(v string) *GetVpcResponseDataBuilder {
	b.s.TenantId = &v
	return b
}

func (b *GetVpcResponseDataBuilder) CreatedTime(v string) *GetVpcResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetVpcResponseDataBuilder) PoolCode(v string) *GetVpcResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetVpcResponseDataBuilder) MopUserId(v string) *GetVpcResponseDataBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *GetVpcResponseDataBuilder) PoolName(v string) *GetVpcResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetVpcResponseDataBuilder) Status(v string) *GetVpcResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *GetVpcResponseDataBuilder) Build() *GetVpcResponseData {
	return b.s
}


