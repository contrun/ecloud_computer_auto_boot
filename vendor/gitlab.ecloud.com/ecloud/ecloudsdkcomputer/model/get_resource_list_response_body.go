// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetResourceListResponseBody struct {

    // 总行数
	TotalSize *int32 `json:"totalSize,omitempty"`
    // 数据
	Data *[]GetResourceListResponseData `json:"data,omitempty"`
}

func (s GetResourceListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetResourceListResponseBody) GoString() string {
	return s.String()
}

func (s GetResourceListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetResourceListResponseBody) SetTotalSize(v int32) *GetResourceListResponseBody {
	s.TotalSize = &v
	return s
}

func (s *GetResourceListResponseBody) SetData(v []GetResourceListResponseData) *GetResourceListResponseBody {
	s.Data = &v
	return s
}


type GetResourceListResponseBodyBuilder struct {
	s *GetResourceListResponseBody
}

func NewGetResourceListResponseBodyBuilder() *GetResourceListResponseBodyBuilder {
	s := &GetResourceListResponseBody{}
	b := &GetResourceListResponseBodyBuilder{s: s}
	return b
}

func (b *GetResourceListResponseBodyBuilder) TotalSize(v int32) *GetResourceListResponseBodyBuilder {
	b.s.TotalSize = &v
	return b
}

func (b *GetResourceListResponseBodyBuilder) Data(v []GetResourceListResponseData) *GetResourceListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetResourceListResponseBodyBuilder) Build() *GetResourceListResponseBody {
	return b.s
}


