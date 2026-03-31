// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchRestartResponseStateEnum string

// List of State
const (
    BatchRestartResponseStateEnumOk BatchRestartResponseStateEnum = "OK"
    BatchRestartResponseStateEnumError BatchRestartResponseStateEnum = "ERROR"
    BatchRestartResponseStateEnumException BatchRestartResponseStateEnum = "EXCEPTION"
    BatchRestartResponseStateEnumForbidden BatchRestartResponseStateEnum = "FORBIDDEN"
    BatchRestartResponseStateEnumRemaining BatchRestartResponseStateEnum = "REMAINING"
    BatchRestartResponseStateEnumTimeout BatchRestartResponseStateEnum = "TIMEOUT"
)

type BatchRestartResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchRestartResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *BatchRestartResponseBody `json:"body,omitempty"`
}

func (s BatchRestartResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchRestartResponse) GoString() string {
	return s.String()
}

func (s BatchRestartResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchRestartResponse) SetRequestId(v string) *BatchRestartResponse {
	s.RequestId = &v
	return s
}

func (s *BatchRestartResponse) SetErrorMessage(v string) *BatchRestartResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchRestartResponse) SetErrorCode(v string) *BatchRestartResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchRestartResponse) SetState(v BatchRestartResponseStateEnum) *BatchRestartResponse {
	s.State = &v
	return s
}

func (s *BatchRestartResponse) SetBody(v *BatchRestartResponseBody) *BatchRestartResponse {
	s.Body = v
	return s
}


type BatchRestartResponseBuilder struct {
	s *BatchRestartResponse
}

func NewBatchRestartResponseBuilder() *BatchRestartResponseBuilder {
	s := &BatchRestartResponse{}
	b := &BatchRestartResponseBuilder{s: s}
	return b
}

func (b *BatchRestartResponseBuilder) RequestId(v string) *BatchRestartResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchRestartResponseBuilder) ErrorMessage(v string) *BatchRestartResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchRestartResponseBuilder) ErrorCode(v string) *BatchRestartResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchRestartResponseBuilder) State(v BatchRestartResponseStateEnum) *BatchRestartResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchRestartResponseBuilder) Body(v *BatchRestartResponseBody) *BatchRestartResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BatchRestartResponseBuilder) Build() *BatchRestartResponse {
	return b.s
}


