// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetUserPolicyStatusRequest struct {


	SetUserPolicyStatusBody *SetUserPolicyStatusBody `json:"setUserPolicyStatusBody,omitempty"`
}

func (s SetUserPolicyStatusRequest) String() string {
	return utils.Beautify(s)
}

func (s SetUserPolicyStatusRequest) GoString() string {
	return s.String()
}

func (s SetUserPolicyStatusRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetUserPolicyStatusRequest) SetSetUserPolicyStatusBody(v *SetUserPolicyStatusBody) *SetUserPolicyStatusRequest {
	s.SetUserPolicyStatusBody = v
	return s
}


type SetUserPolicyStatusRequestBuilder struct {
	s *SetUserPolicyStatusRequest
}

func NewSetUserPolicyStatusRequestBuilder() *SetUserPolicyStatusRequestBuilder {
	s := &SetUserPolicyStatusRequest{}
	b := &SetUserPolicyStatusRequestBuilder{s: s}
	return b
}

func (b *SetUserPolicyStatusRequestBuilder) SetUserPolicyStatusBody(v *SetUserPolicyStatusBody) *SetUserPolicyStatusRequestBuilder {
	b.s.SetUserPolicyStatusBody = v
	return b
}

func (b *SetUserPolicyStatusRequestBuilder) Build() *SetUserPolicyStatusRequest {
	return b.s
}


