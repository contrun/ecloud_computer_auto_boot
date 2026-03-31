// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos struct {

    // 策略ID (非标准uuid)
	PolicyId *string `json:"policyId,omitempty"`
    // 白名单规则
	WhitelistRule *string `json:"whitelistRule,omitempty"`
    // 白名单类型(域名:domainName 进程:process)
	WhitelistType *string `json:"whitelistType,omitempty"`
}

func (s GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos) SetPolicyId(v string) *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos) SetWhitelistRule(v string) *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos {
	s.WhitelistRule = &v
	return s
}

func (s *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos) SetWhitelistType(v string) *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos {
	s.WhitelistType = &v
	return s
}


type GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder struct {
	s *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos
}

func NewGetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder() *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder {
	s := &GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos{}
	b := &GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder) PolicyId(v string) *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder) WhitelistRule(v string) *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder {
	b.s.WhitelistRule = &v
	return b
}

func (b *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder) WhitelistType(v string) *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder {
	b.s.WhitelistType = &v
	return b
}

func (b *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtosBuilder) Build() *GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos {
	return b.s
}


