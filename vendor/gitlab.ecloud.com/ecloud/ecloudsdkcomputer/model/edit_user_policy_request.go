// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EditUserPolicyRequest struct {


	EditUserPolicyBody *EditUserPolicyBody `json:"editUserPolicyBody,omitempty"`
}

func (s EditUserPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s EditUserPolicyRequest) GoString() string {
	return s.String()
}

func (s EditUserPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditUserPolicyRequest) SetEditUserPolicyBody(v *EditUserPolicyBody) *EditUserPolicyRequest {
	s.EditUserPolicyBody = v
	return s
}


type EditUserPolicyRequestBuilder struct {
	s *EditUserPolicyRequest
}

func NewEditUserPolicyRequestBuilder() *EditUserPolicyRequestBuilder {
	s := &EditUserPolicyRequest{}
	b := &EditUserPolicyRequestBuilder{s: s}
	return b
}

func (b *EditUserPolicyRequestBuilder) EditUserPolicyBody(v *EditUserPolicyBody) *EditUserPolicyRequestBuilder {
	b.s.EditUserPolicyBody = v
	return b
}

func (b *EditUserPolicyRequestBuilder) Build() *EditUserPolicyRequest {
	return b.s
}


