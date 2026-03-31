// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetRestartRecordListResponseBody struct {

    // 数据
	Data *[]GetRestartRecordListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetRestartRecordListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetRestartRecordListResponseBody) GoString() string {
	return s.String()
}

func (s GetRestartRecordListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetRestartRecordListResponseBody) SetData(v []GetRestartRecordListResponseData) *GetRestartRecordListResponseBody {
	s.Data = &v
	return s
}

func (s *GetRestartRecordListResponseBody) SetTotalCount(v int32) *GetRestartRecordListResponseBody {
	s.TotalCount = &v
	return s
}


type GetRestartRecordListResponseBodyBuilder struct {
	s *GetRestartRecordListResponseBody
}

func NewGetRestartRecordListResponseBodyBuilder() *GetRestartRecordListResponseBodyBuilder {
	s := &GetRestartRecordListResponseBody{}
	b := &GetRestartRecordListResponseBodyBuilder{s: s}
	return b
}

func (b *GetRestartRecordListResponseBodyBuilder) Data(v []GetRestartRecordListResponseData) *GetRestartRecordListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetRestartRecordListResponseBodyBuilder) TotalCount(v int32) *GetRestartRecordListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetRestartRecordListResponseBodyBuilder) Build() *GetRestartRecordListResponseBody {
	return b.s
}


