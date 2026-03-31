// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyDetailByUserIdQuery struct {
    position.Query
    // 用户ID
	QueryUserId *string `json:"queryUserId,omitempty"`
}

func (s GetPolicyDetailByUserIdQuery) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyDetailByUserIdQuery) GoString() string {
	return s.String()
}

func (s GetPolicyDetailByUserIdQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyDetailByUserIdQuery) SetQueryUserId(v string) *GetPolicyDetailByUserIdQuery {
	s.QueryUserId = &v
	return s
}


type GetPolicyDetailByUserIdQueryBuilder struct {
	s *GetPolicyDetailByUserIdQuery
}

func NewGetPolicyDetailByUserIdQueryBuilder() *GetPolicyDetailByUserIdQueryBuilder {
	s := &GetPolicyDetailByUserIdQuery{}
	b := &GetPolicyDetailByUserIdQueryBuilder{s: s}
	return b
}

func (b *GetPolicyDetailByUserIdQueryBuilder) QueryUserId(v string) *GetPolicyDetailByUserIdQueryBuilder {
	b.s.QueryUserId = &v
	return b
}

func (b *GetPolicyDetailByUserIdQueryBuilder) Build() *GetPolicyDetailByUserIdQuery {
	return b.s
}


