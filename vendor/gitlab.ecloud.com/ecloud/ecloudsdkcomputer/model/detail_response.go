// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DetailResponseStateEnum string

// List of State
const (
    DetailResponseStateEnumOk DetailResponseStateEnum = "OK"
    DetailResponseStateEnumError DetailResponseStateEnum = "ERROR"
    DetailResponseStateEnumException DetailResponseStateEnum = "EXCEPTION"
    DetailResponseStateEnumForbidden DetailResponseStateEnum = "FORBIDDEN"
    DetailResponseStateEnumRemaining DetailResponseStateEnum = "REMAINING"
    DetailResponseStateEnumTimeout DetailResponseStateEnum = "TIMEOUT"
)

type DetailResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *DetailResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *DetailResponseBody `json:"body,omitempty"`
}

func (s DetailResponse) String() string {
	return utils.Beautify(s)
}

func (s DetailResponse) GoString() string {
	return s.String()
}

func (s DetailResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DetailResponse) SetRequestId(v string) *DetailResponse {
	s.RequestId = &v
	return s
}

func (s *DetailResponse) SetErrorMessage(v string) *DetailResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DetailResponse) SetErrorCode(v string) *DetailResponse {
	s.ErrorCode = &v
	return s
}

func (s *DetailResponse) SetState(v DetailResponseStateEnum) *DetailResponse {
	s.State = &v
	return s
}

func (s *DetailResponse) SetBody(v *DetailResponseBody) *DetailResponse {
	s.Body = v
	return s
}


type DetailResponseBuilder struct {
	s *DetailResponse
}

func NewDetailResponseBuilder() *DetailResponseBuilder {
	s := &DetailResponse{}
	b := &DetailResponseBuilder{s: s}
	return b
}

func (b *DetailResponseBuilder) RequestId(v string) *DetailResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DetailResponseBuilder) ErrorMessage(v string) *DetailResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DetailResponseBuilder) ErrorCode(v string) *DetailResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DetailResponseBuilder) State(v DetailResponseStateEnum) *DetailResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DetailResponseBuilder) Body(v *DetailResponseBody) *DetailResponseBuilder {
	b.s.Body = v
	return b
}

func (b *DetailResponseBuilder) Build() *DetailResponse {
	return b.s
}


