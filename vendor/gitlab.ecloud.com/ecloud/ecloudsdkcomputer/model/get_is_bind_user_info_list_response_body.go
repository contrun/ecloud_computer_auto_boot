// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetIsBindUserInfoListResponseBody struct {

    // 总用户数
	Total *int32 `json:"total,omitempty"`
    // 用户列表
	Records *[]GetIsBindUserInfoListResponseRecords `json:"records,omitempty"`
}

func (s GetIsBindUserInfoListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetIsBindUserInfoListResponseBody) GoString() string {
	return s.String()
}

func (s GetIsBindUserInfoListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetIsBindUserInfoListResponseBody) SetTotal(v int32) *GetIsBindUserInfoListResponseBody {
	s.Total = &v
	return s
}

func (s *GetIsBindUserInfoListResponseBody) SetRecords(v []GetIsBindUserInfoListResponseRecords) *GetIsBindUserInfoListResponseBody {
	s.Records = &v
	return s
}


type GetIsBindUserInfoListResponseBodyBuilder struct {
	s *GetIsBindUserInfoListResponseBody
}

func NewGetIsBindUserInfoListResponseBodyBuilder() *GetIsBindUserInfoListResponseBodyBuilder {
	s := &GetIsBindUserInfoListResponseBody{}
	b := &GetIsBindUserInfoListResponseBodyBuilder{s: s}
	return b
}

func (b *GetIsBindUserInfoListResponseBodyBuilder) Total(v int32) *GetIsBindUserInfoListResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *GetIsBindUserInfoListResponseBodyBuilder) Records(v []GetIsBindUserInfoListResponseRecords) *GetIsBindUserInfoListResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *GetIsBindUserInfoListResponseBodyBuilder) Build() *GetIsBindUserInfoListResponseBody {
	return b.s
}


