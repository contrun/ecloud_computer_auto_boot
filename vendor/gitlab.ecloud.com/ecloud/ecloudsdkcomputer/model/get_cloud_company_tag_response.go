// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetCloudCompanyTagResponseStateEnum string

// List of State
const (
    GetCloudCompanyTagResponseStateEnumOk GetCloudCompanyTagResponseStateEnum = "OK"
    GetCloudCompanyTagResponseStateEnumError GetCloudCompanyTagResponseStateEnum = "ERROR"
    GetCloudCompanyTagResponseStateEnumException GetCloudCompanyTagResponseStateEnum = "EXCEPTION"
    GetCloudCompanyTagResponseStateEnumForbidden GetCloudCompanyTagResponseStateEnum = "FORBIDDEN"
    GetCloudCompanyTagResponseStateEnumRemaining GetCloudCompanyTagResponseStateEnum = "REMAINING"
    GetCloudCompanyTagResponseStateEnumTimeout GetCloudCompanyTagResponseStateEnum = "TIMEOUT"
)

type GetCloudCompanyTagResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetCloudCompanyTagResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s GetCloudCompanyTagResponse) String() string {
	return utils.Beautify(s)
}

func (s GetCloudCompanyTagResponse) GoString() string {
	return s.String()
}

func (s GetCloudCompanyTagResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCloudCompanyTagResponse) SetRequestId(v string) *GetCloudCompanyTagResponse {
	s.RequestId = &v
	return s
}

func (s *GetCloudCompanyTagResponse) SetErrorMessage(v string) *GetCloudCompanyTagResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetCloudCompanyTagResponse) SetErrorCode(v string) *GetCloudCompanyTagResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetCloudCompanyTagResponse) SetState(v GetCloudCompanyTagResponseStateEnum) *GetCloudCompanyTagResponse {
	s.State = &v
	return s
}

func (s *GetCloudCompanyTagResponse) SetBody(v string) *GetCloudCompanyTagResponse {
	s.Body = &v
	return s
}


type GetCloudCompanyTagResponseBuilder struct {
	s *GetCloudCompanyTagResponse
}

func NewGetCloudCompanyTagResponseBuilder() *GetCloudCompanyTagResponseBuilder {
	s := &GetCloudCompanyTagResponse{}
	b := &GetCloudCompanyTagResponseBuilder{s: s}
	return b
}

func (b *GetCloudCompanyTagResponseBuilder) RequestId(v string) *GetCloudCompanyTagResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetCloudCompanyTagResponseBuilder) ErrorMessage(v string) *GetCloudCompanyTagResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetCloudCompanyTagResponseBuilder) ErrorCode(v string) *GetCloudCompanyTagResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetCloudCompanyTagResponseBuilder) State(v GetCloudCompanyTagResponseStateEnum) *GetCloudCompanyTagResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetCloudCompanyTagResponseBuilder) Body(v string) *GetCloudCompanyTagResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetCloudCompanyTagResponseBuilder) Build() *GetCloudCompanyTagResponse {
	return b.s
}


