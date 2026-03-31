// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetTaskRecordListResponseStateEnum string

// List of State
const (
    GetTaskRecordListResponseStateEnumOk GetTaskRecordListResponseStateEnum = "OK"
    GetTaskRecordListResponseStateEnumError GetTaskRecordListResponseStateEnum = "ERROR"
    GetTaskRecordListResponseStateEnumException GetTaskRecordListResponseStateEnum = "EXCEPTION"
    GetTaskRecordListResponseStateEnumForbidden GetTaskRecordListResponseStateEnum = "FORBIDDEN"
    GetTaskRecordListResponseStateEnumRemaining GetTaskRecordListResponseStateEnum = "REMAINING"
    GetTaskRecordListResponseStateEnumTimeout GetTaskRecordListResponseStateEnum = "TIMEOUT"
)

type GetTaskRecordListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetTaskRecordListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetTaskRecordListResponseBody `json:"body,omitempty"`
}

func (s GetTaskRecordListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListResponse) GoString() string {
	return s.String()
}

func (s GetTaskRecordListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListResponse) SetRequestId(v string) *GetTaskRecordListResponse {
	s.RequestId = &v
	return s
}

func (s *GetTaskRecordListResponse) SetErrorMessage(v string) *GetTaskRecordListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetTaskRecordListResponse) SetErrorCode(v string) *GetTaskRecordListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskRecordListResponse) SetState(v GetTaskRecordListResponseStateEnum) *GetTaskRecordListResponse {
	s.State = &v
	return s
}

func (s *GetTaskRecordListResponse) SetBody(v *GetTaskRecordListResponseBody) *GetTaskRecordListResponse {
	s.Body = v
	return s
}


type GetTaskRecordListResponseBuilder struct {
	s *GetTaskRecordListResponse
}

func NewGetTaskRecordListResponseBuilder() *GetTaskRecordListResponseBuilder {
	s := &GetTaskRecordListResponse{}
	b := &GetTaskRecordListResponseBuilder{s: s}
	return b
}

func (b *GetTaskRecordListResponseBuilder) RequestId(v string) *GetTaskRecordListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetTaskRecordListResponseBuilder) ErrorMessage(v string) *GetTaskRecordListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetTaskRecordListResponseBuilder) ErrorCode(v string) *GetTaskRecordListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetTaskRecordListResponseBuilder) State(v GetTaskRecordListResponseStateEnum) *GetTaskRecordListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetTaskRecordListResponseBuilder) Body(v *GetTaskRecordListResponseBody) *GetTaskRecordListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetTaskRecordListResponseBuilder) Build() *GetTaskRecordListResponse {
	return b.s
}


