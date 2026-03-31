// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAvailableRecordListResponseBody struct {

    // 数据
	Data *[]GetAvailableRecordListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetAvailableRecordListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetAvailableRecordListResponseBody) GoString() string {
	return s.String()
}

func (s GetAvailableRecordListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAvailableRecordListResponseBody) SetData(v []GetAvailableRecordListResponseData) *GetAvailableRecordListResponseBody {
	s.Data = &v
	return s
}

func (s *GetAvailableRecordListResponseBody) SetTotalCount(v int32) *GetAvailableRecordListResponseBody {
	s.TotalCount = &v
	return s
}


type GetAvailableRecordListResponseBodyBuilder struct {
	s *GetAvailableRecordListResponseBody
}

func NewGetAvailableRecordListResponseBodyBuilder() *GetAvailableRecordListResponseBodyBuilder {
	s := &GetAvailableRecordListResponseBody{}
	b := &GetAvailableRecordListResponseBodyBuilder{s: s}
	return b
}

func (b *GetAvailableRecordListResponseBodyBuilder) Data(v []GetAvailableRecordListResponseData) *GetAvailableRecordListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetAvailableRecordListResponseBodyBuilder) TotalCount(v int32) *GetAvailableRecordListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetAvailableRecordListResponseBodyBuilder) Build() *GetAvailableRecordListResponseBody {
	return b.s
}


