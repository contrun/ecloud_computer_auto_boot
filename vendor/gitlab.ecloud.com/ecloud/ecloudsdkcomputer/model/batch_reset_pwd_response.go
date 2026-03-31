// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchResetPwdResponseStateEnum string

// List of State
const (
    BatchResetPwdResponseStateEnumOk BatchResetPwdResponseStateEnum = "OK"
    BatchResetPwdResponseStateEnumError BatchResetPwdResponseStateEnum = "ERROR"
    BatchResetPwdResponseStateEnumException BatchResetPwdResponseStateEnum = "EXCEPTION"
    BatchResetPwdResponseStateEnumForbidden BatchResetPwdResponseStateEnum = "FORBIDDEN"
    BatchResetPwdResponseStateEnumRemaining BatchResetPwdResponseStateEnum = "REMAINING"
    BatchResetPwdResponseStateEnumTimeout BatchResetPwdResponseStateEnum = "TIMEOUT"
)

type BatchResetPwdResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchResetPwdResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *BatchResetPwdResponseBody `json:"body,omitempty"`
}

func (s BatchResetPwdResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchResetPwdResponse) GoString() string {
	return s.String()
}

func (s BatchResetPwdResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchResetPwdResponse) SetRequestId(v string) *BatchResetPwdResponse {
	s.RequestId = &v
	return s
}

func (s *BatchResetPwdResponse) SetErrorMessage(v string) *BatchResetPwdResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchResetPwdResponse) SetErrorCode(v string) *BatchResetPwdResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchResetPwdResponse) SetState(v BatchResetPwdResponseStateEnum) *BatchResetPwdResponse {
	s.State = &v
	return s
}

func (s *BatchResetPwdResponse) SetBody(v *BatchResetPwdResponseBody) *BatchResetPwdResponse {
	s.Body = v
	return s
}


type BatchResetPwdResponseBuilder struct {
	s *BatchResetPwdResponse
}

func NewBatchResetPwdResponseBuilder() *BatchResetPwdResponseBuilder {
	s := &BatchResetPwdResponse{}
	b := &BatchResetPwdResponseBuilder{s: s}
	return b
}

func (b *BatchResetPwdResponseBuilder) RequestId(v string) *BatchResetPwdResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchResetPwdResponseBuilder) ErrorMessage(v string) *BatchResetPwdResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchResetPwdResponseBuilder) ErrorCode(v string) *BatchResetPwdResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchResetPwdResponseBuilder) State(v BatchResetPwdResponseStateEnum) *BatchResetPwdResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchResetPwdResponseBuilder) Body(v *BatchResetPwdResponseBody) *BatchResetPwdResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BatchResetPwdResponseBuilder) Build() *BatchResetPwdResponse {
	return b.s
}


