// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateShutdownRevertConfigRequest struct {


	UpdateShutdownRevertConfigBody *UpdateShutdownRevertConfigBody `json:"updateShutdownRevertConfigBody,omitempty"`
}

func (s UpdateShutdownRevertConfigRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateShutdownRevertConfigRequest) GoString() string {
	return s.String()
}

func (s UpdateShutdownRevertConfigRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateShutdownRevertConfigRequest) SetUpdateShutdownRevertConfigBody(v *UpdateShutdownRevertConfigBody) *UpdateShutdownRevertConfigRequest {
	s.UpdateShutdownRevertConfigBody = v
	return s
}


type UpdateShutdownRevertConfigRequestBuilder struct {
	s *UpdateShutdownRevertConfigRequest
}

func NewUpdateShutdownRevertConfigRequestBuilder() *UpdateShutdownRevertConfigRequestBuilder {
	s := &UpdateShutdownRevertConfigRequest{}
	b := &UpdateShutdownRevertConfigRequestBuilder{s: s}
	return b
}

func (b *UpdateShutdownRevertConfigRequestBuilder) UpdateShutdownRevertConfigBody(v *UpdateShutdownRevertConfigBody) *UpdateShutdownRevertConfigRequestBuilder {
	b.s.UpdateShutdownRevertConfigBody = v
	return b
}

func (b *UpdateShutdownRevertConfigRequestBuilder) Build() *UpdateShutdownRevertConfigRequest {
	return b.s
}


