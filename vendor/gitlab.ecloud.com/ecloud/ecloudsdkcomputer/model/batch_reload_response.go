// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchReloadResponseStateEnum string

// List of State
const (
    BatchReloadResponseStateEnumOk BatchReloadResponseStateEnum = "OK"
    BatchReloadResponseStateEnumError BatchReloadResponseStateEnum = "ERROR"
    BatchReloadResponseStateEnumException BatchReloadResponseStateEnum = "EXCEPTION"
    BatchReloadResponseStateEnumForbidden BatchReloadResponseStateEnum = "FORBIDDEN"
    BatchReloadResponseStateEnumRemaining BatchReloadResponseStateEnum = "REMAINING"
    BatchReloadResponseStateEnumTimeout BatchReloadResponseStateEnum = "TIMEOUT"
)

type BatchReloadResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchReloadResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *BatchReloadResponseBody `json:"body,omitempty"`
}

func (s BatchReloadResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchReloadResponse) GoString() string {
	return s.String()
}

func (s BatchReloadResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchReloadResponse) SetRequestId(v string) *BatchReloadResponse {
	s.RequestId = &v
	return s
}

func (s *BatchReloadResponse) SetErrorMessage(v string) *BatchReloadResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchReloadResponse) SetErrorCode(v string) *BatchReloadResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchReloadResponse) SetState(v BatchReloadResponseStateEnum) *BatchReloadResponse {
	s.State = &v
	return s
}

func (s *BatchReloadResponse) SetBody(v *BatchReloadResponseBody) *BatchReloadResponse {
	s.Body = v
	return s
}


type BatchReloadResponseBuilder struct {
	s *BatchReloadResponse
}

func NewBatchReloadResponseBuilder() *BatchReloadResponseBuilder {
	s := &BatchReloadResponse{}
	b := &BatchReloadResponseBuilder{s: s}
	return b
}

func (b *BatchReloadResponseBuilder) RequestId(v string) *BatchReloadResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchReloadResponseBuilder) ErrorMessage(v string) *BatchReloadResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchReloadResponseBuilder) ErrorCode(v string) *BatchReloadResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchReloadResponseBuilder) State(v BatchReloadResponseStateEnum) *BatchReloadResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchReloadResponseBuilder) Body(v *BatchReloadResponseBody) *BatchReloadResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BatchReloadResponseBuilder) Build() *BatchReloadResponse {
	return b.s
}


