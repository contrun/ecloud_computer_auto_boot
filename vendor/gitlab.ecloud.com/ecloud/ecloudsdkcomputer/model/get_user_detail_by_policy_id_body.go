// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserDetailByPolicyIdBody struct {
    position.Body
    // 用户策略主键id
	PolicyId *string `json:"policyId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s GetUserDetailByPolicyIdBody) String() string {
	return utils.Beautify(s)
}

func (s GetUserDetailByPolicyIdBody) GoString() string {
	return s.String()
}

func (s GetUserDetailByPolicyIdBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserDetailByPolicyIdBody) SetPolicyId(v string) *GetUserDetailByPolicyIdBody {
	s.PolicyId = &v
	return s
}

func (s *GetUserDetailByPolicyIdBody) SetPageSize(v int32) *GetUserDetailByPolicyIdBody {
	s.PageSize = &v
	return s
}

func (s *GetUserDetailByPolicyIdBody) SetCurrentPage(v int32) *GetUserDetailByPolicyIdBody {
	s.CurrentPage = &v
	return s
}


type GetUserDetailByPolicyIdBodyBuilder struct {
	s *GetUserDetailByPolicyIdBody
}

func NewGetUserDetailByPolicyIdBodyBuilder() *GetUserDetailByPolicyIdBodyBuilder {
	s := &GetUserDetailByPolicyIdBody{}
	b := &GetUserDetailByPolicyIdBodyBuilder{s: s}
	return b
}

func (b *GetUserDetailByPolicyIdBodyBuilder) PolicyId(v string) *GetUserDetailByPolicyIdBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetUserDetailByPolicyIdBodyBuilder) PageSize(v int32) *GetUserDetailByPolicyIdBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetUserDetailByPolicyIdBodyBuilder) CurrentPage(v int32) *GetUserDetailByPolicyIdBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetUserDetailByPolicyIdBodyBuilder) Build() *GetUserDetailByPolicyIdBody {
	return b.s
}


