// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddPolicyRequest struct {


	AddPolicyBody *AddPolicyBody `json:"addPolicyBody,omitempty"`
}

func (s AddPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s AddPolicyRequest) GoString() string {
	return s.String()
}

func (s AddPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddPolicyRequest) SetAddPolicyBody(v *AddPolicyBody) *AddPolicyRequest {
	s.AddPolicyBody = v
	return s
}


type AddPolicyRequestBuilder struct {
	s *AddPolicyRequest
}

func NewAddPolicyRequestBuilder() *AddPolicyRequestBuilder {
	s := &AddPolicyRequest{}
	b := &AddPolicyRequestBuilder{s: s}
	return b
}

func (b *AddPolicyRequestBuilder) AddPolicyBody(v *AddPolicyBody) *AddPolicyRequestBuilder {
	b.s.AddPolicyBody = v
	return b
}

func (b *AddPolicyRequestBuilder) Build() *AddPolicyRequest {
	return b.s
}


