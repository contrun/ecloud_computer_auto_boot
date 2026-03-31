// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ResourceListResponseStateEnum string

// List of State
const (
    ResourceListResponseStateEnumOk ResourceListResponseStateEnum = "OK"
    ResourceListResponseStateEnumError ResourceListResponseStateEnum = "ERROR"
    ResourceListResponseStateEnumException ResourceListResponseStateEnum = "EXCEPTION"
    ResourceListResponseStateEnumForbidden ResourceListResponseStateEnum = "FORBIDDEN"
    ResourceListResponseStateEnumRemaining ResourceListResponseStateEnum = "REMAINING"
    ResourceListResponseStateEnumTimeout ResourceListResponseStateEnum = "TIMEOUT"
)

type ResourceListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ResourceListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ResourceListResponseBody `json:"body,omitempty"`
}

func (s ResourceListResponse) String() string {
	return utils.Beautify(s)
}

func (s ResourceListResponse) GoString() string {
	return s.String()
}

func (s ResourceListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ResourceListResponse) SetRequestId(v string) *ResourceListResponse {
	s.RequestId = &v
	return s
}

func (s *ResourceListResponse) SetErrorMessage(v string) *ResourceListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ResourceListResponse) SetErrorCode(v string) *ResourceListResponse {
	s.ErrorCode = &v
	return s
}

func (s *ResourceListResponse) SetState(v ResourceListResponseStateEnum) *ResourceListResponse {
	s.State = &v
	return s
}

func (s *ResourceListResponse) SetBody(v *ResourceListResponseBody) *ResourceListResponse {
	s.Body = v
	return s
}


type ResourceListResponseBuilder struct {
	s *ResourceListResponse
}

func NewResourceListResponseBuilder() *ResourceListResponseBuilder {
	s := &ResourceListResponse{}
	b := &ResourceListResponseBuilder{s: s}
	return b
}

func (b *ResourceListResponseBuilder) RequestId(v string) *ResourceListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ResourceListResponseBuilder) ErrorMessage(v string) *ResourceListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ResourceListResponseBuilder) ErrorCode(v string) *ResourceListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ResourceListResponseBuilder) State(v ResourceListResponseStateEnum) *ResourceListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ResourceListResponseBuilder) Body(v *ResourceListResponseBody) *ResourceListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ResourceListResponseBuilder) Build() *ResourceListResponse {
	return b.s
}


