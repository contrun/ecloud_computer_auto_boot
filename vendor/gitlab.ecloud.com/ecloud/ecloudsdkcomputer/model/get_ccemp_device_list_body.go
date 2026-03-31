// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetCcempDeviceListBody struct {
    position.Body
    // 设备类型
	DeviceType *string `json:"deviceType,omitempty"`
    // 连接状态 0 未连接；1 已连接
	ConnectStatus *string `json:"connectStatus,omitempty"`
    // 当前页起始下标
	Start *int32 `json:"start,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 用户名
	UserName *string `json:"userName,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
    // 登录租户
	LoginTenant *string `json:"loginTenant,omitempty"`
    // 所属租户
	DeviceTenant *string `json:"deviceTenant,omitempty"`
    // 锁定状态 0 未锁定(默认)；1 已锁定
	LockState *int32 `json:"lockState,omitempty"`
    // 桌面名称
	MachineName *string `json:"machineName,omitempty"`
    // 所属用户
	DeviceOwner *string `json:"deviceOwner,omitempty"`
    // 设备来源 1.定制终端-云能，2.定制终端-终端公司，3.通用终端
	DeviceSource *string `json:"deviceSource,omitempty"`
    // 租户主键id
	TenantId *string `json:"tenantId,omitempty"`
    // 设备型号
	DeviceModel *string `json:"deviceModel,omitempty"`
    // 设备id
	DeviceUid *string `json:"deviceUid,omitempty"`
    // 设备厂商
	DeviceCompany *string `json:"deviceCompany,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s GetCcempDeviceListBody) String() string {
	return utils.Beautify(s)
}

func (s GetCcempDeviceListBody) GoString() string {
	return s.String()
}

func (s GetCcempDeviceListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCcempDeviceListBody) SetDeviceType(v string) *GetCcempDeviceListBody {
	s.DeviceType = &v
	return s
}

func (s *GetCcempDeviceListBody) SetConnectStatus(v string) *GetCcempDeviceListBody {
	s.ConnectStatus = &v
	return s
}

func (s *GetCcempDeviceListBody) SetStart(v int32) *GetCcempDeviceListBody {
	s.Start = &v
	return s
}

func (s *GetCcempDeviceListBody) SetPageSize(v int32) *GetCcempDeviceListBody {
	s.PageSize = &v
	return s
}

func (s *GetCcempDeviceListBody) SetUserName(v string) *GetCcempDeviceListBody {
	s.UserName = &v
	return s
}

func (s *GetCcempDeviceListBody) SetDeviceName(v string) *GetCcempDeviceListBody {
	s.DeviceName = &v
	return s
}

func (s *GetCcempDeviceListBody) SetLoginTenant(v string) *GetCcempDeviceListBody {
	s.LoginTenant = &v
	return s
}

func (s *GetCcempDeviceListBody) SetDeviceTenant(v string) *GetCcempDeviceListBody {
	s.DeviceTenant = &v
	return s
}

func (s *GetCcempDeviceListBody) SetLockState(v int32) *GetCcempDeviceListBody {
	s.LockState = &v
	return s
}

func (s *GetCcempDeviceListBody) SetMachineName(v string) *GetCcempDeviceListBody {
	s.MachineName = &v
	return s
}

func (s *GetCcempDeviceListBody) SetDeviceOwner(v string) *GetCcempDeviceListBody {
	s.DeviceOwner = &v
	return s
}

func (s *GetCcempDeviceListBody) SetDeviceSource(v string) *GetCcempDeviceListBody {
	s.DeviceSource = &v
	return s
}

func (s *GetCcempDeviceListBody) SetTenantId(v string) *GetCcempDeviceListBody {
	s.TenantId = &v
	return s
}

func (s *GetCcempDeviceListBody) SetDeviceModel(v string) *GetCcempDeviceListBody {
	s.DeviceModel = &v
	return s
}

func (s *GetCcempDeviceListBody) SetDeviceUid(v string) *GetCcempDeviceListBody {
	s.DeviceUid = &v
	return s
}

func (s *GetCcempDeviceListBody) SetDeviceCompany(v string) *GetCcempDeviceListBody {
	s.DeviceCompany = &v
	return s
}

func (s *GetCcempDeviceListBody) SetCurrentPage(v int32) *GetCcempDeviceListBody {
	s.CurrentPage = &v
	return s
}


type GetCcempDeviceListBodyBuilder struct {
	s *GetCcempDeviceListBody
}

func NewGetCcempDeviceListBodyBuilder() *GetCcempDeviceListBodyBuilder {
	s := &GetCcempDeviceListBody{}
	b := &GetCcempDeviceListBodyBuilder{s: s}
	return b
}

func (b *GetCcempDeviceListBodyBuilder) DeviceType(v string) *GetCcempDeviceListBodyBuilder {
	b.s.DeviceType = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) ConnectStatus(v string) *GetCcempDeviceListBodyBuilder {
	b.s.ConnectStatus = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) Start(v int32) *GetCcempDeviceListBodyBuilder {
	b.s.Start = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) PageSize(v int32) *GetCcempDeviceListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) UserName(v string) *GetCcempDeviceListBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) DeviceName(v string) *GetCcempDeviceListBodyBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) LoginTenant(v string) *GetCcempDeviceListBodyBuilder {
	b.s.LoginTenant = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) DeviceTenant(v string) *GetCcempDeviceListBodyBuilder {
	b.s.DeviceTenant = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) LockState(v int32) *GetCcempDeviceListBodyBuilder {
	b.s.LockState = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) MachineName(v string) *GetCcempDeviceListBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) DeviceOwner(v string) *GetCcempDeviceListBodyBuilder {
	b.s.DeviceOwner = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) DeviceSource(v string) *GetCcempDeviceListBodyBuilder {
	b.s.DeviceSource = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) TenantId(v string) *GetCcempDeviceListBodyBuilder {
	b.s.TenantId = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) DeviceModel(v string) *GetCcempDeviceListBodyBuilder {
	b.s.DeviceModel = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) DeviceUid(v string) *GetCcempDeviceListBodyBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) DeviceCompany(v string) *GetCcempDeviceListBodyBuilder {
	b.s.DeviceCompany = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) CurrentPage(v int32) *GetCcempDeviceListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetCcempDeviceListBodyBuilder) Build() *GetCcempDeviceListBody {
	return b.s
}


