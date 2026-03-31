// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListImageResponseData struct {

    // 厂商编码，例如:H3C
	CompanyCode *string `json:"companyCode,omitempty"`
    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
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
    // 底层返回的错误信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 底层返回的错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 镜像操作系统，当前版本仅支持(windows)
	ImageOs *string `json:"imageOs,omitempty"`
    // 镜像状态，creating:创建中,create:创建完成,createError:创建失败,deleting:删除中,deleted:已删除,deletedError:删除失败
	ImageStatus *string `json:"imageStatus,omitempty"`
    // 镜像模板UID
	TemplateId *string `json:"templateId,omitempty"`
    // 和底层交互的requestId
	RequestId *string `json:"requestId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 资源池编码,例如:CIDC-RP-35
	PoolCode *string `json:"poolCode,omitempty"`
    // 磁盘性质，仅系统盘(system)
	DiskType *string `json:"diskType,omitempty"`
    // 可用区，例如：CIDC-RP-35-AZ1
	Region *string `json:"region,omitempty"`
    // 系统盘大小，单位GB
	SysDiskSize *string `json:"sysDiskSize,omitempty"`
    // 镜像类型，自定义镜像:custom,共享镜像:share
	ImageType *string `json:"imageType,omitempty"`
    // 镜像模板范围,系统及数据盘(all)、仅系统盘(system)
	ImageTarget *string `json:"imageTarget,omitempty"`
    // dc，例如：ONEDAAS-SERVICE-0002
	Dc *string `json:"dc,omitempty"`
    // 资源池名称,例如：华东-杭州
	PoolName *string `json:"poolName,omitempty"`
    // 底层最近一次交互返回的信息
	ResponseMsg *string `json:"responseMsg,omitempty"`
}

func (s ListImageResponseData) String() string {
	return utils.Beautify(s)
}

func (s ListImageResponseData) GoString() string {
	return s.String()
}

func (s ListImageResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListImageResponseData) SetCompanyCode(v string) *ListImageResponseData {
	s.CompanyCode = &v
	return s
}

func (s *ListImageResponseData) SetModifiedTime(v string) *ListImageResponseData {
	s.ModifiedTime = &v
	return s
}

func (s *ListImageResponseData) SetImageId(v string) *ListImageResponseData {
	s.ImageId = &v
	return s
}

func (s *ListImageResponseData) SetImageName(v string) *ListImageResponseData {
	s.ImageName = &v
	return s
}

func (s *ListImageResponseData) SetImageRemark(v string) *ListImageResponseData {
	s.ImageRemark = &v
	return s
}

func (s *ListImageResponseData) SetCompanyName(v string) *ListImageResponseData {
	s.CompanyName = &v
	return s
}

func (s *ListImageResponseData) SetRegionName(v string) *ListImageResponseData {
	s.RegionName = &v
	return s
}

func (s *ListImageResponseData) SetErrorMessage(v string) *ListImageResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *ListImageResponseData) SetErrorCode(v string) *ListImageResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ListImageResponseData) SetImageOs(v string) *ListImageResponseData {
	s.ImageOs = &v
	return s
}

func (s *ListImageResponseData) SetImageStatus(v string) *ListImageResponseData {
	s.ImageStatus = &v
	return s
}

func (s *ListImageResponseData) SetTemplateId(v string) *ListImageResponseData {
	s.TemplateId = &v
	return s
}

func (s *ListImageResponseData) SetRequestId(v string) *ListImageResponseData {
	s.RequestId = &v
	return s
}

func (s *ListImageResponseData) SetCreatedTime(v string) *ListImageResponseData {
	s.CreatedTime = &v
	return s
}

func (s *ListImageResponseData) SetPoolCode(v string) *ListImageResponseData {
	s.PoolCode = &v
	return s
}

func (s *ListImageResponseData) SetDiskType(v string) *ListImageResponseData {
	s.DiskType = &v
	return s
}

func (s *ListImageResponseData) SetRegion(v string) *ListImageResponseData {
	s.Region = &v
	return s
}

func (s *ListImageResponseData) SetSysDiskSize(v string) *ListImageResponseData {
	s.SysDiskSize = &v
	return s
}

func (s *ListImageResponseData) SetImageType(v string) *ListImageResponseData {
	s.ImageType = &v
	return s
}

func (s *ListImageResponseData) SetImageTarget(v string) *ListImageResponseData {
	s.ImageTarget = &v
	return s
}

func (s *ListImageResponseData) SetDc(v string) *ListImageResponseData {
	s.Dc = &v
	return s
}

func (s *ListImageResponseData) SetPoolName(v string) *ListImageResponseData {
	s.PoolName = &v
	return s
}

func (s *ListImageResponseData) SetResponseMsg(v string) *ListImageResponseData {
	s.ResponseMsg = &v
	return s
}


type ListImageResponseDataBuilder struct {
	s *ListImageResponseData
}

func NewListImageResponseDataBuilder() *ListImageResponseDataBuilder {
	s := &ListImageResponseData{}
	b := &ListImageResponseDataBuilder{s: s}
	return b
}

func (b *ListImageResponseDataBuilder) CompanyCode(v string) *ListImageResponseDataBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *ListImageResponseDataBuilder) ModifiedTime(v string) *ListImageResponseDataBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListImageResponseDataBuilder) ImageId(v string) *ListImageResponseDataBuilder {
	b.s.ImageId = &v
	return b
}

func (b *ListImageResponseDataBuilder) ImageName(v string) *ListImageResponseDataBuilder {
	b.s.ImageName = &v
	return b
}

func (b *ListImageResponseDataBuilder) ImageRemark(v string) *ListImageResponseDataBuilder {
	b.s.ImageRemark = &v
	return b
}

func (b *ListImageResponseDataBuilder) CompanyName(v string) *ListImageResponseDataBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *ListImageResponseDataBuilder) RegionName(v string) *ListImageResponseDataBuilder {
	b.s.RegionName = &v
	return b
}

func (b *ListImageResponseDataBuilder) ErrorMessage(v string) *ListImageResponseDataBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListImageResponseDataBuilder) ErrorCode(v string) *ListImageResponseDataBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListImageResponseDataBuilder) ImageOs(v string) *ListImageResponseDataBuilder {
	b.s.ImageOs = &v
	return b
}

func (b *ListImageResponseDataBuilder) ImageStatus(v string) *ListImageResponseDataBuilder {
	b.s.ImageStatus = &v
	return b
}

func (b *ListImageResponseDataBuilder) TemplateId(v string) *ListImageResponseDataBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ListImageResponseDataBuilder) RequestId(v string) *ListImageResponseDataBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListImageResponseDataBuilder) CreatedTime(v string) *ListImageResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListImageResponseDataBuilder) PoolCode(v string) *ListImageResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *ListImageResponseDataBuilder) DiskType(v string) *ListImageResponseDataBuilder {
	b.s.DiskType = &v
	return b
}

func (b *ListImageResponseDataBuilder) Region(v string) *ListImageResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *ListImageResponseDataBuilder) SysDiskSize(v string) *ListImageResponseDataBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *ListImageResponseDataBuilder) ImageType(v string) *ListImageResponseDataBuilder {
	b.s.ImageType = &v
	return b
}

func (b *ListImageResponseDataBuilder) ImageTarget(v string) *ListImageResponseDataBuilder {
	b.s.ImageTarget = &v
	return b
}

func (b *ListImageResponseDataBuilder) Dc(v string) *ListImageResponseDataBuilder {
	b.s.Dc = &v
	return b
}

func (b *ListImageResponseDataBuilder) PoolName(v string) *ListImageResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListImageResponseDataBuilder) ResponseMsg(v string) *ListImageResponseDataBuilder {
	b.s.ResponseMsg = &v
	return b
}

func (b *ListImageResponseDataBuilder) Build() *ListImageResponseData {
	return b.s
}


