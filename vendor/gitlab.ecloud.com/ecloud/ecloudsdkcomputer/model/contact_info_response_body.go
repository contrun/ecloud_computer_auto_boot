// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ContactInfoResponseBody struct {

    // 联系人
	ContactName *string `json:"contactName,omitempty"`
    // 联系方式
	ContactPhone *string `json:"contactPhone,omitempty"`
}

func (s ContactInfoResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s ContactInfoResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ContactInfoResponseBody) SetContactName(v string) *ContactInfoResponseBody {
	s.ContactName = &v
	return s
}

func (s *ContactInfoResponseBody) SetContactPhone(v string) *ContactInfoResponseBody {
	s.ContactPhone = &v
	return s
}


type ContactInfoResponseBodyBuilder struct {
	s *ContactInfoResponseBody
}

func NewContactInfoResponseBodyBuilder() *ContactInfoResponseBodyBuilder {
	s := &ContactInfoResponseBody{}
	b := &ContactInfoResponseBodyBuilder{s: s}
	return b
}

func (b *ContactInfoResponseBodyBuilder) ContactName(v string) *ContactInfoResponseBodyBuilder {
	b.s.ContactName = &v
	return b
}

func (b *ContactInfoResponseBodyBuilder) ContactPhone(v string) *ContactInfoResponseBodyBuilder {
	b.s.ContactPhone = &v
	return b
}

func (b *ContactInfoResponseBodyBuilder) Build() *ContactInfoResponseBody {
	return b.s
}


