// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteUserPolicyResponseStateEnum string

// List of State
const (
    DeleteUserPolicyResponseStateEnumOk DeleteUserPolicyResponseStateEnum = "OK"
    DeleteUserPolicyResponseStateEnumError DeleteUserPolicyResponseStateEnum = "ERROR"
    DeleteUserPolicyResponseStateEnumException DeleteUserPolicyResponseStateEnum = "EXCEPTION"
    DeleteUserPolicyResponseStateEnumForbidden DeleteUserPolicyResponseStateEnum = "FORBIDDEN"
    DeleteUserPolicyResponseStateEnumRemaining DeleteUserPolicyResponseStateEnum = "REMAINING"
    DeleteUserPolicyResponseStateEnumTimeout DeleteUserPolicyResponseStateEnum = "TIMEOUT"
)

type DeleteUserPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *DeleteUserPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s DeleteUserPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteUserPolicyResponse) GoString() string {
	return s.String()
}

func (s DeleteUserPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteUserPolicyResponse) SetRequestId(v string) *DeleteUserPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteUserPolicyResponse) SetErrorMessage(v string) *DeleteUserPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUserPolicyResponse) SetErrorCode(v string) *DeleteUserPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUserPolicyResponse) SetState(v DeleteUserPolicyResponseStateEnum) *DeleteUserPolicyResponse {
	s.State = &v
	return s
}

func (s *DeleteUserPolicyResponse) SetBody(v string) *DeleteUserPolicyResponse {
	s.Body = &v
	return s
}


type DeleteUserPolicyResponseBuilder struct {
	s *DeleteUserPolicyResponse
}

func NewDeleteUserPolicyResponseBuilder() *DeleteUserPolicyResponseBuilder {
	s := &DeleteUserPolicyResponse{}
	b := &DeleteUserPolicyResponseBuilder{s: s}
	return b
}

func (b *DeleteUserPolicyResponseBuilder) RequestId(v string) *DeleteUserPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteUserPolicyResponseBuilder) ErrorMessage(v string) *DeleteUserPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteUserPolicyResponseBuilder) ErrorCode(v string) *DeleteUserPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteUserPolicyResponseBuilder) State(v DeleteUserPolicyResponseStateEnum) *DeleteUserPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteUserPolicyResponseBuilder) Body(v string) *DeleteUserPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteUserPolicyResponseBuilder) Build() *DeleteUserPolicyResponse {
	return b.s
}


