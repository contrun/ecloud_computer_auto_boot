// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchShutdownResponseStateEnum string

// List of State
const (
    BatchShutdownResponseStateEnumOk BatchShutdownResponseStateEnum = "OK"
    BatchShutdownResponseStateEnumError BatchShutdownResponseStateEnum = "ERROR"
    BatchShutdownResponseStateEnumException BatchShutdownResponseStateEnum = "EXCEPTION"
    BatchShutdownResponseStateEnumForbidden BatchShutdownResponseStateEnum = "FORBIDDEN"
    BatchShutdownResponseStateEnumRemaining BatchShutdownResponseStateEnum = "REMAINING"
    BatchShutdownResponseStateEnumTimeout BatchShutdownResponseStateEnum = "TIMEOUT"
)

type BatchShutdownResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchShutdownResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *BatchShutdownResponseBody `json:"body,omitempty"`
}

func (s BatchShutdownResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchShutdownResponse) GoString() string {
	return s.String()
}

func (s BatchShutdownResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchShutdownResponse) SetRequestId(v string) *BatchShutdownResponse {
	s.RequestId = &v
	return s
}

func (s *BatchShutdownResponse) SetErrorMessage(v string) *BatchShutdownResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchShutdownResponse) SetErrorCode(v string) *BatchShutdownResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchShutdownResponse) SetState(v BatchShutdownResponseStateEnum) *BatchShutdownResponse {
	s.State = &v
	return s
}

func (s *BatchShutdownResponse) SetBody(v *BatchShutdownResponseBody) *BatchShutdownResponse {
	s.Body = v
	return s
}


type BatchShutdownResponseBuilder struct {
	s *BatchShutdownResponse
}

func NewBatchShutdownResponseBuilder() *BatchShutdownResponseBuilder {
	s := &BatchShutdownResponse{}
	b := &BatchShutdownResponseBuilder{s: s}
	return b
}

func (b *BatchShutdownResponseBuilder) RequestId(v string) *BatchShutdownResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchShutdownResponseBuilder) ErrorMessage(v string) *BatchShutdownResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchShutdownResponseBuilder) ErrorCode(v string) *BatchShutdownResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchShutdownResponseBuilder) State(v BatchShutdownResponseStateEnum) *BatchShutdownResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchShutdownResponseBuilder) Body(v *BatchShutdownResponseBody) *BatchShutdownResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BatchShutdownResponseBuilder) Build() *BatchShutdownResponse {
	return b.s
}


