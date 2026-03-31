// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteComputerResponseOrderInfos struct {

    // 实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // mopT
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s DeleteComputerResponseOrderInfos) String() string {
	return utils.Beautify(s)
}

func (s DeleteComputerResponseOrderInfos) GoString() string {
	return s.String()
}

func (s DeleteComputerResponseOrderInfos) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteComputerResponseOrderInfos) SetInstanceId(v string) *DeleteComputerResponseOrderInfos {
	s.InstanceId = &v
	return s
}

func (s *DeleteComputerResponseOrderInfos) SetOrderNo(v string) *DeleteComputerResponseOrderInfos {
	s.OrderNo = &v
	return s
}


type DeleteComputerResponseOrderInfosBuilder struct {
	s *DeleteComputerResponseOrderInfos
}

func NewDeleteComputerResponseOrderInfosBuilder() *DeleteComputerResponseOrderInfosBuilder {
	s := &DeleteComputerResponseOrderInfos{}
	b := &DeleteComputerResponseOrderInfosBuilder{s: s}
	return b
}

func (b *DeleteComputerResponseOrderInfosBuilder) InstanceId(v string) *DeleteComputerResponseOrderInfosBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *DeleteComputerResponseOrderInfosBuilder) OrderNo(v string) *DeleteComputerResponseOrderInfosBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *DeleteComputerResponseOrderInfosBuilder) Build() *DeleteComputerResponseOrderInfos {
	return b.s
}


