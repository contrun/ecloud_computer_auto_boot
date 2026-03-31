// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPublicImagePageResponseData struct {

    // 厂商编码
	CompanyCode *string `json:"companyCode,omitempty"`
    // 公共镜像版本
	ImageVersion *string `json:"imageVersion,omitempty"`
    // 公共镜像名称
	ImageName *string `json:"imageName,omitempty"`
    // 厂商名称
	CompanyName *string `json:"companyName,omitempty"`
    // 可用区名称
	RegionName *string `json:"regionName,omitempty"`
    // 公共镜像操作系统
	ImageOs *string `json:"imageOs,omitempty"`
    // 镜像状态，0可用，1不可用
	ImageStatus *int32 `json:"imageStatus,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 公共镜像系统盘大小限制，单位GB
	SysDiskSize *string `json:"sysDiskSize,omitempty"`
    // 公共镜像uid
	PubImageUid *string `json:"pubImageUid,omitempty"`
    // 镜像模板范围,公共镜像只有仅系统盘(system)
	ImageTarget *string `json:"imageTarget,omitempty"`
    // 可用区名称
	Dc *string `json:"dc,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s GetPublicImagePageResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetPublicImagePageResponseData) GoString() string {
	return s.String()
}

func (s GetPublicImagePageResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPublicImagePageResponseData) SetCompanyCode(v string) *GetPublicImagePageResponseData {
	s.CompanyCode = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetImageVersion(v string) *GetPublicImagePageResponseData {
	s.ImageVersion = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetImageName(v string) *GetPublicImagePageResponseData {
	s.ImageName = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetCompanyName(v string) *GetPublicImagePageResponseData {
	s.CompanyName = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetRegionName(v string) *GetPublicImagePageResponseData {
	s.RegionName = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetImageOs(v string) *GetPublicImagePageResponseData {
	s.ImageOs = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetImageStatus(v int32) *GetPublicImagePageResponseData {
	s.ImageStatus = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetCreatedTime(v string) *GetPublicImagePageResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetPoolCode(v string) *GetPublicImagePageResponseData {
	s.PoolCode = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetRegion(v string) *GetPublicImagePageResponseData {
	s.Region = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetSysDiskSize(v string) *GetPublicImagePageResponseData {
	s.SysDiskSize = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetPubImageUid(v string) *GetPublicImagePageResponseData {
	s.PubImageUid = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetImageTarget(v string) *GetPublicImagePageResponseData {
	s.ImageTarget = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetDc(v string) *GetPublicImagePageResponseData {
	s.Dc = &v
	return s
}

func (s *GetPublicImagePageResponseData) SetPoolName(v string) *GetPublicImagePageResponseData {
	s.PoolName = &v
	return s
}


type GetPublicImagePageResponseDataBuilder struct {
	s *GetPublicImagePageResponseData
}

func NewGetPublicImagePageResponseDataBuilder() *GetPublicImagePageResponseDataBuilder {
	s := &GetPublicImagePageResponseData{}
	b := &GetPublicImagePageResponseDataBuilder{s: s}
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) CompanyCode(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) ImageVersion(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.ImageVersion = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) ImageName(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.ImageName = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) CompanyName(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) RegionName(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.RegionName = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) ImageOs(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.ImageOs = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) ImageStatus(v int32) *GetPublicImagePageResponseDataBuilder {
	b.s.ImageStatus = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) CreatedTime(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) PoolCode(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) Region(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) SysDiskSize(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) PubImageUid(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.PubImageUid = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) ImageTarget(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.ImageTarget = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) Dc(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.Dc = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) PoolName(v string) *GetPublicImagePageResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetPublicImagePageResponseDataBuilder) Build() *GetPublicImagePageResponseData {
	return b.s
}


