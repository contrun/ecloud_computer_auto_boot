// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetApplyImageRecordResponseBody struct {

    // 数据
	Data *[]GetApplyImageRecordResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetApplyImageRecordResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetApplyImageRecordResponseBody) GoString() string {
	return s.String()
}

func (s GetApplyImageRecordResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplyImageRecordResponseBody) SetData(v []GetApplyImageRecordResponseData) *GetApplyImageRecordResponseBody {
	s.Data = &v
	return s
}

func (s *GetApplyImageRecordResponseBody) SetTotalCount(v int32) *GetApplyImageRecordResponseBody {
	s.TotalCount = &v
	return s
}


type GetApplyImageRecordResponseBodyBuilder struct {
	s *GetApplyImageRecordResponseBody
}

func NewGetApplyImageRecordResponseBodyBuilder() *GetApplyImageRecordResponseBodyBuilder {
	s := &GetApplyImageRecordResponseBody{}
	b := &GetApplyImageRecordResponseBodyBuilder{s: s}
	return b
}

func (b *GetApplyImageRecordResponseBodyBuilder) Data(v []GetApplyImageRecordResponseData) *GetApplyImageRecordResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetApplyImageRecordResponseBodyBuilder) TotalCount(v int32) *GetApplyImageRecordResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetApplyImageRecordResponseBodyBuilder) Build() *GetApplyImageRecordResponseBody {
	return b.s
}


