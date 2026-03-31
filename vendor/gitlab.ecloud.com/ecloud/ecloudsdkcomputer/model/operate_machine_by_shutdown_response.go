// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type OperateMachineByShutdownResponseStateEnum string

// List of State
const (
    OperateMachineByShutdownResponseStateEnumOk OperateMachineByShutdownResponseStateEnum = "OK"
    OperateMachineByShutdownResponseStateEnumError OperateMachineByShutdownResponseStateEnum = "ERROR"
    OperateMachineByShutdownResponseStateEnumException OperateMachineByShutdownResponseStateEnum = "EXCEPTION"
    OperateMachineByShutdownResponseStateEnumForbidden OperateMachineByShutdownResponseStateEnum = "FORBIDDEN"
    OperateMachineByShutdownResponseStateEnumRemaining OperateMachineByShutdownResponseStateEnum = "REMAINING"
    OperateMachineByShutdownResponseStateEnumTimeout OperateMachineByShutdownResponseStateEnum = "TIMEOUT"
)

type OperateMachineByShutdownResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *OperateMachineByShutdownResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`
}

func (s OperateMachineByShutdownResponse) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByShutdownResponse) GoString() string {
	return s.String()
}

func (s OperateMachineByShutdownResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByShutdownResponse) SetRequestId(v string) *OperateMachineByShutdownResponse {
	s.RequestId = &v
	return s
}

func (s *OperateMachineByShutdownResponse) SetErrorMessage(v string) *OperateMachineByShutdownResponse {
	s.ErrorMessage = &v
	return s
}

func (s *OperateMachineByShutdownResponse) SetErrorCode(v string) *OperateMachineByShutdownResponse {
	s.ErrorCode = &v
	return s
}

func (s *OperateMachineByShutdownResponse) SetState(v OperateMachineByShutdownResponseStateEnum) *OperateMachineByShutdownResponse {
	s.State = &v
	return s
}

func (s *OperateMachineByShutdownResponse) SetBody(v interface{}) *OperateMachineByShutdownResponse {
	s.Body = v
	return s
}


type OperateMachineByShutdownResponseBuilder struct {
	s *OperateMachineByShutdownResponse
}

func NewOperateMachineByShutdownResponseBuilder() *OperateMachineByShutdownResponseBuilder {
	s := &OperateMachineByShutdownResponse{}
	b := &OperateMachineByShutdownResponseBuilder{s: s}
	return b
}

func (b *OperateMachineByShutdownResponseBuilder) RequestId(v string) *OperateMachineByShutdownResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *OperateMachineByShutdownResponseBuilder) ErrorMessage(v string) *OperateMachineByShutdownResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *OperateMachineByShutdownResponseBuilder) ErrorCode(v string) *OperateMachineByShutdownResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *OperateMachineByShutdownResponseBuilder) State(v OperateMachineByShutdownResponseStateEnum) *OperateMachineByShutdownResponseBuilder {
	b.s.State = &v
	return b
}

func (b *OperateMachineByShutdownResponseBuilder) Body(v interface{}) *OperateMachineByShutdownResponseBuilder {
	b.s.Body = v
	return b
}

func (b *OperateMachineByShutdownResponseBuilder) Build() *OperateMachineByShutdownResponse {
	return b.s
}


