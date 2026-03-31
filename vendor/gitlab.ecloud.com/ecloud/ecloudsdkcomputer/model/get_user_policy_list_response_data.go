// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserPolicyListResponseData struct {

    // 更新时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 策略来源 1 云管 2 BOP 3 CHBN 4企业门户
	SourceType *int32 `json:"sourceType,omitempty"`
    // 是否生效 0 启用 1禁用
	Enable *int32 `json:"enable,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 关联用户数
	BindUserCount *int32 `json:"bindUserCount,omitempty"`
    // 策略Id
	UserPolicyId *string `json:"userPolicyId,omitempty"`
}

func (s GetUserPolicyListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyListResponseData) GoString() string {
	return s.String()
}

func (s GetUserPolicyListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyListResponseData) SetModifiedTime(v string) *GetUserPolicyListResponseData {
	s.ModifiedTime = &v
	return s
}

func (s *GetUserPolicyListResponseData) SetPolicyName(v string) *GetUserPolicyListResponseData {
	s.PolicyName = &v
	return s
}

func (s *GetUserPolicyListResponseData) SetSourceType(v int32) *GetUserPolicyListResponseData {
	s.SourceType = &v
	return s
}

func (s *GetUserPolicyListResponseData) SetEnable(v int32) *GetUserPolicyListResponseData {
	s.Enable = &v
	return s
}

func (s *GetUserPolicyListResponseData) SetPolicyDesc(v string) *GetUserPolicyListResponseData {
	s.PolicyDesc = &v
	return s
}

func (s *GetUserPolicyListResponseData) SetBindUserCount(v int32) *GetUserPolicyListResponseData {
	s.BindUserCount = &v
	return s
}

func (s *GetUserPolicyListResponseData) SetUserPolicyId(v string) *GetUserPolicyListResponseData {
	s.UserPolicyId = &v
	return s
}


type GetUserPolicyListResponseDataBuilder struct {
	s *GetUserPolicyListResponseData
}

func NewGetUserPolicyListResponseDataBuilder() *GetUserPolicyListResponseDataBuilder {
	s := &GetUserPolicyListResponseData{}
	b := &GetUserPolicyListResponseDataBuilder{s: s}
	return b
}

func (b *GetUserPolicyListResponseDataBuilder) ModifiedTime(v string) *GetUserPolicyListResponseDataBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetUserPolicyListResponseDataBuilder) PolicyName(v string) *GetUserPolicyListResponseDataBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetUserPolicyListResponseDataBuilder) SourceType(v int32) *GetUserPolicyListResponseDataBuilder {
	b.s.SourceType = &v
	return b
}

func (b *GetUserPolicyListResponseDataBuilder) Enable(v int32) *GetUserPolicyListResponseDataBuilder {
	b.s.Enable = &v
	return b
}

func (b *GetUserPolicyListResponseDataBuilder) PolicyDesc(v string) *GetUserPolicyListResponseDataBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *GetUserPolicyListResponseDataBuilder) BindUserCount(v int32) *GetUserPolicyListResponseDataBuilder {
	b.s.BindUserCount = &v
	return b
}

func (b *GetUserPolicyListResponseDataBuilder) UserPolicyId(v string) *GetUserPolicyListResponseDataBuilder {
	b.s.UserPolicyId = &v
	return b
}

func (b *GetUserPolicyListResponseDataBuilder) Build() *GetUserPolicyListResponseData {
	return b.s
}


