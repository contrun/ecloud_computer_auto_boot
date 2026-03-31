// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSNatRuleBody struct {
    position.Body
    // NAT网关资源实例id
	NatInstanceId *string `json:"natInstanceId,omitempty"`
    // 弹性公网ipUid
	BandwidthFipUid *string `json:"bandwidthFipUid,omitempty"`
    // 子网uid
	SubnetUid *string `json:"subnetUid,omitempty"`
    // Snat规则名称
	SnatName *string `json:"snatName,omitempty"`
}

func (s CreateSNatRuleBody) String() string {
	return utils.Beautify(s)
}

func (s CreateSNatRuleBody) GoString() string {
	return s.String()
}

func (s CreateSNatRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSNatRuleBody) SetNatInstanceId(v string) *CreateSNatRuleBody {
	s.NatInstanceId = &v
	return s
}

func (s *CreateSNatRuleBody) SetBandwidthFipUid(v string) *CreateSNatRuleBody {
	s.BandwidthFipUid = &v
	return s
}

func (s *CreateSNatRuleBody) SetSubnetUid(v string) *CreateSNatRuleBody {
	s.SubnetUid = &v
	return s
}

func (s *CreateSNatRuleBody) SetSnatName(v string) *CreateSNatRuleBody {
	s.SnatName = &v
	return s
}


type CreateSNatRuleBodyBuilder struct {
	s *CreateSNatRuleBody
}

func NewCreateSNatRuleBodyBuilder() *CreateSNatRuleBodyBuilder {
	s := &CreateSNatRuleBody{}
	b := &CreateSNatRuleBodyBuilder{s: s}
	return b
}

func (b *CreateSNatRuleBodyBuilder) NatInstanceId(v string) *CreateSNatRuleBodyBuilder {
	b.s.NatInstanceId = &v
	return b
}

func (b *CreateSNatRuleBodyBuilder) BandwidthFipUid(v string) *CreateSNatRuleBodyBuilder {
	b.s.BandwidthFipUid = &v
	return b
}

func (b *CreateSNatRuleBodyBuilder) SubnetUid(v string) *CreateSNatRuleBodyBuilder {
	b.s.SubnetUid = &v
	return b
}

func (b *CreateSNatRuleBodyBuilder) SnatName(v string) *CreateSNatRuleBodyBuilder {
	b.s.SnatName = &v
	return b
}

func (b *CreateSNatRuleBodyBuilder) Build() *CreateSNatRuleBody {
	return b.s
}


