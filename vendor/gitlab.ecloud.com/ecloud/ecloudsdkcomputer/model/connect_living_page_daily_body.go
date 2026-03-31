// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectLivingPageDailyBody struct {
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
}

func (s ConnectLivingPageDailyBody) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageDailyBody) GoString() string {
	return s.String()
}

func (s ConnectLivingPageDailyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageDailyBody) SetOrgIdList(v []string) *ConnectLivingPageDailyBody {
	s.OrgIdList = v
	return s
}

func (s *ConnectLivingPageDailyBody) SetPageSize(v int32) *ConnectLivingPageDailyBody {
	s.PageSize = &v
	return s
}

func (s *ConnectLivingPageDailyBody) SetStartTime(v string) *ConnectLivingPageDailyBody {
	s.StartTime = &v
	return s
}

func (s *ConnectLivingPageDailyBody) SetEndTime(v string) *ConnectLivingPageDailyBody {
	s.EndTime = &v
	return s
}

func (s *ConnectLivingPageDailyBody) SetPage(v int32) *ConnectLivingPageDailyBody {
	s.Page = &v
	return s
}


type ConnectLivingPageDailyBodyBuilder struct {
	s *ConnectLivingPageDailyBody
}

func NewConnectLivingPageDailyBodyBuilder() *ConnectLivingPageDailyBodyBuilder {
	s := &ConnectLivingPageDailyBody{}
	b := &ConnectLivingPageDailyBodyBuilder{s: s}
	return b
}

func (b *ConnectLivingPageDailyBodyBuilder) OrgIdList(v []string) *ConnectLivingPageDailyBodyBuilder {
	b.s.OrgIdList = v
	return b
}

func (b *ConnectLivingPageDailyBodyBuilder) PageSize(v int32) *ConnectLivingPageDailyBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ConnectLivingPageDailyBodyBuilder) StartTime(v string) *ConnectLivingPageDailyBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *ConnectLivingPageDailyBodyBuilder) EndTime(v string) *ConnectLivingPageDailyBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *ConnectLivingPageDailyBodyBuilder) Page(v int32) *ConnectLivingPageDailyBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *ConnectLivingPageDailyBodyBuilder) Build() *ConnectLivingPageDailyBody {
	return b.s
}


