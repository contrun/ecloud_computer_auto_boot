// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetOrgListResponseStateEnum string

// List of State
const (
    GetOrgListResponseStateEnumOk GetOrgListResponseStateEnum = "OK"
    GetOrgListResponseStateEnumError GetOrgListResponseStateEnum = "ERROR"
    GetOrgListResponseStateEnumException GetOrgListResponseStateEnum = "EXCEPTION"
    GetOrgListResponseStateEnumForbidden GetOrgListResponseStateEnum = "FORBIDDEN"
    GetOrgListResponseStateEnumRemaining GetOrgListResponseStateEnum = "REMAINING"
    GetOrgListResponseStateEnumTimeout GetOrgListResponseStateEnum = "TIMEOUT"
)

type GetOrgListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetOrgListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetOrgListResponseBody `json:"body,omitempty"`
}

func (s GetOrgListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetOrgListResponse) GoString() string {
	return s.String()
}

func (s GetOrgListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetOrgListResponse) SetRequestId(v string) *GetOrgListResponse {
	s.RequestId = &v
	return s
}

func (s *GetOrgListResponse) SetErrorMessage(v string) *GetOrgListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetOrgListResponse) SetErrorCode(v string) *GetOrgListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetOrgListResponse) SetState(v GetOrgListResponseStateEnum) *GetOrgListResponse {
	s.State = &v
	return s
}

func (s *GetOrgListResponse) SetBody(v *GetOrgListResponseBody) *GetOrgListResponse {
	s.Body = v
	return s
}


type GetOrgListResponseBuilder struct {
	s *GetOrgListResponse
}

func NewGetOrgListResponseBuilder() *GetOrgListResponseBuilder {
	s := &GetOrgListResponse{}
	b := &GetOrgListResponseBuilder{s: s}
	return b
}

func (b *GetOrgListResponseBuilder) RequestId(v string) *GetOrgListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetOrgListResponseBuilder) ErrorMessage(v string) *GetOrgListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetOrgListResponseBuilder) ErrorCode(v string) *GetOrgListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetOrgListResponseBuilder) State(v GetOrgListResponseStateEnum) *GetOrgListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetOrgListResponseBuilder) Body(v *GetOrgListResponseBody) *GetOrgListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetOrgListResponseBuilder) Build() *GetOrgListResponse {
	return b.s
}


