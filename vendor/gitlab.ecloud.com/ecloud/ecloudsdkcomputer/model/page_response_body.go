// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PageResponseBody struct {

    // 总用户数
	Total *int32 `json:"total,omitempty"`
    // 用户列表
	Records *[]PageResponseRecords `json:"records,omitempty"`
}

func (s PageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s PageResponseBody) GoString() string {
	return s.String()
}

func (s PageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PageResponseBody) SetTotal(v int32) *PageResponseBody {
	s.Total = &v
	return s
}

func (s *PageResponseBody) SetRecords(v []PageResponseRecords) *PageResponseBody {
	s.Records = &v
	return s
}


type PageResponseBodyBuilder struct {
	s *PageResponseBody
}

func NewPageResponseBodyBuilder() *PageResponseBodyBuilder {
	s := &PageResponseBody{}
	b := &PageResponseBodyBuilder{s: s}
	return b
}

func (b *PageResponseBodyBuilder) Total(v int32) *PageResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *PageResponseBodyBuilder) Records(v []PageResponseRecords) *PageResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *PageResponseBodyBuilder) Build() *PageResponseBody {
	return b.s
}


