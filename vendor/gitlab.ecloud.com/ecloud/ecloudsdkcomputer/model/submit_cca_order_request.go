// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitCcaOrderRequest struct {


	SubmitCcaOrderBody *SubmitCcaOrderBody `json:"submitCcaOrderBody,omitempty"`
}

func (s SubmitCcaOrderRequest) String() string {
	return utils.Beautify(s)
}

func (s SubmitCcaOrderRequest) GoString() string {
	return s.String()
}

func (s SubmitCcaOrderRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitCcaOrderRequest) SetSubmitCcaOrderBody(v *SubmitCcaOrderBody) *SubmitCcaOrderRequest {
	s.SubmitCcaOrderBody = v
	return s
}


type SubmitCcaOrderRequestBuilder struct {
	s *SubmitCcaOrderRequest
}

func NewSubmitCcaOrderRequestBuilder() *SubmitCcaOrderRequestBuilder {
	s := &SubmitCcaOrderRequest{}
	b := &SubmitCcaOrderRequestBuilder{s: s}
	return b
}

func (b *SubmitCcaOrderRequestBuilder) SubmitCcaOrderBody(v *SubmitCcaOrderBody) *SubmitCcaOrderRequestBuilder {
	b.s.SubmitCcaOrderBody = v
	return b
}

func (b *SubmitCcaOrderRequestBuilder) Build() *SubmitCcaOrderRequest {
	return b.s
}


