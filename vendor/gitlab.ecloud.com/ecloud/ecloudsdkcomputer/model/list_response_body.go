// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListResponseBody struct {

    // 数据
	Data *[]ListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s ListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListResponseBody) GoString() string {
	return s.String()
}

func (s ListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListResponseBody) SetData(v []ListResponseData) *ListResponseBody {
	s.Data = &v
	return s
}

func (s *ListResponseBody) SetTotalCount(v int32) *ListResponseBody {
	s.TotalCount = &v
	return s
}


type ListResponseBodyBuilder struct {
	s *ListResponseBody
}

func NewListResponseBodyBuilder() *ListResponseBodyBuilder {
	s := &ListResponseBody{}
	b := &ListResponseBodyBuilder{s: s}
	return b
}

func (b *ListResponseBodyBuilder) Data(v []ListResponseData) *ListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *ListResponseBodyBuilder) TotalCount(v int32) *ListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *ListResponseBodyBuilder) Build() *ListResponseBody {
	return b.s
}


