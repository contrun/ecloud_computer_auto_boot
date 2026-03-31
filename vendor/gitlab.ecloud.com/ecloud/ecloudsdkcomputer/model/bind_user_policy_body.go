// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindUserPolicyBody struct {
    position.Body
    // 策略Id
	PolicyId *string `json:"policyId,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
}

func (s BindUserPolicyBody) String() string {
	return utils.Beautify(s)
}

func (s BindUserPolicyBody) GoString() string {
	return s.String()
}

func (s BindUserPolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindUserPolicyBody) SetPolicyId(v string) *BindUserPolicyBody {
	s.PolicyId = &v
	return s
}

func (s *BindUserPolicyBody) SetUserName(v string) *BindUserPolicyBody {
	s.UserName = &v
	return s
}


type BindUserPolicyBodyBuilder struct {
	s *BindUserPolicyBody
}

func NewBindUserPolicyBodyBuilder() *BindUserPolicyBodyBuilder {
	s := &BindUserPolicyBody{}
	b := &BindUserPolicyBodyBuilder{s: s}
	return b
}

func (b *BindUserPolicyBodyBuilder) PolicyId(v string) *BindUserPolicyBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *BindUserPolicyBodyBuilder) UserName(v string) *BindUserPolicyBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *BindUserPolicyBodyBuilder) Build() *BindUserPolicyBody {
	return b.s
}


