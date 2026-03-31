// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteComputerPolicyRequest struct {


	DeleteComputerPolicyBody *DeleteComputerPolicyBody `json:"deleteComputerPolicyBody,omitempty"`
}

func (s DeleteComputerPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteComputerPolicyRequest) GoString() string {
	return s.String()
}

func (s DeleteComputerPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteComputerPolicyRequest) SetDeleteComputerPolicyBody(v *DeleteComputerPolicyBody) *DeleteComputerPolicyRequest {
	s.DeleteComputerPolicyBody = v
	return s
}


type DeleteComputerPolicyRequestBuilder struct {
	s *DeleteComputerPolicyRequest
}

func NewDeleteComputerPolicyRequestBuilder() *DeleteComputerPolicyRequestBuilder {
	s := &DeleteComputerPolicyRequest{}
	b := &DeleteComputerPolicyRequestBuilder{s: s}
	return b
}

func (b *DeleteComputerPolicyRequestBuilder) DeleteComputerPolicyBody(v *DeleteComputerPolicyBody) *DeleteComputerPolicyRequestBuilder {
	b.s.DeleteComputerPolicyBody = v
	return b
}

func (b *DeleteComputerPolicyRequestBuilder) Build() *DeleteComputerPolicyRequest {
	return b.s
}


