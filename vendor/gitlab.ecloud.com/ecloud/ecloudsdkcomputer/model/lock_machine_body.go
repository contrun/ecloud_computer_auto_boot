// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LockMachineBody struct {
    position.Body
    // 云电脑ID
	MachineId *string `json:"machineId,omitempty"`
}

func (s LockMachineBody) String() string {
	return utils.Beautify(s)
}

func (s LockMachineBody) GoString() string {
	return s.String()
}

func (s LockMachineBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LockMachineBody) SetMachineId(v string) *LockMachineBody {
	s.MachineId = &v
	return s
}


type LockMachineBodyBuilder struct {
	s *LockMachineBody
}

func NewLockMachineBodyBuilder() *LockMachineBodyBuilder {
	s := &LockMachineBody{}
	b := &LockMachineBodyBuilder{s: s}
	return b
}

func (b *LockMachineBodyBuilder) MachineId(v string) *LockMachineBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *LockMachineBodyBuilder) Build() *LockMachineBody {
	return b.s
}


