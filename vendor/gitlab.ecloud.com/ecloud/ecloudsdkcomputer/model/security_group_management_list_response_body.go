// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SecurityGroupManagementListResponseBody struct {

    // 数据
	Data *[]SecurityGroupManagementListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s SecurityGroupManagementListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s SecurityGroupManagementListResponseBody) GoString() string {
	return s.String()
}

func (s SecurityGroupManagementListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SecurityGroupManagementListResponseBody) SetData(v []SecurityGroupManagementListResponseData) *SecurityGroupManagementListResponseBody {
	s.Data = &v
	return s
}

func (s *SecurityGroupManagementListResponseBody) SetTotalCount(v int32) *SecurityGroupManagementListResponseBody {
	s.TotalCount = &v
	return s
}


type SecurityGroupManagementListResponseBodyBuilder struct {
	s *SecurityGroupManagementListResponseBody
}

func NewSecurityGroupManagementListResponseBodyBuilder() *SecurityGroupManagementListResponseBodyBuilder {
	s := &SecurityGroupManagementListResponseBody{}
	b := &SecurityGroupManagementListResponseBodyBuilder{s: s}
	return b
}

func (b *SecurityGroupManagementListResponseBodyBuilder) Data(v []SecurityGroupManagementListResponseData) *SecurityGroupManagementListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *SecurityGroupManagementListResponseBodyBuilder) TotalCount(v int32) *SecurityGroupManagementListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *SecurityGroupManagementListResponseBodyBuilder) Build() *SecurityGroupManagementListResponseBody {
	return b.s
}


