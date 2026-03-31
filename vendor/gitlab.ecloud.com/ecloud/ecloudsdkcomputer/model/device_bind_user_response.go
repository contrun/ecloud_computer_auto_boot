// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeviceBindUserResponseStateEnum string

// List of State
const (
    DeviceBindUserResponseStateEnumOk DeviceBindUserResponseStateEnum = "OK"
    DeviceBindUserResponseStateEnumError DeviceBindUserResponseStateEnum = "ERROR"
    DeviceBindUserResponseStateEnumException DeviceBindUserResponseStateEnum = "EXCEPTION"
    DeviceBindUserResponseStateEnumForbidden DeviceBindUserResponseStateEnum = "FORBIDDEN"
    DeviceBindUserResponseStateEnumRemaining DeviceBindUserResponseStateEnum = "REMAINING"
    DeviceBindUserResponseStateEnumTimeout DeviceBindUserResponseStateEnum = "TIMEOUT"
)

type DeviceBindUserResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *DeviceBindUserResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s DeviceBindUserResponse) String() string {
	return utils.Beautify(s)
}

func (s DeviceBindUserResponse) GoString() string {
	return s.String()
}

func (s DeviceBindUserResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeviceBindUserResponse) SetRequestId(v string) *DeviceBindUserResponse {
	s.RequestId = &v
	return s
}

func (s *DeviceBindUserResponse) SetErrorMessage(v string) *DeviceBindUserResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeviceBindUserResponse) SetErrorCode(v string) *DeviceBindUserResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeviceBindUserResponse) SetState(v DeviceBindUserResponseStateEnum) *DeviceBindUserResponse {
	s.State = &v
	return s
}

func (s *DeviceBindUserResponse) SetBody(v string) *DeviceBindUserResponse {
	s.Body = &v
	return s
}


type DeviceBindUserResponseBuilder struct {
	s *DeviceBindUserResponse
}

func NewDeviceBindUserResponseBuilder() *DeviceBindUserResponseBuilder {
	s := &DeviceBindUserResponse{}
	b := &DeviceBindUserResponseBuilder{s: s}
	return b
}

func (b *DeviceBindUserResponseBuilder) RequestId(v string) *DeviceBindUserResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeviceBindUserResponseBuilder) ErrorMessage(v string) *DeviceBindUserResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeviceBindUserResponseBuilder) ErrorCode(v string) *DeviceBindUserResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeviceBindUserResponseBuilder) State(v DeviceBindUserResponseStateEnum) *DeviceBindUserResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeviceBindUserResponseBuilder) Body(v string) *DeviceBindUserResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeviceBindUserResponseBuilder) Build() *DeviceBindUserResponse {
	return b.s
}


