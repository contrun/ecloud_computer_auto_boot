// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type LockDeviceByIdResponseStateEnum string

// List of State
const (
    LockDeviceByIdResponseStateEnumOk LockDeviceByIdResponseStateEnum = "OK"
    LockDeviceByIdResponseStateEnumError LockDeviceByIdResponseStateEnum = "ERROR"
    LockDeviceByIdResponseStateEnumException LockDeviceByIdResponseStateEnum = "EXCEPTION"
    LockDeviceByIdResponseStateEnumForbidden LockDeviceByIdResponseStateEnum = "FORBIDDEN"
    LockDeviceByIdResponseStateEnumRemaining LockDeviceByIdResponseStateEnum = "REMAINING"
    LockDeviceByIdResponseStateEnumTimeout LockDeviceByIdResponseStateEnum = "TIMEOUT"
)

type LockDeviceByIdResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *LockDeviceByIdResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s LockDeviceByIdResponse) String() string {
	return utils.Beautify(s)
}

func (s LockDeviceByIdResponse) GoString() string {
	return s.String()
}

func (s LockDeviceByIdResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LockDeviceByIdResponse) SetRequestId(v string) *LockDeviceByIdResponse {
	s.RequestId = &v
	return s
}

func (s *LockDeviceByIdResponse) SetErrorMessage(v string) *LockDeviceByIdResponse {
	s.ErrorMessage = &v
	return s
}

func (s *LockDeviceByIdResponse) SetErrorCode(v string) *LockDeviceByIdResponse {
	s.ErrorCode = &v
	return s
}

func (s *LockDeviceByIdResponse) SetState(v LockDeviceByIdResponseStateEnum) *LockDeviceByIdResponse {
	s.State = &v
	return s
}

func (s *LockDeviceByIdResponse) SetBody(v string) *LockDeviceByIdResponse {
	s.Body = &v
	return s
}


type LockDeviceByIdResponseBuilder struct {
	s *LockDeviceByIdResponse
}

func NewLockDeviceByIdResponseBuilder() *LockDeviceByIdResponseBuilder {
	s := &LockDeviceByIdResponse{}
	b := &LockDeviceByIdResponseBuilder{s: s}
	return b
}

func (b *LockDeviceByIdResponseBuilder) RequestId(v string) *LockDeviceByIdResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *LockDeviceByIdResponseBuilder) ErrorMessage(v string) *LockDeviceByIdResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *LockDeviceByIdResponseBuilder) ErrorCode(v string) *LockDeviceByIdResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *LockDeviceByIdResponseBuilder) State(v LockDeviceByIdResponseStateEnum) *LockDeviceByIdResponseBuilder {
	b.s.State = &v
	return b
}

func (b *LockDeviceByIdResponseBuilder) Body(v string) *LockDeviceByIdResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *LockDeviceByIdResponseBuilder) Build() *LockDeviceByIdResponse {
	return b.s
}


