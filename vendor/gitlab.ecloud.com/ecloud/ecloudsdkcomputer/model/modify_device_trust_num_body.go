// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyDeviceTrustNumBody struct {
    position.Body
    // 可信设备数量
	DeviceTrustNum *int32 `json:"deviceTrustNum,omitempty"`
}

func (s ModifyDeviceTrustNumBody) String() string {
	return utils.Beautify(s)
}

func (s ModifyDeviceTrustNumBody) GoString() string {
	return s.String()
}

func (s ModifyDeviceTrustNumBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyDeviceTrustNumBody) SetDeviceTrustNum(v int32) *ModifyDeviceTrustNumBody {
	s.DeviceTrustNum = &v
	return s
}


type ModifyDeviceTrustNumBodyBuilder struct {
	s *ModifyDeviceTrustNumBody
}

func NewModifyDeviceTrustNumBodyBuilder() *ModifyDeviceTrustNumBodyBuilder {
	s := &ModifyDeviceTrustNumBody{}
	b := &ModifyDeviceTrustNumBodyBuilder{s: s}
	return b
}

func (b *ModifyDeviceTrustNumBodyBuilder) DeviceTrustNum(v int32) *ModifyDeviceTrustNumBodyBuilder {
	b.s.DeviceTrustNum = &v
	return b
}

func (b *ModifyDeviceTrustNumBodyBuilder) Build() *ModifyDeviceTrustNumBody {
	return b.s
}


