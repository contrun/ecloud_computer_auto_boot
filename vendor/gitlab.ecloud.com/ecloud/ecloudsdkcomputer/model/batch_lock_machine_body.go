// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchLockMachineBody struct {
    position.Body
    // 云电脑ID列表
	MachineIds []string `json:"machineIds,omitempty"`
}

func (s BatchLockMachineBody) String() string {
	return utils.Beautify(s)
}

func (s BatchLockMachineBody) GoString() string {
	return s.String()
}

func (s BatchLockMachineBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchLockMachineBody) SetMachineIds(v []string) *BatchLockMachineBody {
	s.MachineIds = v
	return s
}


type BatchLockMachineBodyBuilder struct {
	s *BatchLockMachineBody
}

func NewBatchLockMachineBodyBuilder() *BatchLockMachineBodyBuilder {
	s := &BatchLockMachineBody{}
	b := &BatchLockMachineBodyBuilder{s: s}
	return b
}

func (b *BatchLockMachineBodyBuilder) MachineIds(v []string) *BatchLockMachineBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *BatchLockMachineBodyBuilder) Build() *BatchLockMachineBody {
	return b.s
}


