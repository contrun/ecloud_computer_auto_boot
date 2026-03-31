// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageRecordSharedWithOthersBody struct {
    position.Body
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 镜像UID
	TemplateId *string `json:"templateId,omitempty"`
    // 共享状态
	ShareStatus *string `json:"shareStatus,omitempty"`
}

func (s ImageRecordSharedWithOthersBody) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedWithOthersBody) GoString() string {
	return s.String()
}

func (s ImageRecordSharedWithOthersBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedWithOthersBody) SetPageSize(v int32) *ImageRecordSharedWithOthersBody {
	s.PageSize = &v
	return s
}

func (s *ImageRecordSharedWithOthersBody) SetCurrentPage(v int32) *ImageRecordSharedWithOthersBody {
	s.CurrentPage = &v
	return s
}

func (s *ImageRecordSharedWithOthersBody) SetTemplateId(v string) *ImageRecordSharedWithOthersBody {
	s.TemplateId = &v
	return s
}

func (s *ImageRecordSharedWithOthersBody) SetShareStatus(v string) *ImageRecordSharedWithOthersBody {
	s.ShareStatus = &v
	return s
}


type ImageRecordSharedWithOthersBodyBuilder struct {
	s *ImageRecordSharedWithOthersBody
}

func NewImageRecordSharedWithOthersBodyBuilder() *ImageRecordSharedWithOthersBodyBuilder {
	s := &ImageRecordSharedWithOthersBody{}
	b := &ImageRecordSharedWithOthersBodyBuilder{s: s}
	return b
}

func (b *ImageRecordSharedWithOthersBodyBuilder) PageSize(v int32) *ImageRecordSharedWithOthersBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ImageRecordSharedWithOthersBodyBuilder) CurrentPage(v int32) *ImageRecordSharedWithOthersBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *ImageRecordSharedWithOthersBodyBuilder) TemplateId(v string) *ImageRecordSharedWithOthersBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ImageRecordSharedWithOthersBodyBuilder) ShareStatus(v string) *ImageRecordSharedWithOthersBodyBuilder {
	b.s.ShareStatus = &v
	return b
}

func (b *ImageRecordSharedWithOthersBodyBuilder) Build() *ImageRecordSharedWithOthersBody {
	return b.s
}


