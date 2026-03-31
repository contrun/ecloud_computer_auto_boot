// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserPolicyListResponseBody struct {

    // 数据
	Data *[]GetUserPolicyListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetUserPolicyListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyListResponseBody) GoString() string {
	return s.String()
}

func (s GetUserPolicyListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyListResponseBody) SetData(v []GetUserPolicyListResponseData) *GetUserPolicyListResponseBody {
	s.Data = &v
	return s
}

func (s *GetUserPolicyListResponseBody) SetTotalCount(v int32) *GetUserPolicyListResponseBody {
	s.TotalCount = &v
	return s
}


type GetUserPolicyListResponseBodyBuilder struct {
	s *GetUserPolicyListResponseBody
}

func NewGetUserPolicyListResponseBodyBuilder() *GetUserPolicyListResponseBodyBuilder {
	s := &GetUserPolicyListResponseBody{}
	b := &GetUserPolicyListResponseBodyBuilder{s: s}
	return b
}

func (b *GetUserPolicyListResponseBodyBuilder) Data(v []GetUserPolicyListResponseData) *GetUserPolicyListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetUserPolicyListResponseBodyBuilder) TotalCount(v int32) *GetUserPolicyListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetUserPolicyListResponseBodyBuilder) Build() *GetUserPolicyListResponseBody {
	return b.s
}


