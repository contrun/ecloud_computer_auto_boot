// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNetworkResponseBody struct {

    // 数据
	Data *[]GetNetworkResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetNetworkResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkResponseBody) GoString() string {
	return s.String()
}

func (s GetNetworkResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkResponseBody) SetData(v []GetNetworkResponseData) *GetNetworkResponseBody {
	s.Data = &v
	return s
}

func (s *GetNetworkResponseBody) SetTotalCount(v int32) *GetNetworkResponseBody {
	s.TotalCount = &v
	return s
}


type GetNetworkResponseBodyBuilder struct {
	s *GetNetworkResponseBody
}

func NewGetNetworkResponseBodyBuilder() *GetNetworkResponseBodyBuilder {
	s := &GetNetworkResponseBody{}
	b := &GetNetworkResponseBodyBuilder{s: s}
	return b
}

func (b *GetNetworkResponseBodyBuilder) Data(v []GetNetworkResponseData) *GetNetworkResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetNetworkResponseBodyBuilder) TotalCount(v int32) *GetNetworkResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetNetworkResponseBodyBuilder) Build() *GetNetworkResponseBody {
	return b.s
}


