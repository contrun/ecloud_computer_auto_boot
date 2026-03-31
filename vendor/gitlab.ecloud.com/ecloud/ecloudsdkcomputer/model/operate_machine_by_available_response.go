// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type OperateMachineByAvailableResponseStateEnum string

// List of State
const (
    OperateMachineByAvailableResponseStateEnumOk OperateMachineByAvailableResponseStateEnum = "OK"
    OperateMachineByAvailableResponseStateEnumError OperateMachineByAvailableResponseStateEnum = "ERROR"
    OperateMachineByAvailableResponseStateEnumException OperateMachineByAvailableResponseStateEnum = "EXCEPTION"
    OperateMachineByAvailableResponseStateEnumForbidden OperateMachineByAvailableResponseStateEnum = "FORBIDDEN"
    OperateMachineByAvailableResponseStateEnumRemaining OperateMachineByAvailableResponseStateEnum = "REMAINING"
    OperateMachineByAvailableResponseStateEnumTimeout OperateMachineByAvailableResponseStateEnum = "TIMEOUT"
)

type OperateMachineByAvailableResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *OperateMachineByAvailableResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`
}

func (s OperateMachineByAvailableResponse) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByAvailableResponse) GoString() string {
	return s.String()
}

func (s OperateMachineByAvailableResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByAvailableResponse) SetRequestId(v string) *OperateMachineByAvailableResponse {
	s.RequestId = &v
	return s
}

func (s *OperateMachineByAvailableResponse) SetErrorMessage(v string) *OperateMachineByAvailableResponse {
	s.ErrorMessage = &v
	return s
}

func (s *OperateMachineByAvailableResponse) SetErrorCode(v string) *OperateMachineByAvailableResponse {
	s.ErrorCode = &v
	return s
}

func (s *OperateMachineByAvailableResponse) SetState(v OperateMachineByAvailableResponseStateEnum) *OperateMachineByAvailableResponse {
	s.State = &v
	return s
}

func (s *OperateMachineByAvailableResponse) SetBody(v interface{}) *OperateMachineByAvailableResponse {
	s.Body = v
	return s
}


type OperateMachineByAvailableResponseBuilder struct {
	s *OperateMachineByAvailableResponse
}

func NewOperateMachineByAvailableResponseBuilder() *OperateMachineByAvailableResponseBuilder {
	s := &OperateMachineByAvailableResponse{}
	b := &OperateMachineByAvailableResponseBuilder{s: s}
	return b
}

func (b *OperateMachineByAvailableResponseBuilder) RequestId(v string) *OperateMachineByAvailableResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *OperateMachineByAvailableResponseBuilder) ErrorMessage(v string) *OperateMachineByAvailableResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *OperateMachineByAvailableResponseBuilder) ErrorCode(v string) *OperateMachineByAvailableResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *OperateMachineByAvailableResponseBuilder) State(v OperateMachineByAvailableResponseStateEnum) *OperateMachineByAvailableResponseBuilder {
	b.s.State = &v
	return b
}

func (b *OperateMachineByAvailableResponseBuilder) Body(v interface{}) *OperateMachineByAvailableResponseBuilder {
	b.s.Body = v
	return b
}

func (b *OperateMachineByAvailableResponseBuilder) Build() *OperateMachineByAvailableResponse {
	return b.s
}


