// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type HistoryPageResponseBody struct {

    // 成功集合
	Records *[]HistoryPageResponseRecords `json:"records,omitempty"`
}

func (s HistoryPageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s HistoryPageResponseBody) GoString() string {
	return s.String()
}

func (s HistoryPageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryPageResponseBody) SetRecords(v []HistoryPageResponseRecords) *HistoryPageResponseBody {
	s.Records = &v
	return s
}


type HistoryPageResponseBodyBuilder struct {
	s *HistoryPageResponseBody
}

func NewHistoryPageResponseBodyBuilder() *HistoryPageResponseBodyBuilder {
	s := &HistoryPageResponseBody{}
	b := &HistoryPageResponseBodyBuilder{s: s}
	return b
}

func (b *HistoryPageResponseBodyBuilder) Records(v []HistoryPageResponseRecords) *HistoryPageResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *HistoryPageResponseBodyBuilder) Build() *HistoryPageResponseBody {
	return b.s
}


