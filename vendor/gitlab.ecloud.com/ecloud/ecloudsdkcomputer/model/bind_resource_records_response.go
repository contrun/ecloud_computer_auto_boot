// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BindResourceRecordsResponseStateEnum string

// List of State
const (
    BindResourceRecordsResponseStateEnumOk BindResourceRecordsResponseStateEnum = "OK"
    BindResourceRecordsResponseStateEnumError BindResourceRecordsResponseStateEnum = "ERROR"
    BindResourceRecordsResponseStateEnumException BindResourceRecordsResponseStateEnum = "EXCEPTION"
    BindResourceRecordsResponseStateEnumForbidden BindResourceRecordsResponseStateEnum = "FORBIDDEN"
    BindResourceRecordsResponseStateEnumRemaining BindResourceRecordsResponseStateEnum = "REMAINING"
    BindResourceRecordsResponseStateEnumTimeout BindResourceRecordsResponseStateEnum = "TIMEOUT"
)

type BindResourceRecordsResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BindResourceRecordsResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *BindResourceRecordsResponseBody `json:"body,omitempty"`
}

func (s BindResourceRecordsResponse) String() string {
	return utils.Beautify(s)
}

func (s BindResourceRecordsResponse) GoString() string {
	return s.String()
}

func (s BindResourceRecordsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindResourceRecordsResponse) SetRequestId(v string) *BindResourceRecordsResponse {
	s.RequestId = &v
	return s
}

func (s *BindResourceRecordsResponse) SetErrorMessage(v string) *BindResourceRecordsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BindResourceRecordsResponse) SetErrorCode(v string) *BindResourceRecordsResponse {
	s.ErrorCode = &v
	return s
}

func (s *BindResourceRecordsResponse) SetState(v BindResourceRecordsResponseStateEnum) *BindResourceRecordsResponse {
	s.State = &v
	return s
}

func (s *BindResourceRecordsResponse) SetBody(v *BindResourceRecordsResponseBody) *BindResourceRecordsResponse {
	s.Body = v
	return s
}


type BindResourceRecordsResponseBuilder struct {
	s *BindResourceRecordsResponse
}

func NewBindResourceRecordsResponseBuilder() *BindResourceRecordsResponseBuilder {
	s := &BindResourceRecordsResponse{}
	b := &BindResourceRecordsResponseBuilder{s: s}
	return b
}

func (b *BindResourceRecordsResponseBuilder) RequestId(v string) *BindResourceRecordsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BindResourceRecordsResponseBuilder) ErrorMessage(v string) *BindResourceRecordsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BindResourceRecordsResponseBuilder) ErrorCode(v string) *BindResourceRecordsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BindResourceRecordsResponseBuilder) State(v BindResourceRecordsResponseStateEnum) *BindResourceRecordsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BindResourceRecordsResponseBuilder) Body(v *BindResourceRecordsResponseBody) *BindResourceRecordsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BindResourceRecordsResponseBuilder) Build() *BindResourceRecordsResponse {
	return b.s
}


