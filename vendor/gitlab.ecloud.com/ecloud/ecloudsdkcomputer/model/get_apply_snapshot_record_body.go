// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetApplySnapshotRecordBody struct {
    position.Body
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 快照id
	SnapshotId *string `json:"snapshotId,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s GetApplySnapshotRecordBody) String() string {
	return utils.Beautify(s)
}

func (s GetApplySnapshotRecordBody) GoString() string {
	return s.String()
}

func (s GetApplySnapshotRecordBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplySnapshotRecordBody) SetMachineId(v string) *GetApplySnapshotRecordBody {
	s.MachineId = &v
	return s
}

func (s *GetApplySnapshotRecordBody) SetSnapshotId(v string) *GetApplySnapshotRecordBody {
	s.SnapshotId = &v
	return s
}

func (s *GetApplySnapshotRecordBody) SetPageSize(v int32) *GetApplySnapshotRecordBody {
	s.PageSize = &v
	return s
}

func (s *GetApplySnapshotRecordBody) SetCurrentPage(v int32) *GetApplySnapshotRecordBody {
	s.CurrentPage = &v
	return s
}


type GetApplySnapshotRecordBodyBuilder struct {
	s *GetApplySnapshotRecordBody
}

func NewGetApplySnapshotRecordBodyBuilder() *GetApplySnapshotRecordBodyBuilder {
	s := &GetApplySnapshotRecordBody{}
	b := &GetApplySnapshotRecordBodyBuilder{s: s}
	return b
}

func (b *GetApplySnapshotRecordBodyBuilder) MachineId(v string) *GetApplySnapshotRecordBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetApplySnapshotRecordBodyBuilder) SnapshotId(v string) *GetApplySnapshotRecordBodyBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *GetApplySnapshotRecordBodyBuilder) PageSize(v int32) *GetApplySnapshotRecordBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetApplySnapshotRecordBodyBuilder) CurrentPage(v int32) *GetApplySnapshotRecordBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetApplySnapshotRecordBodyBuilder) Build() *GetApplySnapshotRecordBody {
	return b.s
}


