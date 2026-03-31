// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LockDeviceByIdBody struct {
    position.Body
    // 设备信息表id
	Id *string `json:"id,omitempty"`
}

func (s LockDeviceByIdBody) String() string {
	return utils.Beautify(s)
}

func (s LockDeviceByIdBody) GoString() string {
	return s.String()
}

func (s LockDeviceByIdBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LockDeviceByIdBody) SetId(v string) *LockDeviceByIdBody {
	s.Id = &v
	return s
}


type LockDeviceByIdBodyBuilder struct {
	s *LockDeviceByIdBody
}

func NewLockDeviceByIdBodyBuilder() *LockDeviceByIdBodyBuilder {
	s := &LockDeviceByIdBody{}
	b := &LockDeviceByIdBodyBuilder{s: s}
	return b
}

func (b *LockDeviceByIdBodyBuilder) Id(v string) *LockDeviceByIdBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *LockDeviceByIdBodyBuilder) Build() *LockDeviceByIdBody {
	return b.s
}


