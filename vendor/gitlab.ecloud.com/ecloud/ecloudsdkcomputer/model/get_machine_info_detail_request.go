// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMachineInfoDetailRequest struct {


	GetMachineInfoDetailBody *GetMachineInfoDetailBody `json:"getMachineInfoDetailBody,omitempty"`
}

func (s GetMachineInfoDetailRequest) String() string {
	return utils.Beautify(s)
}

func (s GetMachineInfoDetailRequest) GoString() string {
	return s.String()
}

func (s GetMachineInfoDetailRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMachineInfoDetailRequest) SetGetMachineInfoDetailBody(v *GetMachineInfoDetailBody) *GetMachineInfoDetailRequest {
	s.GetMachineInfoDetailBody = v
	return s
}


type GetMachineInfoDetailRequestBuilder struct {
	s *GetMachineInfoDetailRequest
}

func NewGetMachineInfoDetailRequestBuilder() *GetMachineInfoDetailRequestBuilder {
	s := &GetMachineInfoDetailRequest{}
	b := &GetMachineInfoDetailRequestBuilder{s: s}
	return b
}

func (b *GetMachineInfoDetailRequestBuilder) GetMachineInfoDetailBody(v *GetMachineInfoDetailBody) *GetMachineInfoDetailRequestBuilder {
	b.s.GetMachineInfoDetailBody = v
	return b
}

func (b *GetMachineInfoDetailRequestBuilder) Build() *GetMachineInfoDetailRequest {
	return b.s
}


