// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeletePolicyQuery struct {
    position.Query
    // 定时任务策略uid
	PolicyUid *string `json:"policyUid,omitempty"`
}

func (s DeletePolicyQuery) String() string {
	return utils.Beautify(s)
}

func (s DeletePolicyQuery) GoString() string {
	return s.String()
}

func (s DeletePolicyQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePolicyQuery) SetPolicyUid(v string) *DeletePolicyQuery {
	s.PolicyUid = &v
	return s
}


type DeletePolicyQueryBuilder struct {
	s *DeletePolicyQuery
}

func NewDeletePolicyQueryBuilder() *DeletePolicyQueryBuilder {
	s := &DeletePolicyQuery{}
	b := &DeletePolicyQueryBuilder{s: s}
	return b
}

func (b *DeletePolicyQueryBuilder) PolicyUid(v string) *DeletePolicyQueryBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *DeletePolicyQueryBuilder) Build() *DeletePolicyQuery {
	return b.s
}


