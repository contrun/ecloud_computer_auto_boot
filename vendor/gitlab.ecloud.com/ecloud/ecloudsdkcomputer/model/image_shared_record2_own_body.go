// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageSharedRecord2OwnBody struct {
    position.Body
    // 厂商编码，例如:H3C
	CompanyCode *string `json:"companyCode,omitempty"`
    // 镜像模板名称
	ImageName *string `json:"imageName,omitempty"`
    // 电脑协议（V2.0:自研CMSSZTE，V1.1：新华三，V1.2：中兴，V 1.4：中兴大云融合）
	ComputerProtocol *string `json:"computerProtocol,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 资源池编码,例如:CIDC-RP-35
	PoolCode *string `json:"poolCode,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 镜像模板UID
	TemplateId *string `json:"templateId,omitempty"`
    // 可用区，例如：CIDC-RP-35-AZ1
	Region *string `json:"region,omitempty"`
    // dc，例如：ONEDAAS-SERVICE-0002
	Dc *string `json:"dc,omitempty"`
}

func (s ImageSharedRecord2OwnBody) String() string {
	return utils.Beautify(s)
}

func (s ImageSharedRecord2OwnBody) GoString() string {
	return s.String()
}

func (s ImageSharedRecord2OwnBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageSharedRecord2OwnBody) SetCompanyCode(v string) *ImageSharedRecord2OwnBody {
	s.CompanyCode = &v
	return s
}

func (s *ImageSharedRecord2OwnBody) SetImageName(v string) *ImageSharedRecord2OwnBody {
	s.ImageName = &v
	return s
}

func (s *ImageSharedRecord2OwnBody) SetComputerProtocol(v string) *ImageSharedRecord2OwnBody {
	s.ComputerProtocol = &v
	return s
}

func (s *ImageSharedRecord2OwnBody) SetPageSize(v int32) *ImageSharedRecord2OwnBody {
	s.PageSize = &v
	return s
}

func (s *ImageSharedRecord2OwnBody) SetPoolCode(v string) *ImageSharedRecord2OwnBody {
	s.PoolCode = &v
	return s
}

func (s *ImageSharedRecord2OwnBody) SetCurrentPage(v int32) *ImageSharedRecord2OwnBody {
	s.CurrentPage = &v
	return s
}

func (s *ImageSharedRecord2OwnBody) SetTemplateId(v string) *ImageSharedRecord2OwnBody {
	s.TemplateId = &v
	return s
}

func (s *ImageSharedRecord2OwnBody) SetRegion(v string) *ImageSharedRecord2OwnBody {
	s.Region = &v
	return s
}

func (s *ImageSharedRecord2OwnBody) SetDc(v string) *ImageSharedRecord2OwnBody {
	s.Dc = &v
	return s
}


type ImageSharedRecord2OwnBodyBuilder struct {
	s *ImageSharedRecord2OwnBody
}

func NewImageSharedRecord2OwnBodyBuilder() *ImageSharedRecord2OwnBodyBuilder {
	s := &ImageSharedRecord2OwnBody{}
	b := &ImageSharedRecord2OwnBodyBuilder{s: s}
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) CompanyCode(v string) *ImageSharedRecord2OwnBodyBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) ImageName(v string) *ImageSharedRecord2OwnBodyBuilder {
	b.s.ImageName = &v
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) ComputerProtocol(v string) *ImageSharedRecord2OwnBodyBuilder {
	b.s.ComputerProtocol = &v
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) PageSize(v int32) *ImageSharedRecord2OwnBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) PoolCode(v string) *ImageSharedRecord2OwnBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) CurrentPage(v int32) *ImageSharedRecord2OwnBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) TemplateId(v string) *ImageSharedRecord2OwnBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) Region(v string) *ImageSharedRecord2OwnBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) Dc(v string) *ImageSharedRecord2OwnBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *ImageSharedRecord2OwnBodyBuilder) Build() *ImageSharedRecord2OwnBody {
	return b.s
}


