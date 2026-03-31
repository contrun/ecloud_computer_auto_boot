// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type VerifyAccessTicketBody struct {
    position.Body
    // accessTicket
	AccessTicket *string `json:"accessTicket,omitempty"`
}

func (s VerifyAccessTicketBody) String() string {
	return utils.Beautify(s)
}

func (s VerifyAccessTicketBody) GoString() string {
	return s.String()
}

func (s VerifyAccessTicketBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *VerifyAccessTicketBody) SetAccessTicket(v string) *VerifyAccessTicketBody {
	s.AccessTicket = &v
	return s
}


type VerifyAccessTicketBodyBuilder struct {
	s *VerifyAccessTicketBody
}

func NewVerifyAccessTicketBodyBuilder() *VerifyAccessTicketBodyBuilder {
	s := &VerifyAccessTicketBody{}
	b := &VerifyAccessTicketBodyBuilder{s: s}
	return b
}

func (b *VerifyAccessTicketBodyBuilder) AccessTicket(v string) *VerifyAccessTicketBodyBuilder {
	b.s.AccessTicket = &v
	return b
}

func (b *VerifyAccessTicketBodyBuilder) Build() *VerifyAccessTicketBody {
	return b.s
}


