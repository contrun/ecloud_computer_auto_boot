// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindResourceRecordsResponseBody struct {

    // 数据
	Data *[]BindResourceRecordsResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s BindResourceRecordsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BindResourceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s BindResourceRecordsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindResourceRecordsResponseBody) SetData(v []BindResourceRecordsResponseData) *BindResourceRecordsResponseBody {
	s.Data = &v
	return s
}

func (s *BindResourceRecordsResponseBody) SetTotalCount(v int32) *BindResourceRecordsResponseBody {
	s.TotalCount = &v
	return s
}


type BindResourceRecordsResponseBodyBuilder struct {
	s *BindResourceRecordsResponseBody
}

func NewBindResourceRecordsResponseBodyBuilder() *BindResourceRecordsResponseBodyBuilder {
	s := &BindResourceRecordsResponseBody{}
	b := &BindResourceRecordsResponseBodyBuilder{s: s}
	return b
}

func (b *BindResourceRecordsResponseBodyBuilder) Data(v []BindResourceRecordsResponseData) *BindResourceRecordsResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *BindResourceRecordsResponseBodyBuilder) TotalCount(v int32) *BindResourceRecordsResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *BindResourceRecordsResponseBodyBuilder) Build() *BindResourceRecordsResponseBody {
	return b.s
}


