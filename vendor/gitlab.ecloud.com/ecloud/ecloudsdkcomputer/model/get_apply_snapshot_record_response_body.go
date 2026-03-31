// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetApplySnapshotRecordResponseBody struct {

    // 数据
	Data *[]GetApplySnapshotRecordResponseData `json:"data,omitempty"`
    // 快照还原记录总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetApplySnapshotRecordResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetApplySnapshotRecordResponseBody) GoString() string {
	return s.String()
}

func (s GetApplySnapshotRecordResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplySnapshotRecordResponseBody) SetData(v []GetApplySnapshotRecordResponseData) *GetApplySnapshotRecordResponseBody {
	s.Data = &v
	return s
}

func (s *GetApplySnapshotRecordResponseBody) SetTotalCount(v int32) *GetApplySnapshotRecordResponseBody {
	s.TotalCount = &v
	return s
}


type GetApplySnapshotRecordResponseBodyBuilder struct {
	s *GetApplySnapshotRecordResponseBody
}

func NewGetApplySnapshotRecordResponseBodyBuilder() *GetApplySnapshotRecordResponseBodyBuilder {
	s := &GetApplySnapshotRecordResponseBody{}
	b := &GetApplySnapshotRecordResponseBodyBuilder{s: s}
	return b
}

func (b *GetApplySnapshotRecordResponseBodyBuilder) Data(v []GetApplySnapshotRecordResponseData) *GetApplySnapshotRecordResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetApplySnapshotRecordResponseBodyBuilder) TotalCount(v int32) *GetApplySnapshotRecordResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetApplySnapshotRecordResponseBodyBuilder) Build() *GetApplySnapshotRecordResponseBody {
	return b.s
}


