// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeviceUnBindUserRequest struct {


	DeviceUnBindUserBody *DeviceUnBindUserBody `json:"deviceUnBindUserBody,omitempty"`
}

func (s DeviceUnBindUserRequest) String() string {
	return utils.Beautify(s)
}

func (s DeviceUnBindUserRequest) GoString() string {
	return s.String()
}

func (s DeviceUnBindUserRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeviceUnBindUserRequest) SetDeviceUnBindUserBody(v *DeviceUnBindUserBody) *DeviceUnBindUserRequest {
	s.DeviceUnBindUserBody = v
	return s
}


type DeviceUnBindUserRequestBuilder struct {
	s *DeviceUnBindUserRequest
}

func NewDeviceUnBindUserRequestBuilder() *DeviceUnBindUserRequestBuilder {
	s := &DeviceUnBindUserRequest{}
	b := &DeviceUnBindUserRequestBuilder{s: s}
	return b
}

func (b *DeviceUnBindUserRequestBuilder) DeviceUnBindUserBody(v *DeviceUnBindUserBody) *DeviceUnBindUserRequestBuilder {
	b.s.DeviceUnBindUserBody = v
	return b
}

func (b *DeviceUnBindUserRequestBuilder) Build() *DeviceUnBindUserRequest {
	return b.s
}


