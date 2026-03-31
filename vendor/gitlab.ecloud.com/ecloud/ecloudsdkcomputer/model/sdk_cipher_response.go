// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SdkCipherResponseStateEnum string

// List of State
const (
    SdkCipherResponseStateEnumOk SdkCipherResponseStateEnum = "OK"
    SdkCipherResponseStateEnumError SdkCipherResponseStateEnum = "ERROR"
    SdkCipherResponseStateEnumException SdkCipherResponseStateEnum = "EXCEPTION"
    SdkCipherResponseStateEnumForbidden SdkCipherResponseStateEnum = "FORBIDDEN"
    SdkCipherResponseStateEnumRemaining SdkCipherResponseStateEnum = "REMAINING"
    SdkCipherResponseStateEnumTimeout SdkCipherResponseStateEnum = "TIMEOUT"
)

type SdkCipherResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *SdkCipherResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *SdkCipherResponseBody `json:"body,omitempty"`
}

func (s SdkCipherResponse) String() string {
	return utils.Beautify(s)
}

func (s SdkCipherResponse) GoString() string {
	return s.String()
}

func (s SdkCipherResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SdkCipherResponse) SetRequestId(v string) *SdkCipherResponse {
	s.RequestId = &v
	return s
}

func (s *SdkCipherResponse) SetErrorMessage(v string) *SdkCipherResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SdkCipherResponse) SetErrorCode(v string) *SdkCipherResponse {
	s.ErrorCode = &v
	return s
}

func (s *SdkCipherResponse) SetState(v SdkCipherResponseStateEnum) *SdkCipherResponse {
	s.State = &v
	return s
}

func (s *SdkCipherResponse) SetBody(v *SdkCipherResponseBody) *SdkCipherResponse {
	s.Body = v
	return s
}


type SdkCipherResponseBuilder struct {
	s *SdkCipherResponse
}

func NewSdkCipherResponseBuilder() *SdkCipherResponseBuilder {
	s := &SdkCipherResponse{}
	b := &SdkCipherResponseBuilder{s: s}
	return b
}

func (b *SdkCipherResponseBuilder) RequestId(v string) *SdkCipherResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SdkCipherResponseBuilder) ErrorMessage(v string) *SdkCipherResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SdkCipherResponseBuilder) ErrorCode(v string) *SdkCipherResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SdkCipherResponseBuilder) State(v SdkCipherResponseStateEnum) *SdkCipherResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SdkCipherResponseBuilder) Body(v *SdkCipherResponseBody) *SdkCipherResponseBuilder {
	b.s.Body = v
	return b
}

func (b *SdkCipherResponseBuilder) Build() *SdkCipherResponse {
	return b.s
}


