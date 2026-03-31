// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchUnbindDeviceResponseStateEnum string

// List of State
const (
    BatchUnbindDeviceResponseStateEnumOk BatchUnbindDeviceResponseStateEnum = "OK"
    BatchUnbindDeviceResponseStateEnumError BatchUnbindDeviceResponseStateEnum = "ERROR"
    BatchUnbindDeviceResponseStateEnumException BatchUnbindDeviceResponseStateEnum = "EXCEPTION"
    BatchUnbindDeviceResponseStateEnumForbidden BatchUnbindDeviceResponseStateEnum = "FORBIDDEN"
    BatchUnbindDeviceResponseStateEnumRemaining BatchUnbindDeviceResponseStateEnum = "REMAINING"
    BatchUnbindDeviceResponseStateEnumTimeout BatchUnbindDeviceResponseStateEnum = "TIMEOUT"
)

type BatchUnbindDeviceResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchUnbindDeviceResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s BatchUnbindDeviceResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchUnbindDeviceResponse) GoString() string {
	return s.String()
}

func (s BatchUnbindDeviceResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUnbindDeviceResponse) SetRequestId(v string) *BatchUnbindDeviceResponse {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindDeviceResponse) SetErrorMessage(v string) *BatchUnbindDeviceResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchUnbindDeviceResponse) SetErrorCode(v string) *BatchUnbindDeviceResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchUnbindDeviceResponse) SetState(v BatchUnbindDeviceResponseStateEnum) *BatchUnbindDeviceResponse {
	s.State = &v
	return s
}

func (s *BatchUnbindDeviceResponse) SetBody(v string) *BatchUnbindDeviceResponse {
	s.Body = &v
	return s
}


type BatchUnbindDeviceResponseBuilder struct {
	s *BatchUnbindDeviceResponse
}

func NewBatchUnbindDeviceResponseBuilder() *BatchUnbindDeviceResponseBuilder {
	s := &BatchUnbindDeviceResponse{}
	b := &BatchUnbindDeviceResponseBuilder{s: s}
	return b
}

func (b *BatchUnbindDeviceResponseBuilder) RequestId(v string) *BatchUnbindDeviceResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchUnbindDeviceResponseBuilder) ErrorMessage(v string) *BatchUnbindDeviceResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchUnbindDeviceResponseBuilder) ErrorCode(v string) *BatchUnbindDeviceResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchUnbindDeviceResponseBuilder) State(v BatchUnbindDeviceResponseStateEnum) *BatchUnbindDeviceResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchUnbindDeviceResponseBuilder) Body(v string) *BatchUnbindDeviceResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchUnbindDeviceResponseBuilder) Build() *BatchUnbindDeviceResponse {
	return b.s
}


