// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddUserPolicyRequest struct {


	AddUserPolicyBody *AddUserPolicyBody `json:"addUserPolicyBody,omitempty"`
}

func (s AddUserPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s AddUserPolicyRequest) GoString() string {
	return s.String()
}

func (s AddUserPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddUserPolicyRequest) SetAddUserPolicyBody(v *AddUserPolicyBody) *AddUserPolicyRequest {
	s.AddUserPolicyBody = v
	return s
}


type AddUserPolicyRequestBuilder struct {
	s *AddUserPolicyRequest
}

func NewAddUserPolicyRequestBuilder() *AddUserPolicyRequestBuilder {
	s := &AddUserPolicyRequest{}
	b := &AddUserPolicyRequestBuilder{s: s}
	return b
}

func (b *AddUserPolicyRequestBuilder) AddUserPolicyBody(v *AddUserPolicyBody) *AddUserPolicyRequestBuilder {
	b.s.AddUserPolicyBody = v
	return b
}

func (b *AddUserPolicyRequestBuilder) Build() *AddUserPolicyRequest {
	return b.s
}


