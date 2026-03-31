// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateComputerPolicyRequestExdCloudSideNetRedirectClass struct {

    // 重定向模式: 开关开启状态下必填, 白名单模式:whitelistMode 全局模式:globalMode
	NetRedirectMode *string `json:"netRedirectMode,omitempty"`
    // 白名单类型: 白名单模式下配置
	WhitelistConfigs *[]UpdateComputerPolicyRequestWhitelistConfigs `json:"whitelistConfigs,omitempty"`
    // 网络重定向开关: 类非空情况下必填  关闭(disable)、开启(enable)
	NetRedirectEnable *string `json:"netRedirectEnable,omitempty"`
}

func (s UpdateComputerPolicyRequestExdCloudSideNetRedirectClass) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequestExdCloudSideNetRedirectClass) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequestExdCloudSideNetRedirectClass) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass) SetNetRedirectMode(v string) *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass {
	s.NetRedirectMode = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass) SetWhitelistConfigs(v []UpdateComputerPolicyRequestWhitelistConfigs) *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass {
	s.WhitelistConfigs = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass) SetNetRedirectEnable(v string) *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass {
	s.NetRedirectEnable = &v
	return s
}


type UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder struct {
	s *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass
}

func NewUpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder() *UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder {
	s := &UpdateComputerPolicyRequestExdCloudSideNetRedirectClass{}
	b := &UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder) NetRedirectMode(v string) *UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder {
	b.s.NetRedirectMode = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder) WhitelistConfigs(v []UpdateComputerPolicyRequestWhitelistConfigs) *UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder {
	b.s.WhitelistConfigs = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder) NetRedirectEnable(v string) *UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder {
	b.s.NetRedirectEnable = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdCloudSideNetRedirectClassBuilder) Build() *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass {
	return b.s
}


