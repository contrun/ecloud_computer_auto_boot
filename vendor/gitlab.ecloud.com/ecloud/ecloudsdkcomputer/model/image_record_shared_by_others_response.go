// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ImageRecordSharedByOthersResponseStateEnum string

// List of State
const (
    ImageRecordSharedByOthersResponseStateEnumOk ImageRecordSharedByOthersResponseStateEnum = "OK"
    ImageRecordSharedByOthersResponseStateEnumError ImageRecordSharedByOthersResponseStateEnum = "ERROR"
    ImageRecordSharedByOthersResponseStateEnumException ImageRecordSharedByOthersResponseStateEnum = "EXCEPTION"
    ImageRecordSharedByOthersResponseStateEnumForbidden ImageRecordSharedByOthersResponseStateEnum = "FORBIDDEN"
    ImageRecordSharedByOthersResponseStateEnumRemaining ImageRecordSharedByOthersResponseStateEnum = "REMAINING"
    ImageRecordSharedByOthersResponseStateEnumTimeout ImageRecordSharedByOthersResponseStateEnum = "TIMEOUT"
)

type ImageRecordSharedByOthersResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ImageRecordSharedByOthersResponseStateEnum `json:"state,omitempty"`

	Body *ImageRecordSharedByOthersResponseBody `json:"body,omitempty"`
}

func (s ImageRecordSharedByOthersResponse) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedByOthersResponse) GoString() string {
	return s.String()
}

func (s ImageRecordSharedByOthersResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedByOthersResponse) SetRequestId(v string) *ImageRecordSharedByOthersResponse {
	s.RequestId = &v
	return s
}

func (s *ImageRecordSharedByOthersResponse) SetErrorMessage(v string) *ImageRecordSharedByOthersResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ImageRecordSharedByOthersResponse) SetErrorCode(v string) *ImageRecordSharedByOthersResponse {
	s.ErrorCode = &v
	return s
}

func (s *ImageRecordSharedByOthersResponse) SetState(v ImageRecordSharedByOthersResponseStateEnum) *ImageRecordSharedByOthersResponse {
	s.State = &v
	return s
}

func (s *ImageRecordSharedByOthersResponse) SetBody(v *ImageRecordSharedByOthersResponseBody) *ImageRecordSharedByOthersResponse {
	s.Body = v
	return s
}


type ImageRecordSharedByOthersResponseBuilder struct {
	s *ImageRecordSharedByOthersResponse
}

func NewImageRecordSharedByOthersResponseBuilder() *ImageRecordSharedByOthersResponseBuilder {
	s := &ImageRecordSharedByOthersResponse{}
	b := &ImageRecordSharedByOthersResponseBuilder{s: s}
	return b
}

func (b *ImageRecordSharedByOthersResponseBuilder) RequestId(v string) *ImageRecordSharedByOthersResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseBuilder) ErrorMessage(v string) *ImageRecordSharedByOthersResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseBuilder) ErrorCode(v string) *ImageRecordSharedByOthersResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseBuilder) State(v ImageRecordSharedByOthersResponseStateEnum) *ImageRecordSharedByOthersResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ImageRecordSharedByOthersResponseBuilder) Body(v *ImageRecordSharedByOthersResponseBody) *ImageRecordSharedByOthersResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ImageRecordSharedByOthersResponseBuilder) Build() *ImageRecordSharedByOthersResponse {
	return b.s
}


