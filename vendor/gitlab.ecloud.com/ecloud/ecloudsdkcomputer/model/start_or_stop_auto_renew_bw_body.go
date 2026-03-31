// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartOrStopAutoRenewBwBody struct {
    position.Body
    // 自动续订时长，包月：1-60，包年：1-5，到期后按计费方式一次续订N时长，关闭自动续订请填0
	Duration *int32 `json:"duration,omitempty"`
    // 资源id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
    // 自动续订配置，开启：true，关闭：false
	AutoRenew *bool `json:"autoRenew,omitempty"`
}

func (s StartOrStopAutoRenewBwBody) String() string {
	return utils.Beautify(s)
}

func (s StartOrStopAutoRenewBwBody) GoString() string {
	return s.String()
}

func (s StartOrStopAutoRenewBwBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartOrStopAutoRenewBwBody) SetDuration(v int32) *StartOrStopAutoRenewBwBody {
	s.Duration = &v
	return s
}

func (s *StartOrStopAutoRenewBwBody) SetInstanceId(v string) *StartOrStopAutoRenewBwBody {
	s.InstanceId = &v
	return s
}

func (s *StartOrStopAutoRenewBwBody) SetAutoRenew(v bool) *StartOrStopAutoRenewBwBody {
	s.AutoRenew = &v
	return s
}


type StartOrStopAutoRenewBwBodyBuilder struct {
	s *StartOrStopAutoRenewBwBody
}

func NewStartOrStopAutoRenewBwBodyBuilder() *StartOrStopAutoRenewBwBodyBuilder {
	s := &StartOrStopAutoRenewBwBody{}
	b := &StartOrStopAutoRenewBwBodyBuilder{s: s}
	return b
}

func (b *StartOrStopAutoRenewBwBodyBuilder) Duration(v int32) *StartOrStopAutoRenewBwBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *StartOrStopAutoRenewBwBodyBuilder) InstanceId(v string) *StartOrStopAutoRenewBwBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *StartOrStopAutoRenewBwBodyBuilder) AutoRenew(v bool) *StartOrStopAutoRenewBwBodyBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *StartOrStopAutoRenewBwBodyBuilder) Build() *StartOrStopAutoRenewBwBody {
	return b.s
}


