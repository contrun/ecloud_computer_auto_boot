// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ActivePageResponseRecords struct {

    // 登出时间
	LogoutTime *string `json:"logoutTime,omitempty"`
    // 客户端版本类型
	ClientType *string `json:"clientType,omitempty"`
    // 登录时间
	LoginTime *string `json:"loginTime,omitempty"`
    // 已登录时长
	LoginDuration *string `json:"loginDuration,omitempty"`
    // 租户名称
	MopUserName *string `json:"mopUserName,omitempty"`
    // 客户端版本号
	ClientVersion *string `json:"clientVersion,omitempty"`
    // 用户名
	UserName *string `json:"userName,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
}

func (s ActivePageResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s ActivePageResponseRecords) GoString() string {
	return s.String()
}

func (s ActivePageResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActivePageResponseRecords) SetLogoutTime(v string) *ActivePageResponseRecords {
	s.LogoutTime = &v
	return s
}

func (s *ActivePageResponseRecords) SetClientType(v string) *ActivePageResponseRecords {
	s.ClientType = &v
	return s
}

func (s *ActivePageResponseRecords) SetLoginTime(v string) *ActivePageResponseRecords {
	s.LoginTime = &v
	return s
}

func (s *ActivePageResponseRecords) SetLoginDuration(v string) *ActivePageResponseRecords {
	s.LoginDuration = &v
	return s
}

func (s *ActivePageResponseRecords) SetMopUserName(v string) *ActivePageResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *ActivePageResponseRecords) SetClientVersion(v string) *ActivePageResponseRecords {
	s.ClientVersion = &v
	return s
}

func (s *ActivePageResponseRecords) SetUserName(v string) *ActivePageResponseRecords {
	s.UserName = &v
	return s
}

func (s *ActivePageResponseRecords) SetDeviceName(v string) *ActivePageResponseRecords {
	s.DeviceName = &v
	return s
}


type ActivePageResponseRecordsBuilder struct {
	s *ActivePageResponseRecords
}

func NewActivePageResponseRecordsBuilder() *ActivePageResponseRecordsBuilder {
	s := &ActivePageResponseRecords{}
	b := &ActivePageResponseRecordsBuilder{s: s}
	return b
}

func (b *ActivePageResponseRecordsBuilder) LogoutTime(v string) *ActivePageResponseRecordsBuilder {
	b.s.LogoutTime = &v
	return b
}

func (b *ActivePageResponseRecordsBuilder) ClientType(v string) *ActivePageResponseRecordsBuilder {
	b.s.ClientType = &v
	return b
}

func (b *ActivePageResponseRecordsBuilder) LoginTime(v string) *ActivePageResponseRecordsBuilder {
	b.s.LoginTime = &v
	return b
}

func (b *ActivePageResponseRecordsBuilder) LoginDuration(v string) *ActivePageResponseRecordsBuilder {
	b.s.LoginDuration = &v
	return b
}

func (b *ActivePageResponseRecordsBuilder) MopUserName(v string) *ActivePageResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *ActivePageResponseRecordsBuilder) ClientVersion(v string) *ActivePageResponseRecordsBuilder {
	b.s.ClientVersion = &v
	return b
}

func (b *ActivePageResponseRecordsBuilder) UserName(v string) *ActivePageResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *ActivePageResponseRecordsBuilder) DeviceName(v string) *ActivePageResponseRecordsBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *ActivePageResponseRecordsBuilder) Build() *ActivePageResponseRecords {
	return b.s
}


