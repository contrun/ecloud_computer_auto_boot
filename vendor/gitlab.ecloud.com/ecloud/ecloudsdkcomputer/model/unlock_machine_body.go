// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnlockMachineBody struct {
    position.Body
    // 云电脑ID
	MachineId *string `json:"machineId,omitempty"`
}

func (s UnlockMachineBody) String() string {
	return utils.Beautify(s)
}

func (s UnlockMachineBody) GoString() string {
	return s.String()
}

func (s UnlockMachineBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnlockMachineBody) SetMachineId(v string) *UnlockMachineBody {
	s.MachineId = &v
	return s
}


type UnlockMachineBodyBuilder struct {
	s *UnlockMachineBody
}

func NewUnlockMachineBodyBuilder() *UnlockMachineBodyBuilder {
	s := &UnlockMachineBody{}
	b := &UnlockMachineBodyBuilder{s: s}
	return b
}

func (b *UnlockMachineBodyBuilder) MachineId(v string) *UnlockMachineBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *UnlockMachineBodyBuilder) Build() *UnlockMachineBody {
	return b.s
}


