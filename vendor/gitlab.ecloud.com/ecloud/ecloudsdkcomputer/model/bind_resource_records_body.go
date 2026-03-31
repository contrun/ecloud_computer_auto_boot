// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindResourceRecordsBody struct {
    position.Body
    // 订购资源实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 登录账号，支持模糊搜索
	UserName *string `json:"userName,omitempty"`
    // 管理员账号/分配操作人,支持模糊搜索
	Operator *string `json:"operator,omitempty"`
}

func (s BindResourceRecordsBody) String() string {
	return utils.Beautify(s)
}

func (s BindResourceRecordsBody) GoString() string {
	return s.String()
}

func (s BindResourceRecordsBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindResourceRecordsBody) SetInstanceId(v string) *BindResourceRecordsBody {
	s.InstanceId = &v
	return s
}

func (s *BindResourceRecordsBody) SetPageSize(v int32) *BindResourceRecordsBody {
	s.PageSize = &v
	return s
}

func (s *BindResourceRecordsBody) SetStartTime(v string) *BindResourceRecordsBody {
	s.StartTime = &v
	return s
}

func (s *BindResourceRecordsBody) SetEndTime(v string) *BindResourceRecordsBody {
	s.EndTime = &v
	return s
}

func (s *BindResourceRecordsBody) SetCurrentPage(v int32) *BindResourceRecordsBody {
	s.CurrentPage = &v
	return s
}

func (s *BindResourceRecordsBody) SetUserName(v string) *BindResourceRecordsBody {
	s.UserName = &v
	return s
}

func (s *BindResourceRecordsBody) SetOperator(v string) *BindResourceRecordsBody {
	s.Operator = &v
	return s
}


type BindResourceRecordsBodyBuilder struct {
	s *BindResourceRecordsBody
}

func NewBindResourceRecordsBodyBuilder() *BindResourceRecordsBodyBuilder {
	s := &BindResourceRecordsBody{}
	b := &BindResourceRecordsBodyBuilder{s: s}
	return b
}

func (b *BindResourceRecordsBodyBuilder) InstanceId(v string) *BindResourceRecordsBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *BindResourceRecordsBodyBuilder) PageSize(v int32) *BindResourceRecordsBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *BindResourceRecordsBodyBuilder) StartTime(v string) *BindResourceRecordsBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *BindResourceRecordsBodyBuilder) EndTime(v string) *BindResourceRecordsBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *BindResourceRecordsBodyBuilder) CurrentPage(v int32) *BindResourceRecordsBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *BindResourceRecordsBodyBuilder) UserName(v string) *BindResourceRecordsBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *BindResourceRecordsBodyBuilder) Operator(v string) *BindResourceRecordsBodyBuilder {
	b.s.Operator = &v
	return b
}

func (b *BindResourceRecordsBodyBuilder) Build() *BindResourceRecordsBody {
	return b.s
}


