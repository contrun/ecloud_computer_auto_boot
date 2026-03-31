// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewComputerRequest struct {


	RenewComputerBody *RenewComputerBody `json:"renewComputerBody,omitempty"`
}

func (s RenewComputerRequest) String() string {
	return utils.Beautify(s)
}

func (s RenewComputerRequest) GoString() string {
	return s.String()
}

func (s RenewComputerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewComputerRequest) SetRenewComputerBody(v *RenewComputerBody) *RenewComputerRequest {
	s.RenewComputerBody = v
	return s
}


type RenewComputerRequestBuilder struct {
	s *RenewComputerRequest
}

func NewRenewComputerRequestBuilder() *RenewComputerRequestBuilder {
	s := &RenewComputerRequest{}
	b := &RenewComputerRequestBuilder{s: s}
	return b
}

func (b *RenewComputerRequestBuilder) RenewComputerBody(v *RenewComputerBody) *RenewComputerRequestBuilder {
	b.s.RenewComputerBody = v
	return b
}

func (b *RenewComputerRequestBuilder) Build() *RenewComputerRequest {
	return b.s
}


