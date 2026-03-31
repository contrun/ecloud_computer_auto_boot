// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageRecordSharedByOthersResponseData struct {

    // 厂商编码，例如:H3C
	CompanyCode *string `json:"companyCode,omitempty"`
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

func (s ImageRecordSharedByOthersResponseData) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedByOthersResponseData) GoString() string {
	return s.String()
}

func (s ImageRecordSharedByOthersResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedByOthersResponseData) SetCompanyCode(v string) *ImageRecordSharedByOthersResponseData {
	s.CompanyCode = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetImageName(v string) *ImageRecordSharedByOthersResponseData {
	s.ImageName = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetCompanyName(v string) *ImageRecordSharedByOthersResponseData {
	s.CompanyName = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetRegionName(v string) *ImageRecordSharedByOthersResponseData {
	s.RegionName = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetAcceptorTenantId(v string) *ImageRecordSharedByOthersResponseData {
	s.AcceptorTenantId = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetImageOs(v string) *ImageRecordSharedByOthersResponseData {
	s.ImageOs = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetInitiatorMopUserId(v string) *ImageRecordSharedByOthersResponseData {
	s.InitiatorMopUserId = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetTemplateId(v string) *ImageRecordSharedByOthersResponseData {
	s.TemplateId = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetShareCreatedTime(v string) *ImageRecordSharedByOthersResponseData {
	s.ShareCreatedTime = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetInitiatorMopUserName(v string) *ImageRecordSharedByOthersResponseData {
	s.InitiatorMopUserName = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetInitiatorTenantId(v string) *ImageRecordSharedByOthersResponseData {
	s.InitiatorTenantId = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetPoolCode(v string) *ImageRecordSharedByOthersResponseData {
	s.PoolCode = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetAcceptorMopUserId(v string) *ImageRecordSharedByOthersResponseData {
	s.AcceptorMopUserId = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetRegion(v string) *ImageRecordSharedByOthersResponseData {
	s.Region = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetSysDiskSize(v string) *ImageRecordSharedByOthersResponseData {
	s.SysDiskSize = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetRegionMopName(v string) *ImageRecordSharedByOthersResponseData {
	s.RegionMopName = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetAcceptorMopUserName(v string) *ImageRecordSharedByOthersResponseData {
	s.AcceptorMopUserName = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetImageTarget(v string) *ImageRecordSharedByOthersResponseData {
	s.ImageTarget = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetDc(v string) *ImageRecordSharedByOthersResponseData {
	s.Dc = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetPoolName(v string) *ImageRecordSharedByOthersResponseData {
	s.PoolName = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseData) SetShareStatus(v string) *ImageRecordSharedByOthersResponseData {
	s.ShareStatus = &v
	return s
}


type ImageRecordSharedByOthersResponseDataBuilder struct {
	s *ImageRecordSharedByOthersResponseData
}

func NewImageRecordSharedByOthersResponseDataBuilder() *ImageRecordSharedByOthersResponseDataBuilder {
	s := &ImageRecordSharedByOthersResponseData{}
	b := &ImageRecordSharedByOthersResponseDataBuilder{s: s}
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) CompanyCode(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) ImageName(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.ImageName = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) CompanyName(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) RegionName(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.RegionName = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) AcceptorTenantId(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.AcceptorTenantId = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) ImageOs(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.ImageOs = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) InitiatorMopUserId(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.InitiatorMopUserId = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) TemplateId(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) ShareCreatedTime(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.ShareCreatedTime = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) InitiatorMopUserName(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.InitiatorMopUserName = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) InitiatorTenantId(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.InitiatorTenantId = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) PoolCode(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) AcceptorMopUserId(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.AcceptorMopUserId = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) Region(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) SysDiskSize(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) RegionMopName(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.RegionMopName = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) AcceptorMopUserName(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.AcceptorMopUserName = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) ImageTarget(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.ImageTarget = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) Dc(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.Dc = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) PoolName(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) ShareStatus(v string) *ImageRecordSharedByOthersResponseDataBuilder {
	b.s.ShareStatus = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseDataBuilder) Build() *ImageRecordSharedByOthersResponseData {
	return b.s
}


