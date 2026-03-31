// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindMachineRequest struct {


	BindMachineBody *BindMachineBody `json:"bindMachineBody,omitempty"`
}

func (s BindMachineRequest) String() string {
	return utils.Beautify(s)
}

func (s BindMachineRequest) GoString() string {
	return s.String()
}

func (s BindMachineRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindMachineRequest) SetBindMachineBody(v *BindMachineBody) *BindMachineRequest {
	s.BindMachineBody = v
	return s
}


type BindMachineRequestBuilder struct {
	s *BindMachineRequest
}

func NewBindMachineRequestBuilder() *BindMachineRequestBuilder {
	s := &BindMachineRequest{}
	b := &BindMachineRequestBuilder{s: s}
	return b
}

func (b *BindMachineRequestBuilder) BindMachineBody(v *BindMachineBody) *BindMachineRequestBuilder {
	b.s.BindMachineBody = v
	return b
}

func (b *BindMachineRequestBuilder) Build() *BindMachineRequest {
	return b.s
}


