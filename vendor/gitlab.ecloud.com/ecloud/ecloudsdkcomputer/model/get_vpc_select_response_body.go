// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVpcSelectResponseBody struct {

    // 数据
	Data *[]GetVpcSelectResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetVpcSelectResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetVpcSelectResponseBody) GoString() string {
	return s.String()
}

func (s GetVpcSelectResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcSelectResponseBody) SetData(v []GetVpcSelectResponseData) *GetVpcSelectResponseBody {
	s.Data = &v
	return s
}

func (s *GetVpcSelectResponseBody) SetTotalCount(v int32) *GetVpcSelectResponseBody {
	s.TotalCount = &v
	return s
}


type GetVpcSelectResponseBodyBuilder struct {
	s *GetVpcSelectResponseBody
}

func NewGetVpcSelectResponseBodyBuilder() *GetVpcSelectResponseBodyBuilder {
	s := &GetVpcSelectResponseBody{}
	b := &GetVpcSelectResponseBodyBuilder{s: s}
	return b
}

func (b *GetVpcSelectResponseBodyBuilder) Data(v []GetVpcSelectResponseData) *GetVpcSelectResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetVpcSelectResponseBodyBuilder) TotalCount(v int32) *GetVpcSelectResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetVpcSelectResponseBodyBuilder) Build() *GetVpcSelectResponseBody {
	return b.s
}


