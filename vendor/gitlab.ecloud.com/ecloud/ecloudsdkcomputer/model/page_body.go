// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PageBody struct {
    position.Body
    // 用户名
	AccountName *string `json:"accountName,omitempty"`
    // 操作类型
	OperateType *string `json:"operateType,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s PageBody) String() string {
	return utils.Beautify(s)
}

func (s PageBody) GoString() string {
	return s.String()
}

func (s PageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PageBody) SetAccountName(v string) *PageBody {
	s.AccountName = &v
	return s
}

func (s *PageBody) SetOperateType(v string) *PageBody {
	s.OperateType = &v
	return s
}

func (s *PageBody) SetPageSize(v int32) *PageBody {
	s.PageSize = &v
	return s
}

func (s *PageBody) SetStartTime(v string) *PageBody {
	s.StartTime = &v
	return s
}

func (s *PageBody) SetEndTime(v string) *PageBody {
	s.EndTime = &v
	return s
}

func (s *PageBody) SetCurrentPage(v int32) *PageBody {
	s.CurrentPage = &v
	return s
}


type PageBodyBuilder struct {
	s *PageBody
}

func NewPageBodyBuilder() *PageBodyBuilder {
	s := &PageBody{}
	b := &PageBodyBuilder{s: s}
	return b
}

func (b *PageBodyBuilder) AccountName(v string) *PageBodyBuilder {
	b.s.AccountName = &v
	return b
}

func (b *PageBodyBuilder) OperateType(v string) *PageBodyBuilder {
	b.s.OperateType = &v
	return b
}

func (b *PageBodyBuilder) PageSize(v int32) *PageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *PageBodyBuilder) StartTime(v string) *PageBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *PageBodyBuilder) EndTime(v string) *PageBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *PageBodyBuilder) CurrentPage(v int32) *PageBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *PageBodyBuilder) Build() *PageBody {
	return b.s
}


