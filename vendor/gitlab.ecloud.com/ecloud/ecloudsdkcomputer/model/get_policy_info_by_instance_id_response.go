// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPolicyInfoByInstanceIdResponseStateEnum string

// List of State
const (
    GetPolicyInfoByInstanceIdResponseStateEnumOk GetPolicyInfoByInstanceIdResponseStateEnum = "OK"
    GetPolicyInfoByInstanceIdResponseStateEnumError GetPolicyInfoByInstanceIdResponseStateEnum = "ERROR"
    GetPolicyInfoByInstanceIdResponseStateEnumException GetPolicyInfoByInstanceIdResponseStateEnum = "EXCEPTION"
    GetPolicyInfoByInstanceIdResponseStateEnumForbidden GetPolicyInfoByInstanceIdResponseStateEnum = "FORBIDDEN"
    GetPolicyInfoByInstanceIdResponseStateEnumRemaining GetPolicyInfoByInstanceIdResponseStateEnum = "REMAINING"
    GetPolicyInfoByInstanceIdResponseStateEnumTimeout GetPolicyInfoByInstanceIdResponseStateEnum = "TIMEOUT"
)

type GetPolicyInfoByInstanceIdResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetPolicyInfoByInstanceIdResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetPolicyInfoByInstanceIdResponseBody `json:"body,omitempty"`
}

func (s GetPolicyInfoByInstanceIdResponse) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByInstanceIdResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByInstanceIdResponse) SetRequestId(v string) *GetPolicyInfoByInstanceIdResponse {
	s.RequestId = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponse) SetErrorMessage(v string) *GetPolicyInfoByInstanceIdResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponse) SetErrorCode(v string) *GetPolicyInfoByInstanceIdResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponse) SetState(v GetPolicyInfoByInstanceIdResponseStateEnum) *GetPolicyInfoByInstanceIdResponse {
	s.State = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponse) SetBody(v *GetPolicyInfoByInstanceIdResponseBody) *GetPolicyInfoByInstanceIdResponse {
	s.Body = v
	return s
}


type GetPolicyInfoByInstanceIdResponseBuilder struct {
	s *GetPolicyInfoByInstanceIdResponse
}

func NewGetPolicyInfoByInstanceIdResponseBuilder() *GetPolicyInfoByInstanceIdResponseBuilder {
	s := &GetPolicyInfoByInstanceIdResponse{}
	b := &GetPolicyInfoByInstanceIdResponseBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBuilder) RequestId(v string) *GetPolicyInfoByInstanceIdResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBuilder) ErrorMessage(v string) *GetPolicyInfoByInstanceIdResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBuilder) ErrorCode(v string) *GetPolicyInfoByInstanceIdResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBuilder) State(v GetPolicyInfoByInstanceIdResponseStateEnum) *GetPolicyInfoByInstanceIdResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBuilder) Body(v *GetPolicyInfoByInstanceIdResponseBody) *GetPolicyInfoByInstanceIdResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBuilder) Build() *GetPolicyInfoByInstanceIdResponse {
	return b.s
}


