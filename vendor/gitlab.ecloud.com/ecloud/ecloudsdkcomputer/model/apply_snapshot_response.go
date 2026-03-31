// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ApplySnapshotResponseStateEnum string

// List of State
const (
    ApplySnapshotResponseStateEnumOk ApplySnapshotResponseStateEnum = "OK"
    ApplySnapshotResponseStateEnumError ApplySnapshotResponseStateEnum = "ERROR"
    ApplySnapshotResponseStateEnumException ApplySnapshotResponseStateEnum = "EXCEPTION"
    ApplySnapshotResponseStateEnumForbidden ApplySnapshotResponseStateEnum = "FORBIDDEN"
    ApplySnapshotResponseStateEnumRemaining ApplySnapshotResponseStateEnum = "REMAINING"
    ApplySnapshotResponseStateEnumTimeout ApplySnapshotResponseStateEnum = "TIMEOUT"
)

type ApplySnapshotResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ApplySnapshotResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s ApplySnapshotResponse) String() string {
	return utils.Beautify(s)
}

func (s ApplySnapshotResponse) GoString() string {
	return s.String()
}

func (s ApplySnapshotResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ApplySnapshotResponse) SetRequestId(v string) *ApplySnapshotResponse {
	s.RequestId = &v
	return s
}

func (s *ApplySnapshotResponse) SetErrorMessage(v string) *ApplySnapshotResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ApplySnapshotResponse) SetErrorCode(v string) *ApplySnapshotResponse {
	s.ErrorCode = &v
	return s
}

func (s *ApplySnapshotResponse) SetState(v ApplySnapshotResponseStateEnum) *ApplySnapshotResponse {
	s.State = &v
	return s
}

func (s *ApplySnapshotResponse) SetBody(v string) *ApplySnapshotResponse {
	s.Body = &v
	return s
}


type ApplySnapshotResponseBuilder struct {
	s *ApplySnapshotResponse
}

func NewApplySnapshotResponseBuilder() *ApplySnapshotResponseBuilder {
	s := &ApplySnapshotResponse{}
	b := &ApplySnapshotResponseBuilder{s: s}
	return b
}

func (b *ApplySnapshotResponseBuilder) RequestId(v string) *ApplySnapshotResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ApplySnapshotResponseBuilder) ErrorMessage(v string) *ApplySnapshotResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ApplySnapshotResponseBuilder) ErrorCode(v string) *ApplySnapshotResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ApplySnapshotResponseBuilder) State(v ApplySnapshotResponseStateEnum) *ApplySnapshotResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ApplySnapshotResponseBuilder) Body(v string) *ApplySnapshotResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *ApplySnapshotResponseBuilder) Build() *ApplySnapshotResponse {
	return b.s
}


