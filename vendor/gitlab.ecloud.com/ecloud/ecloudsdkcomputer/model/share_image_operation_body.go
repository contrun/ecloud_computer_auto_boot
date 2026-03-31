// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ShareImageOperationBody struct {
    position.Body
    // 接受者租户表主键IDs
	AcceptorTenantIds []string `json:"acceptorTenantIds,omitempty"`
    // 镜像模板ID
	TemplateIds []string `json:"templateIds,omitempty"`
}

func (s ShareImageOperationBody) String() string {
	return utils.Beautify(s)
}

func (s ShareImageOperationBody) GoString() string {
	return s.String()
}

func (s ShareImageOperationBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ShareImageOperationBody) SetAcceptorTenantIds(v []string) *ShareImageOperationBody {
	s.AcceptorTenantIds = v
	return s
}

func (s *ShareImageOperationBody) SetTemplateIds(v []string) *ShareImageOperationBody {
	s.TemplateIds = v
	return s
}


type ShareImageOperationBodyBuilder struct {
	s *ShareImageOperationBody
}

func NewShareImageOperationBodyBuilder() *ShareImageOperationBodyBuilder {
	s := &ShareImageOperationBody{}
	b := &ShareImageOperationBodyBuilder{s: s}
	return b
}

func (b *ShareImageOperationBodyBuilder) AcceptorTenantIds(v []string) *ShareImageOperationBodyBuilder {
	b.s.AcceptorTenantIds = v
	return b
}

func (b *ShareImageOperationBodyBuilder) TemplateIds(v []string) *ShareImageOperationBodyBuilder {
	b.s.TemplateIds = v
	return b
}

func (b *ShareImageOperationBodyBuilder) Build() *ShareImageOperationBody {
	return b.s
}


