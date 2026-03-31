// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageRecordSharedWithOthersResponseBody struct {

    // 厂商编码，例如:H3C
	CompanyCode *string `json:"companyCode,omitempty"`
    // 镜像模板主键ID
	ImageId *string `json:"imageId,omitempty"`
    // 镜像模板名称
	ImageName *string `json:"imageName,omitempty"`
    // 镜像模板描述
	ImageRemark *string `json:"imageRemark,omitempty"`
    // 厂商名称，例如：新华三
	CompanyName *string `json:"companyName,omitempty"`
    // 可用区名称，例如：可用区一
	RegionName *string `json:"regionName,omitempty"`
    // 镜像操作系统
	ImageOs *string `json:"imageOs,omitempty"`
    // 共享记录
	ImageSharedRecords *ImageRecordSharedWithOthersResponseImageSharedRecords `json:"imageSharedRecords,omitempty"`
    // 镜像模板UID
	TemplateId *string `json:"templateId,omitempty"`
    // 资源池编码,例如:CIDC-RP-35
	PoolCode *string `json:"poolCode,omitempty"`
    // 可用区，例如：CIDC-RP-35-AZ1
	Region *string `json:"region,omitempty"`
    // 系统盘大小，单位GB
	SysDiskSize *string `json:"sysDiskSize,omitempty"`
    // 可用区名称，例如：华东-杭州(可用区一)
	RegionMopName *string `json:"regionMopName,omitempty"`
    // 镜像模板范围,系统及数据盘(all)、仅系统盘(system)
	ImageTarget *string `json:"imageTarget,omitempty"`
    // dc，例如：ONEDAAS-SERVICE-0002
	Dc *string `json:"dc,omitempty"`
    // 资源池名称,例如：华东-杭州
	PoolName *string `json:"poolName,omitempty"`
}

func (s ImageRecordSharedWithOthersResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedWithOthersResponseBody) GoString() string {
	return s.String()
}

func (s ImageRecordSharedWithOthersResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedWithOthersResponseBody) SetCompanyCode(v string) *ImageRecordSharedWithOthersResponseBody {
	s.CompanyCode = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetImageId(v string) *ImageRecordSharedWithOthersResponseBody {
	s.ImageId = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetImageName(v string) *ImageRecordSharedWithOthersResponseBody {
	s.ImageName = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetImageRemark(v string) *ImageRecordSharedWithOthersResponseBody {
	s.ImageRemark = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetCompanyName(v string) *ImageRecordSharedWithOthersResponseBody {
	s.CompanyName = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetRegionName(v string) *ImageRecordSharedWithOthersResponseBody {
	s.RegionName = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetImageOs(v string) *ImageRecordSharedWithOthersResponseBody {
	s.ImageOs = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetImageSharedRecords(v *ImageRecordSharedWithOthersResponseImageSharedRecords) *ImageRecordSharedWithOthersResponseBody {
	s.ImageSharedRecords = v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetTemplateId(v string) *ImageRecordSharedWithOthersResponseBody {
	s.TemplateId = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetPoolCode(v string) *ImageRecordSharedWithOthersResponseBody {
	s.PoolCode = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetRegion(v string) *ImageRecordSharedWithOthersResponseBody {
	s.Region = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetSysDiskSize(v string) *ImageRecordSharedWithOthersResponseBody {
	s.SysDiskSize = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetRegionMopName(v string) *ImageRecordSharedWithOthersResponseBody {
	s.RegionMopName = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetImageTarget(v string) *ImageRecordSharedWithOthersResponseBody {
	s.ImageTarget = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetDc(v string) *ImageRecordSharedWithOthersResponseBody {
	s.Dc = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseBody) SetPoolName(v string) *ImageRecordSharedWithOthersResponseBody {
	s.PoolName = &v
	return s
}


type ImageRecordSharedWithOthersResponseBodyBuilder struct {
	s *ImageRecordSharedWithOthersResponseBody
}

func NewImageRecordSharedWithOthersResponseBodyBuilder() *ImageRecordSharedWithOthersResponseBodyBuilder {
	s := &ImageRecordSharedWithOthersResponseBody{}
	b := &ImageRecordSharedWithOthersResponseBodyBuilder{s: s}
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) CompanyCode(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) ImageId(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.ImageId = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) ImageName(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.ImageName = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) ImageRemark(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.ImageRemark = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) CompanyName(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) RegionName(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.RegionName = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) ImageOs(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.ImageOs = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) ImageSharedRecords(v *ImageRecordSharedWithOthersResponseImageSharedRecords) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.ImageSharedRecords = v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) TemplateId(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) PoolCode(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) Region(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) SysDiskSize(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) RegionMopName(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.RegionMopName = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) ImageTarget(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.ImageTarget = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) Dc(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) PoolName(v string) *ImageRecordSharedWithOthersResponseBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseBodyBuilder) Build() *ImageRecordSharedWithOthersResponseBody {
	return b.s
}


