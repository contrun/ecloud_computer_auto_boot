// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenameResourceBatchBody struct {
    position.Body
    // 云电脑id列表
	MachineIds []string `json:"machineIds,omitempty"`
    // 电脑名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s RenameResourceBatchBody) String() string {
	return utils.Beautify(s)
}

func (s RenameResourceBatchBody) GoString() string {
	return s.String()
}

func (s RenameResourceBatchBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenameResourceBatchBody) SetMachineIds(v []string) *RenameResourceBatchBody {
	s.MachineIds = v
	return s
}

func (s *RenameResourceBatchBody) SetMachineName(v string) *RenameResourceBatchBody {
	s.MachineName = &v
	return s
}


type RenameResourceBatchBodyBuilder struct {
	s *RenameResourceBatchBody
}

func NewRenameResourceBatchBodyBuilder() *RenameResourceBatchBodyBuilder {
	s := &RenameResourceBatchBody{}
	b := &RenameResourceBatchBodyBuilder{s: s}
	return b
}

func (b *RenameResourceBatchBodyBuilder) MachineIds(v []string) *RenameResourceBatchBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *RenameResourceBatchBodyBuilder) MachineName(v string) *RenameResourceBatchBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *RenameResourceBatchBodyBuilder) Build() *RenameResourceBatchBody {
	return b.s
}


