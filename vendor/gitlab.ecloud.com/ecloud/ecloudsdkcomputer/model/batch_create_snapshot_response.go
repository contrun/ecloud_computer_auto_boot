// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchCreateSnapshotResponseStateEnum string

// List of State
const (
    BatchCreateSnapshotResponseStateEnumOk BatchCreateSnapshotResponseStateEnum = "OK"
    BatchCreateSnapshotResponseStateEnumError BatchCreateSnapshotResponseStateEnum = "ERROR"
    BatchCreateSnapshotResponseStateEnumException BatchCreateSnapshotResponseStateEnum = "EXCEPTION"
    BatchCreateSnapshotResponseStateEnumForbidden BatchCreateSnapshotResponseStateEnum = "FORBIDDEN"
    BatchCreateSnapshotResponseStateEnumRemaining BatchCreateSnapshotResponseStateEnum = "REMAINING"
    BatchCreateSnapshotResponseStateEnumTimeout BatchCreateSnapshotResponseStateEnum = "TIMEOUT"
)

type BatchCreateSnapshotResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BatchCreateSnapshotResponseStateEnum `json:"state,omitempty"`
    // 返回快照创建成功的桌面id
	Body []string `json:"body,omitempty"`
}

func (s BatchCreateSnapshotResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchCreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s BatchCreateSnapshotResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchCreateSnapshotResponse) SetRequestId(v string) *BatchCreateSnapshotResponse {
	s.RequestId = &v
	return s
}

func (s *BatchCreateSnapshotResponse) SetErrorMessage(v string) *BatchCreateSnapshotResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchCreateSnapshotResponse) SetErrorCode(v string) *BatchCreateSnapshotResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchCreateSnapshotResponse) SetState(v BatchCreateSnapshotResponseStateEnum) *BatchCreateSnapshotResponse {
	s.State = &v
	return s
}

func (s *BatchCreateSnapshotResponse) SetBody(v []string) *BatchCreateSnapshotResponse {
	s.Body = v
	return s
}


type BatchCreateSnapshotResponseBuilder struct {
	s *BatchCreateSnapshotResponse
}

func NewBatchCreateSnapshotResponseBuilder() *BatchCreateSnapshotResponseBuilder {
	s := &BatchCreateSnapshotResponse{}
	b := &BatchCreateSnapshotResponseBuilder{s: s}
	return b
}

func (b *BatchCreateSnapshotResponseBuilder) RequestId(v string) *BatchCreateSnapshotResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchCreateSnapshotResponseBuilder) ErrorMessage(v string) *BatchCreateSnapshotResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchCreateSnapshotResponseBuilder) ErrorCode(v string) *BatchCreateSnapshotResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchCreateSnapshotResponseBuilder) State(v BatchCreateSnapshotResponseStateEnum) *BatchCreateSnapshotResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchCreateSnapshotResponseBuilder) Body(v []string) *BatchCreateSnapshotResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BatchCreateSnapshotResponseBuilder) Build() *BatchCreateSnapshotResponse {
	return b.s
}


