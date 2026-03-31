// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchUnlockMachineResponseStateEnum string

// List of State
const (
    BatchUnlockMachineResponseStateEnumOk BatchUnlockMachineResponseStateEnum = "OK"
    BatchUnlockMachineResponseStateEnumError BatchUnlockMachineResponseStateEnum = "ERROR"
    BatchUnlockMachineResponseStateEnumException BatchUnlockMachineResponseStateEnum = "EXCEPTION"
    BatchUnlockMachineResponseStateEnumForbidden BatchUnlockMachineResponseStateEnum = "FORBIDDEN"
    BatchUnlockMachineResponseStateEnumRemaining BatchUnlockMachineResponseStateEnum = "REMAINING"
    BatchUnlockMachineResponseStateEnumTimeout BatchUnlockMachineResponseStateEnum = "TIMEOUT"
)

type BatchUnlockMachineResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchUnlockMachineResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s BatchUnlockMachineResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchUnlockMachineResponse) GoString() string {
	return s.String()
}

func (s BatchUnlockMachineResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUnlockMachineResponse) SetRequestId(v string) *BatchUnlockMachineResponse {
	s.RequestId = &v
	return s
}

func (s *BatchUnlockMachineResponse) SetErrorMessage(v string) *BatchUnlockMachineResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchUnlockMachineResponse) SetErrorCode(v string) *BatchUnlockMachineResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchUnlockMachineResponse) SetState(v BatchUnlockMachineResponseStateEnum) *BatchUnlockMachineResponse {
	s.State = &v
	return s
}

func (s *BatchUnlockMachineResponse) SetBody(v bool) *BatchUnlockMachineResponse {
	s.Body = &v
	return s
}


type BatchUnlockMachineResponseBuilder struct {
	s *BatchUnlockMachineResponse
}

func NewBatchUnlockMachineResponseBuilder() *BatchUnlockMachineResponseBuilder {
	s := &BatchUnlockMachineResponse{}
	b := &BatchUnlockMachineResponseBuilder{s: s}
	return b
}

func (b *BatchUnlockMachineResponseBuilder) RequestId(v string) *BatchUnlockMachineResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchUnlockMachineResponseBuilder) ErrorMessage(v string) *BatchUnlockMachineResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchUnlockMachineResponseBuilder) ErrorCode(v string) *BatchUnlockMachineResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchUnlockMachineResponseBuilder) State(v BatchUnlockMachineResponseStateEnum) *BatchUnlockMachineResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchUnlockMachineResponseBuilder) Body(v bool) *BatchUnlockMachineResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchUnlockMachineResponseBuilder) Build() *BatchUnlockMachineResponse {
	return b.s
}


