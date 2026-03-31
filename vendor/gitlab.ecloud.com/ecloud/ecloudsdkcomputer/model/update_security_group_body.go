// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateSecurityGroupBody struct {
    position.Body
    // 安全组名
	SecurityGroupName *string `json:"securityGroupName,omitempty"`
    // 安全组描述
	Description *string `json:"description,omitempty"`
    // 安全组UID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
}

func (s UpdateSecurityGroupBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateSecurityGroupBody) GoString() string {
	return s.String()
}

func (s UpdateSecurityGroupBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSecurityGroupBody) SetSecurityGroupName(v string) *UpdateSecurityGroupBody {
	s.SecurityGroupName = &v
	return s
}

func (s *UpdateSecurityGroupBody) SetDescription(v string) *UpdateSecurityGroupBody {
	s.Description = &v
	return s
}

func (s *UpdateSecurityGroupBody) SetSecurityGroupUid(v string) *UpdateSecurityGroupBody {
	s.SecurityGroupUid = &v
	return s
}


type UpdateSecurityGroupBodyBuilder struct {
	s *UpdateSecurityGroupBody
}

func NewUpdateSecurityGroupBodyBuilder() *UpdateSecurityGroupBodyBuilder {
	s := &UpdateSecurityGroupBody{}
	b := &UpdateSecurityGroupBodyBuilder{s: s}
	return b
}

func (b *UpdateSecurityGroupBodyBuilder) SecurityGroupName(v string) *UpdateSecurityGroupBodyBuilder {
	b.s.SecurityGroupName = &v
	return b
}

func (b *UpdateSecurityGroupBodyBuilder) Description(v string) *UpdateSecurityGroupBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *UpdateSecurityGroupBodyBuilder) SecurityGroupUid(v string) *UpdateSecurityGroupBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *UpdateSecurityGroupBodyBuilder) Build() *UpdateSecurityGroupBody {
	return b.s
}


