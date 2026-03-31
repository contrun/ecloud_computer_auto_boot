// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBindDeviceListBody struct {
    position.Body
    // 设备来源 1.定制终端-云能，2.定制终端-终端公司，3.通用终端
	DeviceSource *string `json:"deviceSource,omitempty"`
    // 最近连接电脑名称
	RecentLoginMachine *string `json:"recentLoginMachine,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 设备型号
	DeviceModel *string `json:"deviceModel,omitempty"`
    // 设备id
	DeviceUid *string `json:"deviceUid,omitempty"`
    // 登录用户
	UserName *string `json:"userName,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 用户id
	UserId *string `json:"userId,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
    // 锁定状态 0 未锁定(默认)；1 已锁定
	LockState *int32 `json:"lockState,omitempty"`
}

func (s GetBindDeviceListBody) String() string {
	return utils.Beautify(s)
}

func (s GetBindDeviceListBody) GoString() string {
	return s.String()
}

func (s GetBindDeviceListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBindDeviceListBody) SetDeviceSource(v string) *GetBindDeviceListBody {
	s.DeviceSource = &v
	return s
}

func (s *GetBindDeviceListBody) SetRecentLoginMachine(v string) *GetBindDeviceListBody {
	s.RecentLoginMachine = &v
	return s
}

func (s *GetBindDeviceListBody) SetPageSize(v int32) *GetBindDeviceListBody {
	s.PageSize = &v
	return s
}

func (s *GetBindDeviceListBody) SetDeviceModel(v string) *GetBindDeviceListBody {
	s.DeviceModel = &v
	return s
}

func (s *GetBindDeviceListBody) SetDeviceUid(v string) *GetBindDeviceListBody {
	s.DeviceUid = &v
	return s
}

func (s *GetBindDeviceListBody) SetUserName(v string) *GetBindDeviceListBody {
	s.UserName = &v
	return s
}

func (s *GetBindDeviceListBody) SetCurrentPage(v int32) *GetBindDeviceListBody {
	s.CurrentPage = &v
	return s
}

func (s *GetBindDeviceListBody) SetUserId(v string) *GetBindDeviceListBody {
	s.UserId = &v
	return s
}

func (s *GetBindDeviceListBody) SetDeviceName(v string) *GetBindDeviceListBody {
	s.DeviceName = &v
	return s
}

func (s *GetBindDeviceListBody) SetLockState(v int32) *GetBindDeviceListBody {
	s.LockState = &v
	return s
}


type GetBindDeviceListBodyBuilder struct {
	s *GetBindDeviceListBody
}

func NewGetBindDeviceListBodyBuilder() *GetBindDeviceListBodyBuilder {
	s := &GetBindDeviceListBody{}
	b := &GetBindDeviceListBodyBuilder{s: s}
	return b
}

func (b *GetBindDeviceListBodyBuilder) DeviceSource(v string) *GetBindDeviceListBodyBuilder {
	b.s.DeviceSource = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) RecentLoginMachine(v string) *GetBindDeviceListBodyBuilder {
	b.s.RecentLoginMachine = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) PageSize(v int32) *GetBindDeviceListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) DeviceModel(v string) *GetBindDeviceListBodyBuilder {
	b.s.DeviceModel = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) DeviceUid(v string) *GetBindDeviceListBodyBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) UserName(v string) *GetBindDeviceListBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) CurrentPage(v int32) *GetBindDeviceListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) UserId(v string) *GetBindDeviceListBodyBuilder {
	b.s.UserId = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) DeviceName(v string) *GetBindDeviceListBodyBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) LockState(v int32) *GetBindDeviceListBodyBuilder {
	b.s.LockState = &v
	return b
}

func (b *GetBindDeviceListBodyBuilder) Build() *GetBindDeviceListBody {
	return b.s
}


