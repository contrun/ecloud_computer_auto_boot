// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ApplySnapshotRequest struct {


	ApplySnapshotBody *ApplySnapshotBody `json:"applySnapshotBody,omitempty"`
}

func (s ApplySnapshotRequest) String() string {
	return utils.Beautify(s)
}

func (s ApplySnapshotRequest) GoString() string {
	return s.String()
}

func (s ApplySnapshotRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ApplySnapshotRequest) SetApplySnapshotBody(v *ApplySnapshotBody) *ApplySnapshotRequest {
	s.ApplySnapshotBody = v
	return s
}


type ApplySnapshotRequestBuilder struct {
	s *ApplySnapshotRequest
}

func NewApplySnapshotRequestBuilder() *ApplySnapshotRequestBuilder {
	s := &ApplySnapshotRequest{}
	b := &ApplySnapshotRequestBuilder{s: s}
	return b
}

func (b *ApplySnapshotRequestBuilder) ApplySnapshotBody(v *ApplySnapshotBody) *ApplySnapshotRequestBuilder {
	b.s.ApplySnapshotBody = v
	return b
}

func (b *ApplySnapshotRequestBuilder) Build() *ApplySnapshotRequest {
	return b.s
}


