// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetClientInfoResponseStateEnum string

// List of State
const (
    GetClientInfoResponseStateEnumOk GetClientInfoResponseStateEnum = "OK"
    GetClientInfoResponseStateEnumError GetClientInfoResponseStateEnum = "ERROR"
    GetClientInfoResponseStateEnumException GetClientInfoResponseStateEnum = "EXCEPTION"
    GetClientInfoResponseStateEnumForbidden GetClientInfoResponseStateEnum = "FORBIDDEN"
    GetClientInfoResponseStateEnumRemaining GetClientInfoResponseStateEnum = "REMAINING"
    GetClientInfoResponseStateEnumTimeout GetClientInfoResponseStateEnum = "TIMEOUT"
)

type GetClientInfoResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetClientInfoResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetClientInfoResponseBody `json:"body,omitempty"`
}

func (s GetClientInfoResponse) String() string {
	return utils.Beautify(s)
}

func (s GetClientInfoResponse) GoString() string {
	return s.String()
}

func (s GetClientInfoResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetClientInfoResponse) SetRequestId(v string) *GetClientInfoResponse {
	s.RequestId = &v
	return s
}

func (s *GetClientInfoResponse) SetErrorMessage(v string) *GetClientInfoResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetClientInfoResponse) SetErrorCode(v string) *GetClientInfoResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetClientInfoResponse) SetState(v GetClientInfoResponseStateEnum) *GetClientInfoResponse {
	s.State = &v
	return s
}

func (s *GetClientInfoResponse) SetBody(v *GetClientInfoResponseBody) *GetClientInfoResponse {
	s.Body = v
	return s
}


type GetClientInfoResponseBuilder struct {
	s *GetClientInfoResponse
}

func NewGetClientInfoResponseBuilder() *GetClientInfoResponseBuilder {
	s := &GetClientInfoResponse{}
	b := &GetClientInfoResponseBuilder{s: s}
	return b
}

func (b *GetClientInfoResponseBuilder) RequestId(v string) *GetClientInfoResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetClientInfoResponseBuilder) ErrorMessage(v string) *GetClientInfoResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetClientInfoResponseBuilder) ErrorCode(v string) *GetClientInfoResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetClientInfoResponseBuilder) State(v GetClientInfoResponseStateEnum) *GetClientInfoResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetClientInfoResponseBuilder) Body(v *GetClientInfoResponseBody) *GetClientInfoResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetClientInfoResponseBuilder) Build() *GetClientInfoResponse {
	return b.s
}


