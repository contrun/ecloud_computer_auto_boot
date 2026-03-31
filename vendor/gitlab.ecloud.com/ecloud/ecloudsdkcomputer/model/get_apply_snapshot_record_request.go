// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetApplySnapshotRecordRequest struct {


	GetApplySnapshotRecordBody *GetApplySnapshotRecordBody `json:"getApplySnapshotRecordBody,omitempty"`
}

func (s GetApplySnapshotRecordRequest) String() string {
	return utils.Beautify(s)
}

func (s GetApplySnapshotRecordRequest) GoString() string {
	return s.String()
}

func (s GetApplySnapshotRecordRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplySnapshotRecordRequest) SetGetApplySnapshotRecordBody(v *GetApplySnapshotRecordBody) *GetApplySnapshotRecordRequest {
	s.GetApplySnapshotRecordBody = v
	return s
}


type GetApplySnapshotRecordRequestBuilder struct {
	s *GetApplySnapshotRecordRequest
}

func NewGetApplySnapshotRecordRequestBuilder() *GetApplySnapshotRecordRequestBuilder {
	s := &GetApplySnapshotRecordRequest{}
	b := &GetApplySnapshotRecordRequestBuilder{s: s}
	return b
}

func (b *GetApplySnapshotRecordRequestBuilder) GetApplySnapshotRecordBody(v *GetApplySnapshotRecordBody) *GetApplySnapshotRecordRequestBuilder {
	b.s.GetApplySnapshotRecordBody = v
	return b
}

func (b *GetApplySnapshotRecordRequestBuilder) Build() *GetApplySnapshotRecordRequest {
	return b.s
}


