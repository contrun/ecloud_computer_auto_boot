// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchBindPolicyResponseStateEnum string

// List of State
const (
    BatchBindPolicyResponseStateEnumOk BatchBindPolicyResponseStateEnum = "OK"
    BatchBindPolicyResponseStateEnumError BatchBindPolicyResponseStateEnum = "ERROR"
    BatchBindPolicyResponseStateEnumException BatchBindPolicyResponseStateEnum = "EXCEPTION"
    BatchBindPolicyResponseStateEnumForbidden BatchBindPolicyResponseStateEnum = "FORBIDDEN"
    BatchBindPolicyResponseStateEnumRemaining BatchBindPolicyResponseStateEnum = "REMAINING"
    BatchBindPolicyResponseStateEnumTimeout BatchBindPolicyResponseStateEnum = "TIMEOUT"
)

type BatchBindPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchBindPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s BatchBindPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchBindPolicyResponse) GoString() string {
	return s.String()
}

func (s BatchBindPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchBindPolicyResponse) SetRequestId(v string) *BatchBindPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *BatchBindPolicyResponse) SetErrorMessage(v string) *BatchBindPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchBindPolicyResponse) SetErrorCode(v string) *BatchBindPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchBindPolicyResponse) SetState(v BatchBindPolicyResponseStateEnum) *BatchBindPolicyResponse {
	s.State = &v
	return s
}

func (s *BatchBindPolicyResponse) SetBody(v string) *BatchBindPolicyResponse {
	s.Body = &v
	return s
}


type BatchBindPolicyResponseBuilder struct {
	s *BatchBindPolicyResponse
}

func NewBatchBindPolicyResponseBuilder() *BatchBindPolicyResponseBuilder {
	s := &BatchBindPolicyResponse{}
	b := &BatchBindPolicyResponseBuilder{s: s}
	return b
}

func (b *BatchBindPolicyResponseBuilder) RequestId(v string) *BatchBindPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchBindPolicyResponseBuilder) ErrorMessage(v string) *BatchBindPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchBindPolicyResponseBuilder) ErrorCode(v string) *BatchBindPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchBindPolicyResponseBuilder) State(v BatchBindPolicyResponseStateEnum) *BatchBindPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchBindPolicyResponseBuilder) Body(v string) *BatchBindPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchBindPolicyResponseBuilder) Build() *BatchBindPolicyResponse {
	return b.s
}


