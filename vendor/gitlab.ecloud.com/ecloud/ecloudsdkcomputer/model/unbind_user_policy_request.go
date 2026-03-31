// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnbindUserPolicyRequest struct {


	UnbindUserPolicyBody *UnbindUserPolicyBody `json:"unbindUserPolicyBody,omitempty"`
}

func (s UnbindUserPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s UnbindUserPolicyRequest) GoString() string {
	return s.String()
}

func (s UnbindUserPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnbindUserPolicyRequest) SetUnbindUserPolicyBody(v *UnbindUserPolicyBody) *UnbindUserPolicyRequest {
	s.UnbindUserPolicyBody = v
	return s
}


type UnbindUserPolicyRequestBuilder struct {
	s *UnbindUserPolicyRequest
}

func NewUnbindUserPolicyRequestBuilder() *UnbindUserPolicyRequestBuilder {
	s := &UnbindUserPolicyRequest{}
	b := &UnbindUserPolicyRequestBuilder{s: s}
	return b
}

func (b *UnbindUserPolicyRequestBuilder) UnbindUserPolicyBody(v *UnbindUserPolicyBody) *UnbindUserPolicyRequestBuilder {
	b.s.UnbindUserPolicyBody = v
	return b
}

func (b *UnbindUserPolicyRequestBuilder) Build() *UnbindUserPolicyRequest {
	return b.s
}


