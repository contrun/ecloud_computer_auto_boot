// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteUserPolicyRequest struct {


	DeleteUserPolicyQuery *DeleteUserPolicyQuery `json:"deleteUserPolicyQuery,omitempty"`
}

func (s DeleteUserPolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteUserPolicyRequest) GoString() string {
	return s.String()
}

func (s DeleteUserPolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteUserPolicyRequest) SetDeleteUserPolicyQuery(v *DeleteUserPolicyQuery) *DeleteUserPolicyRequest {
	s.DeleteUserPolicyQuery = v
	return s
}


type DeleteUserPolicyRequestBuilder struct {
	s *DeleteUserPolicyRequest
}

func NewDeleteUserPolicyRequestBuilder() *DeleteUserPolicyRequestBuilder {
	s := &DeleteUserPolicyRequest{}
	b := &DeleteUserPolicyRequestBuilder{s: s}
	return b
}

func (b *DeleteUserPolicyRequestBuilder) DeleteUserPolicyQuery(v *DeleteUserPolicyQuery) *DeleteUserPolicyRequestBuilder {
	b.s.DeleteUserPolicyQuery = v
	return b
}

func (b *DeleteUserPolicyRequestBuilder) Build() *DeleteUserPolicyRequest {
	return b.s
}


