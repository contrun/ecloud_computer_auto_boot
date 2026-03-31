// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateSNatRuleBody struct {
    position.Body
    // NAT网关资源实例id
	NatInstanceId *string `json:"natInstanceId,omitempty"`
    // 弹性公网ipUid
	BandwidthFipUid *string `json:"bandwidthFipUid,omitempty"`
    // 子网uid
	SubnetUid *string `json:"subnetUid,omitempty"`
    // Snat规则名称
	SnatName *string `json:"snatName,omitempty"`
    // sNAT规则uid
	SnatUid *string `json:"snatUid,omitempty"`
}

func (s UpdateSNatRuleBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateSNatRuleBody) GoString() string {
	return s.String()
}

func (s UpdateSNatRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSNatRuleBody) SetNatInstanceId(v string) *UpdateSNatRuleBody {
	s.NatInstanceId = &v
	return s
}

func (s *UpdateSNatRuleBody) SetBandwidthFipUid(v string) *UpdateSNatRuleBody {
	s.BandwidthFipUid = &v
	return s
}

func (s *UpdateSNatRuleBody) SetSubnetUid(v string) *UpdateSNatRuleBody {
	s.SubnetUid = &v
	return s
}

func (s *UpdateSNatRuleBody) SetSnatName(v string) *UpdateSNatRuleBody {
	s.SnatName = &v
	return s
}

func (s *UpdateSNatRuleBody) SetSnatUid(v string) *UpdateSNatRuleBody {
	s.SnatUid = &v
	return s
}


type UpdateSNatRuleBodyBuilder struct {
	s *UpdateSNatRuleBody
}

func NewUpdateSNatRuleBodyBuilder() *UpdateSNatRuleBodyBuilder {
	s := &UpdateSNatRuleBody{}
	b := &UpdateSNatRuleBodyBuilder{s: s}
	return b
}

func (b *UpdateSNatRuleBodyBuilder) NatInstanceId(v string) *UpdateSNatRuleBodyBuilder {
	b.s.NatInstanceId = &v
	return b
}

func (b *UpdateSNatRuleBodyBuilder) BandwidthFipUid(v string) *UpdateSNatRuleBodyBuilder {
	b.s.BandwidthFipUid = &v
	return b
}

func (b *UpdateSNatRuleBodyBuilder) SubnetUid(v string) *UpdateSNatRuleBodyBuilder {
	b.s.SubnetUid = &v
	return b
}

func (b *UpdateSNatRuleBodyBuilder) SnatName(v string) *UpdateSNatRuleBodyBuilder {
	b.s.SnatName = &v
	return b
}

func (b *UpdateSNatRuleBodyBuilder) SnatUid(v string) *UpdateSNatRuleBodyBuilder {
	b.s.SnatUid = &v
	return b
}

func (b *UpdateSNatRuleBodyBuilder) Build() *UpdateSNatRuleBody {
	return b.s
}


