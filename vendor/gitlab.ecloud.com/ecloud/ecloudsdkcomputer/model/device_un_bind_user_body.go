// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeviceUnBindUserBody struct {
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

func (s DeviceUnBindUserBody) String() string {
	return utils.Beautify(s)
}

func (s DeviceUnBindUserBody) GoString() string {
	return s.String()
}

func (s DeviceUnBindUserBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeviceUnBindUserBody) SetSourceType(v int32) *DeviceUnBindUserBody {
	s.SourceType = &v
	return s
}

func (s *DeviceUnBindUserBody) SetBindType(v string) *DeviceUnBindUserBody {
	s.BindType = &v
	return s
}

func (s *DeviceUnBindUserBody) SetBindUsers(v []string) *DeviceUnBindUserBody {
	s.BindUsers = v
	return s
}

func (s *DeviceUnBindUserBody) SetDeviceModel(v string) *DeviceUnBindUserBody {
	s.DeviceModel = &v
	return s
}

func (s *DeviceUnBindUserBody) SetDeviceUid(v string) *DeviceUnBindUserBody {
	s.DeviceUid = &v
	return s
}

func (s *DeviceUnBindUserBody) SetMopUserId(v string) *DeviceUnBindUserBody {
	s.MopUserId = &v
	return s
}

func (s *DeviceUnBindUserBody) SetDeviceId(v string) *DeviceUnBindUserBody {
	s.DeviceId = &v
	return s
}


type DeviceUnBindUserBodyBuilder struct {
	s *DeviceUnBindUserBody
}

func NewDeviceUnBindUserBodyBuilder() *DeviceUnBindUserBodyBuilder {
	s := &DeviceUnBindUserBody{}
	b := &DeviceUnBindUserBodyBuilder{s: s}
	return b
}

func (b *DeviceUnBindUserBodyBuilder) SourceType(v int32) *DeviceUnBindUserBodyBuilder {
	b.s.SourceType = &v
	return b
}

func (b *DeviceUnBindUserBodyBuilder) BindType(v string) *DeviceUnBindUserBodyBuilder {
	b.s.BindType = &v
	return b
}

func (b *DeviceUnBindUserBodyBuilder) BindUsers(v []string) *DeviceUnBindUserBodyBuilder {
	b.s.BindUsers = v
	return b
}

func (b *DeviceUnBindUserBodyBuilder) DeviceModel(v string) *DeviceUnBindUserBodyBuilder {
	b.s.DeviceModel = &v
	return b
}

func (b *DeviceUnBindUserBodyBuilder) DeviceUid(v string) *DeviceUnBindUserBodyBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *DeviceUnBindUserBodyBuilder) MopUserId(v string) *DeviceUnBindUserBodyBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *DeviceUnBindUserBodyBuilder) DeviceId(v string) *DeviceUnBindUserBodyBuilder {
	b.s.DeviceId = &v
	return b
}

func (b *DeviceUnBindUserBodyBuilder) Build() *DeviceUnBindUserBody {
	return b.s
}


