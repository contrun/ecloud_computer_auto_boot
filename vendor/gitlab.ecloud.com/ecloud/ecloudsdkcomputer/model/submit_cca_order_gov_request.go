// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitCcaOrderGovRequest struct {


	SubmitCcaOrderGovBody *SubmitCcaOrderGovBody `json:"submitCcaOrderGovBody,omitempty"`
}

func (s SubmitCcaOrderGovRequest) String() string {
	return utils.Beautify(s)
}

func (s SubmitCcaOrderGovRequest) GoString() string {
	return s.String()
}

func (s SubmitCcaOrderGovRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitCcaOrderGovRequest) SetSubmitCcaOrderGovBody(v *SubmitCcaOrderGovBody) *SubmitCcaOrderGovRequest {
	s.SubmitCcaOrderGovBody = v
	return s
}


type SubmitCcaOrderGovRequestBuilder struct {
	s *SubmitCcaOrderGovRequest
}

func NewSubmitCcaOrderGovRequestBuilder() *SubmitCcaOrderGovRequestBuilder {
	s := &SubmitCcaOrderGovRequest{}
	b := &SubmitCcaOrderGovRequestBuilder{s: s}
	return b
}

func (b *SubmitCcaOrderGovRequestBuilder) SubmitCcaOrderGovBody(v *SubmitCcaOrderGovBody) *SubmitCcaOrderGovRequestBuilder {
	b.s.SubmitCcaOrderGovBody = v
	return b
}

func (b *SubmitCcaOrderGovRequestBuilder) Build() *SubmitCcaOrderGovRequest {
	return b.s
}


