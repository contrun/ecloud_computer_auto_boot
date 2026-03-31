// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchUnlockMachineRequest struct {


	BatchUnlockMachineBody *BatchUnlockMachineBody `json:"batchUnlockMachineBody,omitempty"`
}

func (s BatchUnlockMachineRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchUnlockMachineRequest) GoString() string {
	return s.String()
}

func (s BatchUnlockMachineRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUnlockMachineRequest) SetBatchUnlockMachineBody(v *BatchUnlockMachineBody) *BatchUnlockMachineRequest {
	s.BatchUnlockMachineBody = v
	return s
}


type BatchUnlockMachineRequestBuilder struct {
	s *BatchUnlockMachineRequest
}

func NewBatchUnlockMachineRequestBuilder() *BatchUnlockMachineRequestBuilder {
	s := &BatchUnlockMachineRequest{}
	b := &BatchUnlockMachineRequestBuilder{s: s}
	return b
}

func (b *BatchUnlockMachineRequestBuilder) BatchUnlockMachineBody(v *BatchUnlockMachineBody) *BatchUnlockMachineRequestBuilder {
	b.s.BatchUnlockMachineBody = v
	return b
}

func (b *BatchUnlockMachineRequestBuilder) Build() *BatchUnlockMachineRequest {
	return b.s
}


