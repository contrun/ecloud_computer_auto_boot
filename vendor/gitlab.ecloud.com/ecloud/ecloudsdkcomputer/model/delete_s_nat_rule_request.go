// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteSNatRuleRequest struct {


	DeleteSNatRuleBody *DeleteSNatRuleBody `json:"deleteSNatRuleBody,omitempty"`
}

func (s DeleteSNatRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteSNatRuleRequest) GoString() string {
	return s.String()
}

func (s DeleteSNatRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSNatRuleRequest) SetDeleteSNatRuleBody(v *DeleteSNatRuleBody) *DeleteSNatRuleRequest {
	s.DeleteSNatRuleBody = v
	return s
}


type DeleteSNatRuleRequestBuilder struct {
	s *DeleteSNatRuleRequest
}

func NewDeleteSNatRuleRequestBuilder() *DeleteSNatRuleRequestBuilder {
	s := &DeleteSNatRuleRequest{}
	b := &DeleteSNatRuleRequestBuilder{s: s}
	return b
}

func (b *DeleteSNatRuleRequestBuilder) DeleteSNatRuleBody(v *DeleteSNatRuleBody) *DeleteSNatRuleRequestBuilder {
	b.s.DeleteSNatRuleBody = v
	return b
}

func (b *DeleteSNatRuleRequestBuilder) Build() *DeleteSNatRuleRequest {
	return b.s
}


