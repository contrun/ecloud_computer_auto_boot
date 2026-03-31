// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CopyPolicyRequest struct {


	CopyPolicyQuery *CopyPolicyQuery `json:"copyPolicyQuery,omitempty"`
}

func (s CopyPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s CopyPolicyRequest) GoString() string {
	return s.String()
}

func (s CopyPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CopyPolicyRequest) SetCopyPolicyQuery(v *CopyPolicyQuery) *CopyPolicyRequest {
	s.CopyPolicyQuery = v
	return s
}


type CopyPolicyRequestBuilder struct {
	s *CopyPolicyRequest
}

func NewCopyPolicyRequestBuilder() *CopyPolicyRequestBuilder {
	s := &CopyPolicyRequest{}
	b := &CopyPolicyRequestBuilder{s: s}
	return b
}

func (b *CopyPolicyRequestBuilder) CopyPolicyQuery(v *CopyPolicyQuery) *CopyPolicyRequestBuilder {
	b.s.CopyPolicyQuery = v
	return b
}

func (b *CopyPolicyRequestBuilder) Build() *CopyPolicyRequest {
	return b.s
}


