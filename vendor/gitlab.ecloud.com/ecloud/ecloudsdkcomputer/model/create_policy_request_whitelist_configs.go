// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreatePolicyRequestWhitelistConfigs struct {

    // 白名单规则: 白名单模式下必填
	WhitelistRule *string `json:"whitelistRule,omitempty"`
    // 白名单类型: 白名单模式下必填, 域名:domainName 进程:process
	WhitelistType *string `json:"whitelistType,omitempty"`
}

func (s CreatePolicyRequestWhitelistConfigs) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestWhitelistConfigs) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestWhitelistConfigs) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestWhitelistConfigs) SetWhitelistRule(v string) *CreatePolicyRequestWhitelistConfigs {
	s.WhitelistRule = &v
	return s
}

func (s *CreatePolicyRequestWhitelistConfigs) SetWhitelistType(v string) *CreatePolicyRequestWhitelistConfigs {
	s.WhitelistType = &v
	return s
}


type CreatePolicyRequestWhitelistConfigsBuilder struct {
	s *CreatePolicyRequestWhitelistConfigs
}

func NewCreatePolicyRequestWhitelistConfigsBuilder() *CreatePolicyRequestWhitelistConfigsBuilder {
	s := &CreatePolicyRequestWhitelistConfigs{}
	b := &CreatePolicyRequestWhitelistConfigsBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestWhitelistConfigsBuilder) WhitelistRule(v string) *CreatePolicyRequestWhitelistConfigsBuilder {
	b.s.WhitelistRule = &v
	return b
}

func (b *CreatePolicyRequestWhitelistConfigsBuilder) WhitelistType(v string) *CreatePolicyRequestWhitelistConfigsBuilder {
	b.s.WhitelistType = &v
	return b
}

func (b *CreatePolicyRequestWhitelistConfigsBuilder) Build() *CreatePolicyRequestWhitelistConfigs {
	return b.s
}


