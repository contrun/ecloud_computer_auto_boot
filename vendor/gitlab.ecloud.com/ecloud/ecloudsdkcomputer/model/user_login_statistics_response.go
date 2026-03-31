// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UserLoginStatisticsResponseStateEnum string

// List of State
const (
    UserLoginStatisticsResponseStateEnumOk UserLoginStatisticsResponseStateEnum = "OK"
    UserLoginStatisticsResponseStateEnumError UserLoginStatisticsResponseStateEnum = "ERROR"
    UserLoginStatisticsResponseStateEnumException UserLoginStatisticsResponseStateEnum = "EXCEPTION"
    UserLoginStatisticsResponseStateEnumForbidden UserLoginStatisticsResponseStateEnum = "FORBIDDEN"
    UserLoginStatisticsResponseStateEnumRemaining UserLoginStatisticsResponseStateEnum = "REMAINING"
    UserLoginStatisticsResponseStateEnumTimeout UserLoginStatisticsResponseStateEnum = "TIMEOUT"
)

type UserLoginStatisticsResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *UserLoginStatisticsResponseStateEnum `json:"state,omitempty"`
    // login：登录 logout：注销 start：开机 shutdown：关机 restart：重启 connect：连接 disconnect：断开 bindPhone：绑定手机号 unbindPhone：解绑手机号 reinstall：重装系统
	Body map[string]*int64 `json:"body,omitempty"`
}

func (s UserLoginStatisticsResponse) String() string {
	return utils.Beautify(s)
}

func (s UserLoginStatisticsResponse) GoString() string {
	return s.String()
}

func (s UserLoginStatisticsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserLoginStatisticsResponse) SetRequestId(v string) *UserLoginStatisticsResponse {
	s.RequestId = &v
	return s
}

func (s *UserLoginStatisticsResponse) SetErrorMessage(v string) *UserLoginStatisticsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UserLoginStatisticsResponse) SetErrorCode(v string) *UserLoginStatisticsResponse {
	s.ErrorCode = &v
	return s
}

func (s *UserLoginStatisticsResponse) SetState(v UserLoginStatisticsResponseStateEnum) *UserLoginStatisticsResponse {
	s.State = &v
	return s
}

func (s *UserLoginStatisticsResponse) SetBody(v map[string]*int64) *UserLoginStatisticsResponse {
	s.Body = v
	return s
}


type UserLoginStatisticsResponseBuilder struct {
	s *UserLoginStatisticsResponse
}

func NewUserLoginStatisticsResponseBuilder() *UserLoginStatisticsResponseBuilder {
	s := &UserLoginStatisticsResponse{}
	b := &UserLoginStatisticsResponseBuilder{s: s}
	return b
}

func (b *UserLoginStatisticsResponseBuilder) RequestId(v string) *UserLoginStatisticsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UserLoginStatisticsResponseBuilder) ErrorMessage(v string) *UserLoginStatisticsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UserLoginStatisticsResponseBuilder) ErrorCode(v string) *UserLoginStatisticsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UserLoginStatisticsResponseBuilder) State(v UserLoginStatisticsResponseStateEnum) *UserLoginStatisticsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UserLoginStatisticsResponseBuilder) Body(v map[string]*int64) *UserLoginStatisticsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *UserLoginStatisticsResponseBuilder) Build() *UserLoginStatisticsResponse {
	return b.s
}


