// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageRecordSharedByOthersBody struct {
    position.Body
    // 所属厂商
	CompanyCode *string `json:"companyCode,omitempty"`
    // 镜像模板名称
	ImageName *string `json:"imageName,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 镜像UID
	TemplateId *string `json:"templateId,omitempty"`
    // 可用区编码
	Region *string `json:"region,omitempty"`
    // 共享状态
	ShareStatus *string `json:"shareStatus,omitempty"`
    // 网络工作区
	Dc *string `json:"dc,omitempty"`
}

func (s ImageRecordSharedByOthersBody) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedByOthersBody) GoString() string {
	return s.String()
}

func (s ImageRecordSharedByOthersBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedByOthersBody) SetCompanyCode(v string) *ImageRecordSharedByOthersBody {
	s.CompanyCode = &v
	return s
}

func (s *ImageRecordSharedByOthersBody) SetImageName(v string) *ImageRecordSharedByOthersBody {
	s.ImageName = &v
	return s
}

func (s *ImageRecordSharedByOthersBody) SetPageSize(v int32) *ImageRecordSharedByOthersBody {
	s.PageSize = &v
	return s
}

func (s *ImageRecordSharedByOthersBody) SetPoolCode(v string) *ImageRecordSharedByOthersBody {
	s.PoolCode = &v
	return s
}

func (s *ImageRecordSharedByOthersBody) SetCurrentPage(v int32) *ImageRecordSharedByOthersBody {
	s.CurrentPage = &v
	return s
}

func (s *ImageRecordSharedByOthersBody) SetTemplateId(v string) *ImageRecordSharedByOthersBody {
	s.TemplateId = &v
	return s
}

func (s *ImageRecordSharedByOthersBody) SetRegion(v string) *ImageRecordSharedByOthersBody {
	s.Region = &v
	return s
}

func (s *ImageRecordSharedByOthersBody) SetShareStatus(v string) *ImageRecordSharedByOthersBody {
	s.ShareStatus = &v
	return s
}

func (s *ImageRecordSharedByOthersBody) SetDc(v string) *ImageRecordSharedByOthersBody {
	s.Dc = &v
	return s
}


type ImageRecordSharedByOthersBodyBuilder struct {
	s *ImageRecordSharedByOthersBody
}

func NewImageRecordSharedByOthersBodyBuilder() *ImageRecordSharedByOthersBodyBuilder {
	s := &ImageRecordSharedByOthersBody{}
	b := &ImageRecordSharedByOthersBodyBuilder{s: s}
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) CompanyCode(v string) *ImageRecordSharedByOthersBodyBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) ImageName(v string) *ImageRecordSharedByOthersBodyBuilder {
	b.s.ImageName = &v
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) PageSize(v int32) *ImageRecordSharedByOthersBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) PoolCode(v string) *ImageRecordSharedByOthersBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) CurrentPage(v int32) *ImageRecordSharedByOthersBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) TemplateId(v string) *ImageRecordSharedByOthersBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) Region(v string) *ImageRecordSharedByOthersBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) ShareStatus(v string) *ImageRecordSharedByOthersBodyBuilder {
	b.s.ShareStatus = &v
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) Dc(v string) *ImageRecordSharedByOthersBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *ImageRecordSharedByOthersBodyBuilder) Build() *ImageRecordSharedByOthersBody {
	return b.s
}


