// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserInfoListResponseBody struct {

    // 总用户数
	Total *int32 `json:"total,omitempty"`
    // 用户列表
	Records *[]GetUserInfoListResponseRecords `json:"records,omitempty"`
}

func (s GetUserInfoListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetUserInfoListResponseBody) GoString() string {
	return s.String()
}

func (s GetUserInfoListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserInfoListResponseBody) SetTotal(v int32) *GetUserInfoListResponseBody {
	s.Total = &v
	return s
}

func (s *GetUserInfoListResponseBody) SetRecords(v []GetUserInfoListResponseRecords) *GetUserInfoListResponseBody {
	s.Records = &v
	return s
}


type GetUserInfoListResponseBodyBuilder struct {
	s *GetUserInfoListResponseBody
}

func NewGetUserInfoListResponseBodyBuilder() *GetUserInfoListResponseBodyBuilder {
	s := &GetUserInfoListResponseBody{}
	b := &GetUserInfoListResponseBodyBuilder{s: s}
	return b
}

func (b *GetUserInfoListResponseBodyBuilder) Total(v int32) *GetUserInfoListResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *GetUserInfoListResponseBodyBuilder) Records(v []GetUserInfoListResponseRecords) *GetUserInfoListResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *GetUserInfoListResponseBodyBuilder) Build() *GetUserInfoListResponseBody {
	return b.s
}


