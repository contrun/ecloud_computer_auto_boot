// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetOrgListResponseBody struct {

    // 成功集合
	Records *[]GetOrgListResponseRecords `json:"records,omitempty"`
}

func (s GetOrgListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetOrgListResponseBody) GoString() string {
	return s.String()
}

func (s GetOrgListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetOrgListResponseBody) SetRecords(v []GetOrgListResponseRecords) *GetOrgListResponseBody {
	s.Records = &v
	return s
}


type GetOrgListResponseBodyBuilder struct {
	s *GetOrgListResponseBody
}

func NewGetOrgListResponseBodyBuilder() *GetOrgListResponseBodyBuilder {
	s := &GetOrgListResponseBody{}
	b := &GetOrgListResponseBodyBuilder{s: s}
	return b
}

func (b *GetOrgListResponseBodyBuilder) Records(v []GetOrgListResponseRecords) *GetOrgListResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *GetOrgListResponseBodyBuilder) Build() *GetOrgListResponseBody {
	return b.s
}


