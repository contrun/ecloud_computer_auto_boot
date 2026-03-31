// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass struct {

    // 自定义壁纸绝对路径
	CustomWallpaperAbsolutePath *string `json:"customWallpaperAbsolutePath,omitempty"`
    // 自定义壁纸相对路径
	CustomWallpaperRelativePath *string `json:"customWallpaperRelativePath,omitempty"`
    // 自定义壁纸开关 开启 enable 关闭 disable，默认 disable
	CustomWallpaperEnable *string `json:"customWallpaperEnable,omitempty"`
}

func (s UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass) SetCustomWallpaperAbsolutePath(v string) *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass {
	s.CustomWallpaperAbsolutePath = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass) SetCustomWallpaperRelativePath(v string) *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass {
	s.CustomWallpaperRelativePath = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass) SetCustomWallpaperEnable(v string) *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass {
	s.CustomWallpaperEnable = &v
	return s
}


type UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder struct {
	s *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass
}

func NewUpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder() *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder {
	s := &UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass{}
	b := &UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder) CustomWallpaperAbsolutePath(v string) *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder {
	b.s.CustomWallpaperAbsolutePath = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder) CustomWallpaperRelativePath(v string) *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder {
	b.s.CustomWallpaperRelativePath = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder) CustomWallpaperEnable(v string) *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder {
	b.s.CustomWallpaperEnable = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClassBuilder) Build() *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass {
	return b.s
}


