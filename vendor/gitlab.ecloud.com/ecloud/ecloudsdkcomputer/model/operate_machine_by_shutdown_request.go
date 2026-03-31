// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OperateMachineByShutdownRequest struct {


	OperateMachineByShutdownQuery *OperateMachineByShutdownQuery `json:"operateMachineByShutdownQuery,omitempty"`
}

func (s OperateMachineByShutdownRequest) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByShutdownRequest) GoString() string {
	return s.String()
}

func (s OperateMachineByShutdownRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByShutdownRequest) SetOperateMachineByShutdownQuery(v *OperateMachineByShutdownQuery) *OperateMachineByShutdownRequest {
	s.OperateMachineByShutdownQuery = v
	return s
}


type OperateMachineByShutdownRequestBuilder struct {
	s *OperateMachineByShutdownRequest
}

func NewOperateMachineByShutdownRequestBuilder() *OperateMachineByShutdownRequestBuilder {
	s := &OperateMachineByShutdownRequest{}
	b := &OperateMachineByShutdownRequestBuilder{s: s}
	return b
}

func (b *OperateMachineByShutdownRequestBuilder) OperateMachineByShutdownQuery(v *OperateMachineByShutdownQuery) *OperateMachineByShutdownRequestBuilder {
	b.s.OperateMachineByShutdownQuery = v
	return b
}

func (b *OperateMachineByShutdownRequestBuilder) Build() *OperateMachineByShutdownRequest {
	return b.s
}


