// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PageResponseRecords struct {

    // 用户名
	AccountName *string `json:"accountName,omitempty"`
    // 操作时间
	OperateTime *string `json:"operateTime,omitempty"`
    // 操作类型
	OperateType *string `json:"operateType,omitempty"`
    // 操作路径
	OperatePath *string `json:"operatePath,omitempty"`
}

func (s PageResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s PageResponseRecords) GoString() string {
	return s.String()
}

func (s PageResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PageResponseRecords) SetAccountName(v string) *PageResponseRecords {
	s.AccountName = &v
	return s
}

func (s *PageResponseRecords) SetOperateTime(v string) *PageResponseRecords {
	s.OperateTime = &v
	return s
}

func (s *PageResponseRecords) SetOperateType(v string) *PageResponseRecords {
	s.OperateType = &v
	return s
}

func (s *PageResponseRecords) SetOperatePath(v string) *PageResponseRecords {
	s.OperatePath = &v
	return s
}


type PageResponseRecordsBuilder struct {
	s *PageResponseRecords
}

func NewPageResponseRecordsBuilder() *PageResponseRecordsBuilder {
	s := &PageResponseRecords{}
	b := &PageResponseRecordsBuilder{s: s}
	return b
}

func (b *PageResponseRecordsBuilder) AccountName(v string) *PageResponseRecordsBuilder {
	b.s.AccountName = &v
	return b
}

func (b *PageResponseRecordsBuilder) OperateTime(v string) *PageResponseRecordsBuilder {
	b.s.OperateTime = &v
	return b
}

func (b *PageResponseRecordsBuilder) OperateType(v string) *PageResponseRecordsBuilder {
	b.s.OperateType = &v
	return b
}

func (b *PageResponseRecordsBuilder) OperatePath(v string) *PageResponseRecordsBuilder {
	b.s.OperatePath = &v
	return b
}

func (b *PageResponseRecordsBuilder) Build() *PageResponseRecords {
	return b.s
}


