// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteOrgBody struct {
    position.Body
    // 删除组织id
	OrgId *string `json:"orgId,omitempty"`
}

func (s DeleteOrgBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteOrgBody) GoString() string {
	return s.String()
}

func (s DeleteOrgBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteOrgBody) SetOrgId(v string) *DeleteOrgBody {
	s.OrgId = &v
	return s
}


type DeleteOrgBodyBuilder struct {
	s *DeleteOrgBody
}

func NewDeleteOrgBodyBuilder() *DeleteOrgBodyBuilder {
	s := &DeleteOrgBody{}
	b := &DeleteOrgBodyBuilder{s: s}
	return b
}

func (b *DeleteOrgBodyBuilder) OrgId(v string) *DeleteOrgBodyBuilder {
	b.s.OrgId = &v
	return b
}

func (b *DeleteOrgBodyBuilder) Build() *DeleteOrgBody {
	return b.s
}


