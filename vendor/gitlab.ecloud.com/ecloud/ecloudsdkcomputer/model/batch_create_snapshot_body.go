// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchCreateSnapshotBody struct {
    position.Body
    // 起始编号
	StartNumber *int32 `json:"startNumber,omitempty"`
    // 批量创建快照信息列表
	SnapshotList *[]BatchCreateSnapshotRequestSnapshotList `json:"snapshotList,omitempty"`
    // 快照名称
	SnapshotName *string `json:"snapshotName,omitempty"`
}

func (s BatchCreateSnapshotBody) String() string {
	return utils.Beautify(s)
}

func (s BatchCreateSnapshotBody) GoString() string {
	return s.String()
}

func (s BatchCreateSnapshotBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchCreateSnapshotBody) SetStartNumber(v int32) *BatchCreateSnapshotBody {
	s.StartNumber = &v
	return s
}

func (s *BatchCreateSnapshotBody) SetSnapshotList(v []BatchCreateSnapshotRequestSnapshotList) *BatchCreateSnapshotBody {
	s.SnapshotList = &v
	return s
}

func (s *BatchCreateSnapshotBody) SetSnapshotName(v string) *BatchCreateSnapshotBody {
	s.SnapshotName = &v
	return s
}


type BatchCreateSnapshotBodyBuilder struct {
	s *BatchCreateSnapshotBody
}

func NewBatchCreateSnapshotBodyBuilder() *BatchCreateSnapshotBodyBuilder {
	s := &BatchCreateSnapshotBody{}
	b := &BatchCreateSnapshotBodyBuilder{s: s}
	return b
}

func (b *BatchCreateSnapshotBodyBuilder) StartNumber(v int32) *BatchCreateSnapshotBodyBuilder {
	b.s.StartNumber = &v
	return b
}

func (b *BatchCreateSnapshotBodyBuilder) SnapshotList(v []BatchCreateSnapshotRequestSnapshotList) *BatchCreateSnapshotBodyBuilder {
	b.s.SnapshotList = &v
	return b
}

func (b *BatchCreateSnapshotBodyBuilder) SnapshotName(v string) *BatchCreateSnapshotBodyBuilder {
	b.s.SnapshotName = &v
	return b
}

func (b *BatchCreateSnapshotBodyBuilder) Build() *BatchCreateSnapshotBody {
	return b.s
}


