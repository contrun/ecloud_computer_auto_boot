// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartAutoRenewBody struct {
    position.Body
    // 自动续订时长，包月：1-60，包年：1-5
	Duration *int32 `json:"duration,omitempty"`
    // 资源id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s StartAutoRenewBody) String() string {
	return utils.Beautify(s)
}

func (s StartAutoRenewBody) GoString() string {
	return s.String()
}

func (s StartAutoRenewBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartAutoRenewBody) SetDuration(v int32) *StartAutoRenewBody {
	s.Duration = &v
	return s
}

func (s *StartAutoRenewBody) SetInstanceId(v string) *StartAutoRenewBody {
	s.InstanceId = &v
	return s
}


type StartAutoRenewBodyBuilder struct {
	s *StartAutoRenewBody
}

func NewStartAutoRenewBodyBuilder() *StartAutoRenewBodyBuilder {
	s := &StartAutoRenewBody{}
	b := &StartAutoRenewBodyBuilder{s: s}
	return b
}

func (b *StartAutoRenewBodyBuilder) Duration(v int32) *StartAutoRenewBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *StartAutoRenewBodyBuilder) InstanceId(v string) *StartAutoRenewBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *StartAutoRenewBodyBuilder) Build() *StartAutoRenewBody {
	return b.s
}


