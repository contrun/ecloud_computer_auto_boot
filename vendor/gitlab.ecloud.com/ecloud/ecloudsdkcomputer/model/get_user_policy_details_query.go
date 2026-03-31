// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserPolicyDetailsQuery struct {
    position.Query
    // 策略ID
	PolicyId *string `json:"policyId,omitempty"`
}

func (s GetUserPolicyDetailsQuery) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyDetailsQuery) GoString() string {
	return s.String()
}

func (s GetUserPolicyDetailsQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyDetailsQuery) SetPolicyId(v string) *GetUserPolicyDetailsQuery {
	s.PolicyId = &v
	return s
}


type GetUserPolicyDetailsQueryBuilder struct {
	s *GetUserPolicyDetailsQuery
}

func NewGetUserPolicyDetailsQueryBuilder() *GetUserPolicyDetailsQueryBuilder {
	s := &GetUserPolicyDetailsQuery{}
	b := &GetUserPolicyDetailsQueryBuilder{s: s}
	return b
}

func (b *GetUserPolicyDetailsQueryBuilder) PolicyId(v string) *GetUserPolicyDetailsQueryBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetUserPolicyDetailsQueryBuilder) Build() *GetUserPolicyDetailsQuery {
	return b.s
}


