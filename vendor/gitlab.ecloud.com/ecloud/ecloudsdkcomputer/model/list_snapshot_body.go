// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSnapshotBody struct {
    position.Body
    // 快照UID
	SnapshotId *string `json:"snapshotId,omitempty"`
    // 云电脑ID
	MachineId *string `json:"machineId,omitempty"`
    // 快照名称
	SnapshotName *string `json:"snapshotName,omitempty"`
    // 快照状态，creating:创建中,create:创建完成,createError:创建失败,deleting:删除中,deleted:已删除,deletedError:删除失败
	SnapshotStatus *string `json:"snapshotStatus,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s ListSnapshotBody) String() string {
	return utils.Beautify(s)
}

func (s ListSnapshotBody) GoString() string {
	return s.String()
}

func (s ListSnapshotBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSnapshotBody) SetSnapshotId(v string) *ListSnapshotBody {
	s.SnapshotId = &v
	return s
}

func (s *ListSnapshotBody) SetMachineId(v string) *ListSnapshotBody {
	s.MachineId = &v
	return s
}

func (s *ListSnapshotBody) SetSnapshotName(v string) *ListSnapshotBody {
	s.SnapshotName = &v
	return s
}

func (s *ListSnapshotBody) SetSnapshotStatus(v string) *ListSnapshotBody {
	s.SnapshotStatus = &v
	return s
}

func (s *ListSnapshotBody) SetPageSize(v int32) *ListSnapshotBody {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotBody) SetCurrentPage(v int32) *ListSnapshotBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSnapshotBody) SetMachineName(v string) *ListSnapshotBody {
	s.MachineName = &v
	return s
}


type ListSnapshotBodyBuilder struct {
	s *ListSnapshotBody
}

func NewListSnapshotBodyBuilder() *ListSnapshotBodyBuilder {
	s := &ListSnapshotBody{}
	b := &ListSnapshotBodyBuilder{s: s}
	return b
}

func (b *ListSnapshotBodyBuilder) SnapshotId(v string) *ListSnapshotBodyBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *ListSnapshotBodyBuilder) MachineId(v string) *ListSnapshotBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ListSnapshotBodyBuilder) SnapshotName(v string) *ListSnapshotBodyBuilder {
	b.s.SnapshotName = &v
	return b
}

func (b *ListSnapshotBodyBuilder) SnapshotStatus(v string) *ListSnapshotBodyBuilder {
	b.s.SnapshotStatus = &v
	return b
}

func (b *ListSnapshotBodyBuilder) PageSize(v int32) *ListSnapshotBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListSnapshotBodyBuilder) CurrentPage(v int32) *ListSnapshotBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *ListSnapshotBodyBuilder) MachineName(v string) *ListSnapshotBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *ListSnapshotBodyBuilder) Build() *ListSnapshotBody {
	return b.s
}


