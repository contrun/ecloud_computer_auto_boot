// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindInstanceListResponseBody struct {

    // 数据
	Data *[]BindInstanceListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s BindInstanceListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BindInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s BindInstanceListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindInstanceListResponseBody) SetData(v []BindInstanceListResponseData) *BindInstanceListResponseBody {
	s.Data = &v
	return s
}

func (s *BindInstanceListResponseBody) SetTotalCount(v int32) *BindInstanceListResponseBody {
	s.TotalCount = &v
	return s
}


type BindInstanceListResponseBodyBuilder struct {
	s *BindInstanceListResponseBody
}

func NewBindInstanceListResponseBodyBuilder() *BindInstanceListResponseBodyBuilder {
	s := &BindInstanceListResponseBody{}
	b := &BindInstanceListResponseBodyBuilder{s: s}
	return b
}

func (b *BindInstanceListResponseBodyBuilder) Data(v []BindInstanceListResponseData) *BindInstanceListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *BindInstanceListResponseBodyBuilder) TotalCount(v int32) *BindInstanceListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *BindInstanceListResponseBodyBuilder) Build() *BindInstanceListResponseBody {
	return b.s
}


