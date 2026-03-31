// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OperateMachineByRestartRequest struct {


	OperateMachineByRestartQuery *OperateMachineByRestartQuery `json:"operateMachineByRestartQuery,omitempty"`
}

func (s OperateMachineByRestartRequest) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByRestartRequest) GoString() string {
	return s.String()
}

func (s OperateMachineByRestartRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByRestartRequest) SetOperateMachineByRestartQuery(v *OperateMachineByRestartQuery) *OperateMachineByRestartRequest {
	s.OperateMachineByRestartQuery = v
	return s
}


type OperateMachineByRestartRequestBuilder struct {
	s *OperateMachineByRestartRequest
}

func NewOperateMachineByRestartRequestBuilder() *OperateMachineByRestartRequestBuilder {
	s := &OperateMachineByRestartRequest{}
	b := &OperateMachineByRestartRequestBuilder{s: s}
	return b
}

func (b *OperateMachineByRestartRequestBuilder) OperateMachineByRestartQuery(v *OperateMachineByRestartQuery) *OperateMachineByRestartRequestBuilder {
	b.s.OperateMachineByRestartQuery = v
	return b
}

func (b *OperateMachineByRestartRequestBuilder) Build() *OperateMachineByRestartRequest {
	return b.s
}


