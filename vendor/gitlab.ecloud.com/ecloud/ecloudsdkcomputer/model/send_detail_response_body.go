// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SendDetailResponseBody struct {

    // 数据
	Data *[]SendDetailResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s SendDetailResponseBody) String() string {
	return utils.Beautify(s)
}

func (s SendDetailResponseBody) GoString() string {
	return s.String()
}

func (s SendDetailResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SendDetailResponseBody) SetData(v []SendDetailResponseData) *SendDetailResponseBody {
	s.Data = &v
	return s
}

func (s *SendDetailResponseBody) SetTotalCount(v int32) *SendDetailResponseBody {
	s.TotalCount = &v
	return s
}


type SendDetailResponseBodyBuilder struct {
	s *SendDetailResponseBody
}

func NewSendDetailResponseBodyBuilder() *SendDetailResponseBodyBuilder {
	s := &SendDetailResponseBody{}
	b := &SendDetailResponseBodyBuilder{s: s}
	return b
}

func (b *SendDetailResponseBodyBuilder) Data(v []SendDetailResponseData) *SendDetailResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *SendDetailResponseBodyBuilder) TotalCount(v int32) *SendDetailResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *SendDetailResponseBodyBuilder) Build() *SendDetailResponseBody {
	return b.s
}


