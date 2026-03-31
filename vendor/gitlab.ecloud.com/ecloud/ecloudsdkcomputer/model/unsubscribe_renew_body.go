// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeRenewBody struct {
    position.Body
    // 退订待生效续订实例id,示例值：[CCA-XXX]
	InstanceIdList []string `json:"instanceIdList,omitempty"`
}

func (s UnsubscribeRenewBody) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeRenewBody) GoString() string {
	return s.String()
}

func (s UnsubscribeRenewBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeRenewBody) SetInstanceIdList(v []string) *UnsubscribeRenewBody {
	s.InstanceIdList = v
	return s
}


type UnsubscribeRenewBodyBuilder struct {
	s *UnsubscribeRenewBody
}

func NewUnsubscribeRenewBodyBuilder() *UnsubscribeRenewBodyBuilder {
	s := &UnsubscribeRenewBody{}
	b := &UnsubscribeRenewBodyBuilder{s: s}
	return b
}

func (b *UnsubscribeRenewBodyBuilder) InstanceIdList(v []string) *UnsubscribeRenewBodyBuilder {
	b.s.InstanceIdList = v
	return b
}

func (b *UnsubscribeRenewBodyBuilder) Build() *UnsubscribeRenewBody {
	return b.s
}


