// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteSnapshotRequest struct {


	DeleteSnapshotBody *DeleteSnapshotBody `json:"deleteSnapshotBody,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s DeleteSnapshotRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSnapshotRequest) SetDeleteSnapshotBody(v *DeleteSnapshotBody) *DeleteSnapshotRequest {
	s.DeleteSnapshotBody = v
	return s
}


type DeleteSnapshotRequestBuilder struct {
	s *DeleteSnapshotRequest
}

func NewDeleteSnapshotRequestBuilder() *DeleteSnapshotRequestBuilder {
	s := &DeleteSnapshotRequest{}
	b := &DeleteSnapshotRequestBuilder{s: s}
	return b
}

func (b *DeleteSnapshotRequestBuilder) DeleteSnapshotBody(v *DeleteSnapshotBody) *DeleteSnapshotRequestBuilder {
	b.s.DeleteSnapshotBody = v
	return b
}

func (b *DeleteSnapshotRequestBuilder) Build() *DeleteSnapshotRequest {
	return b.s
}


