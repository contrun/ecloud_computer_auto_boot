// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSNatRuleResponseBody struct {

    // 数据
	Data *[]ListSNatRuleResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s ListSNatRuleResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListSNatRuleResponseBody) GoString() string {
	return s.String()
}

func (s ListSNatRuleResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSNatRuleResponseBody) SetData(v []ListSNatRuleResponseData) *ListSNatRuleResponseBody {
	s.Data = &v
	return s
}

func (s *ListSNatRuleResponseBody) SetTotalCount(v int32) *ListSNatRuleResponseBody {
	s.TotalCount = &v
	return s
}


type ListSNatRuleResponseBodyBuilder struct {
	s *ListSNatRuleResponseBody
}

func NewListSNatRuleResponseBodyBuilder() *ListSNatRuleResponseBodyBuilder {
	s := &ListSNatRuleResponseBody{}
	b := &ListSNatRuleResponseBodyBuilder{s: s}
	return b
}

func (b *ListSNatRuleResponseBodyBuilder) Data(v []ListSNatRuleResponseData) *ListSNatRuleResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *ListSNatRuleResponseBodyBuilder) TotalCount(v int32) *ListSNatRuleResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *ListSNatRuleResponseBodyBuilder) Build() *ListSNatRuleResponseBody {
	return b.s
}


