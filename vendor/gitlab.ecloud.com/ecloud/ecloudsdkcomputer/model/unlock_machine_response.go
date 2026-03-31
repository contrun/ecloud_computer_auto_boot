// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UnlockMachineResponseStateEnum string

// List of State
const (
    UnlockMachineResponseStateEnumOk UnlockMachineResponseStateEnum = "OK"
    UnlockMachineResponseStateEnumError UnlockMachineResponseStateEnum = "ERROR"
    UnlockMachineResponseStateEnumException UnlockMachineResponseStateEnum = "EXCEPTION"
    UnlockMachineResponseStateEnumForbidden UnlockMachineResponseStateEnum = "FORBIDDEN"
    UnlockMachineResponseStateEnumRemaining UnlockMachineResponseStateEnum = "REMAINING"
    UnlockMachineResponseStateEnumTimeout UnlockMachineResponseStateEnum = "TIMEOUT"
)

type UnlockMachineResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *UnlockMachineResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s UnlockMachineResponse) String() string {
	return utils.Beautify(s)
}

func (s UnlockMachineResponse) GoString() string {
	return s.String()
}

func (s UnlockMachineResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnlockMachineResponse) SetRequestId(v string) *UnlockMachineResponse {
	s.RequestId = &v
	return s
}

func (s *UnlockMachineResponse) SetErrorMessage(v string) *UnlockMachineResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UnlockMachineResponse) SetErrorCode(v string) *UnlockMachineResponse {
	s.ErrorCode = &v
	return s
}

func (s *UnlockMachineResponse) SetState(v UnlockMachineResponseStateEnum) *UnlockMachineResponse {
	s.State = &v
	return s
}

func (s *UnlockMachineResponse) SetBody(v bool) *UnlockMachineResponse {
	s.Body = &v
	return s
}


type UnlockMachineResponseBuilder struct {
	s *UnlockMachineResponse
}

func NewUnlockMachineResponseBuilder() *UnlockMachineResponseBuilder {
	s := &UnlockMachineResponse{}
	b := &UnlockMachineResponseBuilder{s: s}
	return b
}

func (b *UnlockMachineResponseBuilder) RequestId(v string) *UnlockMachineResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UnlockMachineResponseBuilder) ErrorMessage(v string) *UnlockMachineResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UnlockMachineResponseBuilder) ErrorCode(v string) *UnlockMachineResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UnlockMachineResponseBuilder) State(v UnlockMachineResponseStateEnum) *UnlockMachineResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UnlockMachineResponseBuilder) Body(v bool) *UnlockMachineResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UnlockMachineResponseBuilder) Build() *UnlockMachineResponse {
	return b.s
}


