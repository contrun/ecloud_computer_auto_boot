// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnBindMachineRequest struct {


	UnBindMachineBody *UnBindMachineBody `json:"unBindMachineBody,omitempty"`
}

func (s UnBindMachineRequest) String() string {
	return utils.Beautify(s)
}

func (s UnBindMachineRequest) GoString() string {
	return s.String()
}

func (s UnBindMachineRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnBindMachineRequest) SetUnBindMachineBody(v *UnBindMachineBody) *UnBindMachineRequest {
	s.UnBindMachineBody = v
	return s
}


type UnBindMachineRequestBuilder struct {
	s *UnBindMachineRequest
}

func NewUnBindMachineRequestBuilder() *UnBindMachineRequestBuilder {
	s := &UnBindMachineRequest{}
	b := &UnBindMachineRequestBuilder{s: s}
	return b
}

func (b *UnBindMachineRequestBuilder) UnBindMachineBody(v *UnBindMachineBody) *UnBindMachineRequestBuilder {
	b.s.UnBindMachineBody = v
	return b
}

func (b *UnBindMachineRequestBuilder) Build() *UnBindMachineRequest {
	return b.s
}


