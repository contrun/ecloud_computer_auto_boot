// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyContactInfoBody struct {
    position.Body
    // 联系人
	ContactName *string `json:"contactName,omitempty"`
    // 联系方式
	ContactPhone *string `json:"contactPhone,omitempty"`
}

func (s ModifyContactInfoBody) String() string {
	return utils.Beautify(s)
}

func (s ModifyContactInfoBody) GoString() string {
	return s.String()
}

func (s ModifyContactInfoBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyContactInfoBody) SetContactName(v string) *ModifyContactInfoBody {
	s.ContactName = &v
	return s
}

func (s *ModifyContactInfoBody) SetContactPhone(v string) *ModifyContactInfoBody {
	s.ContactPhone = &v
	return s
}


type ModifyContactInfoBodyBuilder struct {
	s *ModifyContactInfoBody
}

func NewModifyContactInfoBodyBuilder() *ModifyContactInfoBodyBuilder {
	s := &ModifyContactInfoBody{}
	b := &ModifyContactInfoBodyBuilder{s: s}
	return b
}

func (b *ModifyContactInfoBodyBuilder) ContactName(v string) *ModifyContactInfoBodyBuilder {
	b.s.ContactName = &v
	return b
}

func (b *ModifyContactInfoBodyBuilder) ContactPhone(v string) *ModifyContactInfoBodyBuilder {
	b.s.ContactPhone = &v
	return b
}

func (b *ModifyContactInfoBodyBuilder) Build() *ModifyContactInfoBody {
	return b.s
}


