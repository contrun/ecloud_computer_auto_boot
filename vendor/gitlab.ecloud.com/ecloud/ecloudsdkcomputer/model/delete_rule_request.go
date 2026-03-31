// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteRuleRequest struct {


	DeleteRuleBody *DeleteRuleBody `json:"deleteRuleBody,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s DeleteRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteRuleRequest) SetDeleteRuleBody(v *DeleteRuleBody) *DeleteRuleRequest {
	s.DeleteRuleBody = v
	return s
}


type DeleteRuleRequestBuilder struct {
	s *DeleteRuleRequest
}

func NewDeleteRuleRequestBuilder() *DeleteRuleRequestBuilder {
	s := &DeleteRuleRequest{}
	b := &DeleteRuleRequestBuilder{s: s}
	return b
}

func (b *DeleteRuleRequestBuilder) DeleteRuleBody(v *DeleteRuleBody) *DeleteRuleRequestBuilder {
	b.s.DeleteRuleBody = v
	return b
}

func (b *DeleteRuleRequestBuilder) Build() *DeleteRuleRequest {
	return b.s
}


