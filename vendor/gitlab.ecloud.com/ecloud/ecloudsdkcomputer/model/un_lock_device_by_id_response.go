// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UnLockDeviceByIdResponseStateEnum string

// List of State
const (
    UnLockDeviceByIdResponseStateEnumOk UnLockDeviceByIdResponseStateEnum = "OK"
    UnLockDeviceByIdResponseStateEnumError UnLockDeviceByIdResponseStateEnum = "ERROR"
    UnLockDeviceByIdResponseStateEnumException UnLockDeviceByIdResponseStateEnum = "EXCEPTION"
    UnLockDeviceByIdResponseStateEnumForbidden UnLockDeviceByIdResponseStateEnum = "FORBIDDEN"
    UnLockDeviceByIdResponseStateEnumRemaining UnLockDeviceByIdResponseStateEnum = "REMAINING"
    UnLockDeviceByIdResponseStateEnumTimeout UnLockDeviceByIdResponseStateEnum = "TIMEOUT"
)

type UnLockDeviceByIdResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *UnLockDeviceByIdResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s UnLockDeviceByIdResponse) String() string {
	return utils.Beautify(s)
}

func (s UnLockDeviceByIdResponse) GoString() string {
	return s.String()
}

func (s UnLockDeviceByIdResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnLockDeviceByIdResponse) SetRequestId(v string) *UnLockDeviceByIdResponse {
	s.RequestId = &v
	return s
}

func (s *UnLockDeviceByIdResponse) SetErrorMessage(v string) *UnLockDeviceByIdResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UnLockDeviceByIdResponse) SetErrorCode(v string) *UnLockDeviceByIdResponse {
	s.ErrorCode = &v
	return s
}

func (s *UnLockDeviceByIdResponse) SetState(v UnLockDeviceByIdResponseStateEnum) *UnLockDeviceByIdResponse {
	s.State = &v
	return s
}

func (s *UnLockDeviceByIdResponse) SetBody(v string) *UnLockDeviceByIdResponse {
	s.Body = &v
	return s
}


type UnLockDeviceByIdResponseBuilder struct {
	s *UnLockDeviceByIdResponse
}

func NewUnLockDeviceByIdResponseBuilder() *UnLockDeviceByIdResponseBuilder {
	s := &UnLockDeviceByIdResponse{}
	b := &UnLockDeviceByIdResponseBuilder{s: s}
	return b
}

func (b *UnLockDeviceByIdResponseBuilder) RequestId(v string) *UnLockDeviceByIdResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UnLockDeviceByIdResponseBuilder) ErrorMessage(v string) *UnLockDeviceByIdResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UnLockDeviceByIdResponseBuilder) ErrorCode(v string) *UnLockDeviceByIdResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UnLockDeviceByIdResponseBuilder) State(v UnLockDeviceByIdResponseStateEnum) *UnLockDeviceByIdResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UnLockDeviceByIdResponseBuilder) Body(v string) *UnLockDeviceByIdResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UnLockDeviceByIdResponseBuilder) Build() *UnLockDeviceByIdResponse {
	return b.s
}


