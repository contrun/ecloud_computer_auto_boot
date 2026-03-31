// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitCcaOrderResponseData struct {

    // 资源实例ID集合
	InstanceIdList []string `json:"instanceIdList,omitempty"`
    // 订单编号
	OrderNo *string `json:"orderNo,omitempty"`
    // 预占流水号
	BatchId *string `json:"batchId,omitempty"`
}

func (s SubmitCcaOrderResponseData) String() string {
	return utils.Beautify(s)
}

func (s SubmitCcaOrderResponseData) GoString() string {
	return s.String()
}

func (s SubmitCcaOrderResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitCcaOrderResponseData) SetInstanceIdList(v []string) *SubmitCcaOrderResponseData {
	s.InstanceIdList = v
	return s
}

func (s *SubmitCcaOrderResponseData) SetOrderNo(v string) *SubmitCcaOrderResponseData {
	s.OrderNo = &v
	return s
}

func (s *SubmitCcaOrderResponseData) SetBatchId(v string) *SubmitCcaOrderResponseData {
	s.BatchId = &v
	return s
}


type SubmitCcaOrderResponseDataBuilder struct {
	s *SubmitCcaOrderResponseData
}

func NewSubmitCcaOrderResponseDataBuilder() *SubmitCcaOrderResponseDataBuilder {
	s := &SubmitCcaOrderResponseData{}
	b := &SubmitCcaOrderResponseDataBuilder{s: s}
	return b
}

func (b *SubmitCcaOrderResponseDataBuilder) InstanceIdList(v []string) *SubmitCcaOrderResponseDataBuilder {
	b.s.InstanceIdList = v
	return b
}

func (b *SubmitCcaOrderResponseDataBuilder) OrderNo(v string) *SubmitCcaOrderResponseDataBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *SubmitCcaOrderResponseDataBuilder) BatchId(v string) *SubmitCcaOrderResponseDataBuilder {
	b.s.BatchId = &v
	return b
}

func (b *SubmitCcaOrderResponseDataBuilder) Build() *SubmitCcaOrderResponseData {
	return b.s
}


