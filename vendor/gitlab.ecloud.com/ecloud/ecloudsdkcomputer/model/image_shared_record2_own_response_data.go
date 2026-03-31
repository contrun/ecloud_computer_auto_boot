// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageSharedRecord2OwnResponseData struct {

    // 厂商编码，例如:H3C
	CompanyCode *string `json:"companyCode,omitempty"`
    // 数据盘大小，单位GB
	DataDiskSize *string `json:"dataDiskSize,omitempty"`
    // 镜像模板名称
	ImageName *string `json:"imageName,omitempty"`
    // 厂商名称，例如：新华三
	CompanyName *string `json:"companyName,omitempty"`
    // 可用区名称，例如：可用区一
	RegionName *string `json:"regionName,omitempty"`
    // 接受者租户主键ID
	AcceptorTenantId *string `json:"acceptorTenantId,omitempty"`
    // 镜像操作系统
	ImageOs *string `json:"imageOs,omitempty"`
    // 发起者租户ID
	InitiatorMopUserId *string `json:"initiatorMopUserId,omitempty"`
    // 镜像模板UID
	TemplateId *string `json:"templateId,omitempty"`
    // 共享时间
	ShareCreatedTime *string `json:"shareCreatedTime,omitempty"`
    // 发起者用户名
	InitiatorMopUserName *string `json:"initiatorMopUserName,omitempty"`
    // 发起者租户主键ID
	InitiatorTenantId *string `json:"initiatorTenantId,omitempty"`
    // 资源池编码,例如:CIDC-RP-35
	PoolCode *string `json:"poolCode,omitempty"`
    // 接受者租户ID
	AcceptorMopUserId *string `json:"acceptorMopUserId,omitempty"`
    // 可用区，例如：CIDC-RP-35-AZ1
	Region *string `json:"region,omitempty"`
    // 系统盘大小，单位GB
	SysDiskSize *string `json:"sysDiskSize,omitempty"`
    // 可用区名称，例如：华东-杭州(可用区一)
	RegionMopName *string `json:"regionMopName,omitempty"`
    // 接受者用户名
	AcceptorMopUserName *string `json:"acceptorMopUserName,omitempty"`
    // 镜像模板范围,系统及数据盘(all)、仅系统盘(system)
	ImageTarget *string `json:"imageTarget,omitempty"`
    // dc，例如：ONEDAAS-SERVICE-0002
	Dc *string `json:"dc,omitempty"`
    // 资源池名称,例如：华东-杭州
	PoolName *string `json:"poolName,omitempty"`
    // 共享状态
	ShareStatus *string `json:"shareStatus,omitempty"`
}

func (s ImageSharedRecord2OwnResponseData) String() string {
	return utils.Beautify(s)
}

func (s ImageSharedRecord2OwnResponseData) GoString() string {
	return s.String()
}

func (s ImageSharedRecord2OwnResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageSharedRecord2OwnResponseData) SetCompanyCode(v string) *ImageSharedRecord2OwnResponseData {
	s.CompanyCode = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetDataDiskSize(v string) *ImageSharedRecord2OwnResponseData {
	s.DataDiskSize = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetImageName(v string) *ImageSharedRecord2OwnResponseData {
	s.ImageName = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetCompanyName(v string) *ImageSharedRecord2OwnResponseData {
	s.CompanyName = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetRegionName(v string) *ImageSharedRecord2OwnResponseData {
	s.RegionName = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetAcceptorTenantId(v string) *ImageSharedRecord2OwnResponseData {
	s.AcceptorTenantId = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetImageOs(v string) *ImageSharedRecord2OwnResponseData {
	s.ImageOs = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetInitiatorMopUserId(v string) *ImageSharedRecord2OwnResponseData {
	s.InitiatorMopUserId = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetTemplateId(v string) *ImageSharedRecord2OwnResponseData {
	s.TemplateId = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetShareCreatedTime(v string) *ImageSharedRecord2OwnResponseData {
	s.ShareCreatedTime = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetInitiatorMopUserName(v string) *ImageSharedRecord2OwnResponseData {
	s.InitiatorMopUserName = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetInitiatorTenantId(v string) *ImageSharedRecord2OwnResponseData {
	s.InitiatorTenantId = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetPoolCode(v string) *ImageSharedRecord2OwnResponseData {
	s.PoolCode = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetAcceptorMopUserId(v string) *ImageSharedRecord2OwnResponseData {
	s.AcceptorMopUserId = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetRegion(v string) *ImageSharedRecord2OwnResponseData {
	s.Region = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetSysDiskSize(v string) *ImageSharedRecord2OwnResponseData {
	s.SysDiskSize = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetRegionMopName(v string) *ImageSharedRecord2OwnResponseData {
	s.RegionMopName = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetAcceptorMopUserName(v string) *ImageSharedRecord2OwnResponseData {
	s.AcceptorMopUserName = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetImageTarget(v string) *ImageSharedRecord2OwnResponseData {
	s.ImageTarget = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetDc(v string) *ImageSharedRecord2OwnResponseData {
	s.Dc = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetPoolName(v string) *ImageSharedRecord2OwnResponseData {
	s.PoolName = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseData) SetShareStatus(v string) *ImageSharedRecord2OwnResponseData {
	s.ShareStatus = &v
	return s
}


type ImageSharedRecord2OwnResponseDataBuilder struct {
	s *ImageSharedRecord2OwnResponseData
}

func NewImageSharedRecord2OwnResponseDataBuilder() *ImageSharedRecord2OwnResponseDataBuilder {
	s := &ImageSharedRecord2OwnResponseData{}
	b := &ImageSharedRecord2OwnResponseDataBuilder{s: s}
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) CompanyCode(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) DataDiskSize(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.DataDiskSize = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) ImageName(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.ImageName = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) CompanyName(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) RegionName(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.RegionName = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) AcceptorTenantId(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.AcceptorTenantId = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) ImageOs(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.ImageOs = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) InitiatorMopUserId(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.InitiatorMopUserId = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) TemplateId(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) ShareCreatedTime(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.ShareCreatedTime = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) InitiatorMopUserName(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.InitiatorMopUserName = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) InitiatorTenantId(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.InitiatorTenantId = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) PoolCode(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) AcceptorMopUserId(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.AcceptorMopUserId = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) Region(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) SysDiskSize(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) RegionMopName(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.RegionMopName = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) AcceptorMopUserName(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.AcceptorMopUserName = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) ImageTarget(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.ImageTarget = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) Dc(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.Dc = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) PoolName(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) ShareStatus(v string) *ImageSharedRecord2OwnResponseDataBuilder {
	b.s.ShareStatus = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseDataBuilder) Build() *ImageSharedRecord2OwnResponseData {
	return b.s
}


