// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateBody struct {
    position.Body
    // 消息查看方式: 0静默消息,1强制消息
	Mode *int32 `json:"mode,omitempty"`
    // 消息主题
	Subject *string `json:"subject,omitempty"`
    // 用户id列表
	UserIds []string `json:"userIds,omitempty"`
    // 发送类型:0:立刻发送 ,1 :定时发送
	SendType *int32 `json:"sendType,omitempty"`
    // 消息id
	MessageId *string `json:"messageId,omitempty"`
    // 消息内容
	Content *string `json:"content,omitempty"`
    // 发送时间
	SendTime *string `json:"sendTime,omitempty"`
}

func (s UpdateBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateBody) GoString() string {
	return s.String()
}

func (s UpdateBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateBody) SetMode(v int32) *UpdateBody {
	s.Mode = &v
	return s
}

func (s *UpdateBody) SetSubject(v string) *UpdateBody {
	s.Subject = &v
	return s
}

func (s *UpdateBody) SetUserIds(v []string) *UpdateBody {
	s.UserIds = v
	return s
}

func (s *UpdateBody) SetSendType(v int32) *UpdateBody {
	s.SendType = &v
	return s
}

func (s *UpdateBody) SetMessageId(v string) *UpdateBody {
	s.MessageId = &v
	return s
}

func (s *UpdateBody) SetContent(v string) *UpdateBody {
	s.Content = &v
	return s
}

func (s *UpdateBody) SetSendTime(v string) *UpdateBody {
	s.SendTime = &v
	return s
}


type UpdateBodyBuilder struct {
	s *UpdateBody
}

func NewUpdateBodyBuilder() *UpdateBodyBuilder {
	s := &UpdateBody{}
	b := &UpdateBodyBuilder{s: s}
	return b
}

func (b *UpdateBodyBuilder) Mode(v int32) *UpdateBodyBuilder {
	b.s.Mode = &v
	return b
}

func (b *UpdateBodyBuilder) Subject(v string) *UpdateBodyBuilder {
	b.s.Subject = &v
	return b
}

func (b *UpdateBodyBuilder) UserIds(v []string) *UpdateBodyBuilder {
	b.s.UserIds = v
	return b
}

func (b *UpdateBodyBuilder) SendType(v int32) *UpdateBodyBuilder {
	b.s.SendType = &v
	return b
}

func (b *UpdateBodyBuilder) MessageId(v string) *UpdateBodyBuilder {
	b.s.MessageId = &v
	return b
}

func (b *UpdateBodyBuilder) Content(v string) *UpdateBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *UpdateBodyBuilder) SendTime(v string) *UpdateBodyBuilder {
	b.s.SendTime = &v
	return b
}

func (b *UpdateBodyBuilder) Build() *UpdateBody {
	return b.s
}


