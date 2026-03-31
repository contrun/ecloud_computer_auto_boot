// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchCreateSnapshotRequest struct {


	BatchCreateSnapshotBody *BatchCreateSnapshotBody `json:"batchCreateSnapshotBody,omitempty"`
}

func (s BatchCreateSnapshotRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchCreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s BatchCreateSnapshotRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchCreateSnapshotRequest) SetBatchCreateSnapshotBody(v *BatchCreateSnapshotBody) *BatchCreateSnapshotRequest {
	s.BatchCreateSnapshotBody = v
	return s
}


type BatchCreateSnapshotRequestBuilder struct {
	s *BatchCreateSnapshotRequest
}

func NewBatchCreateSnapshotRequestBuilder() *BatchCreateSnapshotRequestBuilder {
	s := &BatchCreateSnapshotRequest{}
	b := &BatchCreateSnapshotRequestBuilder{s: s}
	return b
}

func (b *BatchCreateSnapshotRequestBuilder) BatchCreateSnapshotBody(v *BatchCreateSnapshotBody) *BatchCreateSnapshotRequestBuilder {
	b.s.BatchCreateSnapshotBody = v
	return b
}

func (b *BatchCreateSnapshotRequestBuilder) Build() *BatchCreateSnapshotRequest {
	return b.s
}


