// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetShutdownRevertConfigResponseStateEnum string

// List of State
const (
    GetShutdownRevertConfigResponseStateEnumOk GetShutdownRevertConfigResponseStateEnum = "OK"
    GetShutdownRevertConfigResponseStateEnumError GetShutdownRevertConfigResponseStateEnum = "ERROR"
    GetShutdownRevertConfigResponseStateEnumException GetShutdownRevertConfigResponseStateEnum = "EXCEPTION"
    GetShutdownRevertConfigResponseStateEnumForbidden GetShutdownRevertConfigResponseStateEnum = "FORBIDDEN"
    GetShutdownRevertConfigResponseStateEnumRemaining GetShutdownRevertConfigResponseStateEnum = "REMAINING"
    GetShutdownRevertConfigResponseStateEnumTimeout GetShutdownRevertConfigResponseStateEnum = "TIMEOUT"
)

type GetShutdownRevertConfigResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetShutdownRevertConfigResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetShutdownRevertConfigResponseBody `json:"body,omitempty"`
}

func (s GetShutdownRevertConfigResponse) String() string {
	return utils.Beautify(s)
}

func (s GetShutdownRevertConfigResponse) GoString() string {
	return s.String()
}

func (s GetShutdownRevertConfigResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShutdownRevertConfigResponse) SetRequestId(v string) *GetShutdownRevertConfigResponse {
	s.RequestId = &v
	return s
}

func (s *GetShutdownRevertConfigResponse) SetErrorMessage(v string) *GetShutdownRevertConfigResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetShutdownRevertConfigResponse) SetErrorCode(v string) *GetShutdownRevertConfigResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetShutdownRevertConfigResponse) SetState(v GetShutdownRevertConfigResponseStateEnum) *GetShutdownRevertConfigResponse {
	s.State = &v
	return s
}

func (s *GetShutdownRevertConfigResponse) SetBody(v *GetShutdownRevertConfigResponseBody) *GetShutdownRevertConfigResponse {
	s.Body = v
	return s
}


type GetShutdownRevertConfigResponseBuilder struct {
	s *GetShutdownRevertConfigResponse
}

func NewGetShutdownRevertConfigResponseBuilder() *GetShutdownRevertConfigResponseBuilder {
	s := &GetShutdownRevertConfigResponse{}
	b := &GetShutdownRevertConfigResponseBuilder{s: s}
	return b
}

func (b *GetShutdownRevertConfigResponseBuilder) RequestId(v string) *GetShutdownRevertConfigResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBuilder) ErrorMessage(v string) *GetShutdownRevertConfigResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBuilder) ErrorCode(v string) *GetShutdownRevertConfigResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBuilder) State(v GetShutdownRevertConfigResponseStateEnum) *GetShutdownRevertConfigResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBuilder) Body(v *GetShutdownRevertConfigResponseBody) *GetShutdownRevertConfigResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetShutdownRevertConfigResponseBuilder) Build() *GetShutdownRevertConfigResponse {
	return b.s
}


