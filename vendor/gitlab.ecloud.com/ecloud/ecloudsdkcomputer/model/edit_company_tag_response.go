// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type EditCompanyTagResponseStateEnum string

// List of State
const (
    EditCompanyTagResponseStateEnumOk EditCompanyTagResponseStateEnum = "OK"
    EditCompanyTagResponseStateEnumError EditCompanyTagResponseStateEnum = "ERROR"
    EditCompanyTagResponseStateEnumException EditCompanyTagResponseStateEnum = "EXCEPTION"
    EditCompanyTagResponseStateEnumForbidden EditCompanyTagResponseStateEnum = "FORBIDDEN"
    EditCompanyTagResponseStateEnumRemaining EditCompanyTagResponseStateEnum = "REMAINING"
    EditCompanyTagResponseStateEnumTimeout EditCompanyTagResponseStateEnum = "TIMEOUT"
)

type EditCompanyTagResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *EditCompanyTagResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s EditCompanyTagResponse) String() string {
	return utils.Beautify(s)
}

func (s EditCompanyTagResponse) GoString() string {
	return s.String()
}

func (s EditCompanyTagResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditCompanyTagResponse) SetRequestId(v string) *EditCompanyTagResponse {
	s.RequestId = &v
	return s
}

func (s *EditCompanyTagResponse) SetErrorMessage(v string) *EditCompanyTagResponse {
	s.ErrorMessage = &v
	return s
}

func (s *EditCompanyTagResponse) SetErrorCode(v string) *EditCompanyTagResponse {
	s.ErrorCode = &v
	return s
}

func (s *EditCompanyTagResponse) SetState(v EditCompanyTagResponseStateEnum) *EditCompanyTagResponse {
	s.State = &v
	return s
}

func (s *EditCompanyTagResponse) SetBody(v bool) *EditCompanyTagResponse {
	s.Body = &v
	return s
}


type EditCompanyTagResponseBuilder struct {
	s *EditCompanyTagResponse
}

func NewEditCompanyTagResponseBuilder() *EditCompanyTagResponseBuilder {
	s := &EditCompanyTagResponse{}
	b := &EditCompanyTagResponseBuilder{s: s}
	return b
}

func (b *EditCompanyTagResponseBuilder) RequestId(v string) *EditCompanyTagResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *EditCompanyTagResponseBuilder) ErrorMessage(v string) *EditCompanyTagResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *EditCompanyTagResponseBuilder) ErrorCode(v string) *EditCompanyTagResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *EditCompanyTagResponseBuilder) State(v EditCompanyTagResponseStateEnum) *EditCompanyTagResponseBuilder {
	b.s.State = &v
	return b
}

func (b *EditCompanyTagResponseBuilder) Body(v bool) *EditCompanyTagResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *EditCompanyTagResponseBuilder) Build() *EditCompanyTagResponse {
	return b.s
}


