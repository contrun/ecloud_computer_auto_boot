// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchBindPolicyRequest struct {


	BatchBindPolicyBody *BatchBindPolicyBody `json:"batchBindPolicyBody,omitempty"`
}

func (s BatchBindPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchBindPolicyRequest) GoString() string {
	return s.String()
}

func (s BatchBindPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchBindPolicyRequest) SetBatchBindPolicyBody(v *BatchBindPolicyBody) *BatchBindPolicyRequest {
	s.BatchBindPolicyBody = v
	return s
}


type BatchBindPolicyRequestBuilder struct {
	s *BatchBindPolicyRequest
}

func NewBatchBindPolicyRequestBuilder() *BatchBindPolicyRequestBuilder {
	s := &BatchBindPolicyRequest{}
	b := &BatchBindPolicyRequestBuilder{s: s}
	return b
}

func (b *BatchBindPolicyRequestBuilder) BatchBindPolicyBody(v *BatchBindPolicyBody) *BatchBindPolicyRequestBuilder {
	b.s.BatchBindPolicyBody = v
	return b
}

func (b *BatchBindPolicyRequestBuilder) Build() *BatchBindPolicyRequest {
	return b.s
}


