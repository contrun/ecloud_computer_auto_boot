// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type HistoryPageResponseRecords struct {

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

func (s HistoryPageResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s HistoryPageResponseRecords) GoString() string {
	return s.String()
}

func (s HistoryPageResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryPageResponseRecords) SetLogoutTime(v string) *HistoryPageResponseRecords {
	s.LogoutTime = &v
	return s
}

func (s *HistoryPageResponseRecords) SetClientType(v string) *HistoryPageResponseRecords {
	s.ClientType = &v
	return s
}

func (s *HistoryPageResponseRecords) SetLoginTime(v string) *HistoryPageResponseRecords {
	s.LoginTime = &v
	return s
}

func (s *HistoryPageResponseRecords) SetLoginDuration(v string) *HistoryPageResponseRecords {
	s.LoginDuration = &v
	return s
}

func (s *HistoryPageResponseRecords) SetMopUserName(v string) *HistoryPageResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *HistoryPageResponseRecords) SetClientVersion(v string) *HistoryPageResponseRecords {
	s.ClientVersion = &v
	return s
}

func (s *HistoryPageResponseRecords) SetUserName(v string) *HistoryPageResponseRecords {
	s.UserName = &v
	return s
}

func (s *HistoryPageResponseRecords) SetDeviceName(v string) *HistoryPageResponseRecords {
	s.DeviceName = &v
	return s
}


type HistoryPageResponseRecordsBuilder struct {
	s *HistoryPageResponseRecords
}

func NewHistoryPageResponseRecordsBuilder() *HistoryPageResponseRecordsBuilder {
	s := &HistoryPageResponseRecords{}
	b := &HistoryPageResponseRecordsBuilder{s: s}
	return b
}

func (b *HistoryPageResponseRecordsBuilder) LogoutTime(v string) *HistoryPageResponseRecordsBuilder {
	b.s.LogoutTime = &v
	return b
}

func (b *HistoryPageResponseRecordsBuilder) ClientType(v string) *HistoryPageResponseRecordsBuilder {
	b.s.ClientType = &v
	return b
}

func (b *HistoryPageResponseRecordsBuilder) LoginTime(v string) *HistoryPageResponseRecordsBuilder {
	b.s.LoginTime = &v
	return b
}

func (b *HistoryPageResponseRecordsBuilder) LoginDuration(v string) *HistoryPageResponseRecordsBuilder {
	b.s.LoginDuration = &v
	return b
}

func (b *HistoryPageResponseRecordsBuilder) MopUserName(v string) *HistoryPageResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *HistoryPageResponseRecordsBuilder) ClientVersion(v string) *HistoryPageResponseRecordsBuilder {
	b.s.ClientVersion = &v
	return b
}

func (b *HistoryPageResponseRecordsBuilder) UserName(v string) *HistoryPageResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *HistoryPageResponseRecordsBuilder) DeviceName(v string) *HistoryPageResponseRecordsBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *HistoryPageResponseRecordsBuilder) Build() *HistoryPageResponseRecords {
	return b.s
}


