// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddOrgResponseStateEnum string

// List of State
const (
    AddOrgResponseStateEnumOk AddOrgResponseStateEnum = "OK"
    AddOrgResponseStateEnumError AddOrgResponseStateEnum = "ERROR"
    AddOrgResponseStateEnumException AddOrgResponseStateEnum = "EXCEPTION"
    AddOrgResponseStateEnumForbidden AddOrgResponseStateEnum = "FORBIDDEN"
    AddOrgResponseStateEnumRemaining AddOrgResponseStateEnum = "REMAINING"
    AddOrgResponseStateEnumTimeout AddOrgResponseStateEnum = "TIMEOUT"
)

type AddOrgResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *AddOrgResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *AddOrgResponseBody `json:"body,omitempty"`
}

func (s AddOrgResponse) String() string {
	return utils.Beautify(s)
}

func (s AddOrgResponse) GoString() string {
	return s.String()
}

func (s AddOrgResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddOrgResponse) SetRequestId(v string) *AddOrgResponse {
	s.RequestId = &v
	return s
}

func (s *AddOrgResponse) SetErrorMessage(v string) *AddOrgResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddOrgResponse) SetErrorCode(v string) *AddOrgResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddOrgResponse) SetState(v AddOrgResponseStateEnum) *AddOrgResponse {
	s.State = &v
	return s
}

func (s *AddOrgResponse) SetBody(v *AddOrgResponseBody) *AddOrgResponse {
	s.Body = v
	return s
}


type AddOrgResponseBuilder struct {
	s *AddOrgResponse
}

func NewAddOrgResponseBuilder() *AddOrgResponseBuilder {
	s := &AddOrgResponse{}
	b := &AddOrgResponseBuilder{s: s}
	return b
}

func (b *AddOrgResponseBuilder) RequestId(v string) *AddOrgResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddOrgResponseBuilder) ErrorMessage(v string) *AddOrgResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddOrgResponseBuilder) ErrorCode(v string) *AddOrgResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddOrgResponseBuilder) State(v AddOrgResponseStateEnum) *AddOrgResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddOrgResponseBuilder) Body(v *AddOrgResponseBody) *AddOrgResponseBuilder {
	b.s.Body = v
	return b
}

func (b *AddOrgResponseBuilder) Build() *AddOrgResponse {
	return b.s
}


