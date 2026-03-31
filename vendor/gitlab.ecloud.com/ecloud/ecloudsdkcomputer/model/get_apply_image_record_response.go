// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetApplyImageRecordResponseStateEnum string

// List of State
const (
    GetApplyImageRecordResponseStateEnumOk GetApplyImageRecordResponseStateEnum = "OK"
    GetApplyImageRecordResponseStateEnumError GetApplyImageRecordResponseStateEnum = "ERROR"
    GetApplyImageRecordResponseStateEnumException GetApplyImageRecordResponseStateEnum = "EXCEPTION"
    GetApplyImageRecordResponseStateEnumForbidden GetApplyImageRecordResponseStateEnum = "FORBIDDEN"
    GetApplyImageRecordResponseStateEnumRemaining GetApplyImageRecordResponseStateEnum = "REMAINING"
    GetApplyImageRecordResponseStateEnumTimeout GetApplyImageRecordResponseStateEnum = "TIMEOUT"
)

type GetApplyImageRecordResponse struct {


	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetApplyImageRecordResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetApplyImageRecordResponseBody `json:"body,omitempty"`
}

func (s GetApplyImageRecordResponse) String() string {
	return utils.Beautify(s)
}

func (s GetApplyImageRecordResponse) GoString() string {
	return s.String()
}

func (s GetApplyImageRecordResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplyImageRecordResponse) SetRequestId(v string) *GetApplyImageRecordResponse {
	s.RequestId = &v
	return s
}

func (s *GetApplyImageRecordResponse) SetErrorMessage(v string) *GetApplyImageRecordResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetApplyImageRecordResponse) SetErrorCode(v string) *GetApplyImageRecordResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetApplyImageRecordResponse) SetState(v GetApplyImageRecordResponseStateEnum) *GetApplyImageRecordResponse {
	s.State = &v
	return s
}

func (s *GetApplyImageRecordResponse) SetBody(v *GetApplyImageRecordResponseBody) *GetApplyImageRecordResponse {
	s.Body = v
	return s
}


type GetApplyImageRecordResponseBuilder struct {
	s *GetApplyImageRecordResponse
}

func NewGetApplyImageRecordResponseBuilder() *GetApplyImageRecordResponseBuilder {
	s := &GetApplyImageRecordResponse{}
	b := &GetApplyImageRecordResponseBuilder{s: s}
	return b
}

func (b *GetApplyImageRecordResponseBuilder) RequestId(v string) *GetApplyImageRecordResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetApplyImageRecordResponseBuilder) ErrorMessage(v string) *GetApplyImageRecordResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetApplyImageRecordResponseBuilder) ErrorCode(v string) *GetApplyImageRecordResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetApplyImageRecordResponseBuilder) State(v GetApplyImageRecordResponseStateEnum) *GetApplyImageRecordResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetApplyImageRecordResponseBuilder) Body(v *GetApplyImageRecordResponseBody) *GetApplyImageRecordResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetApplyImageRecordResponseBuilder) Build() *GetApplyImageRecordResponse {
	return b.s
}


