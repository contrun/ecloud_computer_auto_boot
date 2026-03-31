// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ReloadByPublicImageResponseStateEnum string

// List of State
const (
    ReloadByPublicImageResponseStateEnumOk ReloadByPublicImageResponseStateEnum = "OK"
    ReloadByPublicImageResponseStateEnumError ReloadByPublicImageResponseStateEnum = "ERROR"
    ReloadByPublicImageResponseStateEnumException ReloadByPublicImageResponseStateEnum = "EXCEPTION"
    ReloadByPublicImageResponseStateEnumForbidden ReloadByPublicImageResponseStateEnum = "FORBIDDEN"
    ReloadByPublicImageResponseStateEnumRemaining ReloadByPublicImageResponseStateEnum = "REMAINING"
    ReloadByPublicImageResponseStateEnumTimeout ReloadByPublicImageResponseStateEnum = "TIMEOUT"
)

type ReloadByPublicImageResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ReloadByPublicImageResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s ReloadByPublicImageResponse) String() string {
	return utils.Beautify(s)
}

func (s ReloadByPublicImageResponse) GoString() string {
	return s.String()
}

func (s ReloadByPublicImageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ReloadByPublicImageResponse) SetRequestId(v string) *ReloadByPublicImageResponse {
	s.RequestId = &v
	return s
}

func (s *ReloadByPublicImageResponse) SetErrorMessage(v string) *ReloadByPublicImageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ReloadByPublicImageResponse) SetErrorCode(v string) *ReloadByPublicImageResponse {
	s.ErrorCode = &v
	return s
}

func (s *ReloadByPublicImageResponse) SetState(v ReloadByPublicImageResponseStateEnum) *ReloadByPublicImageResponse {
	s.State = &v
	return s
}

func (s *ReloadByPublicImageResponse) SetBody(v string) *ReloadByPublicImageResponse {
	s.Body = &v
	return s
}


type ReloadByPublicImageResponseBuilder struct {
	s *ReloadByPublicImageResponse
}

func NewReloadByPublicImageResponseBuilder() *ReloadByPublicImageResponseBuilder {
	s := &ReloadByPublicImageResponse{}
	b := &ReloadByPublicImageResponseBuilder{s: s}
	return b
}

func (b *ReloadByPublicImageResponseBuilder) RequestId(v string) *ReloadByPublicImageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ReloadByPublicImageResponseBuilder) ErrorMessage(v string) *ReloadByPublicImageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ReloadByPublicImageResponseBuilder) ErrorCode(v string) *ReloadByPublicImageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ReloadByPublicImageResponseBuilder) State(v ReloadByPublicImageResponseStateEnum) *ReloadByPublicImageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ReloadByPublicImageResponseBuilder) Body(v string) *ReloadByPublicImageResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *ReloadByPublicImageResponseBuilder) Build() *ReloadByPublicImageResponse {
	return b.s
}


