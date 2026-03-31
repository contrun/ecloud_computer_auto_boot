// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDeviceDetailResponseBody struct {

    // 终端最近一次登录时间
	RecentLoginTime *string `json:"recentLoginTime,omitempty"`
    // 操作系统版本号
	OperatingVersion *string `json:"operatingVersion,omitempty"`
    // 登录用户
	RecentLoginUser *string `json:"recentLoginUser,omitempty"`
    // 系统架构
	SystemArchitecture *string `json:"systemArchitecture,omitempty"`
    // 最近连接桌面名称
	RecentLoginMachine *string `json:"recentLoginMachine,omitempty"`
    // 备注
	Remark *string `json:"remark,omitempty"`
    // 客户端版本
	ClientVersion *string `json:"clientVersion,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
    // 操作系统
	OperatingSystem *string `json:"operatingSystem,omitempty"`
    // 锁定状态 0 未锁定；1 已锁定
	LockState *string `json:"lockState,omitempty"`
    // 客户端类型
	ClientType *string `json:"clientType,omitempty"`
    // '核数
	Cores *string `json:"cores,omitempty"`
    // 绑定用户
	DeviceOwner *string `json:"deviceOwner,omitempty"`
    // 登录租户
	RecentLoginTenant *string `json:"recentLoginTenant,omitempty"`
    // 设备ID(如macId，androidId，uuid等，不可重复)
	DeviceUid *string `json:"deviceUid,omitempty"`
    // 设备表主键ID(唯一)
	Id *string `json:"id,omitempty"`
    // 内存大小
	Ram *string `json:"ram,omitempty"`
    // 设备类型
	DeviceType *string `json:"deviceType,omitempty"`
    // 连接状态 0 未连接；1 已连接
	ConnectStatus *string `json:"connectStatus,omitempty"`
    // IP地址
	IpAddress *string `json:"ipAddress,omitempty"`
    // '处理器'
	Processor *string `json:"processor,omitempty"`
    // 所属租户
	DeviceTenant *string `json:"deviceTenant,omitempty"`
    // 关联订单号
	DeviceOrderId *string `json:"deviceOrderId,omitempty"`
    // 硬盘大小
	Disk *string `json:"disk,omitempty"`
    // MAC地址
	MacAddress *string `json:"macAddress,omitempty"`
    // 终端首次登录时间
	FirstLoginTime *string `json:"firstLoginTime,omitempty"`
    // 设备来源 1.定制终端-云能，2.定制终端-终端公司，3.通用终端
	DeviceSource *string `json:"deviceSource,omitempty"`
    // 设备型号
	DeviceModel *string `json:"deviceModel,omitempty"`
    // 设备厂商
	DeviceCompany *string `json:"deviceCompany,omitempty"`
}

func (s GetDeviceDetailResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetDeviceDetailResponseBody) GoString() string {
	return s.String()
}

func (s GetDeviceDetailResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDeviceDetailResponseBody) SetRecentLoginTime(v string) *GetDeviceDetailResponseBody {
	s.RecentLoginTime = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetOperatingVersion(v string) *GetDeviceDetailResponseBody {
	s.OperatingVersion = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetRecentLoginUser(v string) *GetDeviceDetailResponseBody {
	s.RecentLoginUser = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetSystemArchitecture(v string) *GetDeviceDetailResponseBody {
	s.SystemArchitecture = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetRecentLoginMachine(v string) *GetDeviceDetailResponseBody {
	s.RecentLoginMachine = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetRemark(v string) *GetDeviceDetailResponseBody {
	s.Remark = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetClientVersion(v string) *GetDeviceDetailResponseBody {
	s.ClientVersion = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDeviceName(v string) *GetDeviceDetailResponseBody {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetOperatingSystem(v string) *GetDeviceDetailResponseBody {
	s.OperatingSystem = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetLockState(v string) *GetDeviceDetailResponseBody {
	s.LockState = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetClientType(v string) *GetDeviceDetailResponseBody {
	s.ClientType = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetCores(v string) *GetDeviceDetailResponseBody {
	s.Cores = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDeviceOwner(v string) *GetDeviceDetailResponseBody {
	s.DeviceOwner = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetRecentLoginTenant(v string) *GetDeviceDetailResponseBody {
	s.RecentLoginTenant = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDeviceUid(v string) *GetDeviceDetailResponseBody {
	s.DeviceUid = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetId(v string) *GetDeviceDetailResponseBody {
	s.Id = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetRam(v string) *GetDeviceDetailResponseBody {
	s.Ram = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDeviceType(v string) *GetDeviceDetailResponseBody {
	s.DeviceType = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetConnectStatus(v string) *GetDeviceDetailResponseBody {
	s.ConnectStatus = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetIpAddress(v string) *GetDeviceDetailResponseBody {
	s.IpAddress = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetProcessor(v string) *GetDeviceDetailResponseBody {
	s.Processor = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDeviceTenant(v string) *GetDeviceDetailResponseBody {
	s.DeviceTenant = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDeviceOrderId(v string) *GetDeviceDetailResponseBody {
	s.DeviceOrderId = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDisk(v string) *GetDeviceDetailResponseBody {
	s.Disk = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetMacAddress(v string) *GetDeviceDetailResponseBody {
	s.MacAddress = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetFirstLoginTime(v string) *GetDeviceDetailResponseBody {
	s.FirstLoginTime = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDeviceSource(v string) *GetDeviceDetailResponseBody {
	s.DeviceSource = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDeviceModel(v string) *GetDeviceDetailResponseBody {
	s.DeviceModel = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetDeviceCompany(v string) *GetDeviceDetailResponseBody {
	s.DeviceCompany = &v
	return s
}


type GetDeviceDetailResponseBodyBuilder struct {
	s *GetDeviceDetailResponseBody
}

func NewGetDeviceDetailResponseBodyBuilder() *GetDeviceDetailResponseBodyBuilder {
	s := &GetDeviceDetailResponseBody{}
	b := &GetDeviceDetailResponseBodyBuilder{s: s}
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) RecentLoginTime(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.RecentLoginTime = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) OperatingVersion(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.OperatingVersion = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) RecentLoginUser(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.RecentLoginUser = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) SystemArchitecture(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.SystemArchitecture = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) RecentLoginMachine(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.RecentLoginMachine = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) Remark(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.Remark = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) ClientVersion(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.ClientVersion = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) DeviceName(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) OperatingSystem(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.OperatingSystem = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) LockState(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.LockState = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) ClientType(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.ClientType = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) Cores(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.Cores = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) DeviceOwner(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.DeviceOwner = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) RecentLoginTenant(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.RecentLoginTenant = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) DeviceUid(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) Id(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) Ram(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.Ram = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) DeviceType(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.DeviceType = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) ConnectStatus(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.ConnectStatus = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) IpAddress(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) Processor(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.Processor = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) DeviceTenant(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.DeviceTenant = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) DeviceOrderId(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.DeviceOrderId = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) Disk(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.Disk = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) MacAddress(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.MacAddress = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) FirstLoginTime(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.FirstLoginTime = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) DeviceSource(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.DeviceSource = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) DeviceModel(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.DeviceModel = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) DeviceCompany(v string) *GetDeviceDetailResponseBodyBuilder {
	b.s.DeviceCompany = &v
	return b
}

func (b *GetDeviceDetailResponseBodyBuilder) Build() *GetDeviceDetailResponseBody {
	return b.s
}


