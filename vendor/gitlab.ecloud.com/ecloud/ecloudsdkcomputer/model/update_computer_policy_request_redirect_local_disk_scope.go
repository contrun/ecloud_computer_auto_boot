// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateComputerPolicyRequestRedirectLocalDiskScope struct {

    // 终端系统磁盘（即本地系统盘，通常指C盘） 取值：禁用(disable)、启用(enable)
	TerminalSystemDisk *string `json:"terminalSystemDisk,omitempty"`
    // 可移动驱动器 取值：禁用(disable)、启用(enable)
	RemovableDrive *string `json:"removableDrive,omitempty"`
    // 网络驱动器 取值：禁用(disable)、启用(enable)
	NetworkDrive *string `json:"networkDrive,omitempty"`
    // 重定向本地磁盘， 禁用(disable)、只读(ro)、读写(rw)
	RedirectLocalDisk *string `json:"redirectLocalDisk,omitempty"`
    // 光盘驱动器 取值：禁用(disable)、启用(enable)
	OpticalDiscDrive *string `json:"opticalDiscDrive,omitempty"`
    // 固定驱动器（即本地系统除系统盘外其他磁盘）取值：禁用(disable)、启用(enable)
	FixedDrive *string `json:"fixedDrive,omitempty"`
}

func (s UpdateComputerPolicyRequestRedirectLocalDiskScope) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequestRedirectLocalDiskScope) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequestRedirectLocalDiskScope) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequestRedirectLocalDiskScope) SetTerminalSystemDisk(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScope {
	s.TerminalSystemDisk = &v
	return s
}

func (s *UpdateComputerPolicyRequestRedirectLocalDiskScope) SetRemovableDrive(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScope {
	s.RemovableDrive = &v
	return s
}

func (s *UpdateComputerPolicyRequestRedirectLocalDiskScope) SetNetworkDrive(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScope {
	s.NetworkDrive = &v
	return s
}

func (s *UpdateComputerPolicyRequestRedirectLocalDiskScope) SetRedirectLocalDisk(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScope {
	s.RedirectLocalDisk = &v
	return s
}

func (s *UpdateComputerPolicyRequestRedirectLocalDiskScope) SetOpticalDiscDrive(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScope {
	s.OpticalDiscDrive = &v
	return s
}

func (s *UpdateComputerPolicyRequestRedirectLocalDiskScope) SetFixedDrive(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScope {
	s.FixedDrive = &v
	return s
}


type UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder struct {
	s *UpdateComputerPolicyRequestRedirectLocalDiskScope
}

func NewUpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder() *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder {
	s := &UpdateComputerPolicyRequestRedirectLocalDiskScope{}
	b := &UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder) TerminalSystemDisk(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.TerminalSystemDisk = &v
	return b
}

func (b *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder) RemovableDrive(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.RemovableDrive = &v
	return b
}

func (b *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder) NetworkDrive(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.NetworkDrive = &v
	return b
}

func (b *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder) RedirectLocalDisk(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.RedirectLocalDisk = &v
	return b
}

func (b *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder) OpticalDiscDrive(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.OpticalDiscDrive = &v
	return b
}

func (b *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder) FixedDrive(v string) *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder {
	b.s.FixedDrive = &v
	return b
}

func (b *UpdateComputerPolicyRequestRedirectLocalDiskScopeBuilder) Build() *UpdateComputerPolicyRequestRedirectLocalDiskScope {
	return b.s
}


