// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchRestartBody struct {
    position.Body
    // 桌面id列表
	MachineIds []string `json:"machineIds,omitempty"`
}

func (s BatchRestartBody) String() string {
	return utils.Beautify(s)
}

func (s BatchRestartBody) GoString() string {
	return s.String()
}

func (s BatchRestartBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchRestartBody) SetMachineIds(v []string) *BatchRestartBody {
	s.MachineIds = v
	return s
}


type BatchRestartBodyBuilder struct {
	s *BatchRestartBody
}

func NewBatchRestartBodyBuilder() *BatchRestartBodyBuilder {
	s := &BatchRestartBody{}
	b := &BatchRestartBodyBuilder{s: s}
	return b
}

func (b *BatchRestartBodyBuilder) MachineIds(v []string) *BatchRestartBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *BatchRestartBodyBuilder) Build() *BatchRestartBody {
	return b.s
}


