// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchUnlockMachineBody struct {
    position.Body
    // 云电脑ID列表
	MachineIds []string `json:"machineIds,omitempty"`
}

func (s BatchUnlockMachineBody) String() string {
	return utils.Beautify(s)
}

func (s BatchUnlockMachineBody) GoString() string {
	return s.String()
}

func (s BatchUnlockMachineBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUnlockMachineBody) SetMachineIds(v []string) *BatchUnlockMachineBody {
	s.MachineIds = v
	return s
}


type BatchUnlockMachineBodyBuilder struct {
	s *BatchUnlockMachineBody
}

func NewBatchUnlockMachineBodyBuilder() *BatchUnlockMachineBodyBuilder {
	s := &BatchUnlockMachineBody{}
	b := &BatchUnlockMachineBodyBuilder{s: s}
	return b
}

func (b *BatchUnlockMachineBodyBuilder) MachineIds(v []string) *BatchUnlockMachineBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *BatchUnlockMachineBodyBuilder) Build() *BatchUnlockMachineBody {
	return b.s
}


