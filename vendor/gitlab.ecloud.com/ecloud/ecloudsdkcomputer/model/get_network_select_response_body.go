// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNetworkSelectResponseBody struct {

    // 数据
	Data *[]GetNetworkSelectResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetNetworkSelectResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkSelectResponseBody) GoString() string {
	return s.String()
}

func (s GetNetworkSelectResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkSelectResponseBody) SetData(v []GetNetworkSelectResponseData) *GetNetworkSelectResponseBody {
	s.Data = &v
	return s
}

func (s *GetNetworkSelectResponseBody) SetTotalCount(v int32) *GetNetworkSelectResponseBody {
	s.TotalCount = &v
	return s
}


type GetNetworkSelectResponseBodyBuilder struct {
	s *GetNetworkSelectResponseBody
}

func NewGetNetworkSelectResponseBodyBuilder() *GetNetworkSelectResponseBodyBuilder {
	s := &GetNetworkSelectResponseBody{}
	b := &GetNetworkSelectResponseBodyBuilder{s: s}
	return b
}

func (b *GetNetworkSelectResponseBodyBuilder) Data(v []GetNetworkSelectResponseData) *GetNetworkSelectResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetNetworkSelectResponseBodyBuilder) TotalCount(v int32) *GetNetworkSelectResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetNetworkSelectResponseBodyBuilder) Build() *GetNetworkSelectResponseBody {
	return b.s
}


