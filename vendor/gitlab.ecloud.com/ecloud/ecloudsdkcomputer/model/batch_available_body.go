// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchAvailableBody struct {
    position.Body
    // 云电脑id列表
	MachineIds []string `json:"machineIds,omitempty"`
}

func (s BatchAvailableBody) String() string {
	return utils.Beautify(s)
}

func (s BatchAvailableBody) GoString() string {
	return s.String()
}

func (s BatchAvailableBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchAvailableBody) SetMachineIds(v []string) *BatchAvailableBody {
	s.MachineIds = v
	return s
}


type BatchAvailableBodyBuilder struct {
	s *BatchAvailableBody
}

func NewBatchAvailableBodyBuilder() *BatchAvailableBodyBuilder {
	s := &BatchAvailableBody{}
	b := &BatchAvailableBodyBuilder{s: s}
	return b
}

func (b *BatchAvailableBodyBuilder) MachineIds(v []string) *BatchAvailableBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *BatchAvailableBodyBuilder) Build() *BatchAvailableBody {
	return b.s
}


