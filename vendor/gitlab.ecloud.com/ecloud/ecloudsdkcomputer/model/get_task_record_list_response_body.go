// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTaskRecordListResponseBody struct {

    // 数据
	Data *[]GetTaskRecordListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetTaskRecordListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListResponseBody) GoString() string {
	return s.String()
}

func (s GetTaskRecordListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListResponseBody) SetData(v []GetTaskRecordListResponseData) *GetTaskRecordListResponseBody {
	s.Data = &v
	return s
}

func (s *GetTaskRecordListResponseBody) SetTotalCount(v int32) *GetTaskRecordListResponseBody {
	s.TotalCount = &v
	return s
}


type GetTaskRecordListResponseBodyBuilder struct {
	s *GetTaskRecordListResponseBody
}

func NewGetTaskRecordListResponseBodyBuilder() *GetTaskRecordListResponseBodyBuilder {
	s := &GetTaskRecordListResponseBody{}
	b := &GetTaskRecordListResponseBodyBuilder{s: s}
	return b
}

func (b *GetTaskRecordListResponseBodyBuilder) Data(v []GetTaskRecordListResponseData) *GetTaskRecordListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetTaskRecordListResponseBodyBuilder) TotalCount(v int32) *GetTaskRecordListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetTaskRecordListResponseBodyBuilder) Build() *GetTaskRecordListResponseBody {
	return b.s
}


