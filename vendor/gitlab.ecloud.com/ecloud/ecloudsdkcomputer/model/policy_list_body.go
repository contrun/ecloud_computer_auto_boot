// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PolicyListBody struct {
    position.Body
    // 定时任务uid
	PolicyUid *string `json:"policyUid,omitempty"`
    // 定时任务名称
	PolicyName *string `json:"policyName,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s PolicyListBody) String() string {
	return utils.Beautify(s)
}

func (s PolicyListBody) GoString() string {
	return s.String()
}

func (s PolicyListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyListBody) SetPolicyUid(v string) *PolicyListBody {
	s.PolicyUid = &v
	return s
}

func (s *PolicyListBody) SetPolicyName(v string) *PolicyListBody {
	s.PolicyName = &v
	return s
}

func (s *PolicyListBody) SetPageSize(v int32) *PolicyListBody {
	s.PageSize = &v
	return s
}

func (s *PolicyListBody) SetPoolCode(v string) *PolicyListBody {
	s.PoolCode = &v
	return s
}

func (s *PolicyListBody) SetCurrentPage(v int32) *PolicyListBody {
	s.CurrentPage = &v
	return s
}


type PolicyListBodyBuilder struct {
	s *PolicyListBody
}

func NewPolicyListBodyBuilder() *PolicyListBodyBuilder {
	s := &PolicyListBody{}
	b := &PolicyListBodyBuilder{s: s}
	return b
}

func (b *PolicyListBodyBuilder) PolicyUid(v string) *PolicyListBodyBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *PolicyListBodyBuilder) PolicyName(v string) *PolicyListBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *PolicyListBodyBuilder) PageSize(v int32) *PolicyListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *PolicyListBodyBuilder) PoolCode(v string) *PolicyListBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *PolicyListBodyBuilder) CurrentPage(v int32) *PolicyListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *PolicyListBodyBuilder) Build() *PolicyListBody {
	return b.s
}


