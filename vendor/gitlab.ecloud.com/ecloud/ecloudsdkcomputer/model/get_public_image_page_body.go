// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPublicImagePageBody struct {
    position.Body
    // 镜像模板名称
	ImageName *string `json:"imageName,omitempty"`
    // 电脑协议（V2.0:自研CMSSZTE，V1.1：新华三，V1.2：中兴，V 1.4：中兴大云融合）
	ComputerProtocol *string `json:"computerProtocol,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 资源池
	PoolCode *string `json:"poolCode,omitempty"`
    // 镜像操作系统
	ImageOs *string `json:"imageOs,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 镜像模板ID
	TemplateId *string `json:"templateId,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 系统盘大小
	SysDiskSize *string `json:"sysDiskSize,omitempty"`
    // 数据中心
	Dc *string `json:"dc,omitempty"`
}

func (s GetPublicImagePageBody) String() string {
	return utils.Beautify(s)
}

func (s GetPublicImagePageBody) GoString() string {
	return s.String()
}

func (s GetPublicImagePageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPublicImagePageBody) SetImageName(v string) *GetPublicImagePageBody {
	s.ImageName = &v
	return s
}

func (s *GetPublicImagePageBody) SetComputerProtocol(v string) *GetPublicImagePageBody {
	s.ComputerProtocol = &v
	return s
}

func (s *GetPublicImagePageBody) SetPageSize(v int32) *GetPublicImagePageBody {
	s.PageSize = &v
	return s
}

func (s *GetPublicImagePageBody) SetPoolCode(v string) *GetPublicImagePageBody {
	s.PoolCode = &v
	return s
}

func (s *GetPublicImagePageBody) SetImageOs(v string) *GetPublicImagePageBody {
	s.ImageOs = &v
	return s
}

func (s *GetPublicImagePageBody) SetCurrentPage(v int32) *GetPublicImagePageBody {
	s.CurrentPage = &v
	return s
}

func (s *GetPublicImagePageBody) SetTemplateId(v string) *GetPublicImagePageBody {
	s.TemplateId = &v
	return s
}

func (s *GetPublicImagePageBody) SetRegion(v string) *GetPublicImagePageBody {
	s.Region = &v
	return s
}

func (s *GetPublicImagePageBody) SetSysDiskSize(v string) *GetPublicImagePageBody {
	s.SysDiskSize = &v
	return s
}

func (s *GetPublicImagePageBody) SetDc(v string) *GetPublicImagePageBody {
	s.Dc = &v
	return s
}


type GetPublicImagePageBodyBuilder struct {
	s *GetPublicImagePageBody
}

func NewGetPublicImagePageBodyBuilder() *GetPublicImagePageBodyBuilder {
	s := &GetPublicImagePageBody{}
	b := &GetPublicImagePageBodyBuilder{s: s}
	return b
}

func (b *GetPublicImagePageBodyBuilder) ImageName(v string) *GetPublicImagePageBodyBuilder {
	b.s.ImageName = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) ComputerProtocol(v string) *GetPublicImagePageBodyBuilder {
	b.s.ComputerProtocol = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) PageSize(v int32) *GetPublicImagePageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) PoolCode(v string) *GetPublicImagePageBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) ImageOs(v string) *GetPublicImagePageBodyBuilder {
	b.s.ImageOs = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) CurrentPage(v int32) *GetPublicImagePageBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) TemplateId(v string) *GetPublicImagePageBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) Region(v string) *GetPublicImagePageBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) SysDiskSize(v string) *GetPublicImagePageBodyBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) Dc(v string) *GetPublicImagePageBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *GetPublicImagePageBodyBuilder) Build() *GetPublicImagePageBody {
	return b.s
}


