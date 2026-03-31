// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EditCompanyTagBody struct {
    position.Body
    // 企业ID
	CompanyTag *string `json:"companyTag,omitempty"`
}

func (s EditCompanyTagBody) String() string {
	return utils.Beautify(s)
}

func (s EditCompanyTagBody) GoString() string {
	return s.String()
}

func (s EditCompanyTagBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditCompanyTagBody) SetCompanyTag(v string) *EditCompanyTagBody {
	s.CompanyTag = &v
	return s
}


type EditCompanyTagBodyBuilder struct {
	s *EditCompanyTagBody
}

func NewEditCompanyTagBodyBuilder() *EditCompanyTagBodyBuilder {
	s := &EditCompanyTagBody{}
	b := &EditCompanyTagBodyBuilder{s: s}
	return b
}

func (b *EditCompanyTagBodyBuilder) CompanyTag(v string) *EditCompanyTagBodyBuilder {
	b.s.CompanyTag = &v
	return b
}

func (b *EditCompanyTagBodyBuilder) Build() *EditCompanyTagBody {
	return b.s
}


