// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreatePolicyRequestExdCloudSideNetRedirectClass struct {

    // 重定向模式: 开关开启状态下必填, 白名单模式:whitelistMode 全局模式:globalMode
	NetRedirectMode *string `json:"netRedirectMode,omitempty"`
    // 白名单类型: 白名单模式下配置
	WhitelistConfigs *[]CreatePolicyRequestWhitelistConfigs `json:"whitelistConfigs,omitempty"`
    // 网络重定向开关: 类非空情况下必填  关闭(disable)、开启(enable)
	NetRedirectEnable *string `json:"netRedirectEnable,omitempty"`
}

func (s CreatePolicyRequestExdCloudSideNetRedirectClass) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestExdCloudSideNetRedirectClass) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestExdCloudSideNetRedirectClass) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestExdCloudSideNetRedirectClass) SetNetRedirectMode(v string) *CreatePolicyRequestExdCloudSideNetRedirectClass {
	s.NetRedirectMode = &v
	return s
}

func (s *CreatePolicyRequestExdCloudSideNetRedirectClass) SetWhitelistConfigs(v []CreatePolicyRequestWhitelistConfigs) *CreatePolicyRequestExdCloudSideNetRedirectClass {
	s.WhitelistConfigs = &v
	return s
}

func (s *CreatePolicyRequestExdCloudSideNetRedirectClass) SetNetRedirectEnable(v string) *CreatePolicyRequestExdCloudSideNetRedirectClass {
	s.NetRedirectEnable = &v
	return s
}


type CreatePolicyRequestExdCloudSideNetRedirectClassBuilder struct {
	s *CreatePolicyRequestExdCloudSideNetRedirectClass
}

func NewCreatePolicyRequestExdCloudSideNetRedirectClassBuilder() *CreatePolicyRequestExdCloudSideNetRedirectClassBuilder {
	s := &CreatePolicyRequestExdCloudSideNetRedirectClass{}
	b := &CreatePolicyRequestExdCloudSideNetRedirectClassBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestExdCloudSideNetRedirectClassBuilder) NetRedirectMode(v string) *CreatePolicyRequestExdCloudSideNetRedirectClassBuilder {
	b.s.NetRedirectMode = &v
	return b
}

func (b *CreatePolicyRequestExdCloudSideNetRedirectClassBuilder) WhitelistConfigs(v []CreatePolicyRequestWhitelistConfigs) *CreatePolicyRequestExdCloudSideNetRedirectClassBuilder {
	b.s.WhitelistConfigs = &v
	return b
}

func (b *CreatePolicyRequestExdCloudSideNetRedirectClassBuilder) NetRedirectEnable(v string) *CreatePolicyRequestExdCloudSideNetRedirectClassBuilder {
	b.s.NetRedirectEnable = &v
	return b
}

func (b *CreatePolicyRequestExdCloudSideNetRedirectClassBuilder) Build() *CreatePolicyRequestExdCloudSideNetRedirectClass {
	return b.s
}


