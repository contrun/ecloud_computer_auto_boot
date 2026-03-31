// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LockMachineRequest struct {


	LockMachineBody *LockMachineBody `json:"lockMachineBody,omitempty"`
}

func (s LockMachineRequest) String() string {
	return utils.Beautify(s)
}

func (s LockMachineRequest) GoString() string {
	return s.String()
}

func (s LockMachineRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LockMachineRequest) SetLockMachineBody(v *LockMachineBody) *LockMachineRequest {
	s.LockMachineBody = v
	return s
}


type LockMachineRequestBuilder struct {
	s *LockMachineRequest
}

func NewLockMachineRequestBuilder() *LockMachineRequestBuilder {
	s := &LockMachineRequest{}
	b := &LockMachineRequestBuilder{s: s}
	return b
}

func (b *LockMachineRequestBuilder) LockMachineBody(v *LockMachineBody) *LockMachineRequestBuilder {
	b.s.LockMachineBody = v
	return b
}

func (b *LockMachineRequestBuilder) Build() *LockMachineRequest {
	return b.s
}


