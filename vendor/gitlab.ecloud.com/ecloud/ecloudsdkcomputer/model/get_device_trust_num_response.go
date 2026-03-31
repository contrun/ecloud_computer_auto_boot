// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetDeviceTrustNumResponseStateEnum string

// List of State
const (
    GetDeviceTrustNumResponseStateEnumOk GetDeviceTrustNumResponseStateEnum = "OK"
    GetDeviceTrustNumResponseStateEnumError GetDeviceTrustNumResponseStateEnum = "ERROR"
    GetDeviceTrustNumResponseStateEnumException GetDeviceTrustNumResponseStateEnum = "EXCEPTION"
    GetDeviceTrustNumResponseStateEnumForbidden GetDeviceTrustNumResponseStateEnum = "FORBIDDEN"
    GetDeviceTrustNumResponseStateEnumRemaining GetDeviceTrustNumResponseStateEnum = "REMAINING"
    GetDeviceTrustNumResponseStateEnumTimeout GetDeviceTrustNumResponseStateEnum = "TIMEOUT"
)

type GetDeviceTrustNumResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetDeviceTrustNumResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetDeviceTrustNumResponseBody `json:"body,omitempty"`
}

func (s GetDeviceTrustNumResponse) String() string {
	return utils.Beautify(s)
}

func (s GetDeviceTrustNumResponse) GoString() string {
	return s.String()
}

func (s GetDeviceTrustNumResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDeviceTrustNumResponse) SetRequestId(v string) *GetDeviceTrustNumResponse {
	s.RequestId = &v
	return s
}

func (s *GetDeviceTrustNumResponse) SetErrorMessage(v string) *GetDeviceTrustNumResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeviceTrustNumResponse) SetErrorCode(v string) *GetDeviceTrustNumResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetDeviceTrustNumResponse) SetState(v GetDeviceTrustNumResponseStateEnum) *GetDeviceTrustNumResponse {
	s.State = &v
	return s
}

func (s *GetDeviceTrustNumResponse) SetBody(v *GetDeviceTrustNumResponseBody) *GetDeviceTrustNumResponse {
	s.Body = v
	return s
}


type GetDeviceTrustNumResponseBuilder struct {
	s *GetDeviceTrustNumResponse
}

func NewGetDeviceTrustNumResponseBuilder() *GetDeviceTrustNumResponseBuilder {
	s := &GetDeviceTrustNumResponse{}
	b := &GetDeviceTrustNumResponseBuilder{s: s}
	return b
}

func (b *GetDeviceTrustNumResponseBuilder) RequestId(v string) *GetDeviceTrustNumResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetDeviceTrustNumResponseBuilder) ErrorMessage(v string) *GetDeviceTrustNumResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetDeviceTrustNumResponseBuilder) ErrorCode(v string) *GetDeviceTrustNumResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetDeviceTrustNumResponseBuilder) State(v GetDeviceTrustNumResponseStateEnum) *GetDeviceTrustNumResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetDeviceTrustNumResponseBuilder) Body(v *GetDeviceTrustNumResponseBody) *GetDeviceTrustNumResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetDeviceTrustNumResponseBuilder) Build() *GetDeviceTrustNumResponse {
	return b.s
}


