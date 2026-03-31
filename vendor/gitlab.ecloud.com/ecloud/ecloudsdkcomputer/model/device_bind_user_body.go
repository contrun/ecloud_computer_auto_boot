// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeviceBindUserBody struct {
    position.Body
    // 策略来源 1 云管 2 BOP 3 CHBN 4企业门户 5省二级门户  6 openapi
	SourceType *int32 `json:"sourceType,omitempty"`
    // 操作类型: 0解绑；1绑定
	BindType *string `json:"bindType,omitempty"`
    // 用户账号
	BindUsers []string `json:"bindUsers,omitempty"`
    // 设备型号
	DeviceModel *string `json:"deviceModel,omitempty"`
    // 设备Uid
	DeviceUid *string `json:"deviceUid,omitempty"`
    // mop用户ID
	MopUserId *string `json:"mopUserId,omitempty"`
    // 设备信息device_info表主键id
	DeviceId *string `json:"deviceId,omitempty"`
}

func (s DeviceBindUserBody) String() string {
	return utils.Beautify(s)
}

func (s DeviceBindUserBody) GoString() string {
	return s.String()
}

func (s DeviceBindUserBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeviceBindUserBody) SetSourceType(v int32) *DeviceBindUserBody {
	s.SourceType = &v
	return s
}

func (s *DeviceBindUserBody) SetBindType(v string) *DeviceBindUserBody {
	s.BindType = &v
	return s
}

func (s *DeviceBindUserBody) SetBindUsers(v []string) *DeviceBindUserBody {
	s.BindUsers = v
	return s
}

func (s *DeviceBindUserBody) SetDeviceModel(v string) *DeviceBindUserBody {
	s.DeviceModel = &v
	return s
}

func (s *DeviceBindUserBody) SetDeviceUid(v string) *DeviceBindUserBody {
	s.DeviceUid = &v
	return s
}

func (s *DeviceBindUserBody) SetMopUserId(v string) *DeviceBindUserBody {
	s.MopUserId = &v
	return s
}

func (s *DeviceBindUserBody) SetDeviceId(v string) *DeviceBindUserBody {
	s.DeviceId = &v
	return s
}


type DeviceBindUserBodyBuilder struct {
	s *DeviceBindUserBody
}

func NewDeviceBindUserBodyBuilder() *DeviceBindUserBodyBuilder {
	s := &DeviceBindUserBody{}
	b := &DeviceBindUserBodyBuilder{s: s}
	return b
}

func (b *DeviceBindUserBodyBuilder) SourceType(v int32) *DeviceBindUserBodyBuilder {
	b.s.SourceType = &v
	return b
}

func (b *DeviceBindUserBodyBuilder) BindType(v string) *DeviceBindUserBodyBuilder {
	b.s.BindType = &v
	return b
}

func (b *DeviceBindUserBodyBuilder) BindUsers(v []string) *DeviceBindUserBodyBuilder {
	b.s.BindUsers = v
	return b
}

func (b *DeviceBindUserBodyBuilder) DeviceModel(v string) *DeviceBindUserBodyBuilder {
	b.s.DeviceModel = &v
	return b
}

func (b *DeviceBindUserBodyBuilder) DeviceUid(v string) *DeviceBindUserBodyBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *DeviceBindUserBodyBuilder) MopUserId(v string) *DeviceBindUserBodyBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *DeviceBindUserBodyBuilder) DeviceId(v string) *DeviceBindUserBodyBuilder {
	b.s.DeviceId = &v
	return b
}

func (b *DeviceBindUserBodyBuilder) Build() *DeviceBindUserBody {
	return b.s
}


