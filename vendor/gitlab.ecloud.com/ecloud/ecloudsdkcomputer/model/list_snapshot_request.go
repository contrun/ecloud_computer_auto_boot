// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSnapshotRequest struct {


	ListSnapshotBody *ListSnapshotBody `json:"listSnapshotBody,omitempty"`
}

func (s ListSnapshotRequest) String() string {
	return utils.Beautify(s)
}

func (s ListSnapshotRequest) GoString() string {
	return s.String()
}

func (s ListSnapshotRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSnapshotRequest) SetListSnapshotBody(v *ListSnapshotBody) *ListSnapshotRequest {
	s.ListSnapshotBody = v
	return s
}


type ListSnapshotRequestBuilder struct {
	s *ListSnapshotRequest
}

func NewListSnapshotRequestBuilder() *ListSnapshotRequestBuilder {
	s := &ListSnapshotRequest{}
	b := &ListSnapshotRequestBuilder{s: s}
	return b
}

func (b *ListSnapshotRequestBuilder) ListSnapshotBody(v *ListSnapshotBody) *ListSnapshotRequestBuilder {
	b.s.ListSnapshotBody = v
	return b
}

func (b *ListSnapshotRequestBuilder) Build() *ListSnapshotRequest {
	return b.s
}


