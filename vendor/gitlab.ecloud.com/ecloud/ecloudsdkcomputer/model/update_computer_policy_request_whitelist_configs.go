// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateComputerPolicyRequestWhitelistConfigs struct {

    // 白名单规则: 白名单模式下必填
	WhitelistRule *string `json:"whitelistRule,omitempty"`
    // 白名单类型: 白名单模式下必填, 域名:domainName 进程:process
	WhitelistType *string `json:"whitelistType,omitempty"`
}

func (s UpdateComputerPolicyRequestWhitelistConfigs) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequestWhitelistConfigs) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequestWhitelistConfigs) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequestWhitelistConfigs) SetWhitelistRule(v string) *UpdateComputerPolicyRequestWhitelistConfigs {
	s.WhitelistRule = &v
	return s
}

func (s *UpdateComputerPolicyRequestWhitelistConfigs) SetWhitelistType(v string) *UpdateComputerPolicyRequestWhitelistConfigs {
	s.WhitelistType = &v
	return s
}


type UpdateComputerPolicyRequestWhitelistConfigsBuilder struct {
	s *UpdateComputerPolicyRequestWhitelistConfigs
}

func NewUpdateComputerPolicyRequestWhitelistConfigsBuilder() *UpdateComputerPolicyRequestWhitelistConfigsBuilder {
	s := &UpdateComputerPolicyRequestWhitelistConfigs{}
	b := &UpdateComputerPolicyRequestWhitelistConfigsBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestWhitelistConfigsBuilder) WhitelistRule(v string) *UpdateComputerPolicyRequestWhitelistConfigsBuilder {
	b.s.WhitelistRule = &v
	return b
}

func (b *UpdateComputerPolicyRequestWhitelistConfigsBuilder) WhitelistType(v string) *UpdateComputerPolicyRequestWhitelistConfigsBuilder {
	b.s.WhitelistType = &v
	return b
}

func (b *UpdateComputerPolicyRequestWhitelistConfigsBuilder) Build() *UpdateComputerPolicyRequestWhitelistConfigs {
	return b.s
}


