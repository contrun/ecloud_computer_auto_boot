// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DetailResponseBody struct {

    // 消息方式: 0静默 1强制
	Mode *int32 `json:"mode,omitempty"`
    // 消息主题
	Subject *string `json:"subject,omitempty"`
    // 消息id(UUID)
	MessageId *string `json:"messageId,omitempty"`
    // 消息内容
	Content *string `json:"content,omitempty"`
    // 发送时间
	SendTime *string `json:"sendTime,omitempty"`
    // 发送状态: 0:待发送,1:已发送,2:已撤回
	Status *int32 `json:"status,omitempty"`
}

func (s DetailResponseBody) String() string {
	return utils.Beautify(s)
}

func (s DetailResponseBody) GoString() string {
	return s.String()
}

func (s DetailResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DetailResponseBody) SetMode(v int32) *DetailResponseBody {
	s.Mode = &v
	return s
}

func (s *DetailResponseBody) SetSubject(v string) *DetailResponseBody {
	s.Subject = &v
	return s
}

func (s *DetailResponseBody) SetMessageId(v string) *DetailResponseBody {
	s.MessageId = &v
	return s
}

func (s *DetailResponseBody) SetContent(v string) *DetailResponseBody {
	s.Content = &v
	return s
}

func (s *DetailResponseBody) SetSendTime(v string) *DetailResponseBody {
	s.SendTime = &v
	return s
}

func (s *DetailResponseBody) SetStatus(v int32) *DetailResponseBody {
	s.Status = &v
	return s
}


type DetailResponseBodyBuilder struct {
	s *DetailResponseBody
}

func NewDetailResponseBodyBuilder() *DetailResponseBodyBuilder {
	s := &DetailResponseBody{}
	b := &DetailResponseBodyBuilder{s: s}
	return b
}

func (b *DetailResponseBodyBuilder) Mode(v int32) *DetailResponseBodyBuilder {
	b.s.Mode = &v
	return b
}

func (b *DetailResponseBodyBuilder) Subject(v string) *DetailResponseBodyBuilder {
	b.s.Subject = &v
	return b
}

func (b *DetailResponseBodyBuilder) MessageId(v string) *DetailResponseBodyBuilder {
	b.s.MessageId = &v
	return b
}

func (b *DetailResponseBodyBuilder) Content(v string) *DetailResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *DetailResponseBodyBuilder) SendTime(v string) *DetailResponseBodyBuilder {
	b.s.SendTime = &v
	return b
}

func (b *DetailResponseBodyBuilder) Status(v int32) *DetailResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *DetailResponseBodyBuilder) Build() *DetailResponseBody {
	return b.s
}


