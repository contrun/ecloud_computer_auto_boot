// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindUserPolicyRequest struct {


	BindUserPolicyBody *BindUserPolicyBody `json:"bindUserPolicyBody,omitempty"`
}

func (s BindUserPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s BindUserPolicyRequest) GoString() string {
	return s.String()
}

func (s BindUserPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindUserPolicyRequest) SetBindUserPolicyBody(v *BindUserPolicyBody) *BindUserPolicyRequest {
	s.BindUserPolicyBody = v
	return s
}


type BindUserPolicyRequestBuilder struct {
	s *BindUserPolicyRequest
}

func NewBindUserPolicyRequestBuilder() *BindUserPolicyRequestBuilder {
	s := &BindUserPolicyRequest{}
	b := &BindUserPolicyRequestBuilder{s: s}
	return b
}

func (b *BindUserPolicyRequestBuilder) BindUserPolicyBody(v *BindUserPolicyBody) *BindUserPolicyRequestBuilder {
	b.s.BindUserPolicyBody = v
	return b
}

func (b *BindUserPolicyRequestBuilder) Build() *BindUserPolicyRequest {
	return b.s
}


