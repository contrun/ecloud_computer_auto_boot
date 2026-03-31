// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectLivingPageMonthlyBody struct {
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

func (s ConnectLivingPageMonthlyBody) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageMonthlyBody) GoString() string {
	return s.String()
}

func (s ConnectLivingPageMonthlyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageMonthlyBody) SetOrgIdList(v []string) *ConnectLivingPageMonthlyBody {
	s.OrgIdList = v
	return s
}

func (s *ConnectLivingPageMonthlyBody) SetPageSize(v int32) *ConnectLivingPageMonthlyBody {
	s.PageSize = &v
	return s
}

func (s *ConnectLivingPageMonthlyBody) SetStartTime(v string) *ConnectLivingPageMonthlyBody {
	s.StartTime = &v
	return s
}

func (s *ConnectLivingPageMonthlyBody) SetEndTime(v string) *ConnectLivingPageMonthlyBody {
	s.EndTime = &v
	return s
}

func (s *ConnectLivingPageMonthlyBody) SetPage(v int32) *ConnectLivingPageMonthlyBody {
	s.Page = &v
	return s
}


type ConnectLivingPageMonthlyBodyBuilder struct {
	s *ConnectLivingPageMonthlyBody
}

func NewConnectLivingPageMonthlyBodyBuilder() *ConnectLivingPageMonthlyBodyBuilder {
	s := &ConnectLivingPageMonthlyBody{}
	b := &ConnectLivingPageMonthlyBodyBuilder{s: s}
	return b
}

func (b *ConnectLivingPageMonthlyBodyBuilder) OrgIdList(v []string) *ConnectLivingPageMonthlyBodyBuilder {
	b.s.OrgIdList = v
	return b
}

func (b *ConnectLivingPageMonthlyBodyBuilder) PageSize(v int32) *ConnectLivingPageMonthlyBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ConnectLivingPageMonthlyBodyBuilder) StartTime(v string) *ConnectLivingPageMonthlyBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *ConnectLivingPageMonthlyBodyBuilder) EndTime(v string) *ConnectLivingPageMonthlyBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *ConnectLivingPageMonthlyBodyBuilder) Page(v int32) *ConnectLivingPageMonthlyBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *ConnectLivingPageMonthlyBodyBuilder) Build() *ConnectLivingPageMonthlyBody {
	return b.s
}


