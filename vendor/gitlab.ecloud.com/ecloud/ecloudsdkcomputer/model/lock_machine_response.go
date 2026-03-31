// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type LockMachineResponseStateEnum string

// List of State
const (
    LockMachineResponseStateEnumOk LockMachineResponseStateEnum = "OK"
    LockMachineResponseStateEnumError LockMachineResponseStateEnum = "ERROR"
    LockMachineResponseStateEnumException LockMachineResponseStateEnum = "EXCEPTION"
    LockMachineResponseStateEnumForbidden LockMachineResponseStateEnum = "FORBIDDEN"
    LockMachineResponseStateEnumRemaining LockMachineResponseStateEnum = "REMAINING"
    LockMachineResponseStateEnumTimeout LockMachineResponseStateEnum = "TIMEOUT"
)

type LockMachineResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *LockMachineResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s LockMachineResponse) String() string {
	return utils.Beautify(s)
}

func (s LockMachineResponse) GoString() string {
	return s.String()
}

func (s LockMachineResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LockMachineResponse) SetRequestId(v string) *LockMachineResponse {
	s.RequestId = &v
	return s
}

func (s *LockMachineResponse) SetErrorMessage(v string) *LockMachineResponse {
	s.ErrorMessage = &v
	return s
}

func (s *LockMachineResponse) SetErrorCode(v string) *LockMachineResponse {
	s.ErrorCode = &v
	return s
}

func (s *LockMachineResponse) SetState(v LockMachineResponseStateEnum) *LockMachineResponse {
	s.State = &v
	return s
}

func (s *LockMachineResponse) SetBody(v bool) *LockMachineResponse {
	s.Body = &v
	return s
}


type LockMachineResponseBuilder struct {
	s *LockMachineResponse
}

func NewLockMachineResponseBuilder() *LockMachineResponseBuilder {
	s := &LockMachineResponse{}
	b := &LockMachineResponseBuilder{s: s}
	return b
}

func (b *LockMachineResponseBuilder) RequestId(v string) *LockMachineResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *LockMachineResponseBuilder) ErrorMessage(v string) *LockMachineResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *LockMachineResponseBuilder) ErrorCode(v string) *LockMachineResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *LockMachineResponseBuilder) State(v LockMachineResponseStateEnum) *LockMachineResponseBuilder {
	b.s.State = &v
	return b
}

func (b *LockMachineResponseBuilder) Body(v bool) *LockMachineResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *LockMachineResponseBuilder) Build() *LockMachineResponse {
	return b.s
}


