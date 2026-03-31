// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSnapshotResponseBody struct {

    // 快照uid
	SnapshotId *string `json:"snapshotId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return utils.Beautify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s CreateSnapshotResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}


type CreateSnapshotResponseBodyBuilder struct {
	s *CreateSnapshotResponseBody
}

func NewCreateSnapshotResponseBodyBuilder() *CreateSnapshotResponseBodyBuilder {
	s := &CreateSnapshotResponseBody{}
	b := &CreateSnapshotResponseBodyBuilder{s: s}
	return b
}

func (b *CreateSnapshotResponseBodyBuilder) SnapshotId(v string) *CreateSnapshotResponseBodyBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *CreateSnapshotResponseBodyBuilder) Build() *CreateSnapshotResponseBody {
	return b.s
}


