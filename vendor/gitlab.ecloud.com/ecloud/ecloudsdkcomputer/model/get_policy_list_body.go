// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyListBody struct {
    position.Body
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 所属资源池
	PoolCode *string `json:"poolCode,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 策略状态
	PolicyStatus *string `json:"policyStatus,omitempty"`
    // 所属可用区
	Region *string `json:"region,omitempty"`
    // 网络工作区，V2.0协议非必传
	Dc *string `json:"dc,omitempty"`
}

func (s GetPolicyListBody) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyListBody) GoString() string {
	return s.String()
}

func (s GetPolicyListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyListBody) SetPolicyName(v string) *GetPolicyListBody {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyListBody) SetPageSize(v int32) *GetPolicyListBody {
	s.PageSize = &v
	return s
}

func (s *GetPolicyListBody) SetPoolCode(v string) *GetPolicyListBody {
	s.PoolCode = &v
	return s
}

func (s *GetPolicyListBody) SetCurrentPage(v int32) *GetPolicyListBody {
	s.CurrentPage = &v
	return s
}

func (s *GetPolicyListBody) SetPolicyStatus(v string) *GetPolicyListBody {
	s.PolicyStatus = &v
	return s
}

func (s *GetPolicyListBody) SetRegion(v string) *GetPolicyListBody {
	s.Region = &v
	return s
}

func (s *GetPolicyListBody) SetDc(v string) *GetPolicyListBody {
	s.Dc = &v
	return s
}


type GetPolicyListBodyBuilder struct {
	s *GetPolicyListBody
}

func NewGetPolicyListBodyBuilder() *GetPolicyListBodyBuilder {
	s := &GetPolicyListBody{}
	b := &GetPolicyListBodyBuilder{s: s}
	return b
}

func (b *GetPolicyListBodyBuilder) PolicyName(v string) *GetPolicyListBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetPolicyListBodyBuilder) PageSize(v int32) *GetPolicyListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetPolicyListBodyBuilder) PoolCode(v string) *GetPolicyListBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetPolicyListBodyBuilder) CurrentPage(v int32) *GetPolicyListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetPolicyListBodyBuilder) PolicyStatus(v string) *GetPolicyListBodyBuilder {
	b.s.PolicyStatus = &v
	return b
}

func (b *GetPolicyListBodyBuilder) Region(v string) *GetPolicyListBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *GetPolicyListBodyBuilder) Dc(v string) *GetPolicyListBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *GetPolicyListBodyBuilder) Build() *GetPolicyListBody {
	return b.s
}


