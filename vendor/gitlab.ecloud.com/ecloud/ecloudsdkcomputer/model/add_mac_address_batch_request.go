// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddMacAddressBatchRequest struct {


	AddMacAddressBatchBody *AddMacAddressBatchBody `json:"addMacAddressBatchBody,omitempty"`
}

func (s AddMacAddressBatchRequest) String() string {
	return utils.Beautify(s)
}

func (s AddMacAddressBatchRequest) GoString() string {
	return s.String()
}

func (s AddMacAddressBatchRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddMacAddressBatchRequest) SetAddMacAddressBatchBody(v *AddMacAddressBatchBody) *AddMacAddressBatchRequest {
	s.AddMacAddressBatchBody = v
	return s
}


type AddMacAddressBatchRequestBuilder struct {
	s *AddMacAddressBatchRequest
}

func NewAddMacAddressBatchRequestBuilder() *AddMacAddressBatchRequestBuilder {
	s := &AddMacAddressBatchRequest{}
	b := &AddMacAddressBatchRequestBuilder{s: s}
	return b
}

func (b *AddMacAddressBatchRequestBuilder) AddMacAddressBatchBody(v *AddMacAddressBatchBody) *AddMacAddressBatchRequestBuilder {
	b.s.AddMacAddressBatchBody = v
	return b
}

func (b *AddMacAddressBatchRequestBuilder) Build() *AddMacAddressBatchRequest {
	return b.s
}


