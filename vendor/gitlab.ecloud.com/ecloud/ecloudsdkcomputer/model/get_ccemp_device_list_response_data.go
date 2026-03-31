// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetCcempDeviceListResponseData struct {

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
    // 所属租户名称
	DeviceTenantName *string `json:"deviceTenantName,omitempty"`
    // 客户端类型
	ClientType *string `json:"clientType,omitempty"`
    // '核数
	Cores *string `json:"cores,omitempty"`
    // 绑定用户数量
	BindCount *int32 `json:"bindCount,omitempty"`
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
    // 登录租户名称
	RecentLoginTenantName *string `json:"recentLoginTenantName,omitempty"`
    // 设备厂商
	DeviceCompany *string `json:"deviceCompany,omitempty"`
}

func (s GetCcempDeviceListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetCcempDeviceListResponseData) GoString() string {
	return s.String()
}

func (s GetCcempDeviceListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCcempDeviceListResponseData) SetRecentLoginTime(v string) *GetCcempDeviceListResponseData {
	s.RecentLoginTime = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetOperatingVersion(v string) *GetCcempDeviceListResponseData {
	s.OperatingVersion = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetRecentLoginUser(v string) *GetCcempDeviceListResponseData {
	s.RecentLoginUser = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetSystemArchitecture(v string) *GetCcempDeviceListResponseData {
	s.SystemArchitecture = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetRecentLoginMachine(v string) *GetCcempDeviceListResponseData {
	s.RecentLoginMachine = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetRemark(v string) *GetCcempDeviceListResponseData {
	s.Remark = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetClientVersion(v string) *GetCcempDeviceListResponseData {
	s.ClientVersion = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDeviceName(v string) *GetCcempDeviceListResponseData {
	s.DeviceName = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetOperatingSystem(v string) *GetCcempDeviceListResponseData {
	s.OperatingSystem = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetLockState(v string) *GetCcempDeviceListResponseData {
	s.LockState = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDeviceTenantName(v string) *GetCcempDeviceListResponseData {
	s.DeviceTenantName = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetClientType(v string) *GetCcempDeviceListResponseData {
	s.ClientType = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetCores(v string) *GetCcempDeviceListResponseData {
	s.Cores = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetBindCount(v int32) *GetCcempDeviceListResponseData {
	s.BindCount = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetRecentLoginTenant(v string) *GetCcempDeviceListResponseData {
	s.RecentLoginTenant = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDeviceUid(v string) *GetCcempDeviceListResponseData {
	s.DeviceUid = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetId(v string) *GetCcempDeviceListResponseData {
	s.Id = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetRam(v string) *GetCcempDeviceListResponseData {
	s.Ram = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDeviceType(v string) *GetCcempDeviceListResponseData {
	s.DeviceType = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetConnectStatus(v string) *GetCcempDeviceListResponseData {
	s.ConnectStatus = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetIpAddress(v string) *GetCcempDeviceListResponseData {
	s.IpAddress = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetProcessor(v string) *GetCcempDeviceListResponseData {
	s.Processor = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDeviceTenant(v string) *GetCcempDeviceListResponseData {
	s.DeviceTenant = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDeviceOrderId(v string) *GetCcempDeviceListResponseData {
	s.DeviceOrderId = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDisk(v string) *GetCcempDeviceListResponseData {
	s.Disk = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetMacAddress(v string) *GetCcempDeviceListResponseData {
	s.MacAddress = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetFirstLoginTime(v string) *GetCcempDeviceListResponseData {
	s.FirstLoginTime = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDeviceSource(v string) *GetCcempDeviceListResponseData {
	s.DeviceSource = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDeviceModel(v string) *GetCcempDeviceListResponseData {
	s.DeviceModel = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetRecentLoginTenantName(v string) *GetCcempDeviceListResponseData {
	s.RecentLoginTenantName = &v
	return s
}

func (s *GetCcempDeviceListResponseData) SetDeviceCompany(v string) *GetCcempDeviceListResponseData {
	s.DeviceCompany = &v
	return s
}


type GetCcempDeviceListResponseDataBuilder struct {
	s *GetCcempDeviceListResponseData
}

func NewGetCcempDeviceListResponseDataBuilder() *GetCcempDeviceListResponseDataBuilder {
	s := &GetCcempDeviceListResponseData{}
	b := &GetCcempDeviceListResponseDataBuilder{s: s}
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) RecentLoginTime(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.RecentLoginTime = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) OperatingVersion(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.OperatingVersion = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) RecentLoginUser(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.RecentLoginUser = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) SystemArchitecture(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.SystemArchitecture = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) RecentLoginMachine(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.RecentLoginMachine = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) Remark(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.Remark = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) ClientVersion(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.ClientVersion = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) DeviceName(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) OperatingSystem(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.OperatingSystem = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) LockState(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.LockState = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) DeviceTenantName(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.DeviceTenantName = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) ClientType(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.ClientType = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) Cores(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.Cores = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) BindCount(v int32) *GetCcempDeviceListResponseDataBuilder {
	b.s.BindCount = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) RecentLoginTenant(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.RecentLoginTenant = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) DeviceUid(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) Id(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.Id = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) Ram(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.Ram = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) DeviceType(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.DeviceType = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) ConnectStatus(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.ConnectStatus = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) IpAddress(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) Processor(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.Processor = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) DeviceTenant(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.DeviceTenant = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) DeviceOrderId(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.DeviceOrderId = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) Disk(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.Disk = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) MacAddress(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.MacAddress = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) FirstLoginTime(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.FirstLoginTime = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) DeviceSource(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.DeviceSource = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) DeviceModel(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.DeviceModel = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) RecentLoginTenantName(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.RecentLoginTenantName = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) DeviceCompany(v string) *GetCcempDeviceListResponseDataBuilder {
	b.s.DeviceCompany = &v
	return b
}

func (b *GetCcempDeviceListResponseDataBuilder) Build() *GetCcempDeviceListResponseData {
	return b.s
}


