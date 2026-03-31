// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CopySecurityGroupBody struct {
    position.Body
    // 创建信息
	CreateSecurityGroupInfo *CopySecurityGroupRequestCreateSecurityGroupInfo `json:"createSecurityGroupInfo,omitempty"`
    // 安全组UUID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
}

func (s CopySecurityGroupBody) String() string {
	return utils.Beautify(s)
}

func (s CopySecurityGroupBody) GoString() string {
	return s.String()
}

func (s CopySecurityGroupBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CopySecurityGroupBody) SetCreateSecurityGroupInfo(v *CopySecurityGroupRequestCreateSecurityGroupInfo) *CopySecurityGroupBody {
	s.CreateSecurityGroupInfo = v
	return s
}

func (s *CopySecurityGroupBody) SetSecurityGroupUid(v string) *CopySecurityGroupBody {
	s.SecurityGroupUid = &v
	return s
}


type CopySecurityGroupBodyBuilder struct {
	s *CopySecurityGroupBody
}

func NewCopySecurityGroupBodyBuilder() *CopySecurityGroupBodyBuilder {
	s := &CopySecurityGroupBody{}
	b := &CopySecurityGroupBodyBuilder{s: s}
	return b
}

func (b *CopySecurityGroupBodyBuilder) CreateSecurityGroupInfo(v *CopySecurityGroupRequestCreateSecurityGroupInfo) *CopySecurityGroupBodyBuilder {
	b.s.CreateSecurityGroupInfo = v
	return b
}

func (b *CopySecurityGroupBodyBuilder) SecurityGroupUid(v string) *CopySecurityGroupBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *CopySecurityGroupBodyBuilder) Build() *CopySecurityGroupBody {
	return b.s
}


