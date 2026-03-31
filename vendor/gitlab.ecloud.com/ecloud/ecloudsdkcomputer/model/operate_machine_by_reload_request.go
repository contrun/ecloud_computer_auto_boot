// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OperateMachineByReloadRequest struct {


	OperateMachineByReloadQuery *OperateMachineByReloadQuery `json:"operateMachineByReloadQuery,omitempty"`
}

func (s OperateMachineByReloadRequest) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByReloadRequest) GoString() string {
	return s.String()
}

func (s OperateMachineByReloadRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByReloadRequest) SetOperateMachineByReloadQuery(v *OperateMachineByReloadQuery) *OperateMachineByReloadRequest {
	s.OperateMachineByReloadQuery = v
	return s
}


type OperateMachineByReloadRequestBuilder struct {
	s *OperateMachineByReloadRequest
}

func NewOperateMachineByReloadRequestBuilder() *OperateMachineByReloadRequestBuilder {
	s := &OperateMachineByReloadRequest{}
	b := &OperateMachineByReloadRequestBuilder{s: s}
	return b
}

func (b *OperateMachineByReloadRequestBuilder) OperateMachineByReloadQuery(v *OperateMachineByReloadQuery) *OperateMachineByReloadRequestBuilder {
	b.s.OperateMachineByReloadQuery = v
	return b
}

func (b *OperateMachineByReloadRequestBuilder) Build() *OperateMachineByReloadRequest {
	return b.s
}


