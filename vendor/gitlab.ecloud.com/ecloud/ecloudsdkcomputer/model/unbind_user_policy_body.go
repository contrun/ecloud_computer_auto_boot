// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnbindUserPolicyBody struct {
    position.Body
    // 策略Id
	PolicyId *string `json:"policyId,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
}

func (s UnbindUserPolicyBody) String() string {
	return utils.Beautify(s)
}

func (s UnbindUserPolicyBody) GoString() string {
	return s.String()
}

func (s UnbindUserPolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnbindUserPolicyBody) SetPolicyId(v string) *UnbindUserPolicyBody {
	s.PolicyId = &v
	return s
}

func (s *UnbindUserPolicyBody) SetUserName(v string) *UnbindUserPolicyBody {
	s.UserName = &v
	return s
}


type UnbindUserPolicyBodyBuilder struct {
	s *UnbindUserPolicyBody
}

func NewUnbindUserPolicyBodyBuilder() *UnbindUserPolicyBodyBuilder {
	s := &UnbindUserPolicyBody{}
	b := &UnbindUserPolicyBodyBuilder{s: s}
	return b
}

func (b *UnbindUserPolicyBodyBuilder) PolicyId(v string) *UnbindUserPolicyBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *UnbindUserPolicyBodyBuilder) UserName(v string) *UnbindUserPolicyBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *UnbindUserPolicyBodyBuilder) Build() *UnbindUserPolicyBody {
	return b.s
}


