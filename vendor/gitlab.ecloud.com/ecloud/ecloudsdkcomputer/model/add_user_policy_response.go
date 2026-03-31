// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddUserPolicyResponseStateEnum string

// List of State
const (
    AddUserPolicyResponseStateEnumOk AddUserPolicyResponseStateEnum = "OK"
    AddUserPolicyResponseStateEnumError AddUserPolicyResponseStateEnum = "ERROR"
    AddUserPolicyResponseStateEnumException AddUserPolicyResponseStateEnum = "EXCEPTION"
    AddUserPolicyResponseStateEnumForbidden AddUserPolicyResponseStateEnum = "FORBIDDEN"
    AddUserPolicyResponseStateEnumRemaining AddUserPolicyResponseStateEnum = "REMAINING"
    AddUserPolicyResponseStateEnumTimeout AddUserPolicyResponseStateEnum = "TIMEOUT"
)

type AddUserPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *AddUserPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s AddUserPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s AddUserPolicyResponse) GoString() string {
	return s.String()
}

func (s AddUserPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddUserPolicyResponse) SetRequestId(v string) *AddUserPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *AddUserPolicyResponse) SetErrorMessage(v string) *AddUserPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddUserPolicyResponse) SetErrorCode(v string) *AddUserPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddUserPolicyResponse) SetState(v AddUserPolicyResponseStateEnum) *AddUserPolicyResponse {
	s.State = &v
	return s
}

func (s *AddUserPolicyResponse) SetBody(v string) *AddUserPolicyResponse {
	s.Body = &v
	return s
}


type AddUserPolicyResponseBuilder struct {
	s *AddUserPolicyResponse
}

func NewAddUserPolicyResponseBuilder() *AddUserPolicyResponseBuilder {
	s := &AddUserPolicyResponse{}
	b := &AddUserPolicyResponseBuilder{s: s}
	return b
}

func (b *AddUserPolicyResponseBuilder) RequestId(v string) *AddUserPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddUserPolicyResponseBuilder) ErrorMessage(v string) *AddUserPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddUserPolicyResponseBuilder) ErrorCode(v string) *AddUserPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddUserPolicyResponseBuilder) State(v AddUserPolicyResponseStateEnum) *AddUserPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddUserPolicyResponseBuilder) Body(v string) *AddUserPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AddUserPolicyResponseBuilder) Build() *AddUserPolicyResponse {
	return b.s
}


