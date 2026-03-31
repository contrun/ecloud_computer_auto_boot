// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSnapshotResponseBody struct {

    // 数据
	Data *[]ListSnapshotResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s ListSnapshotResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s ListSnapshotResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSnapshotResponseBody) SetData(v []ListSnapshotResponseData) *ListSnapshotResponseBody {
	s.Data = &v
	return s
}

func (s *ListSnapshotResponseBody) SetTotalCount(v int32) *ListSnapshotResponseBody {
	s.TotalCount = &v
	return s
}


type ListSnapshotResponseBodyBuilder struct {
	s *ListSnapshotResponseBody
}

func NewListSnapshotResponseBodyBuilder() *ListSnapshotResponseBodyBuilder {
	s := &ListSnapshotResponseBody{}
	b := &ListSnapshotResponseBodyBuilder{s: s}
	return b
}

func (b *ListSnapshotResponseBodyBuilder) Data(v []ListSnapshotResponseData) *ListSnapshotResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *ListSnapshotResponseBodyBuilder) TotalCount(v int32) *ListSnapshotResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *ListSnapshotResponseBodyBuilder) Build() *ListSnapshotResponseBody {
	return b.s
}


