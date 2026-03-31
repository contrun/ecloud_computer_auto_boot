// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteUserPolicyQuery struct {
    position.Query
    // 策略ID
	PolicyId *string `json:"policyId,omitempty"`
}

func (s DeleteUserPolicyQuery) String() string {
	return utils.Beautify(s)
}

func (s DeleteUserPolicyQuery) GoString() string {
	return s.String()
}

func (s DeleteUserPolicyQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteUserPolicyQuery) SetPolicyId(v string) *DeleteUserPolicyQuery {
	s.PolicyId = &v
	return s
}


type DeleteUserPolicyQueryBuilder struct {
	s *DeleteUserPolicyQuery
}

func NewDeleteUserPolicyQueryBuilder() *DeleteUserPolicyQueryBuilder {
	s := &DeleteUserPolicyQuery{}
	b := &DeleteUserPolicyQueryBuilder{s: s}
	return b
}

func (b *DeleteUserPolicyQueryBuilder) PolicyId(v string) *DeleteUserPolicyQueryBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *DeleteUserPolicyQueryBuilder) Build() *DeleteUserPolicyQuery {
	return b.s
}


