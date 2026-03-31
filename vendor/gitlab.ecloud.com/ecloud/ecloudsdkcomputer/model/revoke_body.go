// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RevokeBody struct {
    position.Body
    // 消息id
	MessageId *string `json:"messageId,omitempty"`
}

func (s RevokeBody) String() string {
	return utils.Beautify(s)
}

func (s RevokeBody) GoString() string {
	return s.String()
}

func (s RevokeBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RevokeBody) SetMessageId(v string) *RevokeBody {
	s.MessageId = &v
	return s
}


type RevokeBodyBuilder struct {
	s *RevokeBody
}

func NewRevokeBodyBuilder() *RevokeBodyBuilder {
	s := &RevokeBody{}
	b := &RevokeBodyBuilder{s: s}
	return b
}

func (b *RevokeBodyBuilder) MessageId(v string) *RevokeBodyBuilder {
	b.s.MessageId = &v
	return b
}

func (b *RevokeBodyBuilder) Build() *RevokeBody {
	return b.s
}


