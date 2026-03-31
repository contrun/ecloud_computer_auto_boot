// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetShutdownRevertConfigBody struct {
    position.Body
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
}

func (s GetShutdownRevertConfigBody) String() string {
	return utils.Beautify(s)
}

func (s GetShutdownRevertConfigBody) GoString() string {
	return s.String()
}

func (s GetShutdownRevertConfigBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShutdownRevertConfigBody) SetMachineId(v string) *GetShutdownRevertConfigBody {
	s.MachineId = &v
	return s
}


type GetShutdownRevertConfigBodyBuilder struct {
	s *GetShutdownRevertConfigBody
}

func NewGetShutdownRevertConfigBodyBuilder() *GetShutdownRevertConfigBodyBuilder {
	s := &GetShutdownRevertConfigBody{}
	b := &GetShutdownRevertConfigBodyBuilder{s: s}
	return b
}

func (b *GetShutdownRevertConfigBodyBuilder) MachineId(v string) *GetShutdownRevertConfigBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetShutdownRevertConfigBodyBuilder) Build() *GetShutdownRevertConfigBody {
	return b.s
}


