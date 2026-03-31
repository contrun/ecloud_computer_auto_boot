// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdatePolicyRequest struct {


	UpdatePolicyBody *UpdatePolicyBody `json:"updatePolicyBody,omitempty"`
}

func (s UpdatePolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdatePolicyRequest) GoString() string {
	return s.String()
}

func (s UpdatePolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdatePolicyRequest) SetUpdatePolicyBody(v *UpdatePolicyBody) *UpdatePolicyRequest {
	s.UpdatePolicyBody = v
	return s
}


type UpdatePolicyRequestBuilder struct {
	s *UpdatePolicyRequest
}

func NewUpdatePolicyRequestBuilder() *UpdatePolicyRequestBuilder {
	s := &UpdatePolicyRequest{}
	b := &UpdatePolicyRequestBuilder{s: s}
	return b
}

func (b *UpdatePolicyRequestBuilder) UpdatePolicyBody(v *UpdatePolicyBody) *UpdatePolicyRequestBuilder {
	b.s.UpdatePolicyBody = v
	return b
}

func (b *UpdatePolicyRequestBuilder) Build() *UpdatePolicyRequest {
	return b.s
}


