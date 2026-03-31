// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UserLoginInfoBody struct {
    position.Body
    // 设备ip地址
	IpAddress *string `json:"ipAddress,omitempty"`
    // 分页大小
	PageSize *string `json:"pageSize,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 设备Mac地址
	MacAddress *string `json:"macAddress,omitempty"`
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 设备型号
	DeviceModel *string `json:"deviceModel,omitempty"`
    // 设备id
	DeviceUid *string `json:"deviceUid,omitempty"`
    // 租户名称
	MopUserName *string `json:"mopUserName,omitempty"`
    // 操作类型  0 是登录 ；1 是登出
	OperationType *string `json:"operationType,omitempty"`
    // 开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
    // 当前页
	Page *string `json:"page,omitempty"`
    // 操作结果 0 是成功 ；1 是失败
	OperationResult *string `json:"operationResult,omitempty"`
}

func (s UserLoginInfoBody) String() string {
	return utils.Beautify(s)
}

func (s UserLoginInfoBody) GoString() string {
	return s.String()
}

func (s UserLoginInfoBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserLoginInfoBody) SetIpAddress(v string) *UserLoginInfoBody {
	s.IpAddress = &v
	return s
}

func (s *UserLoginInfoBody) SetPageSize(v string) *UserLoginInfoBody {
	s.PageSize = &v
	return s
}

func (s *UserLoginInfoBody) SetUserName(v string) *UserLoginInfoBody {
	s.UserName = &v
	return s
}

func (s *UserLoginInfoBody) SetMacAddress(v string) *UserLoginInfoBody {
	s.MacAddress = &v
	return s
}

func (s *UserLoginInfoBody) SetMachineId(v string) *UserLoginInfoBody {
	s.MachineId = &v
	return s
}

func (s *UserLoginInfoBody) SetDeviceModel(v string) *UserLoginInfoBody {
	s.DeviceModel = &v
	return s
}

func (s *UserLoginInfoBody) SetDeviceUid(v string) *UserLoginInfoBody {
	s.DeviceUid = &v
	return s
}

func (s *UserLoginInfoBody) SetMopUserName(v string) *UserLoginInfoBody {
	s.MopUserName = &v
	return s
}

func (s *UserLoginInfoBody) SetOperationType(v string) *UserLoginInfoBody {
	s.OperationType = &v
	return s
}

func (s *UserLoginInfoBody) SetStartTime(v string) *UserLoginInfoBody {
	s.StartTime = &v
	return s
}

func (s *UserLoginInfoBody) SetEndTime(v string) *UserLoginInfoBody {
	s.EndTime = &v
	return s
}

func (s *UserLoginInfoBody) SetPage(v string) *UserLoginInfoBody {
	s.Page = &v
	return s
}

func (s *UserLoginInfoBody) SetOperationResult(v string) *UserLoginInfoBody {
	s.OperationResult = &v
	return s
}


type UserLoginInfoBodyBuilder struct {
	s *UserLoginInfoBody
}

func NewUserLoginInfoBodyBuilder() *UserLoginInfoBodyBuilder {
	s := &UserLoginInfoBody{}
	b := &UserLoginInfoBodyBuilder{s: s}
	return b
}

func (b *UserLoginInfoBodyBuilder) IpAddress(v string) *UserLoginInfoBodyBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) PageSize(v string) *UserLoginInfoBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) UserName(v string) *UserLoginInfoBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) MacAddress(v string) *UserLoginInfoBodyBuilder {
	b.s.MacAddress = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) MachineId(v string) *UserLoginInfoBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) DeviceModel(v string) *UserLoginInfoBodyBuilder {
	b.s.DeviceModel = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) DeviceUid(v string) *UserLoginInfoBodyBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) MopUserName(v string) *UserLoginInfoBodyBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) OperationType(v string) *UserLoginInfoBodyBuilder {
	b.s.OperationType = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) StartTime(v string) *UserLoginInfoBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) EndTime(v string) *UserLoginInfoBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) Page(v string) *UserLoginInfoBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) OperationResult(v string) *UserLoginInfoBodyBuilder {
	b.s.OperationResult = &v
	return b
}

func (b *UserLoginInfoBodyBuilder) Build() *UserLoginInfoBody {
	return b.s
}


