// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteSnapshotResponseStateEnum string

// List of State
const (
    DeleteSnapshotResponseStateEnumOk DeleteSnapshotResponseStateEnum = "OK"
    DeleteSnapshotResponseStateEnumError DeleteSnapshotResponseStateEnum = "ERROR"
    DeleteSnapshotResponseStateEnumException DeleteSnapshotResponseStateEnum = "EXCEPTION"
    DeleteSnapshotResponseStateEnumForbidden DeleteSnapshotResponseStateEnum = "FORBIDDEN"
    DeleteSnapshotResponseStateEnumRemaining DeleteSnapshotResponseStateEnum = "REMAINING"
    DeleteSnapshotResponseStateEnumTimeout DeleteSnapshotResponseStateEnum = "TIMEOUT"
)

type DeleteSnapshotResponse struct {

    //  每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *DeleteSnapshotResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s DeleteSnapshotResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s DeleteSnapshotResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSnapshotResponse) SetRequestId(v string) *DeleteSnapshotResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotResponse) SetErrorMessage(v string) *DeleteSnapshotResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSnapshotResponse) SetErrorCode(v string) *DeleteSnapshotResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSnapshotResponse) SetState(v DeleteSnapshotResponseStateEnum) *DeleteSnapshotResponse {
	s.State = &v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v string) *DeleteSnapshotResponse {
	s.Body = &v
	return s
}


type DeleteSnapshotResponseBuilder struct {
	s *DeleteSnapshotResponse
}

func NewDeleteSnapshotResponseBuilder() *DeleteSnapshotResponseBuilder {
	s := &DeleteSnapshotResponse{}
	b := &DeleteSnapshotResponseBuilder{s: s}
	return b
}

func (b *DeleteSnapshotResponseBuilder) RequestId(v string) *DeleteSnapshotResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteSnapshotResponseBuilder) ErrorMessage(v string) *DeleteSnapshotResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteSnapshotResponseBuilder) ErrorCode(v string) *DeleteSnapshotResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteSnapshotResponseBuilder) State(v DeleteSnapshotResponseStateEnum) *DeleteSnapshotResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteSnapshotResponseBuilder) Body(v string) *DeleteSnapshotResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteSnapshotResponseBuilder) Build() *DeleteSnapshotResponse {
	return b.s
}


