// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteSnapshotBody struct {
    position.Body
    // 快照id
	SnapshotId *string `json:"snapshotId,omitempty"`
}

func (s DeleteSnapshotBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteSnapshotBody) GoString() string {
	return s.String()
}

func (s DeleteSnapshotBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSnapshotBody) SetSnapshotId(v string) *DeleteSnapshotBody {
	s.SnapshotId = &v
	return s
}


type DeleteSnapshotBodyBuilder struct {
	s *DeleteSnapshotBody
}

func NewDeleteSnapshotBodyBuilder() *DeleteSnapshotBodyBuilder {
	s := &DeleteSnapshotBody{}
	b := &DeleteSnapshotBodyBuilder{s: s}
	return b
}

func (b *DeleteSnapshotBodyBuilder) SnapshotId(v string) *DeleteSnapshotBodyBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *DeleteSnapshotBodyBuilder) Build() *DeleteSnapshotBody {
	return b.s
}


