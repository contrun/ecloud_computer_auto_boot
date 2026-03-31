// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSNatRuleBody struct {
    position.Body
    // nat网关资源实例id
	NatInstanceId *string `json:"natInstanceId,omitempty"`
    // nat网关uid
	NatUid *string `json:"natUid,omitempty"`
}

func (s ListSNatRuleBody) String() string {
	return utils.Beautify(s)
}

func (s ListSNatRuleBody) GoString() string {
	return s.String()
}

func (s ListSNatRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSNatRuleBody) SetNatInstanceId(v string) *ListSNatRuleBody {
	s.NatInstanceId = &v
	return s
}

func (s *ListSNatRuleBody) SetNatUid(v string) *ListSNatRuleBody {
	s.NatUid = &v
	return s
}


type ListSNatRuleBodyBuilder struct {
	s *ListSNatRuleBody
}

func NewListSNatRuleBodyBuilder() *ListSNatRuleBodyBuilder {
	s := &ListSNatRuleBody{}
	b := &ListSNatRuleBodyBuilder{s: s}
	return b
}

func (b *ListSNatRuleBodyBuilder) NatInstanceId(v string) *ListSNatRuleBodyBuilder {
	b.s.NatInstanceId = &v
	return b
}

func (b *ListSNatRuleBodyBuilder) NatUid(v string) *ListSNatRuleBodyBuilder {
	b.s.NatUid = &v
	return b
}

func (b *ListSNatRuleBodyBuilder) Build() *ListSNatRuleBody {
	return b.s
}


