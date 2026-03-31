// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchMoveUserResponseStateEnum string

// List of State
const (
    BatchMoveUserResponseStateEnumOk BatchMoveUserResponseStateEnum = "OK"
    BatchMoveUserResponseStateEnumError BatchMoveUserResponseStateEnum = "ERROR"
    BatchMoveUserResponseStateEnumException BatchMoveUserResponseStateEnum = "EXCEPTION"
    BatchMoveUserResponseStateEnumForbidden BatchMoveUserResponseStateEnum = "FORBIDDEN"
    BatchMoveUserResponseStateEnumRemaining BatchMoveUserResponseStateEnum = "REMAINING"
    BatchMoveUserResponseStateEnumTimeout BatchMoveUserResponseStateEnum = "TIMEOUT"
)

type BatchMoveUserResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchMoveUserResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s BatchMoveUserResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchMoveUserResponse) GoString() string {
	return s.String()
}

func (s BatchMoveUserResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchMoveUserResponse) SetRequestId(v string) *BatchMoveUserResponse {
	s.RequestId = &v
	return s
}

func (s *BatchMoveUserResponse) SetErrorMessage(v string) *BatchMoveUserResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchMoveUserResponse) SetErrorCode(v string) *BatchMoveUserResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchMoveUserResponse) SetState(v BatchMoveUserResponseStateEnum) *BatchMoveUserResponse {
	s.State = &v
	return s
}

func (s *BatchMoveUserResponse) SetBody(v string) *BatchMoveUserResponse {
	s.Body = &v
	return s
}


type BatchMoveUserResponseBuilder struct {
	s *BatchMoveUserResponse
}

func NewBatchMoveUserResponseBuilder() *BatchMoveUserResponseBuilder {
	s := &BatchMoveUserResponse{}
	b := &BatchMoveUserResponseBuilder{s: s}
	return b
}

func (b *BatchMoveUserResponseBuilder) RequestId(v string) *BatchMoveUserResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchMoveUserResponseBuilder) ErrorMessage(v string) *BatchMoveUserResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchMoveUserResponseBuilder) ErrorCode(v string) *BatchMoveUserResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchMoveUserResponseBuilder) State(v BatchMoveUserResponseStateEnum) *BatchMoveUserResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchMoveUserResponseBuilder) Body(v string) *BatchMoveUserResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchMoveUserResponseBuilder) Build() *BatchMoveUserResponse {
	return b.s
}


