// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteComputerPolicyResponseStateEnum string

// List of State
const (
    DeleteComputerPolicyResponseStateEnumOk DeleteComputerPolicyResponseStateEnum = "OK"
    DeleteComputerPolicyResponseStateEnumError DeleteComputerPolicyResponseStateEnum = "ERROR"
    DeleteComputerPolicyResponseStateEnumException DeleteComputerPolicyResponseStateEnum = "EXCEPTION"
    DeleteComputerPolicyResponseStateEnumForbidden DeleteComputerPolicyResponseStateEnum = "FORBIDDEN"
    DeleteComputerPolicyResponseStateEnumRemaining DeleteComputerPolicyResponseStateEnum = "REMAINING"
    DeleteComputerPolicyResponseStateEnumTimeout DeleteComputerPolicyResponseStateEnum = "TIMEOUT"
)

type DeleteComputerPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *DeleteComputerPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s DeleteComputerPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteComputerPolicyResponse) GoString() string {
	return s.String()
}

func (s DeleteComputerPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteComputerPolicyResponse) SetRequestId(v string) *DeleteComputerPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteComputerPolicyResponse) SetErrorMessage(v string) *DeleteComputerPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteComputerPolicyResponse) SetErrorCode(v string) *DeleteComputerPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteComputerPolicyResponse) SetState(v DeleteComputerPolicyResponseStateEnum) *DeleteComputerPolicyResponse {
	s.State = &v
	return s
}

func (s *DeleteComputerPolicyResponse) SetBody(v bool) *DeleteComputerPolicyResponse {
	s.Body = &v
	return s
}


type DeleteComputerPolicyResponseBuilder struct {
	s *DeleteComputerPolicyResponse
}

func NewDeleteComputerPolicyResponseBuilder() *DeleteComputerPolicyResponseBuilder {
	s := &DeleteComputerPolicyResponse{}
	b := &DeleteComputerPolicyResponseBuilder{s: s}
	return b
}

func (b *DeleteComputerPolicyResponseBuilder) RequestId(v string) *DeleteComputerPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteComputerPolicyResponseBuilder) ErrorMessage(v string) *DeleteComputerPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteComputerPolicyResponseBuilder) ErrorCode(v string) *DeleteComputerPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteComputerPolicyResponseBuilder) State(v DeleteComputerPolicyResponseStateEnum) *DeleteComputerPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteComputerPolicyResponseBuilder) Body(v bool) *DeleteComputerPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteComputerPolicyResponseBuilder) Build() *DeleteComputerPolicyResponse {
	return b.s
}


