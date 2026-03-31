// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type VerifyAccessTicketResponseBody struct {

    // accessToken
	AccessToken *string `json:"accessToken,omitempty"`
}

func (s VerifyAccessTicketResponseBody) String() string {
	return utils.Beautify(s)
}

func (s VerifyAccessTicketResponseBody) GoString() string {
	return s.String()
}

func (s VerifyAccessTicketResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *VerifyAccessTicketResponseBody) SetAccessToken(v string) *VerifyAccessTicketResponseBody {
	s.AccessToken = &v
	return s
}


type VerifyAccessTicketResponseBodyBuilder struct {
	s *VerifyAccessTicketResponseBody
}

func NewVerifyAccessTicketResponseBodyBuilder() *VerifyAccessTicketResponseBodyBuilder {
	s := &VerifyAccessTicketResponseBody{}
	b := &VerifyAccessTicketResponseBodyBuilder{s: s}
	return b
}

func (b *VerifyAccessTicketResponseBodyBuilder) AccessToken(v string) *VerifyAccessTicketResponseBodyBuilder {
	b.s.AccessToken = &v
	return b
}

func (b *VerifyAccessTicketResponseBodyBuilder) Build() *VerifyAccessTicketResponseBody {
	return b.s
}


