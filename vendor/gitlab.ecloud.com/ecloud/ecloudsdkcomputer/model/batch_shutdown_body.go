// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchShutdownBody struct {
    position.Body
    // 桌面id列表
	MachineIds []string `json:"machineIds,omitempty"`
}

func (s BatchShutdownBody) String() string {
	return utils.Beautify(s)
}

func (s BatchShutdownBody) GoString() string {
	return s.String()
}

func (s BatchShutdownBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchShutdownBody) SetMachineIds(v []string) *BatchShutdownBody {
	s.MachineIds = v
	return s
}


type BatchShutdownBodyBuilder struct {
	s *BatchShutdownBody
}

func NewBatchShutdownBodyBuilder() *BatchShutdownBodyBuilder {
	s := &BatchShutdownBody{}
	b := &BatchShutdownBodyBuilder{s: s}
	return b
}

func (b *BatchShutdownBodyBuilder) MachineIds(v []string) *BatchShutdownBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *BatchShutdownBodyBuilder) Build() *BatchShutdownBody {
	return b.s
}


