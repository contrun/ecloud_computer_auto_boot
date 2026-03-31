// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTaskRecordListByInstanceIdResponseBody struct {

    // 数据
	Data *[]GetTaskRecordListByInstanceIdResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetTaskRecordListByInstanceIdResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s GetTaskRecordListByInstanceIdResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListByInstanceIdResponseBody) SetData(v []GetTaskRecordListByInstanceIdResponseData) *GetTaskRecordListByInstanceIdResponseBody {
	s.Data = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponseBody) SetTotalCount(v int32) *GetTaskRecordListByInstanceIdResponseBody {
	s.TotalCount = &v
	return s
}


type GetTaskRecordListByInstanceIdResponseBodyBuilder struct {
	s *GetTaskRecordListByInstanceIdResponseBody
}

func NewGetTaskRecordListByInstanceIdResponseBodyBuilder() *GetTaskRecordListByInstanceIdResponseBodyBuilder {
	s := &GetTaskRecordListByInstanceIdResponseBody{}
	b := &GetTaskRecordListByInstanceIdResponseBodyBuilder{s: s}
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseBodyBuilder) Data(v []GetTaskRecordListByInstanceIdResponseData) *GetTaskRecordListByInstanceIdResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseBodyBuilder) TotalCount(v int32) *GetTaskRecordListByInstanceIdResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseBodyBuilder) Build() *GetTaskRecordListByInstanceIdResponseBody {
	return b.s
}


