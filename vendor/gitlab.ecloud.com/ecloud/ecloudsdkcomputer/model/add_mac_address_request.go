// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddMacAddressRequest struct {


	AddMacAddressBody *AddMacAddressBody `json:"addMacAddressBody,omitempty"`
}

func (s AddMacAddressRequest) String() string {
	return utils.Beautify(s)
}

func (s AddMacAddressRequest) GoString() string {
	return s.String()
}

func (s AddMacAddressRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddMacAddressRequest) SetAddMacAddressBody(v *AddMacAddressBody) *AddMacAddressRequest {
	s.AddMacAddressBody = v
	return s
}


type AddMacAddressRequestBuilder struct {
	s *AddMacAddressRequest
}

func NewAddMacAddressRequestBuilder() *AddMacAddressRequestBuilder {
	s := &AddMacAddressRequest{}
	b := &AddMacAddressRequestBuilder{s: s}
	return b
}

func (b *AddMacAddressRequestBuilder) AddMacAddressBody(v *AddMacAddressBody) *AddMacAddressRequestBuilder {
	b.s.AddMacAddressBody = v
	return b
}

func (b *AddMacAddressRequestBuilder) Build() *AddMacAddressRequest {
	return b.s
}


