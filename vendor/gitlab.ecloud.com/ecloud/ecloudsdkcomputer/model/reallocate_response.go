// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ReallocateResponseStateEnum string

// List of State
const (
    ReallocateResponseStateEnumOk ReallocateResponseStateEnum = "OK"
    ReallocateResponseStateEnumError ReallocateResponseStateEnum = "ERROR"
    ReallocateResponseStateEnumException ReallocateResponseStateEnum = "EXCEPTION"
    ReallocateResponseStateEnumForbidden ReallocateResponseStateEnum = "FORBIDDEN"
    ReallocateResponseStateEnumRemaining ReallocateResponseStateEnum = "REMAINING"
    ReallocateResponseStateEnumTimeout ReallocateResponseStateEnum = "TIMEOUT"
)

type ReallocateResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ReallocateResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`
}

func (s ReallocateResponse) String() string {
	return utils.Beautify(s)
}

func (s ReallocateResponse) GoString() string {
	return s.String()
}

func (s ReallocateResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ReallocateResponse) SetRequestId(v string) *ReallocateResponse {
	s.RequestId = &v
	return s
}

func (s *ReallocateResponse) SetErrorMessage(v string) *ReallocateResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ReallocateResponse) SetErrorCode(v string) *ReallocateResponse {
	s.ErrorCode = &v
	return s
}

func (s *ReallocateResponse) SetState(v ReallocateResponseStateEnum) *ReallocateResponse {
	s.State = &v
	return s
}

func (s *ReallocateResponse) SetBody(v interface{}) *ReallocateResponse {
	s.Body = v
	return s
}


type ReallocateResponseBuilder struct {
	s *ReallocateResponse
}

func NewReallocateResponseBuilder() *ReallocateResponseBuilder {
	s := &ReallocateResponse{}
	b := &ReallocateResponseBuilder{s: s}
	return b
}

func (b *ReallocateResponseBuilder) RequestId(v string) *ReallocateResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ReallocateResponseBuilder) ErrorMessage(v string) *ReallocateResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ReallocateResponseBuilder) ErrorCode(v string) *ReallocateResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ReallocateResponseBuilder) State(v ReallocateResponseStateEnum) *ReallocateResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ReallocateResponseBuilder) Body(v interface{}) *ReallocateResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ReallocateResponseBuilder) Build() *ReallocateResponse {
	return b.s
}


