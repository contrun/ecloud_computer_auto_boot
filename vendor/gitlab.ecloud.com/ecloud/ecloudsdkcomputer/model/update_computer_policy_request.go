// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateComputerPolicyRequest struct {


	UpdateComputerPolicyBody *UpdateComputerPolicyBody `json:"updateComputerPolicyBody,omitempty"`
}

func (s UpdateComputerPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequest) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequest) SetUpdateComputerPolicyBody(v *UpdateComputerPolicyBody) *UpdateComputerPolicyRequest {
	s.UpdateComputerPolicyBody = v
	return s
}


type UpdateComputerPolicyRequestBuilder struct {
	s *UpdateComputerPolicyRequest
}

func NewUpdateComputerPolicyRequestBuilder() *UpdateComputerPolicyRequestBuilder {
	s := &UpdateComputerPolicyRequest{}
	b := &UpdateComputerPolicyRequestBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestBuilder) UpdateComputerPolicyBody(v *UpdateComputerPolicyBody) *UpdateComputerPolicyRequestBuilder {
	b.s.UpdateComputerPolicyBody = v
	return b
}

func (b *UpdateComputerPolicyRequestBuilder) Build() *UpdateComputerPolicyRequest {
	return b.s
}


