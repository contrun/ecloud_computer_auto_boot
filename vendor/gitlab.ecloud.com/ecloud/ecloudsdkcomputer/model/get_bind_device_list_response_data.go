// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBindDeviceListResponseData struct {

    // 登录用户
	RecentLoginUser *string `json:"recentLoginUser,omitempty"`
    // 设备来源 1.定制终端-云能，2.定制终端-终端公司，3.通用终端
	DeviceSource *string `json:"deviceSource,omitempty"`
    // 登录租户
	RecentLoginTenant *string `json:"recentLoginTenant,omitempty"`
    // 最近连接桌面名称
	RecentLoginMachine *string `json:"recentLoginMachine,omitempty"`
    // 设备型号
	DeviceModel *string `json:"deviceModel,omitempty"`
    // 设备ID(如macId，androidId，uuid等，不可重复)
	DeviceUid *string `json:"deviceUid,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
    // 锁定状态 0 未锁定；1 已锁定
	LockState *string `json:"lockState,omitempty"`
}

func (s GetBindDeviceListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetBindDeviceListResponseData) GoString() string {
	return s.String()
}

func (s GetBindDeviceListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBindDeviceListResponseData) SetRecentLoginUser(v string) *GetBindDeviceListResponseData {
	s.RecentLoginUser = &v
	return s
}

func (s *GetBindDeviceListResponseData) SetDeviceSource(v string) *GetBindDeviceListResponseData {
	s.DeviceSource = &v
	return s
}

func (s *GetBindDeviceListResponseData) SetRecentLoginTenant(v string) *GetBindDeviceListResponseData {
	s.RecentLoginTenant = &v
	return s
}

func (s *GetBindDeviceListResponseData) SetRecentLoginMachine(v string) *GetBindDeviceListResponseData {
	s.RecentLoginMachine = &v
	return s
}

func (s *GetBindDeviceListResponseData) SetDeviceModel(v string) *GetBindDeviceListResponseData {
	s.DeviceModel = &v
	return s
}

func (s *GetBindDeviceListResponseData) SetDeviceUid(v string) *GetBindDeviceListResponseData {
	s.DeviceUid = &v
	return s
}

func (s *GetBindDeviceListResponseData) SetDeviceName(v string) *GetBindDeviceListResponseData {
	s.DeviceName = &v
	return s
}

func (s *GetBindDeviceListResponseData) SetLockState(v string) *GetBindDeviceListResponseData {
	s.LockState = &v
	return s
}


type GetBindDeviceListResponseDataBuilder struct {
	s *GetBindDeviceListResponseData
}

func NewGetBindDeviceListResponseDataBuilder() *GetBindDeviceListResponseDataBuilder {
	s := &GetBindDeviceListResponseData{}
	b := &GetBindDeviceListResponseDataBuilder{s: s}
	return b
}

func (b *GetBindDeviceListResponseDataBuilder) RecentLoginUser(v string) *GetBindDeviceListResponseDataBuilder {
	b.s.RecentLoginUser = &v
	return b
}

func (b *GetBindDeviceListResponseDataBuilder) DeviceSource(v string) *GetBindDeviceListResponseDataBuilder {
	b.s.DeviceSource = &v
	return b
}

func (b *GetBindDeviceListResponseDataBuilder) RecentLoginTenant(v string) *GetBindDeviceListResponseDataBuilder {
	b.s.RecentLoginTenant = &v
	return b
}

func (b *GetBindDeviceListResponseDataBuilder) RecentLoginMachine(v string) *GetBindDeviceListResponseDataBuilder {
	b.s.RecentLoginMachine = &v
	return b
}

func (b *GetBindDeviceListResponseDataBuilder) DeviceModel(v string) *GetBindDeviceListResponseDataBuilder {
	b.s.DeviceModel = &v
	return b
}

func (b *GetBindDeviceListResponseDataBuilder) DeviceUid(v string) *GetBindDeviceListResponseDataBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *GetBindDeviceListResponseDataBuilder) DeviceName(v string) *GetBindDeviceListResponseDataBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *GetBindDeviceListResponseDataBuilder) LockState(v string) *GetBindDeviceListResponseDataBuilder {
	b.s.LockState = &v
	return b
}

func (b *GetBindDeviceListResponseDataBuilder) Build() *GetBindDeviceListResponseData {
	return b.s
}


