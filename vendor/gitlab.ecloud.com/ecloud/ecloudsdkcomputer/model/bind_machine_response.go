// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BindMachineResponseStateEnum string

// List of State
const (
    BindMachineResponseStateEnumOk BindMachineResponseStateEnum = "OK"
    BindMachineResponseStateEnumError BindMachineResponseStateEnum = "ERROR"
    BindMachineResponseStateEnumException BindMachineResponseStateEnum = "EXCEPTION"
    BindMachineResponseStateEnumForbidden BindMachineResponseStateEnum = "FORBIDDEN"
    BindMachineResponseStateEnumRemaining BindMachineResponseStateEnum = "REMAINING"
    BindMachineResponseStateEnumTimeout BindMachineResponseStateEnum = "TIMEOUT"
)

type BindMachineResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BindMachineResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]BindMachineResponseBody `json:"body,omitempty"`
}

func (s BindMachineResponse) String() string {
	return utils.Beautify(s)
}

func (s BindMachineResponse) GoString() string {
	return s.String()
}

func (s BindMachineResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindMachineResponse) SetRequestId(v string) *BindMachineResponse {
	s.RequestId = &v
	return s
}

func (s *BindMachineResponse) SetErrorMessage(v string) *BindMachineResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BindMachineResponse) SetErrorCode(v string) *BindMachineResponse {
	s.ErrorCode = &v
	return s
}

func (s *BindMachineResponse) SetState(v BindMachineResponseStateEnum) *BindMachineResponse {
	s.State = &v
	return s
}

func (s *BindMachineResponse) SetBody(v []BindMachineResponseBody) *BindMachineResponse {
	s.Body = &v
	return s
}


type BindMachineResponseBuilder struct {
	s *BindMachineResponse
}

func NewBindMachineResponseBuilder() *BindMachineResponseBuilder {
	s := &BindMachineResponse{}
	b := &BindMachineResponseBuilder{s: s}
	return b
}

func (b *BindMachineResponseBuilder) RequestId(v string) *BindMachineResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BindMachineResponseBuilder) ErrorMessage(v string) *BindMachineResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BindMachineResponseBuilder) ErrorCode(v string) *BindMachineResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BindMachineResponseBuilder) State(v BindMachineResponseStateEnum) *BindMachineResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BindMachineResponseBuilder) Body(v []BindMachineResponseBody) *BindMachineResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BindMachineResponseBuilder) Build() *BindMachineResponse {
	return b.s
}


