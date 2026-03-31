// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchApplySnapshotResponseStateEnum string

// List of State
const (
    BatchApplySnapshotResponseStateEnumOk BatchApplySnapshotResponseStateEnum = "OK"
    BatchApplySnapshotResponseStateEnumError BatchApplySnapshotResponseStateEnum = "ERROR"
    BatchApplySnapshotResponseStateEnumException BatchApplySnapshotResponseStateEnum = "EXCEPTION"
    BatchApplySnapshotResponseStateEnumForbidden BatchApplySnapshotResponseStateEnum = "FORBIDDEN"
    BatchApplySnapshotResponseStateEnumRemaining BatchApplySnapshotResponseStateEnum = "REMAINING"
    BatchApplySnapshotResponseStateEnumTimeout BatchApplySnapshotResponseStateEnum = "TIMEOUT"
)

type BatchApplySnapshotResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchApplySnapshotResponseStateEnum `json:"state,omitempty"`
    // 返回还原快照成功的桌面id
	Body []string `json:"body,omitempty"`
}

func (s BatchApplySnapshotResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchApplySnapshotResponse) GoString() string {
	return s.String()
}

func (s BatchApplySnapshotResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchApplySnapshotResponse) SetRequestId(v string) *BatchApplySnapshotResponse {
	s.RequestId = &v
	return s
}

func (s *BatchApplySnapshotResponse) SetErrorMessage(v string) *BatchApplySnapshotResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchApplySnapshotResponse) SetErrorCode(v string) *BatchApplySnapshotResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchApplySnapshotResponse) SetState(v BatchApplySnapshotResponseStateEnum) *BatchApplySnapshotResponse {
	s.State = &v
	return s
}

func (s *BatchApplySnapshotResponse) SetBody(v []string) *BatchApplySnapshotResponse {
	s.Body = v
	return s
}


type BatchApplySnapshotResponseBuilder struct {
	s *BatchApplySnapshotResponse
}

func NewBatchApplySnapshotResponseBuilder() *BatchApplySnapshotResponseBuilder {
	s := &BatchApplySnapshotResponse{}
	b := &BatchApplySnapshotResponseBuilder{s: s}
	return b
}

func (b *BatchApplySnapshotResponseBuilder) RequestId(v string) *BatchApplySnapshotResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchApplySnapshotResponseBuilder) ErrorMessage(v string) *BatchApplySnapshotResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchApplySnapshotResponseBuilder) ErrorCode(v string) *BatchApplySnapshotResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchApplySnapshotResponseBuilder) State(v BatchApplySnapshotResponseStateEnum) *BatchApplySnapshotResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchApplySnapshotResponseBuilder) Body(v []string) *BatchApplySnapshotResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BatchApplySnapshotResponseBuilder) Build() *BatchApplySnapshotResponse {
	return b.s
}


