// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type HistoryConnectPageResponseBody struct {

    // 总会话数
	Total *int32 `json:"total,omitempty"`
    // 会话列表
	Records *[]HistoryConnectPageResponseRecords `json:"records,omitempty"`
}

func (s HistoryConnectPageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s HistoryConnectPageResponseBody) GoString() string {
	return s.String()
}

func (s HistoryConnectPageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryConnectPageResponseBody) SetTotal(v int32) *HistoryConnectPageResponseBody {
	s.Total = &v
	return s
}

func (s *HistoryConnectPageResponseBody) SetRecords(v []HistoryConnectPageResponseRecords) *HistoryConnectPageResponseBody {
	s.Records = &v
	return s
}


type HistoryConnectPageResponseBodyBuilder struct {
	s *HistoryConnectPageResponseBody
}

func NewHistoryConnectPageResponseBodyBuilder() *HistoryConnectPageResponseBodyBuilder {
	s := &HistoryConnectPageResponseBody{}
	b := &HistoryConnectPageResponseBodyBuilder{s: s}
	return b
}

func (b *HistoryConnectPageResponseBodyBuilder) Total(v int32) *HistoryConnectPageResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *HistoryConnectPageResponseBodyBuilder) Records(v []HistoryConnectPageResponseRecords) *HistoryConnectPageResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *HistoryConnectPageResponseBodyBuilder) Build() *HistoryConnectPageResponseBody {
	return b.s
}


