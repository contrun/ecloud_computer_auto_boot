// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitNatBody struct {
    position.Body
    // 订购时长，包月：1-60，包年：1-5，非免费时必填，免费时此字段不生效
	PeriodNum *int32 `json:"periodNum,omitempty"`
    // 规格编码，2409121114115168：小型网关，仅仅支持免费赠送；2409121114535170：中型网关；2409121115165172：大型NAT网关；2409121115375174：超大型NAT网关。示例值：2409121114115168
	ChaId *string `json:"chaId,omitempty"`
    // NAT网关名称，长度2-128位，字母、中划线和下划线的组合（必须以字母开头），示例值：nat1
	NatName *string `json:"natName,omitempty"`
    // vpc的id,示例值：vpc-xxxx
	VpcId *string `json:"vpcId,omitempty"`
    // 是否自动续订标识，true：是，false：否，默认为false
	AutoRenew *bool `json:"autoRenew,omitempty"`
    // 计费方式，month：包月，year：包年，非免费时必填，免费时此字段不生效
	MeasureUnit *string `json:"measureUnit,omitempty"`
    // 可用区编码。示例值：CIDC-CORE-00
	ZoneCode *string `json:"zoneCode,omitempty"`
}

func (s SubmitNatBody) String() string {
	return utils.Beautify(s)
}

func (s SubmitNatBody) GoString() string {
	return s.String()
}

func (s SubmitNatBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitNatBody) SetPeriodNum(v int32) *SubmitNatBody {
	s.PeriodNum = &v
	return s
}

func (s *SubmitNatBody) SetChaId(v string) *SubmitNatBody {
	s.ChaId = &v
	return s
}

func (s *SubmitNatBody) SetNatName(v string) *SubmitNatBody {
	s.NatName = &v
	return s
}

func (s *SubmitNatBody) SetVpcId(v string) *SubmitNatBody {
	s.VpcId = &v
	return s
}

func (s *SubmitNatBody) SetAutoRenew(v bool) *SubmitNatBody {
	s.AutoRenew = &v
	return s
}

func (s *SubmitNatBody) SetMeasureUnit(v string) *SubmitNatBody {
	s.MeasureUnit = &v
	return s
}

func (s *SubmitNatBody) SetZoneCode(v string) *SubmitNatBody {
	s.ZoneCode = &v
	return s
}


type SubmitNatBodyBuilder struct {
	s *SubmitNatBody
}

func NewSubmitNatBodyBuilder() *SubmitNatBodyBuilder {
	s := &SubmitNatBody{}
	b := &SubmitNatBodyBuilder{s: s}
	return b
}

func (b *SubmitNatBodyBuilder) PeriodNum(v int32) *SubmitNatBodyBuilder {
	b.s.PeriodNum = &v
	return b
}

func (b *SubmitNatBodyBuilder) ChaId(v string) *SubmitNatBodyBuilder {
	b.s.ChaId = &v
	return b
}

func (b *SubmitNatBodyBuilder) NatName(v string) *SubmitNatBodyBuilder {
	b.s.NatName = &v
	return b
}

func (b *SubmitNatBodyBuilder) VpcId(v string) *SubmitNatBodyBuilder {
	b.s.VpcId = &v
	return b
}

func (b *SubmitNatBodyBuilder) AutoRenew(v bool) *SubmitNatBodyBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *SubmitNatBodyBuilder) MeasureUnit(v string) *SubmitNatBodyBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *SubmitNatBodyBuilder) ZoneCode(v string) *SubmitNatBodyBuilder {
	b.s.ZoneCode = &v
	return b
}

func (b *SubmitNatBodyBuilder) Build() *SubmitNatBody {
	return b.s
}


