// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartOrStopAutoRenewBwResponseData struct {

    // 资源id
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s StartOrStopAutoRenewBwResponseData) String() string {
	return utils.Beautify(s)
}

func (s StartOrStopAutoRenewBwResponseData) GoString() string {
	return s.String()
}

func (s StartOrStopAutoRenewBwResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartOrStopAutoRenewBwResponseData) SetInstanceId(v string) *StartOrStopAutoRenewBwResponseData {
	s.InstanceId = &v
	return s
}


type StartOrStopAutoRenewBwResponseDataBuilder struct {
	s *StartOrStopAutoRenewBwResponseData
}

func NewStartOrStopAutoRenewBwResponseDataBuilder() *StartOrStopAutoRenewBwResponseDataBuilder {
	s := &StartOrStopAutoRenewBwResponseData{}
	b := &StartOrStopAutoRenewBwResponseDataBuilder{s: s}
	return b
}

func (b *StartOrStopAutoRenewBwResponseDataBuilder) InstanceId(v string) *StartOrStopAutoRenewBwResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *StartOrStopAutoRenewBwResponseDataBuilder) Build() *StartOrStopAutoRenewBwResponseData {
	return b.s
}


