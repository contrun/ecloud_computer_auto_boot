// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetUserPolicyStatusBody struct {
    position.Body
    // 策略是否开启, 0：启用，1：禁用
	Enable *int32 `json:"enable,omitempty"`
    // 用户策略id
	UserPolicyId *string `json:"userPolicyId,omitempty"`
}

func (s SetUserPolicyStatusBody) String() string {
	return utils.Beautify(s)
}

func (s SetUserPolicyStatusBody) GoString() string {
	return s.String()
}

func (s SetUserPolicyStatusBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetUserPolicyStatusBody) SetEnable(v int32) *SetUserPolicyStatusBody {
	s.Enable = &v
	return s
}

func (s *SetUserPolicyStatusBody) SetUserPolicyId(v string) *SetUserPolicyStatusBody {
	s.UserPolicyId = &v
	return s
}


type SetUserPolicyStatusBodyBuilder struct {
	s *SetUserPolicyStatusBody
}

func NewSetUserPolicyStatusBodyBuilder() *SetUserPolicyStatusBodyBuilder {
	s := &SetUserPolicyStatusBody{}
	b := &SetUserPolicyStatusBodyBuilder{s: s}
	return b
}

func (b *SetUserPolicyStatusBodyBuilder) Enable(v int32) *SetUserPolicyStatusBodyBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetUserPolicyStatusBodyBuilder) UserPolicyId(v string) *SetUserPolicyStatusBodyBuilder {
	b.s.UserPolicyId = &v
	return b
}

func (b *SetUserPolicyStatusBodyBuilder) Build() *SetUserPolicyStatusBody {
	return b.s
}


