// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnlockMachineRequest struct {


	UnlockMachineBody *UnlockMachineBody `json:"unlockMachineBody,omitempty"`
}

func (s UnlockMachineRequest) String() string {
	return utils.Beautify(s)
}

func (s UnlockMachineRequest) GoString() string {
	return s.String()
}

func (s UnlockMachineRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnlockMachineRequest) SetUnlockMachineBody(v *UnlockMachineBody) *UnlockMachineRequest {
	s.UnlockMachineBody = v
	return s
}


type UnlockMachineRequestBuilder struct {
	s *UnlockMachineRequest
}

func NewUnlockMachineRequestBuilder() *UnlockMachineRequestBuilder {
	s := &UnlockMachineRequest{}
	b := &UnlockMachineRequestBuilder{s: s}
	return b
}

func (b *UnlockMachineRequestBuilder) UnlockMachineBody(v *UnlockMachineBody) *UnlockMachineRequestBuilder {
	b.s.UnlockMachineBody = v
	return b
}

func (b *UnlockMachineRequestBuilder) Build() *UnlockMachineRequest {
	return b.s
}


