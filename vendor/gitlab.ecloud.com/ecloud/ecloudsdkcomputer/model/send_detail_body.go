// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SendDetailBody struct {
    position.Body
    // 消息id
	MessageId *string `json:"messageId,omitempty"`
}

func (s SendDetailBody) String() string {
	return utils.Beautify(s)
}

func (s SendDetailBody) GoString() string {
	return s.String()
}

func (s SendDetailBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SendDetailBody) SetMessageId(v string) *SendDetailBody {
	s.MessageId = &v
	return s
}


type SendDetailBodyBuilder struct {
	s *SendDetailBody
}

func NewSendDetailBodyBuilder() *SendDetailBodyBuilder {
	s := &SendDetailBody{}
	b := &SendDetailBodyBuilder{s: s}
	return b
}

func (b *SendDetailBodyBuilder) MessageId(v string) *SendDetailBodyBuilder {
	b.s.MessageId = &v
	return b
}

func (b *SendDetailBodyBuilder) Build() *SendDetailBody {
	return b.s
}


