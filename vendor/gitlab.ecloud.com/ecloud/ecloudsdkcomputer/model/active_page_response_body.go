// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ActivePageResponseBody struct {

    // 成功集合
	Records *[]ActivePageResponseRecords `json:"records,omitempty"`
}

func (s ActivePageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ActivePageResponseBody) GoString() string {
	return s.String()
}

func (s ActivePageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActivePageResponseBody) SetRecords(v []ActivePageResponseRecords) *ActivePageResponseBody {
	s.Records = &v
	return s
}


type ActivePageResponseBodyBuilder struct {
	s *ActivePageResponseBody
}

func NewActivePageResponseBodyBuilder() *ActivePageResponseBodyBuilder {
	s := &ActivePageResponseBody{}
	b := &ActivePageResponseBodyBuilder{s: s}
	return b
}

func (b *ActivePageResponseBodyBuilder) Records(v []ActivePageResponseRecords) *ActivePageResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *ActivePageResponseBodyBuilder) Build() *ActivePageResponseBody {
	return b.s
}


