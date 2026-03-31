// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchApplySnapshotBody struct {
    position.Body
    // 批量应用快照信息
	SnapshotList *[]BatchApplySnapshotRequestSnapshotList `json:"snapshotList,omitempty"`
}

func (s BatchApplySnapshotBody) String() string {
	return utils.Beautify(s)
}

func (s BatchApplySnapshotBody) GoString() string {
	return s.String()
}

func (s BatchApplySnapshotBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchApplySnapshotBody) SetSnapshotList(v []BatchApplySnapshotRequestSnapshotList) *BatchApplySnapshotBody {
	s.SnapshotList = &v
	return s
}


type BatchApplySnapshotBodyBuilder struct {
	s *BatchApplySnapshotBody
}

func NewBatchApplySnapshotBodyBuilder() *BatchApplySnapshotBodyBuilder {
	s := &BatchApplySnapshotBody{}
	b := &BatchApplySnapshotBodyBuilder{s: s}
	return b
}

func (b *BatchApplySnapshotBodyBuilder) SnapshotList(v []BatchApplySnapshotRequestSnapshotList) *BatchApplySnapshotBodyBuilder {
	b.s.SnapshotList = &v
	return b
}

func (b *BatchApplySnapshotBodyBuilder) Build() *BatchApplySnapshotBody {
	return b.s
}


