// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetIsBindUserInfoListResponseStateEnum string

// List of State
const (
    GetIsBindUserInfoListResponseStateEnumOk GetIsBindUserInfoListResponseStateEnum = "OK"
    GetIsBindUserInfoListResponseStateEnumError GetIsBindUserInfoListResponseStateEnum = "ERROR"
    GetIsBindUserInfoListResponseStateEnumException GetIsBindUserInfoListResponseStateEnum = "EXCEPTION"
    GetIsBindUserInfoListResponseStateEnumForbidden GetIsBindUserInfoListResponseStateEnum = "FORBIDDEN"
    GetIsBindUserInfoListResponseStateEnumRemaining GetIsBindUserInfoListResponseStateEnum = "REMAINING"
    GetIsBindUserInfoListResponseStateEnumTimeout GetIsBindUserInfoListResponseStateEnum = "TIMEOUT"
)

type GetIsBindUserInfoListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetIsBindUserInfoListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetIsBindUserInfoListResponseBody `json:"body,omitempty"`
}

func (s GetIsBindUserInfoListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetIsBindUserInfoListResponse) GoString() string {
	return s.String()
}

func (s GetIsBindUserInfoListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetIsBindUserInfoListResponse) SetRequestId(v string) *GetIsBindUserInfoListResponse {
	s.RequestId = &v
	return s
}

func (s *GetIsBindUserInfoListResponse) SetErrorMessage(v string) *GetIsBindUserInfoListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetIsBindUserInfoListResponse) SetErrorCode(v string) *GetIsBindUserInfoListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetIsBindUserInfoListResponse) SetState(v GetIsBindUserInfoListResponseStateEnum) *GetIsBindUserInfoListResponse {
	s.State = &v
	return s
}

func (s *GetIsBindUserInfoListResponse) SetBody(v *GetIsBindUserInfoListResponseBody) *GetIsBindUserInfoListResponse {
	s.Body = v
	return s
}


type GetIsBindUserInfoListResponseBuilder struct {
	s *GetIsBindUserInfoListResponse
}

func NewGetIsBindUserInfoListResponseBuilder() *GetIsBindUserInfoListResponseBuilder {
	s := &GetIsBindUserInfoListResponse{}
	b := &GetIsBindUserInfoListResponseBuilder{s: s}
	return b
}

func (b *GetIsBindUserInfoListResponseBuilder) RequestId(v string) *GetIsBindUserInfoListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetIsBindUserInfoListResponseBuilder) ErrorMessage(v string) *GetIsBindUserInfoListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetIsBindUserInfoListResponseBuilder) ErrorCode(v string) *GetIsBindUserInfoListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetIsBindUserInfoListResponseBuilder) State(v GetIsBindUserInfoListResponseStateEnum) *GetIsBindUserInfoListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetIsBindUserInfoListResponseBuilder) Body(v *GetIsBindUserInfoListResponseBody) *GetIsBindUserInfoListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetIsBindUserInfoListResponseBuilder) Build() *GetIsBindUserInfoListResponse {
	return b.s
}


