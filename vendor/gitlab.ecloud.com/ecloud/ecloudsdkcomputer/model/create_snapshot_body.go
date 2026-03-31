// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSnapshotBody struct {
    position.Body
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 快照名称
	SnapshotName *string `json:"snapshotName,omitempty"`
    // 快照描述
	SnapshotRemark *string `json:"snapshotRemark,omitempty"`
    // 快照范围,目前仅支持系统盘创建快照，传参system
	SnapshotTarget *string `json:"snapshotTarget,omitempty"`
}

func (s CreateSnapshotBody) String() string {
	return utils.Beautify(s)
}

func (s CreateSnapshotBody) GoString() string {
	return s.String()
}

func (s CreateSnapshotBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSnapshotBody) SetMachineId(v string) *CreateSnapshotBody {
	s.MachineId = &v
	return s
}

func (s *CreateSnapshotBody) SetSnapshotName(v string) *CreateSnapshotBody {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotBody) SetSnapshotRemark(v string) *CreateSnapshotBody {
	s.SnapshotRemark = &v
	return s
}

func (s *CreateSnapshotBody) SetSnapshotTarget(v string) *CreateSnapshotBody {
	s.SnapshotTarget = &v
	return s
}


type CreateSnapshotBodyBuilder struct {
	s *CreateSnapshotBody
}

func NewCreateSnapshotBodyBuilder() *CreateSnapshotBodyBuilder {
	s := &CreateSnapshotBody{}
	b := &CreateSnapshotBodyBuilder{s: s}
	return b
}

func (b *CreateSnapshotBodyBuilder) MachineId(v string) *CreateSnapshotBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *CreateSnapshotBodyBuilder) SnapshotName(v string) *CreateSnapshotBodyBuilder {
	b.s.SnapshotName = &v
	return b
}

func (b *CreateSnapshotBodyBuilder) SnapshotRemark(v string) *CreateSnapshotBodyBuilder {
	b.s.SnapshotRemark = &v
	return b
}

func (b *CreateSnapshotBodyBuilder) SnapshotTarget(v string) *CreateSnapshotBodyBuilder {
	b.s.SnapshotTarget = &v
	return b
}

func (b *CreateSnapshotBodyBuilder) Build() *CreateSnapshotBody {
	return b.s
}


