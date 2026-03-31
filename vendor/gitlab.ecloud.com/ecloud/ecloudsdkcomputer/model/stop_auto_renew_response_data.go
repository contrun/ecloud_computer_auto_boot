// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StopAutoRenewResponseData struct {

    // 资源id
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s StopAutoRenewResponseData) String() string {
	return utils.Beautify(s)
}

func (s StopAutoRenewResponseData) GoString() string {
	return s.String()
}

func (s StopAutoRenewResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StopAutoRenewResponseData) SetInstanceId(v string) *StopAutoRenewResponseData {
	s.InstanceId = &v
	return s
}


type StopAutoRenewResponseDataBuilder struct {
	s *StopAutoRenewResponseData
}

func NewStopAutoRenewResponseDataBuilder() *StopAutoRenewResponseDataBuilder {
	s := &StopAutoRenewResponseData{}
	b := &StopAutoRenewResponseDataBuilder{s: s}
	return b
}

func (b *StopAutoRenewResponseDataBuilder) InstanceId(v string) *StopAutoRenewResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *StopAutoRenewResponseDataBuilder) Build() *StopAutoRenewResponseData {
	return b.s
}


