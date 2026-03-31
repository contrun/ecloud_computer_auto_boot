// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyDeviceTrustNumRequest struct {


	ModifyDeviceTrustNumBody *ModifyDeviceTrustNumBody `json:"modifyDeviceTrustNumBody,omitempty"`
}

func (s ModifyDeviceTrustNumRequest) String() string {
	return utils.Beautify(s)
}

func (s ModifyDeviceTrustNumRequest) GoString() string {
	return s.String()
}

func (s ModifyDeviceTrustNumRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyDeviceTrustNumRequest) SetModifyDeviceTrustNumBody(v *ModifyDeviceTrustNumBody) *ModifyDeviceTrustNumRequest {
	s.ModifyDeviceTrustNumBody = v
	return s
}


type ModifyDeviceTrustNumRequestBuilder struct {
	s *ModifyDeviceTrustNumRequest
}

func NewModifyDeviceTrustNumRequestBuilder() *ModifyDeviceTrustNumRequestBuilder {
	s := &ModifyDeviceTrustNumRequest{}
	b := &ModifyDeviceTrustNumRequestBuilder{s: s}
	return b
}

func (b *ModifyDeviceTrustNumRequestBuilder) ModifyDeviceTrustNumBody(v *ModifyDeviceTrustNumBody) *ModifyDeviceTrustNumRequestBuilder {
	b.s.ModifyDeviceTrustNumBody = v
	return b
}

func (b *ModifyDeviceTrustNumRequestBuilder) Build() *ModifyDeviceTrustNumRequest {
	return b.s
}


