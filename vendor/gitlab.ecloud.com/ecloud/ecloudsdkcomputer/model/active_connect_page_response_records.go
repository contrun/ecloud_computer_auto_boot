// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ActiveConnectPageResponseRecords struct {

    // 已连接时长
	ConnectedDuration *string `json:"connectedDuration,omitempty"`
    // 连接状态
	IsConnect *string `json:"isConnect,omitempty"`
    // 客户端版本
	ClientVersion *string `json:"clientVersion,omitempty"`
    // 用户名
	UserName *string `json:"userName,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
    // 用户Id
	UserId *string `json:"userId,omitempty"`
    // 电脑名称
	MachineName *string `json:"machineName,omitempty"`
    // 断开原因
	DisConnectReason *string `json:"disConnectReason,omitempty"`
    // 客户端类型
	ClientType *string `json:"clientType,omitempty"`
    // 登录时间
	LoginTime *string `json:"loginTime,omitempty"`
    // 电脑ID
	MachineId *string `json:"machineId,omitempty"`
    // 规格Id
	ResourceTypeId *int64 `json:"resourceTypeId,omitempty"`
    // 所属渠道code
	OrderSourceCode *string `json:"orderSourceCode,omitempty"`
    // 租户Id
	TenantId *int64 `json:"tenantId,omitempty"`
    // 设备Id
	DeviceUid *string `json:"deviceUid,omitempty"`
    // 租户名称
	MopUserName *string `json:"mopUserName,omitempty"`
    // 会话开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 会话结束时间
	EndTime *string `json:"endTime,omitempty"`
}

func (s ActiveConnectPageResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s ActiveConnectPageResponseRecords) GoString() string {
	return s.String()
}

func (s ActiveConnectPageResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActiveConnectPageResponseRecords) SetConnectedDuration(v string) *ActiveConnectPageResponseRecords {
	s.ConnectedDuration = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetIsConnect(v string) *ActiveConnectPageResponseRecords {
	s.IsConnect = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetClientVersion(v string) *ActiveConnectPageResponseRecords {
	s.ClientVersion = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetUserName(v string) *ActiveConnectPageResponseRecords {
	s.UserName = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetDeviceName(v string) *ActiveConnectPageResponseRecords {
	s.DeviceName = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetUserId(v string) *ActiveConnectPageResponseRecords {
	s.UserId = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetMachineName(v string) *ActiveConnectPageResponseRecords {
	s.MachineName = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetDisConnectReason(v string) *ActiveConnectPageResponseRecords {
	s.DisConnectReason = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetClientType(v string) *ActiveConnectPageResponseRecords {
	s.ClientType = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetLoginTime(v string) *ActiveConnectPageResponseRecords {
	s.LoginTime = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetMachineId(v string) *ActiveConnectPageResponseRecords {
	s.MachineId = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetResourceTypeId(v int64) *ActiveConnectPageResponseRecords {
	s.ResourceTypeId = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetOrderSourceCode(v string) *ActiveConnectPageResponseRecords {
	s.OrderSourceCode = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetTenantId(v int64) *ActiveConnectPageResponseRecords {
	s.TenantId = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetDeviceUid(v string) *ActiveConnectPageResponseRecords {
	s.DeviceUid = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetMopUserName(v string) *ActiveConnectPageResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetStartTime(v string) *ActiveConnectPageResponseRecords {
	s.StartTime = &v
	return s
}

func (s *ActiveConnectPageResponseRecords) SetEndTime(v string) *ActiveConnectPageResponseRecords {
	s.EndTime = &v
	return s
}


type ActiveConnectPageResponseRecordsBuilder struct {
	s *ActiveConnectPageResponseRecords
}

func NewActiveConnectPageResponseRecordsBuilder() *ActiveConnectPageResponseRecordsBuilder {
	s := &ActiveConnectPageResponseRecords{}
	b := &ActiveConnectPageResponseRecordsBuilder{s: s}
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) ConnectedDuration(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.ConnectedDuration = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) IsConnect(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.IsConnect = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) ClientVersion(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.ClientVersion = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) UserName(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) DeviceName(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) UserId(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.UserId = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) MachineName(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.MachineName = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) DisConnectReason(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.DisConnectReason = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) ClientType(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.ClientType = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) LoginTime(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.LoginTime = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) MachineId(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) ResourceTypeId(v int64) *ActiveConnectPageResponseRecordsBuilder {
	b.s.ResourceTypeId = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) OrderSourceCode(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.OrderSourceCode = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) TenantId(v int64) *ActiveConnectPageResponseRecordsBuilder {
	b.s.TenantId = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) DeviceUid(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) MopUserName(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) StartTime(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.StartTime = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) EndTime(v string) *ActiveConnectPageResponseRecordsBuilder {
	b.s.EndTime = &v
	return b
}

func (b *ActiveConnectPageResponseRecordsBuilder) Build() *ActiveConnectPageResponseRecords {
	return b.s
}


