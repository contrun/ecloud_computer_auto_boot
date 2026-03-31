// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetApplySnapshotRecordResponseStateEnum string

// List of State
const (
    GetApplySnapshotRecordResponseStateEnumOk GetApplySnapshotRecordResponseStateEnum = "OK"
    GetApplySnapshotRecordResponseStateEnumError GetApplySnapshotRecordResponseStateEnum = "ERROR"
    GetApplySnapshotRecordResponseStateEnumException GetApplySnapshotRecordResponseStateEnum = "EXCEPTION"
    GetApplySnapshotRecordResponseStateEnumForbidden GetApplySnapshotRecordResponseStateEnum = "FORBIDDEN"
    GetApplySnapshotRecordResponseStateEnumRemaining GetApplySnapshotRecordResponseStateEnum = "REMAINING"
    GetApplySnapshotRecordResponseStateEnumTimeout GetApplySnapshotRecordResponseStateEnum = "TIMEOUT"
)

type GetApplySnapshotRecordResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetApplySnapshotRecordResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetApplySnapshotRecordResponseBody `json:"body,omitempty"`
}

func (s GetApplySnapshotRecordResponse) String() string {
	return utils.Beautify(s)
}

func (s GetApplySnapshotRecordResponse) GoString() string {
	return s.String()
}

func (s GetApplySnapshotRecordResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplySnapshotRecordResponse) SetRequestId(v string) *GetApplySnapshotRecordResponse {
	s.RequestId = &v
	return s
}

func (s *GetApplySnapshotRecordResponse) SetErrorMessage(v string) *GetApplySnapshotRecordResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetApplySnapshotRecordResponse) SetErrorCode(v string) *GetApplySnapshotRecordResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetApplySnapshotRecordResponse) SetState(v GetApplySnapshotRecordResponseStateEnum) *GetApplySnapshotRecordResponse {
	s.State = &v
	return s
}

func (s *GetApplySnapshotRecordResponse) SetBody(v *GetApplySnapshotRecordResponseBody) *GetApplySnapshotRecordResponse {
	s.Body = v
	return s
}


type GetApplySnapshotRecordResponseBuilder struct {
	s *GetApplySnapshotRecordResponse
}

func NewGetApplySnapshotRecordResponseBuilder() *GetApplySnapshotRecordResponseBuilder {
	s := &GetApplySnapshotRecordResponse{}
	b := &GetApplySnapshotRecordResponseBuilder{s: s}
	return b
}

func (b *GetApplySnapshotRecordResponseBuilder) RequestId(v string) *GetApplySnapshotRecordResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetApplySnapshotRecordResponseBuilder) ErrorMessage(v string) *GetApplySnapshotRecordResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetApplySnapshotRecordResponseBuilder) ErrorCode(v string) *GetApplySnapshotRecordResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetApplySnapshotRecordResponseBuilder) State(v GetApplySnapshotRecordResponseStateEnum) *GetApplySnapshotRecordResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetApplySnapshotRecordResponseBuilder) Body(v *GetApplySnapshotRecordResponseBody) *GetApplySnapshotRecordResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetApplySnapshotRecordResponseBuilder) Build() *GetApplySnapshotRecordResponse {
	return b.s
}


