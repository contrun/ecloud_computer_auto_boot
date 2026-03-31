// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVpcResponseBody struct {

    // 数据
	Data *[]GetVpcResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetVpcResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetVpcResponseBody) GoString() string {
	return s.String()
}

func (s GetVpcResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcResponseBody) SetData(v []GetVpcResponseData) *GetVpcResponseBody {
	s.Data = &v
	return s
}

func (s *GetVpcResponseBody) SetTotalCount(v int32) *GetVpcResponseBody {
	s.TotalCount = &v
	return s
}


type GetVpcResponseBodyBuilder struct {
	s *GetVpcResponseBody
}

func NewGetVpcResponseBodyBuilder() *GetVpcResponseBodyBuilder {
	s := &GetVpcResponseBody{}
	b := &GetVpcResponseBodyBuilder{s: s}
	return b
}

func (b *GetVpcResponseBodyBuilder) Data(v []GetVpcResponseData) *GetVpcResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetVpcResponseBodyBuilder) TotalCount(v int32) *GetVpcResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetVpcResponseBodyBuilder) Build() *GetVpcResponseBody {
	return b.s
}


