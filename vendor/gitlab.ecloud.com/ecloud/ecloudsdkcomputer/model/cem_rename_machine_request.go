// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CemRenameMachineRequest struct {


	CemRenameMachineQuery *CemRenameMachineQuery `json:"cemRenameMachineQuery,omitempty"`
}

func (s CemRenameMachineRequest) String() string {
	return utils.Beautify(s)
}

func (s CemRenameMachineRequest) GoString() string {
	return s.String()
}

func (s CemRenameMachineRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CemRenameMachineRequest) SetCemRenameMachineQuery(v *CemRenameMachineQuery) *CemRenameMachineRequest {
	s.CemRenameMachineQuery = v
	return s
}


type CemRenameMachineRequestBuilder struct {
	s *CemRenameMachineRequest
}

func NewCemRenameMachineRequestBuilder() *CemRenameMachineRequestBuilder {
	s := &CemRenameMachineRequest{}
	b := &CemRenameMachineRequestBuilder{s: s}
	return b
}

func (b *CemRenameMachineRequestBuilder) CemRenameMachineQuery(v *CemRenameMachineQuery) *CemRenameMachineRequestBuilder {
	b.s.CemRenameMachineQuery = v
	return b
}

func (b *CemRenameMachineRequestBuilder) Build() *CemRenameMachineRequest {
	return b.s
}


