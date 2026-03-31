// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type HistoryConnectPageResponseRecords struct {

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

func (s HistoryConnectPageResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s HistoryConnectPageResponseRecords) GoString() string {
	return s.String()
}

func (s HistoryConnectPageResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryConnectPageResponseRecords) SetConnectedDuration(v string) *HistoryConnectPageResponseRecords {
	s.ConnectedDuration = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetIsConnect(v string) *HistoryConnectPageResponseRecords {
	s.IsConnect = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetClientVersion(v string) *HistoryConnectPageResponseRecords {
	s.ClientVersion = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetUserName(v string) *HistoryConnectPageResponseRecords {
	s.UserName = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetDeviceName(v string) *HistoryConnectPageResponseRecords {
	s.DeviceName = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetUserId(v string) *HistoryConnectPageResponseRecords {
	s.UserId = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetMachineName(v string) *HistoryConnectPageResponseRecords {
	s.MachineName = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetDisConnectReason(v string) *HistoryConnectPageResponseRecords {
	s.DisConnectReason = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetClientType(v string) *HistoryConnectPageResponseRecords {
	s.ClientType = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetLoginTime(v string) *HistoryConnectPageResponseRecords {
	s.LoginTime = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetMachineId(v string) *HistoryConnectPageResponseRecords {
	s.MachineId = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetResourceTypeId(v int64) *HistoryConnectPageResponseRecords {
	s.ResourceTypeId = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetOrderSourceCode(v string) *HistoryConnectPageResponseRecords {
	s.OrderSourceCode = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetTenantId(v int64) *HistoryConnectPageResponseRecords {
	s.TenantId = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetDeviceUid(v string) *HistoryConnectPageResponseRecords {
	s.DeviceUid = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetMopUserName(v string) *HistoryConnectPageResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetStartTime(v string) *HistoryConnectPageResponseRecords {
	s.StartTime = &v
	return s
}

func (s *HistoryConnectPageResponseRecords) SetEndTime(v string) *HistoryConnectPageResponseRecords {
	s.EndTime = &v
	return s
}


type HistoryConnectPageResponseRecordsBuilder struct {
	s *HistoryConnectPageResponseRecords
}

func NewHistoryConnectPageResponseRecordsBuilder() *HistoryConnectPageResponseRecordsBuilder {
	s := &HistoryConnectPageResponseRecords{}
	b := &HistoryConnectPageResponseRecordsBuilder{s: s}
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) ConnectedDuration(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.ConnectedDuration = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) IsConnect(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.IsConnect = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) ClientVersion(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.ClientVersion = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) UserName(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) DeviceName(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) UserId(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.UserId = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) MachineName(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.MachineName = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) DisConnectReason(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.DisConnectReason = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) ClientType(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.ClientType = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) LoginTime(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.LoginTime = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) MachineId(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.MachineId = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) ResourceTypeId(v int64) *HistoryConnectPageResponseRecordsBuilder {
	b.s.ResourceTypeId = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) OrderSourceCode(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.OrderSourceCode = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) TenantId(v int64) *HistoryConnectPageResponseRecordsBuilder {
	b.s.TenantId = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) DeviceUid(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) MopUserName(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) StartTime(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.StartTime = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) EndTime(v string) *HistoryConnectPageResponseRecordsBuilder {
	b.s.EndTime = &v
	return b
}

func (b *HistoryConnectPageResponseRecordsBuilder) Build() *HistoryConnectPageResponseRecords {
	return b.s
}


