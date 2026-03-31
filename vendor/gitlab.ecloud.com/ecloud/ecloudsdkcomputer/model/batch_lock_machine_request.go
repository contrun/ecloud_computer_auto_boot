// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchLockMachineRequest struct {


	BatchLockMachineBody *BatchLockMachineBody `json:"batchLockMachineBody,omitempty"`
}

func (s BatchLockMachineRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchLockMachineRequest) GoString() string {
	return s.String()
}

func (s BatchLockMachineRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchLockMachineRequest) SetBatchLockMachineBody(v *BatchLockMachineBody) *BatchLockMachineRequest {
	s.BatchLockMachineBody = v
	return s
}


type BatchLockMachineRequestBuilder struct {
	s *BatchLockMachineRequest
}

func NewBatchLockMachineRequestBuilder() *BatchLockMachineRequestBuilder {
	s := &BatchLockMachineRequest{}
	b := &BatchLockMachineRequestBuilder{s: s}
	return b
}

func (b *BatchLockMachineRequestBuilder) BatchLockMachineBody(v *BatchLockMachineBody) *BatchLockMachineRequestBuilder {
	b.s.BatchLockMachineBody = v
	return b
}

func (b *BatchLockMachineRequestBuilder) Build() *BatchLockMachineRequest {
	return b.s
}


