// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserPolicyListBody struct {
    position.Body
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	PageNum *int32 `json:"pageNum,omitempty"`
}

func (s GetUserPolicyListBody) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyListBody) GoString() string {
	return s.String()
}

func (s GetUserPolicyListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyListBody) SetPolicyName(v string) *GetUserPolicyListBody {
	s.PolicyName = &v
	return s
}

func (s *GetUserPolicyListBody) SetPolicyDesc(v string) *GetUserPolicyListBody {
	s.PolicyDesc = &v
	return s
}

func (s *GetUserPolicyListBody) SetPageSize(v int32) *GetUserPolicyListBody {
	s.PageSize = &v
	return s
}

func (s *GetUserPolicyListBody) SetPageNum(v int32) *GetUserPolicyListBody {
	s.PageNum = &v
	return s
}


type GetUserPolicyListBodyBuilder struct {
	s *GetUserPolicyListBody
}

func NewGetUserPolicyListBodyBuilder() *GetUserPolicyListBodyBuilder {
	s := &GetUserPolicyListBody{}
	b := &GetUserPolicyListBodyBuilder{s: s}
	return b
}

func (b *GetUserPolicyListBodyBuilder) PolicyName(v string) *GetUserPolicyListBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetUserPolicyListBodyBuilder) PolicyDesc(v string) *GetUserPolicyListBodyBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *GetUserPolicyListBodyBuilder) PageSize(v int32) *GetUserPolicyListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetUserPolicyListBodyBuilder) PageNum(v int32) *GetUserPolicyListBodyBuilder {
	b.s.PageNum = &v
	return b
}

func (b *GetUserPolicyListBodyBuilder) Build() *GetUserPolicyListBody {
	return b.s
}


