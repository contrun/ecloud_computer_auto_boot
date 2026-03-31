// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelInstanceListResponseOrderInfos struct {

    // 实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // mopT
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s CancelInstanceListResponseOrderInfos) String() string {
	return utils.Beautify(s)
}

func (s CancelInstanceListResponseOrderInfos) GoString() string {
	return s.String()
}

func (s CancelInstanceListResponseOrderInfos) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelInstanceListResponseOrderInfos) SetInstanceId(v string) *CancelInstanceListResponseOrderInfos {
	s.InstanceId = &v
	return s
}

func (s *CancelInstanceListResponseOrderInfos) SetOrderNo(v string) *CancelInstanceListResponseOrderInfos {
	s.OrderNo = &v
	return s
}


type CancelInstanceListResponseOrderInfosBuilder struct {
	s *CancelInstanceListResponseOrderInfos
}

func NewCancelInstanceListResponseOrderInfosBuilder() *CancelInstanceListResponseOrderInfosBuilder {
	s := &CancelInstanceListResponseOrderInfos{}
	b := &CancelInstanceListResponseOrderInfosBuilder{s: s}
	return b
}

func (b *CancelInstanceListResponseOrderInfosBuilder) InstanceId(v string) *CancelInstanceListResponseOrderInfosBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *CancelInstanceListResponseOrderInfosBuilder) OrderNo(v string) *CancelInstanceListResponseOrderInfosBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *CancelInstanceListResponseOrderInfosBuilder) Build() *CancelInstanceListResponseOrderInfos {
	return b.s
}


