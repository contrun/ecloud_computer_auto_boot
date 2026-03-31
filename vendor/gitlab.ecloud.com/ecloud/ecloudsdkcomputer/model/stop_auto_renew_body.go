// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StopAutoRenewBody struct {
    position.Body
    // 资源id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s StopAutoRenewBody) String() string {
	return utils.Beautify(s)
}

func (s StopAutoRenewBody) GoString() string {
	return s.String()
}

func (s StopAutoRenewBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StopAutoRenewBody) SetInstanceId(v string) *StopAutoRenewBody {
	s.InstanceId = &v
	return s
}


type StopAutoRenewBodyBuilder struct {
	s *StopAutoRenewBody
}

func NewStopAutoRenewBodyBuilder() *StopAutoRenewBodyBuilder {
	s := &StopAutoRenewBody{}
	b := &StopAutoRenewBodyBuilder{s: s}
	return b
}

func (b *StopAutoRenewBodyBuilder) InstanceId(v string) *StopAutoRenewBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *StopAutoRenewBodyBuilder) Build() *StopAutoRenewBody {
	return b.s
}


