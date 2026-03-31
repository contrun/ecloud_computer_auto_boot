// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartAutoRenewNatResponseData struct {

    // 资源id
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s StartAutoRenewNatResponseData) String() string {
	return utils.Beautify(s)
}

func (s StartAutoRenewNatResponseData) GoString() string {
	return s.String()
}

func (s StartAutoRenewNatResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartAutoRenewNatResponseData) SetInstanceId(v string) *StartAutoRenewNatResponseData {
	s.InstanceId = &v
	return s
}


type StartAutoRenewNatResponseDataBuilder struct {
	s *StartAutoRenewNatResponseData
}

func NewStartAutoRenewNatResponseDataBuilder() *StartAutoRenewNatResponseDataBuilder {
	s := &StartAutoRenewNatResponseData{}
	b := &StartAutoRenewNatResponseDataBuilder{s: s}
	return b
}

func (b *StartAutoRenewNatResponseDataBuilder) InstanceId(v string) *StartAutoRenewNatResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *StartAutoRenewNatResponseDataBuilder) Build() *StartAutoRenewNatResponseData {
	return b.s
}


