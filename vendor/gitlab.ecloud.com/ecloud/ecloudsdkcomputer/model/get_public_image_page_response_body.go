// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPublicImagePageResponseBody struct {

    // 数据
	Data *[]GetPublicImagePageResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetPublicImagePageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetPublicImagePageResponseBody) GoString() string {
	return s.String()
}

func (s GetPublicImagePageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPublicImagePageResponseBody) SetData(v []GetPublicImagePageResponseData) *GetPublicImagePageResponseBody {
	s.Data = &v
	return s
}

func (s *GetPublicImagePageResponseBody) SetTotalCount(v int32) *GetPublicImagePageResponseBody {
	s.TotalCount = &v
	return s
}


type GetPublicImagePageResponseBodyBuilder struct {
	s *GetPublicImagePageResponseBody
}

func NewGetPublicImagePageResponseBodyBuilder() *GetPublicImagePageResponseBodyBuilder {
	s := &GetPublicImagePageResponseBody{}
	b := &GetPublicImagePageResponseBodyBuilder{s: s}
	return b
}

func (b *GetPublicImagePageResponseBodyBuilder) Data(v []GetPublicImagePageResponseData) *GetPublicImagePageResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetPublicImagePageResponseBodyBuilder) TotalCount(v int32) *GetPublicImagePageResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetPublicImagePageResponseBodyBuilder) Build() *GetPublicImagePageResponseBody {
	return b.s
}


