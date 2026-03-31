// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListImageResponseBody struct {

    // 数据
	Data *[]ListImageResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s ListImageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListImageResponseBody) GoString() string {
	return s.String()
}

func (s ListImageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListImageResponseBody) SetData(v []ListImageResponseData) *ListImageResponseBody {
	s.Data = &v
	return s
}

func (s *ListImageResponseBody) SetTotalCount(v int32) *ListImageResponseBody {
	s.TotalCount = &v
	return s
}


type ListImageResponseBodyBuilder struct {
	s *ListImageResponseBody
}

func NewListImageResponseBodyBuilder() *ListImageResponseBodyBuilder {
	s := &ListImageResponseBody{}
	b := &ListImageResponseBodyBuilder{s: s}
	return b
}

func (b *ListImageResponseBodyBuilder) Data(v []ListImageResponseData) *ListImageResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *ListImageResponseBodyBuilder) TotalCount(v int32) *ListImageResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *ListImageResponseBodyBuilder) Build() *ListImageResponseBody {
	return b.s
}


