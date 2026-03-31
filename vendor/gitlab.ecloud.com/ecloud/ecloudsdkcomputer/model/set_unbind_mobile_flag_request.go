// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetUnbindMobileFlagRequest struct {


	SetUnbindMobileFlagBody *SetUnbindMobileFlagBody `json:"setUnbindMobileFlagBody,omitempty"`
}

func (s SetUnbindMobileFlagRequest) String() string {
	return utils.Beautify(s)
}

func (s SetUnbindMobileFlagRequest) GoString() string {
	return s.String()
}

func (s SetUnbindMobileFlagRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetUnbindMobileFlagRequest) SetSetUnbindMobileFlagBody(v *SetUnbindMobileFlagBody) *SetUnbindMobileFlagRequest {
	s.SetUnbindMobileFlagBody = v
	return s
}


type SetUnbindMobileFlagRequestBuilder struct {
	s *SetUnbindMobileFlagRequest
}

func NewSetUnbindMobileFlagRequestBuilder() *SetUnbindMobileFlagRequestBuilder {
	s := &SetUnbindMobileFlagRequest{}
	b := &SetUnbindMobileFlagRequestBuilder{s: s}
	return b
}

func (b *SetUnbindMobileFlagRequestBuilder) SetUnbindMobileFlagBody(v *SetUnbindMobileFlagBody) *SetUnbindMobileFlagRequestBuilder {
	b.s.SetUnbindMobileFlagBody = v
	return b
}

func (b *SetUnbindMobileFlagRequestBuilder) Build() *SetUnbindMobileFlagRequest {
	return b.s
}


