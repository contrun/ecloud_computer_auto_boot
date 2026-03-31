// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSNatRuleResponseData struct {

    // vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
    // 弹性公网ip
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
    // 弹性公网uid
	BandwidthFipUid *string `json:"bandwidthFipUid,omitempty"`
    // 子网uid
	SubnetUid *string `json:"subnetUid,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 来源
	RequestSource *string `json:"requestSource,omitempty"`
    // snat规则名称
	SnatName *string `json:"snatName,omitempty"`
    // snat规则uid
	SnatUid *string `json:"snatUid,omitempty"`
    // 子网名称
	SubnetName *string `json:"subnetName,omitempty"`
    // snat规则名状态
	Status *string `json:"status,omitempty"`
}

func (s ListSNatRuleResponseData) String() string {
	return utils.Beautify(s)
}

func (s ListSNatRuleResponseData) GoString() string {
	return s.String()
}

func (s ListSNatRuleResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSNatRuleResponseData) SetVpcUid(v string) *ListSNatRuleResponseData {
	s.VpcUid = &v
	return s
}

func (s *ListSNatRuleResponseData) SetPublicIpAddress(v string) *ListSNatRuleResponseData {
	s.PublicIpAddress = &v
	return s
}

func (s *ListSNatRuleResponseData) SetBandwidthFipUid(v string) *ListSNatRuleResponseData {
	s.BandwidthFipUid = &v
	return s
}

func (s *ListSNatRuleResponseData) SetSubnetUid(v string) *ListSNatRuleResponseData {
	s.SubnetUid = &v
	return s
}

func (s *ListSNatRuleResponseData) SetCreatedTime(v string) *ListSNatRuleResponseData {
	s.CreatedTime = &v
	return s
}

func (s *ListSNatRuleResponseData) SetRequestSource(v string) *ListSNatRuleResponseData {
	s.RequestSource = &v
	return s
}

func (s *ListSNatRuleResponseData) SetSnatName(v string) *ListSNatRuleResponseData {
	s.SnatName = &v
	return s
}

func (s *ListSNatRuleResponseData) SetSnatUid(v string) *ListSNatRuleResponseData {
	s.SnatUid = &v
	return s
}

func (s *ListSNatRuleResponseData) SetSubnetName(v string) *ListSNatRuleResponseData {
	s.SubnetName = &v
	return s
}

func (s *ListSNatRuleResponseData) SetStatus(v string) *ListSNatRuleResponseData {
	s.Status = &v
	return s
}


type ListSNatRuleResponseDataBuilder struct {
	s *ListSNatRuleResponseData
}

func NewListSNatRuleResponseDataBuilder() *ListSNatRuleResponseDataBuilder {
	s := &ListSNatRuleResponseData{}
	b := &ListSNatRuleResponseDataBuilder{s: s}
	return b
}

func (b *ListSNatRuleResponseDataBuilder) VpcUid(v string) *ListSNatRuleResponseDataBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) PublicIpAddress(v string) *ListSNatRuleResponseDataBuilder {
	b.s.PublicIpAddress = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) BandwidthFipUid(v string) *ListSNatRuleResponseDataBuilder {
	b.s.BandwidthFipUid = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) SubnetUid(v string) *ListSNatRuleResponseDataBuilder {
	b.s.SubnetUid = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) CreatedTime(v string) *ListSNatRuleResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) RequestSource(v string) *ListSNatRuleResponseDataBuilder {
	b.s.RequestSource = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) SnatName(v string) *ListSNatRuleResponseDataBuilder {
	b.s.SnatName = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) SnatUid(v string) *ListSNatRuleResponseDataBuilder {
	b.s.SnatUid = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) SubnetName(v string) *ListSNatRuleResponseDataBuilder {
	b.s.SubnetName = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) Status(v string) *ListSNatRuleResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *ListSNatRuleResponseDataBuilder) Build() *ListSNatRuleResponseData {
	return b.s
}


