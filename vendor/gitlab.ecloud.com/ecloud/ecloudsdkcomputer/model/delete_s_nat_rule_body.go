// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteSNatRuleBody struct {
    position.Body
    // NAT网关资源实例id
	NatInstanceId *string `json:"natInstanceId,omitempty"`
    // sNAT规则uid
	SnatUid *string `json:"snatUid,omitempty"`
}

func (s DeleteSNatRuleBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteSNatRuleBody) GoString() string {
	return s.String()
}

func (s DeleteSNatRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSNatRuleBody) SetNatInstanceId(v string) *DeleteSNatRuleBody {
	s.NatInstanceId = &v
	return s
}

func (s *DeleteSNatRuleBody) SetSnatUid(v string) *DeleteSNatRuleBody {
	s.SnatUid = &v
	return s
}


type DeleteSNatRuleBodyBuilder struct {
	s *DeleteSNatRuleBody
}

func NewDeleteSNatRuleBodyBuilder() *DeleteSNatRuleBodyBuilder {
	s := &DeleteSNatRuleBody{}
	b := &DeleteSNatRuleBodyBuilder{s: s}
	return b
}

func (b *DeleteSNatRuleBodyBuilder) NatInstanceId(v string) *DeleteSNatRuleBodyBuilder {
	b.s.NatInstanceId = &v
	return b
}

func (b *DeleteSNatRuleBodyBuilder) SnatUid(v string) *DeleteSNatRuleBodyBuilder {
	b.s.SnatUid = &v
	return b
}

func (b *DeleteSNatRuleBodyBuilder) Build() *DeleteSNatRuleBody {
	return b.s
}


