// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ActivePageBody struct {
    position.Body
    // 分页大小
	PageSize *string `json:"pageSize,omitempty"`
    // 当前页
	Page *string `json:"page,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
}

func (s ActivePageBody) String() string {
	return utils.Beautify(s)
}

func (s ActivePageBody) GoString() string {
	return s.String()
}

func (s ActivePageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActivePageBody) SetPageSize(v string) *ActivePageBody {
	s.PageSize = &v
	return s
}

func (s *ActivePageBody) SetPage(v string) *ActivePageBody {
	s.Page = &v
	return s
}

func (s *ActivePageBody) SetUserName(v string) *ActivePageBody {
	s.UserName = &v
	return s
}

func (s *ActivePageBody) SetDeviceName(v string) *ActivePageBody {
	s.DeviceName = &v
	return s
}


type ActivePageBodyBuilder struct {
	s *ActivePageBody
}

func NewActivePageBodyBuilder() *ActivePageBodyBuilder {
	s := &ActivePageBody{}
	b := &ActivePageBodyBuilder{s: s}
	return b
}

func (b *ActivePageBodyBuilder) PageSize(v string) *ActivePageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ActivePageBodyBuilder) Page(v string) *ActivePageBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *ActivePageBodyBuilder) UserName(v string) *ActivePageBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *ActivePageBodyBuilder) DeviceName(v string) *ActivePageBodyBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *ActivePageBodyBuilder) Build() *ActivePageBody {
	return b.s
}


