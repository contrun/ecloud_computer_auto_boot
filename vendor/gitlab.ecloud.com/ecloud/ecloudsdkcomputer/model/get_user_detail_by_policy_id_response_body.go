// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserDetailByPolicyIdResponseBody struct {

    // 总用户数
	Total *int32 `json:"total,omitempty"`
    // 用户列表
	Records *[]GetUserDetailByPolicyIdResponseRecords `json:"records,omitempty"`
}

func (s GetUserDetailByPolicyIdResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetUserDetailByPolicyIdResponseBody) GoString() string {
	return s.String()
}

func (s GetUserDetailByPolicyIdResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserDetailByPolicyIdResponseBody) SetTotal(v int32) *GetUserDetailByPolicyIdResponseBody {
	s.Total = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseBody) SetRecords(v []GetUserDetailByPolicyIdResponseRecords) *GetUserDetailByPolicyIdResponseBody {
	s.Records = &v
	return s
}


type GetUserDetailByPolicyIdResponseBodyBuilder struct {
	s *GetUserDetailByPolicyIdResponseBody
}

func NewGetUserDetailByPolicyIdResponseBodyBuilder() *GetUserDetailByPolicyIdResponseBodyBuilder {
	s := &GetUserDetailByPolicyIdResponseBody{}
	b := &GetUserDetailByPolicyIdResponseBodyBuilder{s: s}
	return b
}

func (b *GetUserDetailByPolicyIdResponseBodyBuilder) Total(v int32) *GetUserDetailByPolicyIdResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseBodyBuilder) Records(v []GetUserDetailByPolicyIdResponseRecords) *GetUserDetailByPolicyIdResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseBodyBuilder) Build() *GetUserDetailByPolicyIdResponseBody {
	return b.s
}


