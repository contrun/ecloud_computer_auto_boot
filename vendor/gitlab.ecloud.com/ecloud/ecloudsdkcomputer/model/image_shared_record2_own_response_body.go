// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageSharedRecord2OwnResponseBody struct {

    // 数据
	Data *[]ImageSharedRecord2OwnResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s ImageSharedRecord2OwnResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ImageSharedRecord2OwnResponseBody) GoString() string {
	return s.String()
}

func (s ImageSharedRecord2OwnResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageSharedRecord2OwnResponseBody) SetData(v []ImageSharedRecord2OwnResponseData) *ImageSharedRecord2OwnResponseBody {
	s.Data = &v
	return s
}

func (s *ImageSharedRecord2OwnResponseBody) SetTotalCount(v int32) *ImageSharedRecord2OwnResponseBody {
	s.TotalCount = &v
	return s
}


type ImageSharedRecord2OwnResponseBodyBuilder struct {
	s *ImageSharedRecord2OwnResponseBody
}

func NewImageSharedRecord2OwnResponseBodyBuilder() *ImageSharedRecord2OwnResponseBodyBuilder {
	s := &ImageSharedRecord2OwnResponseBody{}
	b := &ImageSharedRecord2OwnResponseBodyBuilder{s: s}
	return b
}

func (b *ImageSharedRecord2OwnResponseBodyBuilder) Data(v []ImageSharedRecord2OwnResponseData) *ImageSharedRecord2OwnResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseBodyBuilder) TotalCount(v int32) *ImageSharedRecord2OwnResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseBodyBuilder) Build() *ImageSharedRecord2OwnResponseBody {
	return b.s
}


