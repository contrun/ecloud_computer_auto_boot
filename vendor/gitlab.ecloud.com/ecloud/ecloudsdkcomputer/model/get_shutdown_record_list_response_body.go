// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetShutdownRecordListResponseBody struct {

    // 数据
	Data *[]GetShutdownRecordListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetShutdownRecordListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetShutdownRecordListResponseBody) GoString() string {
	return s.String()
}

func (s GetShutdownRecordListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShutdownRecordListResponseBody) SetData(v []GetShutdownRecordListResponseData) *GetShutdownRecordListResponseBody {
	s.Data = &v
	return s
}

func (s *GetShutdownRecordListResponseBody) SetTotalCount(v int32) *GetShutdownRecordListResponseBody {
	s.TotalCount = &v
	return s
}


type GetShutdownRecordListResponseBodyBuilder struct {
	s *GetShutdownRecordListResponseBody
}

func NewGetShutdownRecordListResponseBodyBuilder() *GetShutdownRecordListResponseBodyBuilder {
	s := &GetShutdownRecordListResponseBody{}
	b := &GetShutdownRecordListResponseBodyBuilder{s: s}
	return b
}

func (b *GetShutdownRecordListResponseBodyBuilder) Data(v []GetShutdownRecordListResponseData) *GetShutdownRecordListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetShutdownRecordListResponseBodyBuilder) TotalCount(v int32) *GetShutdownRecordListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetShutdownRecordListResponseBodyBuilder) Build() *GetShutdownRecordListResponseBody {
	return b.s
}


