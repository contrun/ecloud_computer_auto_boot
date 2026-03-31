// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetClientInfoResponseBody struct {

    // accessTicket
	AccessTicket *string `json:"accessTicket,omitempty"`
}

func (s GetClientInfoResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetClientInfoResponseBody) GoString() string {
	return s.String()
}

func (s GetClientInfoResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetClientInfoResponseBody) SetAccessTicket(v string) *GetClientInfoResponseBody {
	s.AccessTicket = &v
	return s
}


type GetClientInfoResponseBodyBuilder struct {
	s *GetClientInfoResponseBody
}

func NewGetClientInfoResponseBodyBuilder() *GetClientInfoResponseBodyBuilder {
	s := &GetClientInfoResponseBody{}
	b := &GetClientInfoResponseBodyBuilder{s: s}
	return b
}

func (b *GetClientInfoResponseBodyBuilder) AccessTicket(v string) *GetClientInfoResponseBodyBuilder {
	b.s.AccessTicket = &v
	return b
}

func (b *GetClientInfoResponseBodyBuilder) Build() *GetClientInfoResponseBody {
	return b.s
}


