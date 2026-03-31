// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchApplySnapshotRequest struct {


	BatchApplySnapshotBody *BatchApplySnapshotBody `json:"batchApplySnapshotBody,omitempty"`
}

func (s BatchApplySnapshotRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchApplySnapshotRequest) GoString() string {
	return s.String()
}

func (s BatchApplySnapshotRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchApplySnapshotRequest) SetBatchApplySnapshotBody(v *BatchApplySnapshotBody) *BatchApplySnapshotRequest {
	s.BatchApplySnapshotBody = v
	return s
}


type BatchApplySnapshotRequestBuilder struct {
	s *BatchApplySnapshotRequest
}

func NewBatchApplySnapshotRequestBuilder() *BatchApplySnapshotRequestBuilder {
	s := &BatchApplySnapshotRequest{}
	b := &BatchApplySnapshotRequestBuilder{s: s}
	return b
}

func (b *BatchApplySnapshotRequestBuilder) BatchApplySnapshotBody(v *BatchApplySnapshotBody) *BatchApplySnapshotRequestBuilder {
	b.s.BatchApplySnapshotBody = v
	return b
}

func (b *BatchApplySnapshotRequestBuilder) Build() *BatchApplySnapshotRequest {
	return b.s
}


