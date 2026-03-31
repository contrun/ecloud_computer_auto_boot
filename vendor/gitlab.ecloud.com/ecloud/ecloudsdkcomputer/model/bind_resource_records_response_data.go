// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindResourceRecordsResponseData struct {

    // 分配时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 新用户/登录账号
	UserName *string `json:"userName,omitempty"`
    // 管理员账号/分配操作人
	Operator *string `json:"operator,omitempty"`
}

func (s BindResourceRecordsResponseData) String() string {
	return utils.Beautify(s)
}

func (s BindResourceRecordsResponseData) GoString() string {
	return s.String()
}

func (s BindResourceRecordsResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindResourceRecordsResponseData) SetCreatedTime(v string) *BindResourceRecordsResponseData {
	s.CreatedTime = &v
	return s
}

func (s *BindResourceRecordsResponseData) SetUserName(v string) *BindResourceRecordsResponseData {
	s.UserName = &v
	return s
}

func (s *BindResourceRecordsResponseData) SetOperator(v string) *BindResourceRecordsResponseData {
	s.Operator = &v
	return s
}


type BindResourceRecordsResponseDataBuilder struct {
	s *BindResourceRecordsResponseData
}

func NewBindResourceRecordsResponseDataBuilder() *BindResourceRecordsResponseDataBuilder {
	s := &BindResourceRecordsResponseData{}
	b := &BindResourceRecordsResponseDataBuilder{s: s}
	return b
}

func (b *BindResourceRecordsResponseDataBuilder) CreatedTime(v string) *BindResourceRecordsResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *BindResourceRecordsResponseDataBuilder) UserName(v string) *BindResourceRecordsResponseDataBuilder {
	b.s.UserName = &v
	return b
}

func (b *BindResourceRecordsResponseDataBuilder) Operator(v string) *BindResourceRecordsResponseDataBuilder {
	b.s.Operator = &v
	return b
}

func (b *BindResourceRecordsResponseDataBuilder) Build() *BindResourceRecordsResponseData {
	return b.s
}


