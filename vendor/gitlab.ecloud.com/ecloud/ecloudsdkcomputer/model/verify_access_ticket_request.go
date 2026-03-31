// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type VerifyAccessTicketRequest struct {


	VerifyAccessTicketBody *VerifyAccessTicketBody `json:"verifyAccessTicketBody,omitempty"`
}

func (s VerifyAccessTicketRequest) String() string {
	return utils.Beautify(s)
}

func (s VerifyAccessTicketRequest) GoString() string {
	return s.String()
}

func (s VerifyAccessTicketRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *VerifyAccessTicketRequest) SetVerifyAccessTicketBody(v *VerifyAccessTicketBody) *VerifyAccessTicketRequest {
	s.VerifyAccessTicketBody = v
	return s
}


type VerifyAccessTicketRequestBuilder struct {
	s *VerifyAccessTicketRequest
}

func NewVerifyAccessTicketRequestBuilder() *VerifyAccessTicketRequestBuilder {
	s := &VerifyAccessTicketRequest{}
	b := &VerifyAccessTicketRequestBuilder{s: s}
	return b
}

func (b *VerifyAccessTicketRequestBuilder) VerifyAccessTicketBody(v *VerifyAccessTicketBody) *VerifyAccessTicketRequestBuilder {
	b.s.VerifyAccessTicketBody = v
	return b
}

func (b *VerifyAccessTicketRequestBuilder) Build() *VerifyAccessTicketRequest {
	return b.s
}


