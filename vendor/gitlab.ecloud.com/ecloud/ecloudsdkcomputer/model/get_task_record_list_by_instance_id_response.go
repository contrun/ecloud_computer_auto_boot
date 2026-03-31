// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetTaskRecordListByInstanceIdResponseStateEnum string

// List of State
const (
    GetTaskRecordListByInstanceIdResponseStateEnumOk GetTaskRecordListByInstanceIdResponseStateEnum = "OK"
    GetTaskRecordListByInstanceIdResponseStateEnumError GetTaskRecordListByInstanceIdResponseStateEnum = "ERROR"
    GetTaskRecordListByInstanceIdResponseStateEnumException GetTaskRecordListByInstanceIdResponseStateEnum = "EXCEPTION"
    GetTaskRecordListByInstanceIdResponseStateEnumForbidden GetTaskRecordListByInstanceIdResponseStateEnum = "FORBIDDEN"
    GetTaskRecordListByInstanceIdResponseStateEnumRemaining GetTaskRecordListByInstanceIdResponseStateEnum = "REMAINING"
    GetTaskRecordListByInstanceIdResponseStateEnumTimeout GetTaskRecordListByInstanceIdResponseStateEnum = "TIMEOUT"
)

type GetTaskRecordListByInstanceIdResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetTaskRecordListByInstanceIdResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetTaskRecordListByInstanceIdResponseBody `json:"body,omitempty"`
}

func (s GetTaskRecordListByInstanceIdResponse) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s GetTaskRecordListByInstanceIdResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListByInstanceIdResponse) SetRequestId(v string) *GetTaskRecordListByInstanceIdResponse {
	s.RequestId = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponse) SetErrorMessage(v string) *GetTaskRecordListByInstanceIdResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponse) SetErrorCode(v string) *GetTaskRecordListByInstanceIdResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponse) SetState(v GetTaskRecordListByInstanceIdResponseStateEnum) *GetTaskRecordListByInstanceIdResponse {
	s.State = &v
	return s
}

func (s *GetTaskRecordListByInstanceIdResponse) SetBody(v *GetTaskRecordListByInstanceIdResponseBody) *GetTaskRecordListByInstanceIdResponse {
	s.Body = v
	return s
}


type GetTaskRecordListByInstanceIdResponseBuilder struct {
	s *GetTaskRecordListByInstanceIdResponse
}

func NewGetTaskRecordListByInstanceIdResponseBuilder() *GetTaskRecordListByInstanceIdResponseBuilder {
	s := &GetTaskRecordListByInstanceIdResponse{}
	b := &GetTaskRecordListByInstanceIdResponseBuilder{s: s}
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseBuilder) RequestId(v string) *GetTaskRecordListByInstanceIdResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseBuilder) ErrorMessage(v string) *GetTaskRecordListByInstanceIdResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseBuilder) ErrorCode(v string) *GetTaskRecordListByInstanceIdResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseBuilder) State(v GetTaskRecordListByInstanceIdResponseStateEnum) *GetTaskRecordListByInstanceIdResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseBuilder) Body(v *GetTaskRecordListByInstanceIdResponseBody) *GetTaskRecordListByInstanceIdResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetTaskRecordListByInstanceIdResponseBuilder) Build() *GetTaskRecordListByInstanceIdResponse {
	return b.s
}


