// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchBindDeviceBody struct {
    position.Body
    // 设备id列表
	DeviceIds []string `json:"deviceIds,omitempty"`
    // 用户信息user_info表主键id
	UserId *string `json:"userId,omitempty"`
}

func (s BatchBindDeviceBody) String() string {
	return utils.Beautify(s)
}

func (s BatchBindDeviceBody) GoString() string {
	return s.String()
}

func (s BatchBindDeviceBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchBindDeviceBody) SetDeviceIds(v []string) *BatchBindDeviceBody {
	s.DeviceIds = v
	return s
}

func (s *BatchBindDeviceBody) SetUserId(v string) *BatchBindDeviceBody {
	s.UserId = &v
	return s
}


type BatchBindDeviceBodyBuilder struct {
	s *BatchBindDeviceBody
}

func NewBatchBindDeviceBodyBuilder() *BatchBindDeviceBodyBuilder {
	s := &BatchBindDeviceBody{}
	b := &BatchBindDeviceBodyBuilder{s: s}
	return b
}

func (b *BatchBindDeviceBodyBuilder) DeviceIds(v []string) *BatchBindDeviceBodyBuilder {
	b.s.DeviceIds = v
	return b
}

func (b *BatchBindDeviceBodyBuilder) UserId(v string) *BatchBindDeviceBodyBuilder {
	b.s.UserId = &v
	return b
}

func (b *BatchBindDeviceBodyBuilder) Build() *BatchBindDeviceBody {
	return b.s
}


