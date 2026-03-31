// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeletePolicyResponseStateEnum string

// List of State
const (
    DeletePolicyResponseStateEnumOk DeletePolicyResponseStateEnum = "OK"
    DeletePolicyResponseStateEnumError DeletePolicyResponseStateEnum = "ERROR"
    DeletePolicyResponseStateEnumException DeletePolicyResponseStateEnum = "EXCEPTION"
    DeletePolicyResponseStateEnumForbidden DeletePolicyResponseStateEnum = "FORBIDDEN"
    DeletePolicyResponseStateEnumRemaining DeletePolicyResponseStateEnum = "REMAINING"
    DeletePolicyResponseStateEnumTimeout DeletePolicyResponseStateEnum = "TIMEOUT"
)

type DeletePolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *DeletePolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s DeletePolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s DeletePolicyResponse) GoString() string {
	return s.String()
}

func (s DeletePolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePolicyResponse) SetRequestId(v string) *DeletePolicyResponse {
	s.RequestId = &v
	return s
}

func (s *DeletePolicyResponse) SetErrorMessage(v string) *DeletePolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePolicyResponse) SetErrorCode(v string) *DeletePolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeletePolicyResponse) SetState(v DeletePolicyResponseStateEnum) *DeletePolicyResponse {
	s.State = &v
	return s
}

func (s *DeletePolicyResponse) SetBody(v string) *DeletePolicyResponse {
	s.Body = &v
	return s
}


type DeletePolicyResponseBuilder struct {
	s *DeletePolicyResponse
}

func NewDeletePolicyResponseBuilder() *DeletePolicyResponseBuilder {
	s := &DeletePolicyResponse{}
	b := &DeletePolicyResponseBuilder{s: s}
	return b
}

func (b *DeletePolicyResponseBuilder) RequestId(v string) *DeletePolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeletePolicyResponseBuilder) ErrorMessage(v string) *DeletePolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeletePolicyResponseBuilder) ErrorCode(v string) *DeletePolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeletePolicyResponseBuilder) State(v DeletePolicyResponseStateEnum) *DeletePolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeletePolicyResponseBuilder) Body(v string) *DeletePolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeletePolicyResponseBuilder) Build() *DeletePolicyResponse {
	return b.s
}


