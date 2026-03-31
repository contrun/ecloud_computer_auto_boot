// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchUnbindPolicyRequest struct {


	BatchUnbindPolicyBody *BatchUnbindPolicyBody `json:"batchUnbindPolicyBody,omitempty"`
}

func (s BatchUnbindPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchUnbindPolicyRequest) GoString() string {
	return s.String()
}

func (s BatchUnbindPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUnbindPolicyRequest) SetBatchUnbindPolicyBody(v *BatchUnbindPolicyBody) *BatchUnbindPolicyRequest {
	s.BatchUnbindPolicyBody = v
	return s
}


type BatchUnbindPolicyRequestBuilder struct {
	s *BatchUnbindPolicyRequest
}

func NewBatchUnbindPolicyRequestBuilder() *BatchUnbindPolicyRequestBuilder {
	s := &BatchUnbindPolicyRequest{}
	b := &BatchUnbindPolicyRequestBuilder{s: s}
	return b
}

func (b *BatchUnbindPolicyRequestBuilder) BatchUnbindPolicyBody(v *BatchUnbindPolicyBody) *BatchUnbindPolicyRequestBuilder {
	b.s.BatchUnbindPolicyBody = v
	return b
}

func (b *BatchUnbindPolicyRequestBuilder) Build() *BatchUnbindPolicyRequest {
	return b.s
}


