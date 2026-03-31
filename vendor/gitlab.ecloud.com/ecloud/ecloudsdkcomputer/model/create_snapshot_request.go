// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSnapshotRequest struct {


	CreateSnapshotBody *CreateSnapshotBody `json:"createSnapshotBody,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s CreateSnapshotRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSnapshotRequest) SetCreateSnapshotBody(v *CreateSnapshotBody) *CreateSnapshotRequest {
	s.CreateSnapshotBody = v
	return s
}


type CreateSnapshotRequestBuilder struct {
	s *CreateSnapshotRequest
}

func NewCreateSnapshotRequestBuilder() *CreateSnapshotRequestBuilder {
	s := &CreateSnapshotRequest{}
	b := &CreateSnapshotRequestBuilder{s: s}
	return b
}

func (b *CreateSnapshotRequestBuilder) CreateSnapshotBody(v *CreateSnapshotBody) *CreateSnapshotRequestBuilder {
	b.s.CreateSnapshotBody = v
	return b
}

func (b *CreateSnapshotRequestBuilder) Build() *CreateSnapshotRequest {
	return b.s
}


