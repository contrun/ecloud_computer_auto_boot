// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PolicyListResponseBody struct {

    // 数据
	Data *[]PolicyListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s PolicyListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s PolicyListResponseBody) GoString() string {
	return s.String()
}

func (s PolicyListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyListResponseBody) SetData(v []PolicyListResponseData) *PolicyListResponseBody {
	s.Data = &v
	return s
}

func (s *PolicyListResponseBody) SetTotalCount(v int32) *PolicyListResponseBody {
	s.TotalCount = &v
	return s
}


type PolicyListResponseBodyBuilder struct {
	s *PolicyListResponseBody
}

func NewPolicyListResponseBodyBuilder() *PolicyListResponseBodyBuilder {
	s := &PolicyListResponseBody{}
	b := &PolicyListResponseBodyBuilder{s: s}
	return b
}

func (b *PolicyListResponseBodyBuilder) Data(v []PolicyListResponseData) *PolicyListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *PolicyListResponseBodyBuilder) TotalCount(v int32) *PolicyListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *PolicyListResponseBodyBuilder) Build() *PolicyListResponseBody {
	return b.s
}


