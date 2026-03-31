// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchCreateImageResponseStateEnum string

// List of State
const (
    BatchCreateImageResponseStateEnumOk BatchCreateImageResponseStateEnum = "OK"
    BatchCreateImageResponseStateEnumError BatchCreateImageResponseStateEnum = "ERROR"
    BatchCreateImageResponseStateEnumException BatchCreateImageResponseStateEnum = "EXCEPTION"
    BatchCreateImageResponseStateEnumForbidden BatchCreateImageResponseStateEnum = "FORBIDDEN"
    BatchCreateImageResponseStateEnumRemaining BatchCreateImageResponseStateEnum = "REMAINING"
    BatchCreateImageResponseStateEnumTimeout BatchCreateImageResponseStateEnum = "TIMEOUT"
)

type BatchCreateImageResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchCreateImageResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]BatchCreateImageResponseBody `json:"body,omitempty"`
}

func (s BatchCreateImageResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchCreateImageResponse) GoString() string {
	return s.String()
}

func (s BatchCreateImageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchCreateImageResponse) SetRequestId(v string) *BatchCreateImageResponse {
	s.RequestId = &v
	return s
}

func (s *BatchCreateImageResponse) SetErrorMessage(v string) *BatchCreateImageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchCreateImageResponse) SetErrorCode(v string) *BatchCreateImageResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchCreateImageResponse) SetState(v BatchCreateImageResponseStateEnum) *BatchCreateImageResponse {
	s.State = &v
	return s
}

func (s *BatchCreateImageResponse) SetBody(v []BatchCreateImageResponseBody) *BatchCreateImageResponse {
	s.Body = &v
	return s
}


type BatchCreateImageResponseBuilder struct {
	s *BatchCreateImageResponse
}

func NewBatchCreateImageResponseBuilder() *BatchCreateImageResponseBuilder {
	s := &BatchCreateImageResponse{}
	b := &BatchCreateImageResponseBuilder{s: s}
	return b
}

func (b *BatchCreateImageResponseBuilder) RequestId(v string) *BatchCreateImageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchCreateImageResponseBuilder) ErrorMessage(v string) *BatchCreateImageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchCreateImageResponseBuilder) ErrorCode(v string) *BatchCreateImageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchCreateImageResponseBuilder) State(v BatchCreateImageResponseStateEnum) *BatchCreateImageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchCreateImageResponseBuilder) Body(v []BatchCreateImageResponseBody) *BatchCreateImageResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchCreateImageResponseBuilder) Build() *BatchCreateImageResponse {
	return b.s
}


