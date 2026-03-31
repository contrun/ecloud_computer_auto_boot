// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPublicImagePageResponseStateEnum string

// List of State
const (
    GetPublicImagePageResponseStateEnumOk GetPublicImagePageResponseStateEnum = "OK"
    GetPublicImagePageResponseStateEnumError GetPublicImagePageResponseStateEnum = "ERROR"
    GetPublicImagePageResponseStateEnumException GetPublicImagePageResponseStateEnum = "EXCEPTION"
    GetPublicImagePageResponseStateEnumForbidden GetPublicImagePageResponseStateEnum = "FORBIDDEN"
    GetPublicImagePageResponseStateEnumRemaining GetPublicImagePageResponseStateEnum = "REMAINING"
    GetPublicImagePageResponseStateEnumTimeout GetPublicImagePageResponseStateEnum = "TIMEOUT"
)

type GetPublicImagePageResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    //  统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetPublicImagePageResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetPublicImagePageResponseBody `json:"body,omitempty"`
}

func (s GetPublicImagePageResponse) String() string {
	return utils.Beautify(s)
}

func (s GetPublicImagePageResponse) GoString() string {
	return s.String()
}

func (s GetPublicImagePageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPublicImagePageResponse) SetRequestId(v string) *GetPublicImagePageResponse {
	s.RequestId = &v
	return s
}

func (s *GetPublicImagePageResponse) SetErrorMessage(v string) *GetPublicImagePageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetPublicImagePageResponse) SetErrorCode(v string) *GetPublicImagePageResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetPublicImagePageResponse) SetState(v GetPublicImagePageResponseStateEnum) *GetPublicImagePageResponse {
	s.State = &v
	return s
}

func (s *GetPublicImagePageResponse) SetBody(v *GetPublicImagePageResponseBody) *GetPublicImagePageResponse {
	s.Body = v
	return s
}


type GetPublicImagePageResponseBuilder struct {
	s *GetPublicImagePageResponse
}

func NewGetPublicImagePageResponseBuilder() *GetPublicImagePageResponseBuilder {
	s := &GetPublicImagePageResponse{}
	b := &GetPublicImagePageResponseBuilder{s: s}
	return b
}

func (b *GetPublicImagePageResponseBuilder) RequestId(v string) *GetPublicImagePageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetPublicImagePageResponseBuilder) ErrorMessage(v string) *GetPublicImagePageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetPublicImagePageResponseBuilder) ErrorCode(v string) *GetPublicImagePageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetPublicImagePageResponseBuilder) State(v GetPublicImagePageResponseStateEnum) *GetPublicImagePageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetPublicImagePageResponseBuilder) Body(v *GetPublicImagePageResponseBody) *GetPublicImagePageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetPublicImagePageResponseBuilder) Build() *GetPublicImagePageResponse {
	return b.s
}


