// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ApplySnapshotBody struct {
    position.Body
    // 快照id
	SnapshotId *string `json:"snapshotId,omitempty"`
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
}

func (s ApplySnapshotBody) String() string {
	return utils.Beautify(s)
}

func (s ApplySnapshotBody) GoString() string {
	return s.String()
}

func (s ApplySnapshotBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ApplySnapshotBody) SetSnapshotId(v string) *ApplySnapshotBody {
	s.SnapshotId = &v
	return s
}

func (s *ApplySnapshotBody) SetMachineId(v string) *ApplySnapshotBody {
	s.MachineId = &v
	return s
}


type ApplySnapshotBodyBuilder struct {
	s *ApplySnapshotBody
}

func NewApplySnapshotBodyBuilder() *ApplySnapshotBodyBuilder {
	s := &ApplySnapshotBody{}
	b := &ApplySnapshotBodyBuilder{s: s}
	return b
}

func (b *ApplySnapshotBodyBuilder) SnapshotId(v string) *ApplySnapshotBodyBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *ApplySnapshotBodyBuilder) MachineId(v string) *ApplySnapshotBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ApplySnapshotBodyBuilder) Build() *ApplySnapshotBody {
	return b.s
}


