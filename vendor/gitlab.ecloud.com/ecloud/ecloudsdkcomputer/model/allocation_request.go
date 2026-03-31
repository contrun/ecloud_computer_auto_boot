// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AllocationRequest struct {


	AllocationBody *AllocationBody `json:"allocationBody,omitempty"`
}

func (s AllocationRequest) String() string {
	return utils.Beautify(s)
}

func (s AllocationRequest) GoString() string {
	return s.String()
}

func (s AllocationRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AllocationRequest) SetAllocationBody(v *AllocationBody) *AllocationRequest {
	s.AllocationBody = v
	return s
}


type AllocationRequestBuilder struct {
	s *AllocationRequest
}

func NewAllocationRequestBuilder() *AllocationRequestBuilder {
	s := &AllocationRequest{}
	b := &AllocationRequestBuilder{s: s}
	return b
}

func (b *AllocationRequestBuilder) AllocationBody(v *AllocationBody) *AllocationRequestBuilder {
	b.s.AllocationBody = v
	return b
}

func (b *AllocationRequestBuilder) Build() *AllocationRequest {
	return b.s
}


