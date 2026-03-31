// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdResponseQosInfoDetail struct {

    // 带宽类型，当前仅上下行带宽
	BandwidthType *string `json:"bandwidthType,omitempty"`
    // qos名称
	QosName *string `json:"qosName,omitempty"`
    // 带宽数值
	QosSize *int32 `json:"qosSize,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // qosUid
	QosUid *string `json:"qosUid,omitempty"`
}

func (s GetPolicyInfoByIdResponseQosInfoDetail) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseQosInfoDetail) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseQosInfoDetail) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseQosInfoDetail) SetBandwidthType(v string) *GetPolicyInfoByIdResponseQosInfoDetail {
	s.BandwidthType = &v
	return s
}

func (s *GetPolicyInfoByIdResponseQosInfoDetail) SetQosName(v string) *GetPolicyInfoByIdResponseQosInfoDetail {
	s.QosName = &v
	return s
}

func (s *GetPolicyInfoByIdResponseQosInfoDetail) SetQosSize(v int32) *GetPolicyInfoByIdResponseQosInfoDetail {
	s.QosSize = &v
	return s
}

func (s *GetPolicyInfoByIdResponseQosInfoDetail) SetDescription(v string) *GetPolicyInfoByIdResponseQosInfoDetail {
	s.Description = &v
	return s
}

func (s *GetPolicyInfoByIdResponseQosInfoDetail) SetQosUid(v string) *GetPolicyInfoByIdResponseQosInfoDetail {
	s.QosUid = &v
	return s
}


type GetPolicyInfoByIdResponseQosInfoDetailBuilder struct {
	s *GetPolicyInfoByIdResponseQosInfoDetail
}

func NewGetPolicyInfoByIdResponseQosInfoDetailBuilder() *GetPolicyInfoByIdResponseQosInfoDetailBuilder {
	s := &GetPolicyInfoByIdResponseQosInfoDetail{}
	b := &GetPolicyInfoByIdResponseQosInfoDetailBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseQosInfoDetailBuilder) BandwidthType(v string) *GetPolicyInfoByIdResponseQosInfoDetailBuilder {
	b.s.BandwidthType = &v
	return b
}

func (b *GetPolicyInfoByIdResponseQosInfoDetailBuilder) QosName(v string) *GetPolicyInfoByIdResponseQosInfoDetailBuilder {
	b.s.QosName = &v
	return b
}

func (b *GetPolicyInfoByIdResponseQosInfoDetailBuilder) QosSize(v int32) *GetPolicyInfoByIdResponseQosInfoDetailBuilder {
	b.s.QosSize = &v
	return b
}

func (b *GetPolicyInfoByIdResponseQosInfoDetailBuilder) Description(v string) *GetPolicyInfoByIdResponseQosInfoDetailBuilder {
	b.s.Description = &v
	return b
}

func (b *GetPolicyInfoByIdResponseQosInfoDetailBuilder) QosUid(v string) *GetPolicyInfoByIdResponseQosInfoDetailBuilder {
	b.s.QosUid = &v
	return b
}

func (b *GetPolicyInfoByIdResponseQosInfoDetailBuilder) Build() *GetPolicyInfoByIdResponseQosInfoDetail {
	return b.s
}


