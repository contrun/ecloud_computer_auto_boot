// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartAutoRenewResponseData struct {

    // 资源id
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s StartAutoRenewResponseData) String() string {
	return utils.Beautify(s)
}

func (s StartAutoRenewResponseData) GoString() string {
	return s.String()
}

func (s StartAutoRenewResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartAutoRenewResponseData) SetInstanceId(v string) *StartAutoRenewResponseData {
	s.InstanceId = &v
	return s
}


type StartAutoRenewResponseDataBuilder struct {
	s *StartAutoRenewResponseData
}

func NewStartAutoRenewResponseDataBuilder() *StartAutoRenewResponseDataBuilder {
	s := &StartAutoRenewResponseData{}
	b := &StartAutoRenewResponseDataBuilder{s: s}
	return b
}

func (b *StartAutoRenewResponseDataBuilder) InstanceId(v string) *StartAutoRenewResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *StartAutoRenewResponseDataBuilder) Build() *StartAutoRenewResponseData {
	return b.s
}


