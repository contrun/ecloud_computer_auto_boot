// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitCcaOrderGovResponseData struct {

    // 资源实例ID集合
	InstanceIdList []string `json:"instanceIdList,omitempty"`
    // 订单编号
	OrderNo *string `json:"orderNo,omitempty"`
    // 预占流水号
	BatchId *string `json:"batchId,omitempty"`
}

func (s SubmitCcaOrderGovResponseData) String() string {
	return utils.Beautify(s)
}

func (s SubmitCcaOrderGovResponseData) GoString() string {
	return s.String()
}

func (s SubmitCcaOrderGovResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitCcaOrderGovResponseData) SetInstanceIdList(v []string) *SubmitCcaOrderGovResponseData {
	s.InstanceIdList = v
	return s
}

func (s *SubmitCcaOrderGovResponseData) SetOrderNo(v string) *SubmitCcaOrderGovResponseData {
	s.OrderNo = &v
	return s
}

func (s *SubmitCcaOrderGovResponseData) SetBatchId(v string) *SubmitCcaOrderGovResponseData {
	s.BatchId = &v
	return s
}


type SubmitCcaOrderGovResponseDataBuilder struct {
	s *SubmitCcaOrderGovResponseData
}

func NewSubmitCcaOrderGovResponseDataBuilder() *SubmitCcaOrderGovResponseDataBuilder {
	s := &SubmitCcaOrderGovResponseData{}
	b := &SubmitCcaOrderGovResponseDataBuilder{s: s}
	return b
}

func (b *SubmitCcaOrderGovResponseDataBuilder) InstanceIdList(v []string) *SubmitCcaOrderGovResponseDataBuilder {
	b.s.InstanceIdList = v
	return b
}

func (b *SubmitCcaOrderGovResponseDataBuilder) OrderNo(v string) *SubmitCcaOrderGovResponseDataBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *SubmitCcaOrderGovResponseDataBuilder) BatchId(v string) *SubmitCcaOrderGovResponseDataBuilder {
	b.s.BatchId = &v
	return b
}

func (b *SubmitCcaOrderGovResponseDataBuilder) Build() *SubmitCcaOrderGovResponseData {
	return b.s
}


