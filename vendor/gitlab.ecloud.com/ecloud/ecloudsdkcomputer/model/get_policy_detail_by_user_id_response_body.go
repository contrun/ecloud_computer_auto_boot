// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyDetailByUserIdResponseBody struct {

    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 企业门户组织id
	CcempOrganizeId *string `json:"ccempOrganizeId,omitempty"`
    // 是否生效,0:启用,1:禁用
	Enable *int32 `json:"enable,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 策略id
	UserPolicyId *string `json:"userPolicyId,omitempty"`
}

func (s GetPolicyDetailByUserIdResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyDetailByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s GetPolicyDetailByUserIdResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyDetailByUserIdResponseBody) SetPolicyName(v string) *GetPolicyDetailByUserIdResponseBody {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyDetailByUserIdResponseBody) SetCcempOrganizeId(v string) *GetPolicyDetailByUserIdResponseBody {
	s.CcempOrganizeId = &v
	return s
}

func (s *GetPolicyDetailByUserIdResponseBody) SetEnable(v int32) *GetPolicyDetailByUserIdResponseBody {
	s.Enable = &v
	return s
}

func (s *GetPolicyDetailByUserIdResponseBody) SetPolicyDesc(v string) *GetPolicyDetailByUserIdResponseBody {
	s.PolicyDesc = &v
	return s
}

func (s *GetPolicyDetailByUserIdResponseBody) SetUserPolicyId(v string) *GetPolicyDetailByUserIdResponseBody {
	s.UserPolicyId = &v
	return s
}


type GetPolicyDetailByUserIdResponseBodyBuilder struct {
	s *GetPolicyDetailByUserIdResponseBody
}

func NewGetPolicyDetailByUserIdResponseBodyBuilder() *GetPolicyDetailByUserIdResponseBodyBuilder {
	s := &GetPolicyDetailByUserIdResponseBody{}
	b := &GetPolicyDetailByUserIdResponseBodyBuilder{s: s}
	return b
}

func (b *GetPolicyDetailByUserIdResponseBodyBuilder) PolicyName(v string) *GetPolicyDetailByUserIdResponseBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBodyBuilder) CcempOrganizeId(v string) *GetPolicyDetailByUserIdResponseBodyBuilder {
	b.s.CcempOrganizeId = &v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBodyBuilder) Enable(v int32) *GetPolicyDetailByUserIdResponseBodyBuilder {
	b.s.Enable = &v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBodyBuilder) PolicyDesc(v string) *GetPolicyDetailByUserIdResponseBodyBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBodyBuilder) UserPolicyId(v string) *GetPolicyDetailByUserIdResponseBodyBuilder {
	b.s.UserPolicyId = &v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBodyBuilder) Build() *GetPolicyDetailByUserIdResponseBody {
	return b.s
}


