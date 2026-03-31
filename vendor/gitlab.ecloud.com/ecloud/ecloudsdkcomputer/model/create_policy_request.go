// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreatePolicyRequest struct {


	CreatePolicyBody *CreatePolicyBody `json:"createPolicyBody,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s CreatePolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequest) SetCreatePolicyBody(v *CreatePolicyBody) *CreatePolicyRequest {
	s.CreatePolicyBody = v
	return s
}


type CreatePolicyRequestBuilder struct {
	s *CreatePolicyRequest
}

func NewCreatePolicyRequestBuilder() *CreatePolicyRequestBuilder {
	s := &CreatePolicyRequest{}
	b := &CreatePolicyRequestBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestBuilder) CreatePolicyBody(v *CreatePolicyBody) *CreatePolicyRequestBuilder {
	b.s.CreatePolicyBody = v
	return b
}

func (b *CreatePolicyRequestBuilder) Build() *CreatePolicyRequest {
	return b.s
}


