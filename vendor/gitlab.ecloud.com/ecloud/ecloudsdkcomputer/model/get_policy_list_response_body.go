// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyListResponseBody struct {

    // 数据
	Data *[]GetPolicyListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetPolicyListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyListResponseBody) GoString() string {
	return s.String()
}

func (s GetPolicyListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyListResponseBody) SetData(v []GetPolicyListResponseData) *GetPolicyListResponseBody {
	s.Data = &v
	return s
}

func (s *GetPolicyListResponseBody) SetTotalCount(v int32) *GetPolicyListResponseBody {
	s.TotalCount = &v
	return s
}


type GetPolicyListResponseBodyBuilder struct {
	s *GetPolicyListResponseBody
}

func NewGetPolicyListResponseBodyBuilder() *GetPolicyListResponseBodyBuilder {
	s := &GetPolicyListResponseBody{}
	b := &GetPolicyListResponseBodyBuilder{s: s}
	return b
}

func (b *GetPolicyListResponseBodyBuilder) Data(v []GetPolicyListResponseData) *GetPolicyListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetPolicyListResponseBodyBuilder) TotalCount(v int32) *GetPolicyListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetPolicyListResponseBodyBuilder) Build() *GetPolicyListResponseBody {
	return b.s
}


