// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetShutdownRevertConfigRequest struct {


	GetShutdownRevertConfigBody *GetShutdownRevertConfigBody `json:"getShutdownRevertConfigBody,omitempty"`
}

func (s GetShutdownRevertConfigRequest) String() string {
	return utils.Beautify(s)
}

func (s GetShutdownRevertConfigRequest) GoString() string {
	return s.String()
}

func (s GetShutdownRevertConfigRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShutdownRevertConfigRequest) SetGetShutdownRevertConfigBody(v *GetShutdownRevertConfigBody) *GetShutdownRevertConfigRequest {
	s.GetShutdownRevertConfigBody = v
	return s
}


type GetShutdownRevertConfigRequestBuilder struct {
	s *GetShutdownRevertConfigRequest
}

func NewGetShutdownRevertConfigRequestBuilder() *GetShutdownRevertConfigRequestBuilder {
	s := &GetShutdownRevertConfigRequest{}
	b := &GetShutdownRevertConfigRequestBuilder{s: s}
	return b
}

func (b *GetShutdownRevertConfigRequestBuilder) GetShutdownRevertConfigBody(v *GetShutdownRevertConfigBody) *GetShutdownRevertConfigRequestBuilder {
	b.s.GetShutdownRevertConfigBody = v
	return b
}

func (b *GetShutdownRevertConfigRequestBuilder) Build() *GetShutdownRevertConfigRequest {
	return b.s
}


