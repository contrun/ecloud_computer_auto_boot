// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartAutoRenewNatBody struct {
    position.Body
    // 自动续订时长，包月：1-60，包年：1-5，到期后按计费方式一次续订N时长，关闭自动续订请填0
	Duration *int32 `json:"duration,omitempty"`
    // 资源id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
    // 自动续订配置，开启：true，关闭：false
	AutoRenew *bool `json:"autoRenew,omitempty"`
}

func (s StartAutoRenewNatBody) String() string {
	return utils.Beautify(s)
}

func (s StartAutoRenewNatBody) GoString() string {
	return s.String()
}

func (s StartAutoRenewNatBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartAutoRenewNatBody) SetDuration(v int32) *StartAutoRenewNatBody {
	s.Duration = &v
	return s
}

func (s *StartAutoRenewNatBody) SetInstanceId(v string) *StartAutoRenewNatBody {
	s.InstanceId = &v
	return s
}

func (s *StartAutoRenewNatBody) SetAutoRenew(v bool) *StartAutoRenewNatBody {
	s.AutoRenew = &v
	return s
}


type StartAutoRenewNatBodyBuilder struct {
	s *StartAutoRenewNatBody
}

func NewStartAutoRenewNatBodyBuilder() *StartAutoRenewNatBodyBuilder {
	s := &StartAutoRenewNatBody{}
	b := &StartAutoRenewNatBodyBuilder{s: s}
	return b
}

func (b *StartAutoRenewNatBodyBuilder) Duration(v int32) *StartAutoRenewNatBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *StartAutoRenewNatBodyBuilder) InstanceId(v string) *StartAutoRenewNatBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *StartAutoRenewNatBodyBuilder) AutoRenew(v bool) *StartAutoRenewNatBodyBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *StartAutoRenewNatBodyBuilder) Build() *StartAutoRenewNatBody {
	return b.s
}


