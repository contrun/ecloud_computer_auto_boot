// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteSecurityGroupBody struct {
    position.Body
    // 安全组UID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
}

func (s DeleteSecurityGroupBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteSecurityGroupBody) GoString() string {
	return s.String()
}

func (s DeleteSecurityGroupBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSecurityGroupBody) SetSecurityGroupUid(v string) *DeleteSecurityGroupBody {
	s.SecurityGroupUid = &v
	return s
}


type DeleteSecurityGroupBodyBuilder struct {
	s *DeleteSecurityGroupBody
}

func NewDeleteSecurityGroupBodyBuilder() *DeleteSecurityGroupBodyBuilder {
	s := &DeleteSecurityGroupBody{}
	b := &DeleteSecurityGroupBodyBuilder{s: s}
	return b
}

func (b *DeleteSecurityGroupBodyBuilder) SecurityGroupUid(v string) *DeleteSecurityGroupBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *DeleteSecurityGroupBodyBuilder) Build() *DeleteSecurityGroupBody {
	return b.s
}


