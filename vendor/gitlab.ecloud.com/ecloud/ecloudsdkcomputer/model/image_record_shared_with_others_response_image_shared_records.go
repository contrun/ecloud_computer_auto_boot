// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageRecordSharedWithOthersResponseImageSharedRecords struct {

    // 数据
	Data *[]ImageRecordSharedWithOthersResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s ImageRecordSharedWithOthersResponseImageSharedRecords) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedWithOthersResponseImageSharedRecords) GoString() string {
	return s.String()
}

func (s ImageRecordSharedWithOthersResponseImageSharedRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedWithOthersResponseImageSharedRecords) SetData(v []ImageRecordSharedWithOthersResponseData) *ImageRecordSharedWithOthersResponseImageSharedRecords {
	s.Data = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseImageSharedRecords) SetTotalCount(v int32) *ImageRecordSharedWithOthersResponseImageSharedRecords {
	s.TotalCount = &v
	return s
}


type ImageRecordSharedWithOthersResponseImageSharedRecordsBuilder struct {
	s *ImageRecordSharedWithOthersResponseImageSharedRecords
}

func NewImageRecordSharedWithOthersResponseImageSharedRecordsBuilder() *ImageRecordSharedWithOthersResponseImageSharedRecordsBuilder {
	s := &ImageRecordSharedWithOthersResponseImageSharedRecords{}
	b := &ImageRecordSharedWithOthersResponseImageSharedRecordsBuilder{s: s}
	return b
}

func (b *ImageRecordSharedWithOthersResponseImageSharedRecordsBuilder) Data(v []ImageRecordSharedWithOthersResponseData) *ImageRecordSharedWithOthersResponseImageSharedRecordsBuilder {
	b.s.Data = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseImageSharedRecordsBuilder) TotalCount(v int32) *ImageRecordSharedWithOthersResponseImageSharedRecordsBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseImageSharedRecordsBuilder) Build() *ImageRecordSharedWithOthersResponseImageSharedRecords {
	return b.s
}


