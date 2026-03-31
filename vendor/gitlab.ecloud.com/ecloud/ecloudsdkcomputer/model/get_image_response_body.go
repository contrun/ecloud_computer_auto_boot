// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetImageResponseBody struct {

    // 数据
	Data *[]GetImageResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s GetImageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetImageResponseBody) SetData(v []GetImageResponseData) *GetImageResponseBody {
	s.Data = &v
	return s
}

func (s *GetImageResponseBody) SetTotalCount(v int32) *GetImageResponseBody {
	s.TotalCount = &v
	return s
}


type GetImageResponseBodyBuilder struct {
	s *GetImageResponseBody
}

func NewGetImageResponseBodyBuilder() *GetImageResponseBodyBuilder {
	s := &GetImageResponseBody{}
	b := &GetImageResponseBodyBuilder{s: s}
	return b
}

func (b *GetImageResponseBodyBuilder) Data(v []GetImageResponseData) *GetImageResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetImageResponseBodyBuilder) TotalCount(v int32) *GetImageResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetImageResponseBodyBuilder) Build() *GetImageResponseBody {
	return b.s
}


