// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetAvailableRecordListResponseStateEnum string

// List of State
const (
    GetAvailableRecordListResponseStateEnumOk GetAvailableRecordListResponseStateEnum = "OK"
    GetAvailableRecordListResponseStateEnumError GetAvailableRecordListResponseStateEnum = "ERROR"
    GetAvailableRecordListResponseStateEnumException GetAvailableRecordListResponseStateEnum = "EXCEPTION"
    GetAvailableRecordListResponseStateEnumForbidden GetAvailableRecordListResponseStateEnum = "FORBIDDEN"
    GetAvailableRecordListResponseStateEnumRemaining GetAvailableRecordListResponseStateEnum = "REMAINING"
    GetAvailableRecordListResponseStateEnumTimeout GetAvailableRecordListResponseStateEnum = "TIMEOUT"
)

type GetAvailableRecordListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetAvailableRecordListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetAvailableRecordListResponseBody `json:"body,omitempty"`
}

func (s GetAvailableRecordListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetAvailableRecordListResponse) GoString() string {
	return s.String()
}

func (s GetAvailableRecordListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAvailableRecordListResponse) SetRequestId(v string) *GetAvailableRecordListResponse {
	s.RequestId = &v
	return s
}

func (s *GetAvailableRecordListResponse) SetErrorMessage(v string) *GetAvailableRecordListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetAvailableRecordListResponse) SetErrorCode(v string) *GetAvailableRecordListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetAvailableRecordListResponse) SetState(v GetAvailableRecordListResponseStateEnum) *GetAvailableRecordListResponse {
	s.State = &v
	return s
}

func (s *GetAvailableRecordListResponse) SetBody(v *GetAvailableRecordListResponseBody) *GetAvailableRecordListResponse {
	s.Body = v
	return s
}


type GetAvailableRecordListResponseBuilder struct {
	s *GetAvailableRecordListResponse
}

func NewGetAvailableRecordListResponseBuilder() *GetAvailableRecordListResponseBuilder {
	s := &GetAvailableRecordListResponse{}
	b := &GetAvailableRecordListResponseBuilder{s: s}
	return b
}

func (b *GetAvailableRecordListResponseBuilder) RequestId(v string) *GetAvailableRecordListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetAvailableRecordListResponseBuilder) ErrorMessage(v string) *GetAvailableRecordListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetAvailableRecordListResponseBuilder) ErrorCode(v string) *GetAvailableRecordListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetAvailableRecordListResponseBuilder) State(v GetAvailableRecordListResponseStateEnum) *GetAvailableRecordListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetAvailableRecordListResponseBuilder) Body(v *GetAvailableRecordListResponseBody) *GetAvailableRecordListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetAvailableRecordListResponseBuilder) Build() *GetAvailableRecordListResponse {
	return b.s
}


