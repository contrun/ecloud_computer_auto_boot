// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ImageSharedRecord2OwnResponseStateEnum string

// List of State
const (
    ImageSharedRecord2OwnResponseStateEnumOk ImageSharedRecord2OwnResponseStateEnum = "OK"
    ImageSharedRecord2OwnResponseStateEnumError ImageSharedRecord2OwnResponseStateEnum = "ERROR"
    ImageSharedRecord2OwnResponseStateEnumException ImageSharedRecord2OwnResponseStateEnum = "EXCEPTION"
    ImageSharedRecord2OwnResponseStateEnumForbidden ImageSharedRecord2OwnResponseStateEnum = "FORBIDDEN"
    ImageSharedRecord2OwnResponseStateEnumRemaining ImageSharedRecord2OwnResponseStateEnum = "REMAINING"
    ImageSharedRecord2OwnResponseStateEnumTimeout ImageSharedRecord2OwnResponseStateEnum = "TIMEOUT"
)

type ImageSharedRecord2OwnResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ImageSharedRecord2OwnResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ImageSharedRecord2OwnResponseBody `json:"body,omitempty"`
}

func (s ImageSharedRecord2OwnResponse) String() string {
	return utils.Beautify(s)
}

func (s ImageSharedRecord2OwnResponse) GoString() string {
	return s.String()
}

func (s ImageSharedRecord2OwnResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageSharedRecord2OwnResponse) SetRequestId(v string) *ImageSharedRecord2OwnResponse {
	s.RequestId = &v
	return s
}

func (s *ImageSharedRecord2OwnResponse) SetErrorMessage(v string) *ImageSharedRecord2OwnResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ImageSharedRecord2OwnResponse) SetErrorCode(v string) *ImageSharedRecord2OwnResponse {
	s.ErrorCode = &v
	return s
}

func (s *ImageSharedRecord2OwnResponse) SetState(v ImageSharedRecord2OwnResponseStateEnum) *ImageSharedRecord2OwnResponse {
	s.State = &v
	return s
}

func (s *ImageSharedRecord2OwnResponse) SetBody(v *ImageSharedRecord2OwnResponseBody) *ImageSharedRecord2OwnResponse {
	s.Body = v
	return s
}


type ImageSharedRecord2OwnResponseBuilder struct {
	s *ImageSharedRecord2OwnResponse
}

func NewImageSharedRecord2OwnResponseBuilder() *ImageSharedRecord2OwnResponseBuilder {
	s := &ImageSharedRecord2OwnResponse{}
	b := &ImageSharedRecord2OwnResponseBuilder{s: s}
	return b
}

func (b *ImageSharedRecord2OwnResponseBuilder) RequestId(v string) *ImageSharedRecord2OwnResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseBuilder) ErrorMessage(v string) *ImageSharedRecord2OwnResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseBuilder) ErrorCode(v string) *ImageSharedRecord2OwnResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseBuilder) State(v ImageSharedRecord2OwnResponseStateEnum) *ImageSharedRecord2OwnResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ImageSharedRecord2OwnResponseBuilder) Body(v *ImageSharedRecord2OwnResponseBody) *ImageSharedRecord2OwnResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ImageSharedRecord2OwnResponseBuilder) Build() *ImageSharedRecord2OwnResponse {
	return b.s
}


