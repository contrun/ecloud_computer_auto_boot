// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListBody struct {
    position.Body
    // 查看方式: 0静默消息,1强制消息
	Mode *int32 `json:"mode,omitempty"`
    // 创建开始时间
	CreatedBeginTime *string `json:"createdBeginTime,omitempty"`
    // 创建结束时间
	CreatedEndTime *string `json:"createdEndTime,omitempty"`
    // 消息主题
	Subject *string `json:"subject,omitempty"`
    // 发送开始时间
	SendBeginTime *string `json:"sendBeginTime,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 发送状态: 0:待发送,1:已发送,2:已撤回
	Status *int32 `json:"status,omitempty"`
    // 发送结束时间
	SendEndTime *string `json:"sendEndTime,omitempty"`
}

func (s ListBody) String() string {
	return utils.Beautify(s)
}

func (s ListBody) GoString() string {
	return s.String()
}

func (s ListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListBody) SetMode(v int32) *ListBody {
	s.Mode = &v
	return s
}

func (s *ListBody) SetCreatedBeginTime(v string) *ListBody {
	s.CreatedBeginTime = &v
	return s
}

func (s *ListBody) SetCreatedEndTime(v string) *ListBody {
	s.CreatedEndTime = &v
	return s
}

func (s *ListBody) SetSubject(v string) *ListBody {
	s.Subject = &v
	return s
}

func (s *ListBody) SetSendBeginTime(v string) *ListBody {
	s.SendBeginTime = &v
	return s
}

func (s *ListBody) SetPageSize(v int32) *ListBody {
	s.PageSize = &v
	return s
}

func (s *ListBody) SetCurrentPage(v int32) *ListBody {
	s.CurrentPage = &v
	return s
}

func (s *ListBody) SetStatus(v int32) *ListBody {
	s.Status = &v
	return s
}

func (s *ListBody) SetSendEndTime(v string) *ListBody {
	s.SendEndTime = &v
	return s
}


type ListBodyBuilder struct {
	s *ListBody
}

func NewListBodyBuilder() *ListBodyBuilder {
	s := &ListBody{}
	b := &ListBodyBuilder{s: s}
	return b
}

func (b *ListBodyBuilder) Mode(v int32) *ListBodyBuilder {
	b.s.Mode = &v
	return b
}

func (b *ListBodyBuilder) CreatedBeginTime(v string) *ListBodyBuilder {
	b.s.CreatedBeginTime = &v
	return b
}

func (b *ListBodyBuilder) CreatedEndTime(v string) *ListBodyBuilder {
	b.s.CreatedEndTime = &v
	return b
}

func (b *ListBodyBuilder) Subject(v string) *ListBodyBuilder {
	b.s.Subject = &v
	return b
}

func (b *ListBodyBuilder) SendBeginTime(v string) *ListBodyBuilder {
	b.s.SendBeginTime = &v
	return b
}

func (b *ListBodyBuilder) PageSize(v int32) *ListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListBodyBuilder) CurrentPage(v int32) *ListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *ListBodyBuilder) Status(v int32) *ListBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *ListBodyBuilder) SendEndTime(v string) *ListBodyBuilder {
	b.s.SendEndTime = &v
	return b
}

func (b *ListBodyBuilder) Build() *ListBody {
	return b.s
}


