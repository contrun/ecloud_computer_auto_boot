// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchApplySnapshotRequestSnapshotList struct {

    // 快照id
	SnapshotId *string `json:"snapshotId,omitempty"`
    // 云电脑ID
	MachineId *string `json:"machineId,omitempty"`
}

func (s BatchApplySnapshotRequestSnapshotList) String() string {
	return utils.Beautify(s)
}

func (s BatchApplySnapshotRequestSnapshotList) GoString() string {
	return s.String()
}

func (s BatchApplySnapshotRequestSnapshotList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchApplySnapshotRequestSnapshotList) SetSnapshotId(v string) *BatchApplySnapshotRequestSnapshotList {
	s.SnapshotId = &v
	return s
}

func (s *BatchApplySnapshotRequestSnapshotList) SetMachineId(v string) *BatchApplySnapshotRequestSnapshotList {
	s.MachineId = &v
	return s
}


type BatchApplySnapshotRequestSnapshotListBuilder struct {
	s *BatchApplySnapshotRequestSnapshotList
}

func NewBatchApplySnapshotRequestSnapshotListBuilder() *BatchApplySnapshotRequestSnapshotListBuilder {
	s := &BatchApplySnapshotRequestSnapshotList{}
	b := &BatchApplySnapshotRequestSnapshotListBuilder{s: s}
	return b
}

func (b *BatchApplySnapshotRequestSnapshotListBuilder) SnapshotId(v string) *BatchApplySnapshotRequestSnapshotListBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *BatchApplySnapshotRequestSnapshotListBuilder) MachineId(v string) *BatchApplySnapshotRequestSnapshotListBuilder {
	b.s.MachineId = &v
	return b
}

func (b *BatchApplySnapshotRequestSnapshotListBuilder) Build() *BatchApplySnapshotRequestSnapshotList {
	return b.s
}


