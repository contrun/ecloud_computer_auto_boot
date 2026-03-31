// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListResponseData struct {

    // 消息主题
	Subject *string `json:"subject,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 消息ID(UUID)
	MessageId *string `json:"messageId,omitempty"`
    // 已阅用户数
	ReadUser *int32 `json:"readUser,omitempty"`
    // 目标用户数
	AllUser *int32 `json:"allUser,omitempty"`
    // 发送时间
	SendTime *string `json:"sendTime,omitempty"`
    // 发送状态: 0:待发送,1:已发送,2:已撤回
	Status *int32 `json:"status,omitempty"`
}

func (s ListResponseData) String() string {
	return utils.Beautify(s)
}

func (s ListResponseData) GoString() string {
	return s.String()
}

func (s ListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListResponseData) SetSubject(v string) *ListResponseData {
	s.Subject = &v
	return s
}

func (s *ListResponseData) SetCreatedTime(v string) *ListResponseData {
	s.CreatedTime = &v
	return s
}

func (s *ListResponseData) SetMessageId(v string) *ListResponseData {
	s.MessageId = &v
	return s
}

func (s *ListResponseData) SetReadUser(v int32) *ListResponseData {
	s.ReadUser = &v
	return s
}

func (s *ListResponseData) SetAllUser(v int32) *ListResponseData {
	s.AllUser = &v
	return s
}

func (s *ListResponseData) SetSendTime(v string) *ListResponseData {
	s.SendTime = &v
	return s
}

func (s *ListResponseData) SetStatus(v int32) *ListResponseData {
	s.Status = &v
	return s
}


type ListResponseDataBuilder struct {
	s *ListResponseData
}

func NewListResponseDataBuilder() *ListResponseDataBuilder {
	s := &ListResponseData{}
	b := &ListResponseDataBuilder{s: s}
	return b
}

func (b *ListResponseDataBuilder) Subject(v string) *ListResponseDataBuilder {
	b.s.Subject = &v
	return b
}

func (b *ListResponseDataBuilder) CreatedTime(v string) *ListResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListResponseDataBuilder) MessageId(v string) *ListResponseDataBuilder {
	b.s.MessageId = &v
	return b
}

func (b *ListResponseDataBuilder) ReadUser(v int32) *ListResponseDataBuilder {
	b.s.ReadUser = &v
	return b
}

func (b *ListResponseDataBuilder) AllUser(v int32) *ListResponseDataBuilder {
	b.s.AllUser = &v
	return b
}

func (b *ListResponseDataBuilder) SendTime(v string) *ListResponseDataBuilder {
	b.s.SendTime = &v
	return b
}

func (b *ListResponseDataBuilder) Status(v int32) *ListResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *ListResponseDataBuilder) Build() *ListResponseData {
	return b.s
}


