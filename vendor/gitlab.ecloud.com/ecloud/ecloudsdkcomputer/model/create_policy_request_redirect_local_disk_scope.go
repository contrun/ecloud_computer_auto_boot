// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreatePolicyRequestRedirectLocalDiskScope struct {

    // 终端系统磁盘（即本地系统盘，通常指C盘） 取值：禁用(disable)、启用(enable)
	TerminalSystemDisk *string `json:"terminalSystemDisk,omitempty"`
    // 可移动驱动器 取值：禁用(disable)、启用(enable)
	RemovableDrive *string `json:"removableDrive,omitempty"`
    // 网络驱动器 取值：禁用(disable)、启用(enable)
	NetworkDrive *string `json:"networkDrive,omitempty"`
    // 重定向本地磁盘 默认读写， 禁用(disable)、只读(ro)、读写(rw)
	RedirectLocalDisk *string `json:"redirectLocalDisk,omitempty"`
    // 光盘驱动器 取值：禁用(disable)、启用(enable)
	OpticalDiscDrive *string `json:"opticalDiscDrive,omitempty"`
    // 固定驱动器（即本地系统除系统盘外其他磁盘）取值：禁用(disable)、启用(enable)
	FixedDrive *string `json:"fixedDrive,omitempty"`
}

func (s CreatePolicyRequestRedirectLocalDiskScope) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestRedirectLocalDiskScope) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestRedirectLocalDiskScope) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestRedirectLocalDiskScope) SetTerminalSystemDisk(v string) *CreatePolicyRequestRedirectLocalDiskScope {
	s.TerminalSystemDisk = &v
	return s
}

func (s *CreatePolicyRequestRedirectLocalDiskScope) SetRemovableDrive(v string) *CreatePolicyRequestRedirectLocalDiskScope {
	s.RemovableDrive = &v
	return s
}

func (s *CreatePolicyRequestRedirectLocalDiskScope) SetNetworkDrive(v string) *CreatePolicyRequestRedirectLocalDiskScope {
	s.NetworkDrive = &v
	return s
}

func (s *CreatePolicyRequestRedirectLocalDiskScope) SetRedirectLocalDisk(v string) *CreatePolicyRequestRedirectLocalDiskScope {
	s.RedirectLocalDisk = &v
	return s
}

func (s *CreatePolicyRequestRedirectLocalDiskScope) SetOpticalDiscDrive(v string) *CreatePolicyRequestRedirectLocalDiskScope {
	s.OpticalDiscDrive = &v
	return s
}

func (s *CreatePolicyRequestRedirectLocalDiskScope) SetFixedDrive(v string) *CreatePolicyRequestRedirectLocalDiskScope {
	s.FixedDrive = &v
	return s
}


type CreatePolicyRequestRedirectLocalDiskScopeBuilder struct {
	s *CreatePolicyRequestRedirectLocalDiskScope
}

func NewCreatePolicyRequestRedirectLocalDiskScopeBuilder() *CreatePolicyRequestRedirectLocalDiskScopeBuilder {
	s := &CreatePolicyRequestRedirectLocalDiskScope{}
	b := &CreatePolicyRequestRedirectLocalDiskScopeBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestRedirectLocalDiskScopeBuilder) TerminalSystemDisk(v string) *CreatePolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.TerminalSystemDisk = &v
	return b
}

func (b *CreatePolicyRequestRedirectLocalDiskScopeBuilder) RemovableDrive(v string) *CreatePolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.RemovableDrive = &v
	return b
}

func (b *CreatePolicyRequestRedirectLocalDiskScopeBuilder) NetworkDrive(v string) *CreatePolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.NetworkDrive = &v
	return b
}

func (b *CreatePolicyRequestRedirectLocalDiskScopeBuilder) RedirectLocalDisk(v string) *CreatePolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.RedirectLocalDisk = &v
	return b
}

func (b *CreatePolicyRequestRedirectLocalDiskScopeBuilder) OpticalDiscDrive(v string) *CreatePolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.OpticalDiscDrive = &v
	return b
}

func (b *CreatePolicyRequestRedirectLocalDiskScopeBuilder) FixedDrive(v string) *CreatePolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.FixedDrive = &v
	return b
}

func (b *CreatePolicyRequestRedirectLocalDiskScopeBuilder) Build() *CreatePolicyRequestRedirectLocalDiskScope {
	return b.s
}


