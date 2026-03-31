// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchReloadBody struct {
    position.Body
    // 桌面id列表
	MachineIds []string `json:"machineIds,omitempty"`
}

func (s BatchReloadBody) String() string {
	return utils.Beautify(s)
}

func (s BatchReloadBody) GoString() string {
	return s.String()
}

func (s BatchReloadBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchReloadBody) SetMachineIds(v []string) *BatchReloadBody {
	s.MachineIds = v
	return s
}


type BatchReloadBodyBuilder struct {
	s *BatchReloadBody
}

func NewBatchReloadBodyBuilder() *BatchReloadBodyBuilder {
	s := &BatchReloadBody{}
	b := &BatchReloadBodyBuilder{s: s}
	return b
}

func (b *BatchReloadBodyBuilder) MachineIds(v []string) *BatchReloadBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *BatchReloadBodyBuilder) Build() *BatchReloadBody {
	return b.s
}


