// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnLockDeviceByIdBody struct {
    position.Body
    // 设备信息表id
	Id *string `json:"id,omitempty"`
}

func (s UnLockDeviceByIdBody) String() string {
	return utils.Beautify(s)
}

func (s UnLockDeviceByIdBody) GoString() string {
	return s.String()
}

func (s UnLockDeviceByIdBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnLockDeviceByIdBody) SetId(v string) *UnLockDeviceByIdBody {
	s.Id = &v
	return s
}


type UnLockDeviceByIdBodyBuilder struct {
	s *UnLockDeviceByIdBody
}

func NewUnLockDeviceByIdBodyBuilder() *UnLockDeviceByIdBodyBuilder {
	s := &UnLockDeviceByIdBody{}
	b := &UnLockDeviceByIdBodyBuilder{s: s}
	return b
}

func (b *UnLockDeviceByIdBodyBuilder) Id(v string) *UnLockDeviceByIdBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UnLockDeviceByIdBodyBuilder) Build() *UnLockDeviceByIdBody {
	return b.s
}


