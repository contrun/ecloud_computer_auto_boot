// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetRestartRecordListResponseStateEnum string

// List of State
const (
    GetRestartRecordListResponseStateEnumOk GetRestartRecordListResponseStateEnum = "OK"
    GetRestartRecordListResponseStateEnumError GetRestartRecordListResponseStateEnum = "ERROR"
    GetRestartRecordListResponseStateEnumException GetRestartRecordListResponseStateEnum = "EXCEPTION"
    GetRestartRecordListResponseStateEnumForbidden GetRestartRecordListResponseStateEnum = "FORBIDDEN"
    GetRestartRecordListResponseStateEnumRemaining GetRestartRecordListResponseStateEnum = "REMAINING"
    GetRestartRecordListResponseStateEnumTimeout GetRestartRecordListResponseStateEnum = "TIMEOUT"
)

type GetRestartRecordListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetRestartRecordListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetRestartRecordListResponseBody `json:"body,omitempty"`
}

func (s GetRestartRecordListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetRestartRecordListResponse) GoString() string {
	return s.String()
}

func (s GetRestartRecordListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetRestartRecordListResponse) SetRequestId(v string) *GetRestartRecordListResponse {
	s.RequestId = &v
	return s
}

func (s *GetRestartRecordListResponse) SetErrorMessage(v string) *GetRestartRecordListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetRestartRecordListResponse) SetErrorCode(v string) *GetRestartRecordListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetRestartRecordListResponse) SetState(v GetRestartRecordListResponseStateEnum) *GetRestartRecordListResponse {
	s.State = &v
	return s
}

func (s *GetRestartRecordListResponse) SetBody(v *GetRestartRecordListResponseBody) *GetRestartRecordListResponse {
	s.Body = v
	return s
}


type GetRestartRecordListResponseBuilder struct {
	s *GetRestartRecordListResponse
}

func NewGetRestartRecordListResponseBuilder() *GetRestartRecordListResponseBuilder {
	s := &GetRestartRecordListResponse{}
	b := &GetRestartRecordListResponseBuilder{s: s}
	return b
}

func (b *GetRestartRecordListResponseBuilder) RequestId(v string) *GetRestartRecordListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetRestartRecordListResponseBuilder) ErrorMessage(v string) *GetRestartRecordListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetRestartRecordListResponseBuilder) ErrorCode(v string) *GetRestartRecordListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetRestartRecordListResponseBuilder) State(v GetRestartRecordListResponseStateEnum) *GetRestartRecordListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetRestartRecordListResponseBuilder) Body(v *GetRestartRecordListResponseBody) *GetRestartRecordListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetRestartRecordListResponseBuilder) Build() *GetRestartRecordListResponse {
	return b.s
}


