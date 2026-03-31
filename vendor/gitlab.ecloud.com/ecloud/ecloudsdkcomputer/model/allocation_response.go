// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AllocationResponseStateEnum string

// List of State
const (
    AllocationResponseStateEnumOk AllocationResponseStateEnum = "OK"
    AllocationResponseStateEnumError AllocationResponseStateEnum = "ERROR"
    AllocationResponseStateEnumException AllocationResponseStateEnum = "EXCEPTION"
    AllocationResponseStateEnumForbidden AllocationResponseStateEnum = "FORBIDDEN"
    AllocationResponseStateEnumRemaining AllocationResponseStateEnum = "REMAINING"
    AllocationResponseStateEnumTimeout AllocationResponseStateEnum = "TIMEOUT"
)

type AllocationResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *AllocationResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`
}

func (s AllocationResponse) String() string {
	return utils.Beautify(s)
}

func (s AllocationResponse) GoString() string {
	return s.String()
}

func (s AllocationResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AllocationResponse) SetRequestId(v string) *AllocationResponse {
	s.RequestId = &v
	return s
}

func (s *AllocationResponse) SetErrorMessage(v string) *AllocationResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AllocationResponse) SetErrorCode(v string) *AllocationResponse {
	s.ErrorCode = &v
	return s
}

func (s *AllocationResponse) SetState(v AllocationResponseStateEnum) *AllocationResponse {
	s.State = &v
	return s
}

func (s *AllocationResponse) SetBody(v interface{}) *AllocationResponse {
	s.Body = v
	return s
}


type AllocationResponseBuilder struct {
	s *AllocationResponse
}

func NewAllocationResponseBuilder() *AllocationResponseBuilder {
	s := &AllocationResponse{}
	b := &AllocationResponseBuilder{s: s}
	return b
}

func (b *AllocationResponseBuilder) RequestId(v string) *AllocationResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AllocationResponseBuilder) ErrorMessage(v string) *AllocationResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AllocationResponseBuilder) ErrorCode(v string) *AllocationResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AllocationResponseBuilder) State(v AllocationResponseStateEnum) *AllocationResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AllocationResponseBuilder) Body(v interface{}) *AllocationResponseBuilder {
	b.s.Body = v
	return b
}

func (b *AllocationResponseBuilder) Build() *AllocationResponse {
	return b.s
}


