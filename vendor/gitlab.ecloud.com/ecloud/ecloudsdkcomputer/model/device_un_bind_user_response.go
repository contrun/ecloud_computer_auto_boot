// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeviceUnBindUserResponseStateEnum string

// List of State
const (
    DeviceUnBindUserResponseStateEnumOk DeviceUnBindUserResponseStateEnum = "OK"
    DeviceUnBindUserResponseStateEnumError DeviceUnBindUserResponseStateEnum = "ERROR"
    DeviceUnBindUserResponseStateEnumException DeviceUnBindUserResponseStateEnum = "EXCEPTION"
    DeviceUnBindUserResponseStateEnumForbidden DeviceUnBindUserResponseStateEnum = "FORBIDDEN"
    DeviceUnBindUserResponseStateEnumRemaining DeviceUnBindUserResponseStateEnum = "REMAINING"
    DeviceUnBindUserResponseStateEnumTimeout DeviceUnBindUserResponseStateEnum = "TIMEOUT"
)

type DeviceUnBindUserResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *DeviceUnBindUserResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s DeviceUnBindUserResponse) String() string {
	return utils.Beautify(s)
}

func (s DeviceUnBindUserResponse) GoString() string {
	return s.String()
}

func (s DeviceUnBindUserResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeviceUnBindUserResponse) SetRequestId(v string) *DeviceUnBindUserResponse {
	s.RequestId = &v
	return s
}

func (s *DeviceUnBindUserResponse) SetErrorMessage(v string) *DeviceUnBindUserResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeviceUnBindUserResponse) SetErrorCode(v string) *DeviceUnBindUserResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeviceUnBindUserResponse) SetState(v DeviceUnBindUserResponseStateEnum) *DeviceUnBindUserResponse {
	s.State = &v
	return s
}

func (s *DeviceUnBindUserResponse) SetBody(v string) *DeviceUnBindUserResponse {
	s.Body = &v
	return s
}


type DeviceUnBindUserResponseBuilder struct {
	s *DeviceUnBindUserResponse
}

func NewDeviceUnBindUserResponseBuilder() *DeviceUnBindUserResponseBuilder {
	s := &DeviceUnBindUserResponse{}
	b := &DeviceUnBindUserResponseBuilder{s: s}
	return b
}

func (b *DeviceUnBindUserResponseBuilder) RequestId(v string) *DeviceUnBindUserResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeviceUnBindUserResponseBuilder) ErrorMessage(v string) *DeviceUnBindUserResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeviceUnBindUserResponseBuilder) ErrorCode(v string) *DeviceUnBindUserResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeviceUnBindUserResponseBuilder) State(v DeviceUnBindUserResponseStateEnum) *DeviceUnBindUserResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeviceUnBindUserResponseBuilder) Body(v string) *DeviceUnBindUserResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeviceUnBindUserResponseBuilder) Build() *DeviceUnBindUserResponse {
	return b.s
}


