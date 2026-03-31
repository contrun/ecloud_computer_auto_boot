// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchCreateSnapshotRequestSnapshotList struct {

    // 桌面id
	MachineId *string `json:"machineId,omitempty"`
    // 快照描述
	SnapshotRemark *string `json:"snapshotRemark,omitempty"`
}

func (s BatchCreateSnapshotRequestSnapshotList) String() string {
	return utils.Beautify(s)
}

func (s BatchCreateSnapshotRequestSnapshotList) GoString() string {
	return s.String()
}

func (s BatchCreateSnapshotRequestSnapshotList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchCreateSnapshotRequestSnapshotList) SetMachineId(v string) *BatchCreateSnapshotRequestSnapshotList {
	s.MachineId = &v
	return s
}

func (s *BatchCreateSnapshotRequestSnapshotList) SetSnapshotRemark(v string) *BatchCreateSnapshotRequestSnapshotList {
	s.SnapshotRemark = &v
	return s
}


type BatchCreateSnapshotRequestSnapshotListBuilder struct {
	s *BatchCreateSnapshotRequestSnapshotList
}

func NewBatchCreateSnapshotRequestSnapshotListBuilder() *BatchCreateSnapshotRequestSnapshotListBuilder {
	s := &BatchCreateSnapshotRequestSnapshotList{}
	b := &BatchCreateSnapshotRequestSnapshotListBuilder{s: s}
	return b
}

func (b *BatchCreateSnapshotRequestSnapshotListBuilder) MachineId(v string) *BatchCreateSnapshotRequestSnapshotListBuilder {
	b.s.MachineId = &v
	return b
}

func (b *BatchCreateSnapshotRequestSnapshotListBuilder) SnapshotRemark(v string) *BatchCreateSnapshotRequestSnapshotListBuilder {
	b.s.SnapshotRemark = &v
	return b
}

func (b *BatchCreateSnapshotRequestSnapshotListBuilder) Build() *BatchCreateSnapshotRequestSnapshotList {
	return b.s
}


