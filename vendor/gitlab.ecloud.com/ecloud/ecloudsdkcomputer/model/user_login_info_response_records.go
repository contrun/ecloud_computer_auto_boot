// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UserLoginInfoResponseRecords struct {

    // 设备ip地址
	IpAddress *string `json:"ipAddress,omitempty"`
    // 用户名
	UserName *string `json:"userName,omitempty"`
    // 客户端版本号
	ClientVersion *string `json:"clientVersion,omitempty"`
    // 用户Id
	UserId *string `json:"userId,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
    // 客户端ip
	IpAddr *string `json:"ipAddr,omitempty"`
    // 操作时间
	OperationTime *int64 `json:"operationTime,omitempty"`
    // 登出时间
	LogoutTime *int64 `json:"logoutTime,omitempty"`
    // 设备Mac地址
	MacAddress *string `json:"macAddress,omitempty"`
    // 客户端版本类型
	ClientType *string `json:"clientType,omitempty"`
    // 失败原因
	OperateFailReason *string `json:"operateFailReason,omitempty"`
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 登录时间
	LoginTime *int64 `json:"loginTime,omitempty"`
    // 租户名称
	MopUserName *string `json:"mopUserName,omitempty"`
    // 操作类型: 0 是登录 ；1 是登出
	OperationType *string `json:"operationType,omitempty"`
    // 设备型号
	DeviceModel *string `json:"deviceModel,omitempty"`
    // 设备id
	DeviceUid *string `json:"deviceUid,omitempty"`
    // 用户类型: 0 是公众版用户 ；1 是政企版用户
	UserType *string `json:"userType,omitempty"`
    // 操作时间 yyyy-MM-dd HH:mm:ss
	OperateDate *string `json:"operateDate,omitempty"`
    // 操作结果: 0 是成功 ；1 是失败
	OperationResult *string `json:"operationResult,omitempty"`
}

func (s UserLoginInfoResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s UserLoginInfoResponseRecords) GoString() string {
	return s.String()
}

func (s UserLoginInfoResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserLoginInfoResponseRecords) SetIpAddress(v string) *UserLoginInfoResponseRecords {
	s.IpAddress = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetUserName(v string) *UserLoginInfoResponseRecords {
	s.UserName = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetClientVersion(v string) *UserLoginInfoResponseRecords {
	s.ClientVersion = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetUserId(v string) *UserLoginInfoResponseRecords {
	s.UserId = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetDeviceName(v string) *UserLoginInfoResponseRecords {
	s.DeviceName = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetIpAddr(v string) *UserLoginInfoResponseRecords {
	s.IpAddr = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetOperationTime(v int64) *UserLoginInfoResponseRecords {
	s.OperationTime = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetLogoutTime(v int64) *UserLoginInfoResponseRecords {
	s.LogoutTime = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetMacAddress(v string) *UserLoginInfoResponseRecords {
	s.MacAddress = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetClientType(v string) *UserLoginInfoResponseRecords {
	s.ClientType = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetOperateFailReason(v string) *UserLoginInfoResponseRecords {
	s.OperateFailReason = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetMachineId(v string) *UserLoginInfoResponseRecords {
	s.MachineId = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetLoginTime(v int64) *UserLoginInfoResponseRecords {
	s.LoginTime = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetMopUserName(v string) *UserLoginInfoResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetOperationType(v string) *UserLoginInfoResponseRecords {
	s.OperationType = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetDeviceModel(v string) *UserLoginInfoResponseRecords {
	s.DeviceModel = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetDeviceUid(v string) *UserLoginInfoResponseRecords {
	s.DeviceUid = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetUserType(v string) *UserLoginInfoResponseRecords {
	s.UserType = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetOperateDate(v string) *UserLoginInfoResponseRecords {
	s.OperateDate = &v
	return s
}

func (s *UserLoginInfoResponseRecords) SetOperationResult(v string) *UserLoginInfoResponseRecords {
	s.OperationResult = &v
	return s
}


type UserLoginInfoResponseRecordsBuilder struct {
	s *UserLoginInfoResponseRecords
}

func NewUserLoginInfoResponseRecordsBuilder() *UserLoginInfoResponseRecordsBuilder {
	s := &UserLoginInfoResponseRecords{}
	b := &UserLoginInfoResponseRecordsBuilder{s: s}
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) IpAddress(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) UserName(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) ClientVersion(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.ClientVersion = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) UserId(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.UserId = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) DeviceName(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) IpAddr(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.IpAddr = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) OperationTime(v int64) *UserLoginInfoResponseRecordsBuilder {
	b.s.OperationTime = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) LogoutTime(v int64) *UserLoginInfoResponseRecordsBuilder {
	b.s.LogoutTime = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) MacAddress(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.MacAddress = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) ClientType(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.ClientType = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) OperateFailReason(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.OperateFailReason = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) MachineId(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.MachineId = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) LoginTime(v int64) *UserLoginInfoResponseRecordsBuilder {
	b.s.LoginTime = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) MopUserName(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) OperationType(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.OperationType = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) DeviceModel(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.DeviceModel = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) DeviceUid(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) UserType(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.UserType = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) OperateDate(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.OperateDate = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) OperationResult(v string) *UserLoginInfoResponseRecordsBuilder {
	b.s.OperationResult = &v
	return b
}

func (b *UserLoginInfoResponseRecordsBuilder) Build() *UserLoginInfoResponseRecords {
	return b.s
}


