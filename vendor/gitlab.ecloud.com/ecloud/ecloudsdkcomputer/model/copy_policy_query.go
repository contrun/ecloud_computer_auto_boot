// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CopyPolicyQuery struct {
    position.Query
    // 定时任务策略uuid
	PolicyUid *string `json:"policyUid,omitempty"`
}

func (s CopyPolicyQuery) String() string {
	return utils.Beautify(s)
}

func (s CopyPolicyQuery) GoString() string {
	return s.String()
}

func (s CopyPolicyQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CopyPolicyQuery) SetPolicyUid(v string) *CopyPolicyQuery {
	s.PolicyUid = &v
	return s
}


type CopyPolicyQueryBuilder struct {
	s *CopyPolicyQuery
}

func NewCopyPolicyQueryBuilder() *CopyPolicyQueryBuilder {
	s := &CopyPolicyQuery{}
	b := &CopyPolicyQueryBuilder{s: s}
	return b
}

func (b *CopyPolicyQueryBuilder) PolicyUid(v string) *CopyPolicyQueryBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *CopyPolicyQueryBuilder) Build() *CopyPolicyQuery {
	return b.s
}


