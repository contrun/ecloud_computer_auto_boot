// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DetailBody struct {
    position.Body
    // 消息id
	MessageId *string `json:"messageId,omitempty"`
}

func (s DetailBody) String() string {
	return utils.Beautify(s)
}

func (s DetailBody) GoString() string {
	return s.String()
}

func (s DetailBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DetailBody) SetMessageId(v string) *DetailBody {
	s.MessageId = &v
	return s
}


type DetailBodyBuilder struct {
	s *DetailBody
}

func NewDetailBodyBuilder() *DetailBodyBuilder {
	s := &DetailBody{}
	b := &DetailBodyBuilder{s: s}
	return b
}

func (b *DetailBodyBuilder) MessageId(v string) *DetailBodyBuilder {
	b.s.MessageId = &v
	return b
}

func (b *DetailBodyBuilder) Build() *DetailBody {
	return b.s
}


