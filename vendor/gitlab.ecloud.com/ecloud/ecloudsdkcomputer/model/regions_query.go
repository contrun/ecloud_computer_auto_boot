// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RegionsQuery struct {
    position.Query
    // 编码，查询省份无需传值，查询市区传上级编码
	Code *string `json:"code,omitempty"`
}

func (s RegionsQuery) String() string {
	return utils.Beautify(s)
}

func (s RegionsQuery) GoString() string {
	return s.String()
}

func (s RegionsQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RegionsQuery) SetCode(v string) *RegionsQuery {
	s.Code = &v
	return s
}


type RegionsQueryBuilder struct {
	s *RegionsQuery
}

func NewRegionsQueryBuilder() *RegionsQueryBuilder {
	s := &RegionsQuery{}
	b := &RegionsQueryBuilder{s: s}
	return b
}

func (b *RegionsQueryBuilder) Code(v string) *RegionsQueryBuilder {
	b.s.Code = &v
	return b
}

func (b *RegionsQueryBuilder) Build() *RegionsQuery {
	return b.s
}


