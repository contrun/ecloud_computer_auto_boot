// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UnBindMachineResponseStateEnum string

// List of State
const (
    UnBindMachineResponseStateEnumOk UnBindMachineResponseStateEnum = "OK"
    UnBindMachineResponseStateEnumError UnBindMachineResponseStateEnum = "ERROR"
    UnBindMachineResponseStateEnumException UnBindMachineResponseStateEnum = "EXCEPTION"
    UnBindMachineResponseStateEnumForbidden UnBindMachineResponseStateEnum = "FORBIDDEN"
    UnBindMachineResponseStateEnumRemaining UnBindMachineResponseStateEnum = "REMAINING"
    UnBindMachineResponseStateEnumTimeout UnBindMachineResponseStateEnum = "TIMEOUT"
)

type UnBindMachineResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *UnBindMachineResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]UnBindMachineResponseBody `json:"body,omitempty"`
}

func (s UnBindMachineResponse) String() string {
	return utils.Beautify(s)
}

func (s UnBindMachineResponse) GoString() string {
	return s.String()
}

func (s UnBindMachineResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnBindMachineResponse) SetRequestId(v string) *UnBindMachineResponse {
	s.RequestId = &v
	return s
}

func (s *UnBindMachineResponse) SetErrorMessage(v string) *UnBindMachineResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UnBindMachineResponse) SetErrorCode(v string) *UnBindMachineResponse {
	s.ErrorCode = &v
	return s
}

func (s *UnBindMachineResponse) SetState(v UnBindMachineResponseStateEnum) *UnBindMachineResponse {
	s.State = &v
	return s
}

func (s *UnBindMachineResponse) SetBody(v []UnBindMachineResponseBody) *UnBindMachineResponse {
	s.Body = &v
	return s
}


type UnBindMachineResponseBuilder struct {
	s *UnBindMachineResponse
}

func NewUnBindMachineResponseBuilder() *UnBindMachineResponseBuilder {
	s := &UnBindMachineResponse{}
	b := &UnBindMachineResponseBuilder{s: s}
	return b
}

func (b *UnBindMachineResponseBuilder) RequestId(v string) *UnBindMachineResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UnBindMachineResponseBuilder) ErrorMessage(v string) *UnBindMachineResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UnBindMachineResponseBuilder) ErrorCode(v string) *UnBindMachineResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UnBindMachineResponseBuilder) State(v UnBindMachineResponseStateEnum) *UnBindMachineResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UnBindMachineResponseBuilder) Body(v []UnBindMachineResponseBody) *UnBindMachineResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UnBindMachineResponseBuilder) Build() *UnBindMachineResponse {
	return b.s
}


