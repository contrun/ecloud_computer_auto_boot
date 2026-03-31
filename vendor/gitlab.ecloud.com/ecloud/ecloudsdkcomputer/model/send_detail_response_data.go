// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SendDetailResponseData struct {

    // 消息已读: 0未读,1已读
	Readed *int32 `json:"readed,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 授权状态: 0:未授权,1:已授权
	AuthorizationState *string `json:"authorizationState,omitempty"`
}

func (s SendDetailResponseData) String() string {
	return utils.Beautify(s)
}

func (s SendDetailResponseData) GoString() string {
	return s.String()
}

func (s SendDetailResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SendDetailResponseData) SetReaded(v int32) *SendDetailResponseData {
	s.Readed = &v
	return s
}

func (s *SendDetailResponseData) SetUserName(v string) *SendDetailResponseData {
	s.UserName = &v
	return s
}

func (s *SendDetailResponseData) SetAuthorizationState(v string) *SendDetailResponseData {
	s.AuthorizationState = &v
	return s
}


type SendDetailResponseDataBuilder struct {
	s *SendDetailResponseData
}

func NewSendDetailResponseDataBuilder() *SendDetailResponseDataBuilder {
	s := &SendDetailResponseData{}
	b := &SendDetailResponseDataBuilder{s: s}
	return b
}

func (b *SendDetailResponseDataBuilder) Readed(v int32) *SendDetailResponseDataBuilder {
	b.s.Readed = &v
	return b
}

func (b *SendDetailResponseDataBuilder) UserName(v string) *SendDetailResponseDataBuilder {
	b.s.UserName = &v
	return b
}

func (b *SendDetailResponseDataBuilder) AuthorizationState(v string) *SendDetailResponseDataBuilder {
	b.s.AuthorizationState = &v
	return b
}

func (b *SendDetailResponseDataBuilder) Build() *SendDetailResponseData {
	return b.s
}


