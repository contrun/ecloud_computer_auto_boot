// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeviceBindUserRequest struct {


	DeviceBindUserBody *DeviceBindUserBody `json:"deviceBindUserBody,omitempty"`
}

func (s DeviceBindUserRequest) String() string {
	return utils.Beautify(s)
}

func (s DeviceBindUserRequest) GoString() string {
	return s.String()
}

func (s DeviceBindUserRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeviceBindUserRequest) SetDeviceBindUserBody(v *DeviceBindUserBody) *DeviceBindUserRequest {
	s.DeviceBindUserBody = v
	return s
}


type DeviceBindUserRequestBuilder struct {
	s *DeviceBindUserRequest
}

func NewDeviceBindUserRequestBuilder() *DeviceBindUserRequestBuilder {
	s := &DeviceBindUserRequest{}
	b := &DeviceBindUserRequestBuilder{s: s}
	return b
}

func (b *DeviceBindUserRequestBuilder) DeviceBindUserBody(v *DeviceBindUserBody) *DeviceBindUserRequestBuilder {
	b.s.DeviceBindUserBody = v
	return b
}

func (b *DeviceBindUserRequestBuilder) Build() *DeviceBindUserRequest {
	return b.s
}


