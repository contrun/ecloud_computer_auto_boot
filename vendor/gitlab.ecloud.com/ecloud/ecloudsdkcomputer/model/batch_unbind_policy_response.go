// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchUnbindPolicyResponseStateEnum string

// List of State
const (
    BatchUnbindPolicyResponseStateEnumOk BatchUnbindPolicyResponseStateEnum = "OK"
    BatchUnbindPolicyResponseStateEnumError BatchUnbindPolicyResponseStateEnum = "ERROR"
    BatchUnbindPolicyResponseStateEnumException BatchUnbindPolicyResponseStateEnum = "EXCEPTION"
    BatchUnbindPolicyResponseStateEnumForbidden BatchUnbindPolicyResponseStateEnum = "FORBIDDEN"
    BatchUnbindPolicyResponseStateEnumRemaining BatchUnbindPolicyResponseStateEnum = "REMAINING"
    BatchUnbindPolicyResponseStateEnumTimeout BatchUnbindPolicyResponseStateEnum = "TIMEOUT"
)

type BatchUnbindPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchUnbindPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s BatchUnbindPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchUnbindPolicyResponse) GoString() string {
	return s.String()
}

func (s BatchUnbindPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUnbindPolicyResponse) SetRequestId(v string) *BatchUnbindPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindPolicyResponse) SetErrorMessage(v string) *BatchUnbindPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchUnbindPolicyResponse) SetErrorCode(v string) *BatchUnbindPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchUnbindPolicyResponse) SetState(v BatchUnbindPolicyResponseStateEnum) *BatchUnbindPolicyResponse {
	s.State = &v
	return s
}

func (s *BatchUnbindPolicyResponse) SetBody(v string) *BatchUnbindPolicyResponse {
	s.Body = &v
	return s
}


type BatchUnbindPolicyResponseBuilder struct {
	s *BatchUnbindPolicyResponse
}

func NewBatchUnbindPolicyResponseBuilder() *BatchUnbindPolicyResponseBuilder {
	s := &BatchUnbindPolicyResponse{}
	b := &BatchUnbindPolicyResponseBuilder{s: s}
	return b
}

func (b *BatchUnbindPolicyResponseBuilder) RequestId(v string) *BatchUnbindPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchUnbindPolicyResponseBuilder) ErrorMessage(v string) *BatchUnbindPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchUnbindPolicyResponseBuilder) ErrorCode(v string) *BatchUnbindPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchUnbindPolicyResponseBuilder) State(v BatchUnbindPolicyResponseStateEnum) *BatchUnbindPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchUnbindPolicyResponseBuilder) Body(v string) *BatchUnbindPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchUnbindPolicyResponseBuilder) Build() *BatchUnbindPolicyResponse {
	return b.s
}


