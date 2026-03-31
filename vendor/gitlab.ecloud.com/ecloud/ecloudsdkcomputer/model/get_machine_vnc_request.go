// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMachineVNCRequest struct {


	GetMachineVNCQuery *GetMachineVNCQuery `json:"getMachineVNCQuery,omitempty"`
}

func (s GetMachineVNCRequest) String() string {
	return utils.Beautify(s)
}

func (s GetMachineVNCRequest) GoString() string {
	return s.String()
}

func (s GetMachineVNCRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMachineVNCRequest) SetGetMachineVNCQuery(v *GetMachineVNCQuery) *GetMachineVNCRequest {
	s.GetMachineVNCQuery = v
	return s
}


type GetMachineVNCRequestBuilder struct {
	s *GetMachineVNCRequest
}

func NewGetMachineVNCRequestBuilder() *GetMachineVNCRequestBuilder {
	s := &GetMachineVNCRequest{}
	b := &GetMachineVNCRequestBuilder{s: s}
	return b
}

func (b *GetMachineVNCRequestBuilder) GetMachineVNCQuery(v *GetMachineVNCQuery) *GetMachineVNCRequestBuilder {
	b.s.GetMachineVNCQuery = v
	return b
}

func (b *GetMachineVNCRequestBuilder) Build() *GetMachineVNCRequest {
	return b.s
}


