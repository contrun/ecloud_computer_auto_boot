// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchReloadByImageResponseStateEnum string

// List of State
const (
    BatchReloadByImageResponseStateEnumOk BatchReloadByImageResponseStateEnum = "OK"
    BatchReloadByImageResponseStateEnumError BatchReloadByImageResponseStateEnum = "ERROR"
    BatchReloadByImageResponseStateEnumException BatchReloadByImageResponseStateEnum = "EXCEPTION"
    BatchReloadByImageResponseStateEnumForbidden BatchReloadByImageResponseStateEnum = "FORBIDDEN"
    BatchReloadByImageResponseStateEnumRemaining BatchReloadByImageResponseStateEnum = "REMAINING"
    BatchReloadByImageResponseStateEnumTimeout BatchReloadByImageResponseStateEnum = "TIMEOUT"
)

type BatchReloadByImageResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchReloadByImageResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]BatchReloadByImageResponseBody `json:"body,omitempty"`
}

func (s BatchReloadByImageResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchReloadByImageResponse) GoString() string {
	return s.String()
}

func (s BatchReloadByImageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchReloadByImageResponse) SetRequestId(v string) *BatchReloadByImageResponse {
	s.RequestId = &v
	return s
}

func (s *BatchReloadByImageResponse) SetErrorMessage(v string) *BatchReloadByImageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchReloadByImageResponse) SetErrorCode(v string) *BatchReloadByImageResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchReloadByImageResponse) SetState(v BatchReloadByImageResponseStateEnum) *BatchReloadByImageResponse {
	s.State = &v
	return s
}

func (s *BatchReloadByImageResponse) SetBody(v []BatchReloadByImageResponseBody) *BatchReloadByImageResponse {
	s.Body = &v
	return s
}


type BatchReloadByImageResponseBuilder struct {
	s *BatchReloadByImageResponse
}

func NewBatchReloadByImageResponseBuilder() *BatchReloadByImageResponseBuilder {
	s := &BatchReloadByImageResponse{}
	b := &BatchReloadByImageResponseBuilder{s: s}
	return b
}

func (b *BatchReloadByImageResponseBuilder) RequestId(v string) *BatchReloadByImageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchReloadByImageResponseBuilder) ErrorMessage(v string) *BatchReloadByImageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchReloadByImageResponseBuilder) ErrorCode(v string) *BatchReloadByImageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchReloadByImageResponseBuilder) State(v BatchReloadByImageResponseStateEnum) *BatchReloadByImageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchReloadByImageResponseBuilder) Body(v []BatchReloadByImageResponseBody) *BatchReloadByImageResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchReloadByImageResponseBuilder) Build() *BatchReloadByImageResponse {
	return b.s
}


