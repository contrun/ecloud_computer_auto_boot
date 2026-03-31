// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddMacAddressBody struct {
    position.Body
    // mac地址
	MacAddreess *string `json:"macAddreess,omitempty"`
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 备注
	Remarks *string `json:"remarks,omitempty"`
}

func (s AddMacAddressBody) String() string {
	return utils.Beautify(s)
}

func (s AddMacAddressBody) GoString() string {
	return s.String()
}

func (s AddMacAddressBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddMacAddressBody) SetMacAddreess(v string) *AddMacAddressBody {
	s.MacAddreess = &v
	return s
}

func (s *AddMacAddressBody) SetMachineId(v string) *AddMacAddressBody {
	s.MachineId = &v
	return s
}

func (s *AddMacAddressBody) SetRemarks(v string) *AddMacAddressBody {
	s.Remarks = &v
	return s
}


type AddMacAddressBodyBuilder struct {
	s *AddMacAddressBody
}

func NewAddMacAddressBodyBuilder() *AddMacAddressBodyBuilder {
	s := &AddMacAddressBody{}
	b := &AddMacAddressBodyBuilder{s: s}
	return b
}

func (b *AddMacAddressBodyBuilder) MacAddreess(v string) *AddMacAddressBodyBuilder {
	b.s.MacAddreess = &v
	return b
}

func (b *AddMacAddressBodyBuilder) MachineId(v string) *AddMacAddressBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *AddMacAddressBodyBuilder) Remarks(v string) *AddMacAddressBodyBuilder {
	b.s.Remarks = &v
	return b
}

func (b *AddMacAddressBodyBuilder) Build() *AddMacAddressBody {
	return b.s
}


