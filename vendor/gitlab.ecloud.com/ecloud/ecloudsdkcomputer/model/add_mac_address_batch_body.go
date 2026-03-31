// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddMacAddressBatchBody struct {
    position.Body
    // mac地址
	MacAddreess *string `json:"macAddreess,omitempty"`
    // 云电脑id列表
	MachineIds []string `json:"machineIds,omitempty"`
    // 备注
	Remarks *string `json:"remarks,omitempty"`
}

func (s AddMacAddressBatchBody) String() string {
	return utils.Beautify(s)
}

func (s AddMacAddressBatchBody) GoString() string {
	return s.String()
}

func (s AddMacAddressBatchBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddMacAddressBatchBody) SetMacAddreess(v string) *AddMacAddressBatchBody {
	s.MacAddreess = &v
	return s
}

func (s *AddMacAddressBatchBody) SetMachineIds(v []string) *AddMacAddressBatchBody {
	s.MachineIds = v
	return s
}

func (s *AddMacAddressBatchBody) SetRemarks(v string) *AddMacAddressBatchBody {
	s.Remarks = &v
	return s
}


type AddMacAddressBatchBodyBuilder struct {
	s *AddMacAddressBatchBody
}

func NewAddMacAddressBatchBodyBuilder() *AddMacAddressBatchBodyBuilder {
	s := &AddMacAddressBatchBody{}
	b := &AddMacAddressBatchBodyBuilder{s: s}
	return b
}

func (b *AddMacAddressBatchBodyBuilder) MacAddreess(v string) *AddMacAddressBatchBodyBuilder {
	b.s.MacAddreess = &v
	return b
}

func (b *AddMacAddressBatchBodyBuilder) MachineIds(v []string) *AddMacAddressBatchBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *AddMacAddressBatchBodyBuilder) Remarks(v string) *AddMacAddressBatchBodyBuilder {
	b.s.Remarks = &v
	return b
}

func (b *AddMacAddressBatchBodyBuilder) Build() *AddMacAddressBatchBody {
	return b.s
}


