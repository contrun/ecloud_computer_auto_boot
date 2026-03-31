// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitBwResponseData struct {

    // 资源实例ID集合
	InstanceIdList []string `json:"instanceIdList,omitempty"`
    // 订单编号
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s SubmitBwResponseData) String() string {
	return utils.Beautify(s)
}

func (s SubmitBwResponseData) GoString() string {
	return s.String()
}

func (s SubmitBwResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitBwResponseData) SetInstanceIdList(v []string) *SubmitBwResponseData {
	s.InstanceIdList = v
	return s
}

func (s *SubmitBwResponseData) SetOrderNo(v string) *SubmitBwResponseData {
	s.OrderNo = &v
	return s
}


type SubmitBwResponseDataBuilder struct {
	s *SubmitBwResponseData
}

func NewSubmitBwResponseDataBuilder() *SubmitBwResponseDataBuilder {
	s := &SubmitBwResponseData{}
	b := &SubmitBwResponseDataBuilder{s: s}
	return b
}

func (b *SubmitBwResponseDataBuilder) InstanceIdList(v []string) *SubmitBwResponseDataBuilder {
	b.s.InstanceIdList = v
	return b
}

func (b *SubmitBwResponseDataBuilder) OrderNo(v string) *SubmitBwResponseDataBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *SubmitBwResponseDataBuilder) Build() *SubmitBwResponseData {
	return b.s
}


