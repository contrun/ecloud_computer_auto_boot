// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginUsageStatisticsPageBody struct {
    position.Body
    // 组织id列表
	OrgIdList []string `json:"orgIdList,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
    // 当前页
	Page *int32 `json:"page,omitempty"`
    // 用户名
	UserName *string `json:"userName,omitempty"`
}

func (s LoginUsageStatisticsPageBody) String() string {
	return utils.Beautify(s)
}

func (s LoginUsageStatisticsPageBody) GoString() string {
	return s.String()
}

func (s LoginUsageStatisticsPageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginUsageStatisticsPageBody) SetOrgIdList(v []string) *LoginUsageStatisticsPageBody {
	s.OrgIdList = v
	return s
}

func (s *LoginUsageStatisticsPageBody) SetPageSize(v int32) *LoginUsageStatisticsPageBody {
	s.PageSize = &v
	return s
}

func (s *LoginUsageStatisticsPageBody) SetStartTime(v string) *LoginUsageStatisticsPageBody {
	s.StartTime = &v
	return s
}

func (s *LoginUsageStatisticsPageBody) SetEndTime(v string) *LoginUsageStatisticsPageBody {
	s.EndTime = &v
	return s
}

func (s *LoginUsageStatisticsPageBody) SetPage(v int32) *LoginUsageStatisticsPageBody {
	s.Page = &v
	return s
}

func (s *LoginUsageStatisticsPageBody) SetUserName(v string) *LoginUsageStatisticsPageBody {
	s.UserName = &v
	return s
}


type LoginUsageStatisticsPageBodyBuilder struct {
	s *LoginUsageStatisticsPageBody
}

func NewLoginUsageStatisticsPageBodyBuilder() *LoginUsageStatisticsPageBodyBuilder {
	s := &LoginUsageStatisticsPageBody{}
	b := &LoginUsageStatisticsPageBodyBuilder{s: s}
	return b
}

func (b *LoginUsageStatisticsPageBodyBuilder) OrgIdList(v []string) *LoginUsageStatisticsPageBodyBuilder {
	b.s.OrgIdList = v
	return b
}

func (b *LoginUsageStatisticsPageBodyBuilder) PageSize(v int32) *LoginUsageStatisticsPageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *LoginUsageStatisticsPageBodyBuilder) StartTime(v string) *LoginUsageStatisticsPageBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *LoginUsageStatisticsPageBodyBuilder) EndTime(v string) *LoginUsageStatisticsPageBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *LoginUsageStatisticsPageBodyBuilder) Page(v int32) *LoginUsageStatisticsPageBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *LoginUsageStatisticsPageBodyBuilder) UserName(v string) *LoginUsageStatisticsPageBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *LoginUsageStatisticsPageBodyBuilder) Build() *LoginUsageStatisticsPageBody {
	return b.s
}


