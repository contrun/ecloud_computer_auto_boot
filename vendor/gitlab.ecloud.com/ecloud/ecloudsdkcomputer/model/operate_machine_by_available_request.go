// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OperateMachineByAvailableRequest struct {


	OperateMachineByAvailableQuery *OperateMachineByAvailableQuery `json:"operateMachineByAvailableQuery,omitempty"`
}

func (s OperateMachineByAvailableRequest) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByAvailableRequest) GoString() string {
	return s.String()
}

func (s OperateMachineByAvailableRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByAvailableRequest) SetOperateMachineByAvailableQuery(v *OperateMachineByAvailableQuery) *OperateMachineByAvailableRequest {
	s.OperateMachineByAvailableQuery = v
	return s
}


type OperateMachineByAvailableRequestBuilder struct {
	s *OperateMachineByAvailableRequest
}

func NewOperateMachineByAvailableRequestBuilder() *OperateMachineByAvailableRequestBuilder {
	s := &OperateMachineByAvailableRequest{}
	b := &OperateMachineByAvailableRequestBuilder{s: s}
	return b
}

func (b *OperateMachineByAvailableRequestBuilder) OperateMachineByAvailableQuery(v *OperateMachineByAvailableQuery) *OperateMachineByAvailableRequestBuilder {
	b.s.OperateMachineByAvailableQuery = v
	return b
}

func (b *OperateMachineByAvailableRequestBuilder) Build() *OperateMachineByAvailableRequest {
	return b.s
}


