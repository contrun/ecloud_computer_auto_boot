// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteBody struct {
    position.Body
    // 消息id
	MessageId *string `json:"messageId,omitempty"`
}

func (s DeleteBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteBody) GoString() string {
	return s.String()
}

func (s DeleteBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteBody) SetMessageId(v string) *DeleteBody {
	s.MessageId = &v
	return s
}


type DeleteBodyBuilder struct {
	s *DeleteBody
}

func NewDeleteBodyBuilder() *DeleteBodyBuilder {
	s := &DeleteBody{}
	b := &DeleteBodyBuilder{s: s}
	return b
}

func (b *DeleteBodyBuilder) MessageId(v string) *DeleteBodyBuilder {
	b.s.MessageId = &v
	return b
}

func (b *DeleteBodyBuilder) Build() *DeleteBody {
	return b.s
}


