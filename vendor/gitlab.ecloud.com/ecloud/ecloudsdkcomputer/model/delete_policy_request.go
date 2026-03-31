// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeletePolicyRequest struct {


	DeletePolicyQuery *DeletePolicyQuery `json:"deletePolicyQuery,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s DeletePolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePolicyRequest) SetDeletePolicyQuery(v *DeletePolicyQuery) *DeletePolicyRequest {
	s.DeletePolicyQuery = v
	return s
}


type DeletePolicyRequestBuilder struct {
	s *DeletePolicyRequest
}

func NewDeletePolicyRequestBuilder() *DeletePolicyRequestBuilder {
	s := &DeletePolicyRequest{}
	b := &DeletePolicyRequestBuilder{s: s}
	return b
}

func (b *DeletePolicyRequestBuilder) DeletePolicyQuery(v *DeletePolicyQuery) *DeletePolicyRequestBuilder {
	b.s.DeletePolicyQuery = v
	return b
}

func (b *DeletePolicyRequestBuilder) Build() *DeletePolicyRequest {
	return b.s
}


