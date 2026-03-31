// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdQuery struct {
    position.Query
    // 电脑策略id
	PolicyId *string `json:"policyId,omitempty"`
}

func (s GetPolicyInfoByIdQuery) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdQuery) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdQuery) SetPolicyId(v string) *GetPolicyInfoByIdQuery {
	s.PolicyId = &v
	return s
}


type GetPolicyInfoByIdQueryBuilder struct {
	s *GetPolicyInfoByIdQuery
}

func NewGetPolicyInfoByIdQueryBuilder() *GetPolicyInfoByIdQueryBuilder {
	s := &GetPolicyInfoByIdQuery{}
	b := &GetPolicyInfoByIdQueryBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdQueryBuilder) PolicyId(v string) *GetPolicyInfoByIdQueryBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetPolicyInfoByIdQueryBuilder) Build() *GetPolicyInfoByIdQuery {
	return b.s
}


