// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchUserReActiveNotifyResponseStateEnum string

// List of State
const (
    BatchUserReActiveNotifyResponseStateEnumOk BatchUserReActiveNotifyResponseStateEnum = "OK"
    BatchUserReActiveNotifyResponseStateEnumError BatchUserReActiveNotifyResponseStateEnum = "ERROR"
    BatchUserReActiveNotifyResponseStateEnumException BatchUserReActiveNotifyResponseStateEnum = "EXCEPTION"
    BatchUserReActiveNotifyResponseStateEnumForbidden BatchUserReActiveNotifyResponseStateEnum = "FORBIDDEN"
    BatchUserReActiveNotifyResponseStateEnumRemaining BatchUserReActiveNotifyResponseStateEnum = "REMAINING"
    BatchUserReActiveNotifyResponseStateEnumTimeout BatchUserReActiveNotifyResponseStateEnum = "TIMEOUT"
)

type BatchUserReActiveNotifyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchUserReActiveNotifyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s BatchUserReActiveNotifyResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchUserReActiveNotifyResponse) GoString() string {
	return s.String()
}

func (s BatchUserReActiveNotifyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUserReActiveNotifyResponse) SetRequestId(v string) *BatchUserReActiveNotifyResponse {
	s.RequestId = &v
	return s
}

func (s *BatchUserReActiveNotifyResponse) SetErrorMessage(v string) *BatchUserReActiveNotifyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchUserReActiveNotifyResponse) SetErrorCode(v string) *BatchUserReActiveNotifyResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchUserReActiveNotifyResponse) SetState(v BatchUserReActiveNotifyResponseStateEnum) *BatchUserReActiveNotifyResponse {
	s.State = &v
	return s
}

func (s *BatchUserReActiveNotifyResponse) SetBody(v bool) *BatchUserReActiveNotifyResponse {
	s.Body = &v
	return s
}


type BatchUserReActiveNotifyResponseBuilder struct {
	s *BatchUserReActiveNotifyResponse
}

func NewBatchUserReActiveNotifyResponseBuilder() *BatchUserReActiveNotifyResponseBuilder {
	s := &BatchUserReActiveNotifyResponse{}
	b := &BatchUserReActiveNotifyResponseBuilder{s: s}
	return b
}

func (b *BatchUserReActiveNotifyResponseBuilder) RequestId(v string) *BatchUserReActiveNotifyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchUserReActiveNotifyResponseBuilder) ErrorMessage(v string) *BatchUserReActiveNotifyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchUserReActiveNotifyResponseBuilder) ErrorCode(v string) *BatchUserReActiveNotifyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchUserReActiveNotifyResponseBuilder) State(v BatchUserReActiveNotifyResponseStateEnum) *BatchUserReActiveNotifyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchUserReActiveNotifyResponseBuilder) Body(v bool) *BatchUserReActiveNotifyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchUserReActiveNotifyResponseBuilder) Build() *BatchUserReActiveNotifyResponse {
	return b.s
}


