// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteComputerResponseData struct {

    // 订单列表
	OrderInfos *[]DeleteComputerResponseOrderInfos `json:"orderInfos,omitempty"`
}

func (s DeleteComputerResponseData) String() string {
	return utils.Beautify(s)
}

func (s DeleteComputerResponseData) GoString() string {
	return s.String()
}

func (s DeleteComputerResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteComputerResponseData) SetOrderInfos(v []DeleteComputerResponseOrderInfos) *DeleteComputerResponseData {
	s.OrderInfos = &v
	return s
}


type DeleteComputerResponseDataBuilder struct {
	s *DeleteComputerResponseData
}

func NewDeleteComputerResponseDataBuilder() *DeleteComputerResponseDataBuilder {
	s := &DeleteComputerResponseData{}
	b := &DeleteComputerResponseDataBuilder{s: s}
	return b
}

func (b *DeleteComputerResponseDataBuilder) OrderInfos(v []DeleteComputerResponseOrderInfos) *DeleteComputerResponseDataBuilder {
	b.s.OrderInfos = &v
	return b
}

func (b *DeleteComputerResponseDataBuilder) Build() *DeleteComputerResponseData {
	return b.s
}


