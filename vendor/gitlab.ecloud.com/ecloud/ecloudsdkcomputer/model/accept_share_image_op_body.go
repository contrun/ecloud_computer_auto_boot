// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AcceptShareImageOpBody struct {
    position.Body
    // 发起者租户表主键ID
	InitiatorTenantId *string `json:"initiatorTenantId,omitempty"`
    // 镜像UID
	TemplateId *string `json:"templateId,omitempty"`
    // 操作状态,未处理：unAccepted，已接收：accepted,已拒绝：rejected
	OpStatus *string `json:"opStatus,omitempty"`
}

func (s AcceptShareImageOpBody) String() string {
	return utils.Beautify(s)
}

func (s AcceptShareImageOpBody) GoString() string {
	return s.String()
}

func (s AcceptShareImageOpBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AcceptShareImageOpBody) SetInitiatorTenantId(v string) *AcceptShareImageOpBody {
	s.InitiatorTenantId = &v
	return s
}

func (s *AcceptShareImageOpBody) SetTemplateId(v string) *AcceptShareImageOpBody {
	s.TemplateId = &v
	return s
}

func (s *AcceptShareImageOpBody) SetOpStatus(v string) *AcceptShareImageOpBody {
	s.OpStatus = &v
	return s
}


type AcceptShareImageOpBodyBuilder struct {
	s *AcceptShareImageOpBody
}

func NewAcceptShareImageOpBodyBuilder() *AcceptShareImageOpBodyBuilder {
	s := &AcceptShareImageOpBody{}
	b := &AcceptShareImageOpBodyBuilder{s: s}
	return b
}

func (b *AcceptShareImageOpBodyBuilder) InitiatorTenantId(v string) *AcceptShareImageOpBodyBuilder {
	b.s.InitiatorTenantId = &v
	return b
}

func (b *AcceptShareImageOpBodyBuilder) TemplateId(v string) *AcceptShareImageOpBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *AcceptShareImageOpBodyBuilder) OpStatus(v string) *AcceptShareImageOpBodyBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *AcceptShareImageOpBodyBuilder) Build() *AcceptShareImageOpBody {
	return b.s
}


