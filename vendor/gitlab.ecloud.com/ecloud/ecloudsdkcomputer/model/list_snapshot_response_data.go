// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSnapshotResponseData struct {

    // 云电脑ID
	MachineId *string `json:"machineId,omitempty"`
    // 快照UID
	SnapshotId *string `json:"snapshotId,omitempty"`
    // 快照名称
	SnapshotName *string `json:"snapshotName,omitempty"`
    // 快照描述
	SnapshotRemark *string `json:"snapshotRemark,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 快照状态，creating:创建中,create:创建完成,createError:创建失败,deleting:删除中,deleted:已删除,deletedError:删除失败
	SnapshotStatus *string `json:"snapshotStatus,omitempty"`
    // 快照范围,系统及数据盘(all)、仅系统盘(system)、仅数据盘(data)
	SnapshotTarget *string `json:"snapshotTarget,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
    // 底层最近一次交互返回的信息
	ResponseMsg *string `json:"responseMsg,omitempty"`
}

func (s ListSnapshotResponseData) String() string {
	return utils.Beautify(s)
}

func (s ListSnapshotResponseData) GoString() string {
	return s.String()
}

func (s ListSnapshotResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSnapshotResponseData) SetMachineId(v string) *ListSnapshotResponseData {
	s.MachineId = &v
	return s
}

func (s *ListSnapshotResponseData) SetSnapshotId(v string) *ListSnapshotResponseData {
	s.SnapshotId = &v
	return s
}

func (s *ListSnapshotResponseData) SetSnapshotName(v string) *ListSnapshotResponseData {
	s.SnapshotName = &v
	return s
}

func (s *ListSnapshotResponseData) SetSnapshotRemark(v string) *ListSnapshotResponseData {
	s.SnapshotRemark = &v
	return s
}

func (s *ListSnapshotResponseData) SetCreatedTime(v string) *ListSnapshotResponseData {
	s.CreatedTime = &v
	return s
}

func (s *ListSnapshotResponseData) SetSnapshotStatus(v string) *ListSnapshotResponseData {
	s.SnapshotStatus = &v
	return s
}

func (s *ListSnapshotResponseData) SetSnapshotTarget(v string) *ListSnapshotResponseData {
	s.SnapshotTarget = &v
	return s
}

func (s *ListSnapshotResponseData) SetMachineName(v string) *ListSnapshotResponseData {
	s.MachineName = &v
	return s
}

func (s *ListSnapshotResponseData) SetResponseMsg(v string) *ListSnapshotResponseData {
	s.ResponseMsg = &v
	return s
}


type ListSnapshotResponseDataBuilder struct {
	s *ListSnapshotResponseData
}

func NewListSnapshotResponseDataBuilder() *ListSnapshotResponseDataBuilder {
	s := &ListSnapshotResponseData{}
	b := &ListSnapshotResponseDataBuilder{s: s}
	return b
}

func (b *ListSnapshotResponseDataBuilder) MachineId(v string) *ListSnapshotResponseDataBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ListSnapshotResponseDataBuilder) SnapshotId(v string) *ListSnapshotResponseDataBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *ListSnapshotResponseDataBuilder) SnapshotName(v string) *ListSnapshotResponseDataBuilder {
	b.s.SnapshotName = &v
	return b
}

func (b *ListSnapshotResponseDataBuilder) SnapshotRemark(v string) *ListSnapshotResponseDataBuilder {
	b.s.SnapshotRemark = &v
	return b
}

func (b *ListSnapshotResponseDataBuilder) CreatedTime(v string) *ListSnapshotResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListSnapshotResponseDataBuilder) SnapshotStatus(v string) *ListSnapshotResponseDataBuilder {
	b.s.SnapshotStatus = &v
	return b
}

func (b *ListSnapshotResponseDataBuilder) SnapshotTarget(v string) *ListSnapshotResponseDataBuilder {
	b.s.SnapshotTarget = &v
	return b
}

func (b *ListSnapshotResponseDataBuilder) MachineName(v string) *ListSnapshotResponseDataBuilder {
	b.s.MachineName = &v
	return b
}

func (b *ListSnapshotResponseDataBuilder) ResponseMsg(v string) *ListSnapshotResponseDataBuilder {
	b.s.ResponseMsg = &v
	return b
}

func (b *ListSnapshotResponseDataBuilder) Build() *ListSnapshotResponseData {
	return b.s
}


