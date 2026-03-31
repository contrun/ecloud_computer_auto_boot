// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListImageBody struct {
    position.Body
    // 镜像模板名称
	ImageName *string `json:"imageName,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 资源池，例如：CIDC-RP-33
	PoolCode *string `json:"poolCode,omitempty"`
    // 镜像状态，creating:创建中,create:创建完成,createError:创建失败,deleting:删除中,deleted:已删除,deletedError:删除失败
	ImageStatus *string `json:"imageStatus,omitempty"`
    // 可用区，例如：CIDC-RP-33-AZ1
	Region *string `json:"region,omitempty"`
    // 镜像模板ID
	TemplateId *string `json:"templateId,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 可用区，例如：ONEDAAS-SERVICE-0002
	Dc *string `json:"dc,omitempty"`
}

func (s ListImageBody) String() string {
	return utils.Beautify(s)
}

func (s ListImageBody) GoString() string {
	return s.String()
}

func (s ListImageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListImageBody) SetImageName(v string) *ListImageBody {
	s.ImageName = &v
	return s
}

func (s *ListImageBody) SetPageSize(v int32) *ListImageBody {
	s.PageSize = &v
	return s
}

func (s *ListImageBody) SetPoolCode(v string) *ListImageBody {
	s.PoolCode = &v
	return s
}

func (s *ListImageBody) SetImageStatus(v string) *ListImageBody {
	s.ImageStatus = &v
	return s
}

func (s *ListImageBody) SetRegion(v string) *ListImageBody {
	s.Region = &v
	return s
}

func (s *ListImageBody) SetTemplateId(v string) *ListImageBody {
	s.TemplateId = &v
	return s
}

func (s *ListImageBody) SetCurrentPage(v int32) *ListImageBody {
	s.CurrentPage = &v
	return s
}

func (s *ListImageBody) SetDc(v string) *ListImageBody {
	s.Dc = &v
	return s
}


type ListImageBodyBuilder struct {
	s *ListImageBody
}

func NewListImageBodyBuilder() *ListImageBodyBuilder {
	s := &ListImageBody{}
	b := &ListImageBodyBuilder{s: s}
	return b
}

func (b *ListImageBodyBuilder) ImageName(v string) *ListImageBodyBuilder {
	b.s.ImageName = &v
	return b
}

func (b *ListImageBodyBuilder) PageSize(v int32) *ListImageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListImageBodyBuilder) PoolCode(v string) *ListImageBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *ListImageBodyBuilder) ImageStatus(v string) *ListImageBodyBuilder {
	b.s.ImageStatus = &v
	return b
}

func (b *ListImageBodyBuilder) Region(v string) *ListImageBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *ListImageBodyBuilder) TemplateId(v string) *ListImageBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ListImageBodyBuilder) CurrentPage(v int32) *ListImageBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *ListImageBodyBuilder) Dc(v string) *ListImageBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *ListImageBodyBuilder) Build() *ListImageBody {
	return b.s
}


