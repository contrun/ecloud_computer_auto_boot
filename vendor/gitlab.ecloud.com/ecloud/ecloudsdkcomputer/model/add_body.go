// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddBody struct {
    position.Body
    // 消息查看方式: 0静默消息,1强制消息
	Mode *int32 `json:"mode,omitempty"`
    // 消息主题
	Subject *string `json:"subject,omitempty"`
    // 用户id列表
	UserIds []string `json:"userIds,omitempty"`
    // 发送类型:0:立刻发送 ,1 :定时发送
	SendType *int32 `json:"sendType,omitempty"`
    // 消息内容
	Content *string `json:"content,omitempty"`
    // 发送时间
	SendTime *string `json:"sendTime,omitempty"`
}

func (s AddBody) String() string {
	return utils.Beautify(s)
}

func (s AddBody) GoString() string {
	return s.String()
}

func (s AddBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddBody) SetMode(v int32) *AddBody {
	s.Mode = &v
	return s
}

func (s *AddBody) SetSubject(v string) *AddBody {
	s.Subject = &v
	return s
}

func (s *AddBody) SetUserIds(v []string) *AddBody {
	s.UserIds = v
	return s
}

func (s *AddBody) SetSendType(v int32) *AddBody {
	s.SendType = &v
	return s
}

func (s *AddBody) SetContent(v string) *AddBody {
	s.Content = &v
	return s
}

func (s *AddBody) SetSendTime(v string) *AddBody {
	s.SendTime = &v
	return s
}


type AddBodyBuilder struct {
	s *AddBody
}

func NewAddBodyBuilder() *AddBodyBuilder {
	s := &AddBody{}
	b := &AddBodyBuilder{s: s}
	return b
}

func (b *AddBodyBuilder) Mode(v int32) *AddBodyBuilder {
	b.s.Mode = &v
	return b
}

func (b *AddBodyBuilder) Subject(v string) *AddBodyBuilder {
	b.s.Subject = &v
	return b
}

func (b *AddBodyBuilder) UserIds(v []string) *AddBodyBuilder {
	b.s.UserIds = v
	return b
}

func (b *AddBodyBuilder) SendType(v int32) *AddBodyBuilder {
	b.s.SendType = &v
	return b
}

func (b *AddBodyBuilder) Content(v string) *AddBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *AddBodyBuilder) SendTime(v string) *AddBodyBuilder {
	b.s.SendTime = &v
	return b
}

func (b *AddBodyBuilder) Build() *AddBody {
	return b.s
}


