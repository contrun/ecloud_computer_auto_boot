// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchDeleteUserResponseStateEnum string

// List of State
const (
    BatchDeleteUserResponseStateEnumOk BatchDeleteUserResponseStateEnum = "OK"
    BatchDeleteUserResponseStateEnumError BatchDeleteUserResponseStateEnum = "ERROR"
    BatchDeleteUserResponseStateEnumException BatchDeleteUserResponseStateEnum = "EXCEPTION"
    BatchDeleteUserResponseStateEnumForbidden BatchDeleteUserResponseStateEnum = "FORBIDDEN"
    BatchDeleteUserResponseStateEnumRemaining BatchDeleteUserResponseStateEnum = "REMAINING"
    BatchDeleteUserResponseStateEnumTimeout BatchDeleteUserResponseStateEnum = "TIMEOUT"
)

type BatchDeleteUserResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchDeleteUserResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`
}

func (s BatchDeleteUserResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteUserResponse) GoString() string {
	return s.String()
}

func (s BatchDeleteUserResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteUserResponse) SetRequestId(v string) *BatchDeleteUserResponse {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteUserResponse) SetErrorMessage(v string) *BatchDeleteUserResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchDeleteUserResponse) SetErrorCode(v string) *BatchDeleteUserResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchDeleteUserResponse) SetState(v BatchDeleteUserResponseStateEnum) *BatchDeleteUserResponse {
	s.State = &v
	return s
}

func (s *BatchDeleteUserResponse) SetBody(v interface{}) *BatchDeleteUserResponse {
	s.Body = v
	return s
}


type BatchDeleteUserResponseBuilder struct {
	s *BatchDeleteUserResponse
}

func NewBatchDeleteUserResponseBuilder() *BatchDeleteUserResponseBuilder {
	s := &BatchDeleteUserResponse{}
	b := &BatchDeleteUserResponseBuilder{s: s}
	return b
}

func (b *BatchDeleteUserResponseBuilder) RequestId(v string) *BatchDeleteUserResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchDeleteUserResponseBuilder) ErrorMessage(v string) *BatchDeleteUserResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchDeleteUserResponseBuilder) ErrorCode(v string) *BatchDeleteUserResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchDeleteUserResponseBuilder) State(v BatchDeleteUserResponseStateEnum) *BatchDeleteUserResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchDeleteUserResponseBuilder) Body(v interface{}) *BatchDeleteUserResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BatchDeleteUserResponseBuilder) Build() *BatchDeleteUserResponse {
	return b.s
}


