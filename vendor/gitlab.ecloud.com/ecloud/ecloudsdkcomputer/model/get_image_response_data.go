// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetImageResponseData struct {

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

func (s GetImageResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetImageResponseData) GoString() string {
	return s.String()
}

func (s GetImageResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetImageResponseData) SetModifiedTime(v string) *GetImageResponseData {
	s.ModifiedTime = &v
	return s
}

func (s *GetImageResponseData) SetImageId(v string) *GetImageResponseData {
	s.ImageId = &v
	return s
}

func (s *GetImageResponseData) SetImageName(v string) *GetImageResponseData {
	s.ImageName = &v
	return s
}

func (s *GetImageResponseData) SetImageRemark(v string) *GetImageResponseData {
	s.ImageRemark = &v
	return s
}

func (s *GetImageResponseData) SetCompanyName(v string) *GetImageResponseData {
	s.CompanyName = &v
	return s
}

func (s *GetImageResponseData) SetRegionName(v string) *GetImageResponseData {
	s.RegionName = &v
	return s
}

func (s *GetImageResponseData) SetErrorMessage(v string) *GetImageResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *GetImageResponseData) SetErrorCode(v string) *GetImageResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GetImageResponseData) SetImageOs(v string) *GetImageResponseData {
	s.ImageOs = &v
	return s
}

func (s *GetImageResponseData) SetImageStatus(v string) *GetImageResponseData {
	s.ImageStatus = &v
	return s
}

func (s *GetImageResponseData) SetTemplateId(v string) *GetImageResponseData {
	s.TemplateId = &v
	return s
}

func (s *GetImageResponseData) SetRequestId(v string) *GetImageResponseData {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseData) SetCreatedTime(v string) *GetImageResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetImageResponseData) SetPoolCode(v string) *GetImageResponseData {
	s.PoolCode = &v
	return s
}

func (s *GetImageResponseData) SetDiskType(v string) *GetImageResponseData {
	s.DiskType = &v
	return s
}

func (s *GetImageResponseData) SetRegion(v string) *GetImageResponseData {
	s.Region = &v
	return s
}

func (s *GetImageResponseData) SetSysDiskSize(v string) *GetImageResponseData {
	s.SysDiskSize = &v
	return s
}

func (s *GetImageResponseData) SetImageType(v string) *GetImageResponseData {
	s.ImageType = &v
	return s
}

func (s *GetImageResponseData) SetImageTarget(v string) *GetImageResponseData {
	s.ImageTarget = &v
	return s
}

func (s *GetImageResponseData) SetDc(v string) *GetImageResponseData {
	s.Dc = &v
	return s
}

func (s *GetImageResponseData) SetPoolName(v string) *GetImageResponseData {
	s.PoolName = &v
	return s
}

func (s *GetImageResponseData) SetResponseMsg(v string) *GetImageResponseData {
	s.ResponseMsg = &v
	return s
}


type GetImageResponseDataBuilder struct {
	s *GetImageResponseData
}

func NewGetImageResponseDataBuilder() *GetImageResponseDataBuilder {
	s := &GetImageResponseData{}
	b := &GetImageResponseDataBuilder{s: s}
	return b
}

func (b *GetImageResponseDataBuilder) ModifiedTime(v string) *GetImageResponseDataBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetImageResponseDataBuilder) ImageId(v string) *GetImageResponseDataBuilder {
	b.s.ImageId = &v
	return b
}

func (b *GetImageResponseDataBuilder) ImageName(v string) *GetImageResponseDataBuilder {
	b.s.ImageName = &v
	return b
}

func (b *GetImageResponseDataBuilder) ImageRemark(v string) *GetImageResponseDataBuilder {
	b.s.ImageRemark = &v
	return b
}

func (b *GetImageResponseDataBuilder) CompanyName(v string) *GetImageResponseDataBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *GetImageResponseDataBuilder) RegionName(v string) *GetImageResponseDataBuilder {
	b.s.RegionName = &v
	return b
}

func (b *GetImageResponseDataBuilder) ErrorMessage(v string) *GetImageResponseDataBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetImageResponseDataBuilder) ErrorCode(v string) *GetImageResponseDataBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetImageResponseDataBuilder) ImageOs(v string) *GetImageResponseDataBuilder {
	b.s.ImageOs = &v
	return b
}

func (b *GetImageResponseDataBuilder) ImageStatus(v string) *GetImageResponseDataBuilder {
	b.s.ImageStatus = &v
	return b
}

func (b *GetImageResponseDataBuilder) TemplateId(v string) *GetImageResponseDataBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *GetImageResponseDataBuilder) RequestId(v string) *GetImageResponseDataBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetImageResponseDataBuilder) CreatedTime(v string) *GetImageResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetImageResponseDataBuilder) PoolCode(v string) *GetImageResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetImageResponseDataBuilder) DiskType(v string) *GetImageResponseDataBuilder {
	b.s.DiskType = &v
	return b
}

func (b *GetImageResponseDataBuilder) Region(v string) *GetImageResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *GetImageResponseDataBuilder) SysDiskSize(v string) *GetImageResponseDataBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *GetImageResponseDataBuilder) ImageType(v string) *GetImageResponseDataBuilder {
	b.s.ImageType = &v
	return b
}

func (b *GetImageResponseDataBuilder) ImageTarget(v string) *GetImageResponseDataBuilder {
	b.s.ImageTarget = &v
	return b
}

func (b *GetImageResponseDataBuilder) Dc(v string) *GetImageResponseDataBuilder {
	b.s.Dc = &v
	return b
}

func (b *GetImageResponseDataBuilder) PoolName(v string) *GetImageResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetImageResponseDataBuilder) ResponseMsg(v string) *GetImageResponseDataBuilder {
	b.s.ResponseMsg = &v
	return b
}

func (b *GetImageResponseDataBuilder) Build() *GetImageResponseData {
	return b.s
}


