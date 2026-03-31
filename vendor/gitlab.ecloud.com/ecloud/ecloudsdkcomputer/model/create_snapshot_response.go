// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateSnapshotResponseStateEnum string

// List of State
const (
    CreateSnapshotResponseStateEnumOk CreateSnapshotResponseStateEnum = "OK"
    CreateSnapshotResponseStateEnumError CreateSnapshotResponseStateEnum = "ERROR"
    CreateSnapshotResponseStateEnumException CreateSnapshotResponseStateEnum = "EXCEPTION"
    CreateSnapshotResponseStateEnumForbidden CreateSnapshotResponseStateEnum = "FORBIDDEN"
    CreateSnapshotResponseStateEnumRemaining CreateSnapshotResponseStateEnum = "REMAINING"
    CreateSnapshotResponseStateEnumTimeout CreateSnapshotResponseStateEnum = "TIMEOUT"
)

type CreateSnapshotResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *CreateSnapshotResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *CreateSnapshotResponseBody `json:"body,omitempty"`
}

func (s CreateSnapshotResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s CreateSnapshotResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSnapshotResponse) SetRequestId(v string) *CreateSnapshotResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponse) SetErrorMessage(v string) *CreateSnapshotResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSnapshotResponse) SetErrorCode(v string) *CreateSnapshotResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateSnapshotResponse) SetState(v CreateSnapshotResponseStateEnum) *CreateSnapshotResponse {
	s.State = &v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}


type CreateSnapshotResponseBuilder struct {
	s *CreateSnapshotResponse
}

func NewCreateSnapshotResponseBuilder() *CreateSnapshotResponseBuilder {
	s := &CreateSnapshotResponse{}
	b := &CreateSnapshotResponseBuilder{s: s}
	return b
}

func (b *CreateSnapshotResponseBuilder) RequestId(v string) *CreateSnapshotResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateSnapshotResponseBuilder) ErrorMessage(v string) *CreateSnapshotResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateSnapshotResponseBuilder) ErrorCode(v string) *CreateSnapshotResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateSnapshotResponseBuilder) State(v CreateSnapshotResponseStateEnum) *CreateSnapshotResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateSnapshotResponseBuilder) Body(v *CreateSnapshotResponseBody) *CreateSnapshotResponseBuilder {
	b.s.Body = v
	return b
}

func (b *CreateSnapshotResponseBuilder) Build() *CreateSnapshotResponse {
	return b.s
}


