// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetShutdownRecordListResponseStateEnum string

// List of State
const (
    GetShutdownRecordListResponseStateEnumOk GetShutdownRecordListResponseStateEnum = "OK"
    GetShutdownRecordListResponseStateEnumError GetShutdownRecordListResponseStateEnum = "ERROR"
    GetShutdownRecordListResponseStateEnumException GetShutdownRecordListResponseStateEnum = "EXCEPTION"
    GetShutdownRecordListResponseStateEnumForbidden GetShutdownRecordListResponseStateEnum = "FORBIDDEN"
    GetShutdownRecordListResponseStateEnumRemaining GetShutdownRecordListResponseStateEnum = "REMAINING"
    GetShutdownRecordListResponseStateEnumTimeout GetShutdownRecordListResponseStateEnum = "TIMEOUT"
)

type GetShutdownRecordListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetShutdownRecordListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetShutdownRecordListResponseBody `json:"body,omitempty"`
}

func (s GetShutdownRecordListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetShutdownRecordListResponse) GoString() string {
	return s.String()
}

func (s GetShutdownRecordListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShutdownRecordListResponse) SetRequestId(v string) *GetShutdownRecordListResponse {
	s.RequestId = &v
	return s
}

func (s *GetShutdownRecordListResponse) SetErrorMessage(v string) *GetShutdownRecordListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetShutdownRecordListResponse) SetErrorCode(v string) *GetShutdownRecordListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetShutdownRecordListResponse) SetState(v GetShutdownRecordListResponseStateEnum) *GetShutdownRecordListResponse {
	s.State = &v
	return s
}

func (s *GetShutdownRecordListResponse) SetBody(v *GetShutdownRecordListResponseBody) *GetShutdownRecordListResponse {
	s.Body = v
	return s
}


type GetShutdownRecordListResponseBuilder struct {
	s *GetShutdownRecordListResponse
}

func NewGetShutdownRecordListResponseBuilder() *GetShutdownRecordListResponseBuilder {
	s := &GetShutdownRecordListResponse{}
	b := &GetShutdownRecordListResponseBuilder{s: s}
	return b
}

func (b *GetShutdownRecordListResponseBuilder) RequestId(v string) *GetShutdownRecordListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetShutdownRecordListResponseBuilder) ErrorMessage(v string) *GetShutdownRecordListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetShutdownRecordListResponseBuilder) ErrorCode(v string) *GetShutdownRecordListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetShutdownRecordListResponseBuilder) State(v GetShutdownRecordListResponseStateEnum) *GetShutdownRecordListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetShutdownRecordListResponseBuilder) Body(v *GetShutdownRecordListResponseBody) *GetShutdownRecordListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetShutdownRecordListResponseBuilder) Build() *GetShutdownRecordListResponse {
	return b.s
}


