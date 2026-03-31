// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageRecordSharedByOthersResponseBody struct {

    // 数据
	Data *[]ImageRecordSharedByOthersResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s ImageRecordSharedByOthersResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedByOthersResponseBody) GoString() string {
	return s.String()
}

func (s ImageRecordSharedByOthersResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedByOthersResponseBody) SetData(v []ImageRecordSharedByOthersResponseData) *ImageRecordSharedByOthersResponseBody {
	s.Data = &v
	return s
}

func (s *ImageRecordSharedByOthersResponseBody) SetTotalCount(v int32) *ImageRecordSharedByOthersResponseBody {
	s.TotalCount = &v
	return s
}


type ImageRecordSharedByOthersResponseBodyBuilder struct {
	s *ImageRecordSharedByOthersResponseBody
}

func NewImageRecordSharedByOthersResponseBodyBuilder() *ImageRecordSharedByOthersResponseBodyBuilder {
	s := &ImageRecordSharedByOthersResponseBody{}
	b := &ImageRecordSharedByOthersResponseBodyBuilder{s: s}
	return b
}

func (b *ImageRecordSharedByOthersResponseBodyBuilder) Data(v []ImageRecordSharedByOthersResponseData) *ImageRecordSharedByOthersResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseBodyBuilder) TotalCount(v int32) *ImageRecordSharedByOthersResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseBodyBuilder) Build() *ImageRecordSharedByOthersResponseBody {
	return b.s
}


