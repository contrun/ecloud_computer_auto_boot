// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PolicyDescQuery struct {
    position.Query
    // 定时任务策略uuid
	PolicyUid *string `json:"policyUid,omitempty"`
}

func (s PolicyDescQuery) String() string {
	return utils.Beautify(s)
}

func (s PolicyDescQuery) GoString() string {
	return s.String()
}

func (s PolicyDescQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyDescQuery) SetPolicyUid(v string) *PolicyDescQuery {
	s.PolicyUid = &v
	return s
}


type PolicyDescQueryBuilder struct {
	s *PolicyDescQuery
}

func NewPolicyDescQueryBuilder() *PolicyDescQueryBuilder {
	s := &PolicyDescQuery{}
	b := &PolicyDescQueryBuilder{s: s}
	return b
}

func (b *PolicyDescQueryBuilder) PolicyUid(v string) *PolicyDescQueryBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *PolicyDescQueryBuilder) Build() *PolicyDescQuery {
	return b.s
}


