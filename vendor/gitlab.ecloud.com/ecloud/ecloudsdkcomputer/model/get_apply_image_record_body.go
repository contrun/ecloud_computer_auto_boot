// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetApplyImageRecordBody struct {
    position.Body
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 镜像模板ID
	TemplateId *string `json:"templateId,omitempty"`
    // 镜像类型标识,自定义:custom,共享:share,公共:public
	ImageTag *string `json:"imageTag,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s GetApplyImageRecordBody) String() string {
	return utils.Beautify(s)
}

func (s GetApplyImageRecordBody) GoString() string {
	return s.String()
}

func (s GetApplyImageRecordBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplyImageRecordBody) SetPageSize(v int32) *GetApplyImageRecordBody {
	s.PageSize = &v
	return s
}

func (s *GetApplyImageRecordBody) SetTemplateId(v string) *GetApplyImageRecordBody {
	s.TemplateId = &v
	return s
}

func (s *GetApplyImageRecordBody) SetImageTag(v string) *GetApplyImageRecordBody {
	s.ImageTag = &v
	return s
}

func (s *GetApplyImageRecordBody) SetCurrentPage(v int32) *GetApplyImageRecordBody {
	s.CurrentPage = &v
	return s
}


type GetApplyImageRecordBodyBuilder struct {
	s *GetApplyImageRecordBody
}

func NewGetApplyImageRecordBodyBuilder() *GetApplyImageRecordBodyBuilder {
	s := &GetApplyImageRecordBody{}
	b := &GetApplyImageRecordBodyBuilder{s: s}
	return b
}

func (b *GetApplyImageRecordBodyBuilder) PageSize(v int32) *GetApplyImageRecordBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetApplyImageRecordBodyBuilder) TemplateId(v string) *GetApplyImageRecordBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *GetApplyImageRecordBodyBuilder) ImageTag(v string) *GetApplyImageRecordBodyBuilder {
	b.s.ImageTag = &v
	return b
}

func (b *GetApplyImageRecordBodyBuilder) CurrentPage(v int32) *GetApplyImageRecordBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetApplyImageRecordBodyBuilder) Build() *GetApplyImageRecordBody {
	return b.s
}


