// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchUnbindDeviceBody struct {
    position.Body
    // 设备id列表
	DeviceIds []string `json:"deviceIds,omitempty"`
    // 用户信息user_info表主键id
	UserId *string `json:"userId,omitempty"`
}

func (s BatchUnbindDeviceBody) String() string {
	return utils.Beautify(s)
}

func (s BatchUnbindDeviceBody) GoString() string {
	return s.String()
}

func (s BatchUnbindDeviceBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUnbindDeviceBody) SetDeviceIds(v []string) *BatchUnbindDeviceBody {
	s.DeviceIds = v
	return s
}

func (s *BatchUnbindDeviceBody) SetUserId(v string) *BatchUnbindDeviceBody {
	s.UserId = &v
	return s
}


type BatchUnbindDeviceBodyBuilder struct {
	s *BatchUnbindDeviceBody
}

func NewBatchUnbindDeviceBodyBuilder() *BatchUnbindDeviceBodyBuilder {
	s := &BatchUnbindDeviceBody{}
	b := &BatchUnbindDeviceBodyBuilder{s: s}
	return b
}

func (b *BatchUnbindDeviceBodyBuilder) DeviceIds(v []string) *BatchUnbindDeviceBodyBuilder {
	b.s.DeviceIds = v
	return b
}

func (b *BatchUnbindDeviceBodyBuilder) UserId(v string) *BatchUnbindDeviceBodyBuilder {
	b.s.UserId = &v
	return b
}

func (b *BatchUnbindDeviceBodyBuilder) Build() *BatchUnbindDeviceBody {
	return b.s
}


