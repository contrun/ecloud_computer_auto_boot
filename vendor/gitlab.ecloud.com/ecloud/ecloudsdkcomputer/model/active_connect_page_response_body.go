// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ActiveConnectPageResponseBody struct {

    // 总会话数
	Total *int32 `json:"total,omitempty"`
    // 会话列表
	Records *[]ActiveConnectPageResponseRecords `json:"records,omitempty"`
}

func (s ActiveConnectPageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ActiveConnectPageResponseBody) GoString() string {
	return s.String()
}

func (s ActiveConnectPageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActiveConnectPageResponseBody) SetTotal(v int32) *ActiveConnectPageResponseBody {
	s.Total = &v
	return s
}

func (s *ActiveConnectPageResponseBody) SetRecords(v []ActiveConnectPageResponseRecords) *ActiveConnectPageResponseBody {
	s.Records = &v
	return s
}


type ActiveConnectPageResponseBodyBuilder struct {
	s *ActiveConnectPageResponseBody
}

func NewActiveConnectPageResponseBodyBuilder() *ActiveConnectPageResponseBodyBuilder {
	s := &ActiveConnectPageResponseBody{}
	b := &ActiveConnectPageResponseBodyBuilder{s: s}
	return b
}

func (b *ActiveConnectPageResponseBodyBuilder) Total(v int32) *ActiveConnectPageResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ActiveConnectPageResponseBodyBuilder) Records(v []ActiveConnectPageResponseRecords) *ActiveConnectPageResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *ActiveConnectPageResponseBodyBuilder) Build() *ActiveConnectPageResponseBody {
	return b.s
}


