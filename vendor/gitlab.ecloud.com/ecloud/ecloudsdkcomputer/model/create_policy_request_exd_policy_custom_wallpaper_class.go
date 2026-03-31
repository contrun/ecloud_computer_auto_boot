// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreatePolicyRequestExdPolicyCustomWallpaperClass struct {

    // 自定义壁纸绝对路径
	CustomWallpaperAbsolutePath *string `json:"customWallpaperAbsolutePath,omitempty"`
    // 自定义壁纸相对路径
	CustomWallpaperRelativePath *string `json:"customWallpaperRelativePath,omitempty"`
    // 自定义壁纸开关 开启 enable 关闭 disable，默认 disable
	CustomWallpaperEnable *string `json:"customWallpaperEnable,omitempty"`
}

func (s CreatePolicyRequestExdPolicyCustomWallpaperClass) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestExdPolicyCustomWallpaperClass) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestExdPolicyCustomWallpaperClass) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestExdPolicyCustomWallpaperClass) SetCustomWallpaperAbsolutePath(v string) *CreatePolicyRequestExdPolicyCustomWallpaperClass {
	s.CustomWallpaperAbsolutePath = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyCustomWallpaperClass) SetCustomWallpaperRelativePath(v string) *CreatePolicyRequestExdPolicyCustomWallpaperClass {
	s.CustomWallpaperRelativePath = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyCustomWallpaperClass) SetCustomWallpaperEnable(v string) *CreatePolicyRequestExdPolicyCustomWallpaperClass {
	s.CustomWallpaperEnable = &v
	return s
}


type CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder struct {
	s *CreatePolicyRequestExdPolicyCustomWallpaperClass
}

func NewCreatePolicyRequestExdPolicyCustomWallpaperClassBuilder() *CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder {
	s := &CreatePolicyRequestExdPolicyCustomWallpaperClass{}
	b := &CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder) CustomWallpaperAbsolutePath(v string) *CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder {
	b.s.CustomWallpaperAbsolutePath = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder) CustomWallpaperRelativePath(v string) *CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder {
	b.s.CustomWallpaperRelativePath = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder) CustomWallpaperEnable(v string) *CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder {
	b.s.CustomWallpaperEnable = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyCustomWallpaperClassBuilder) Build() *CreatePolicyRequestExdPolicyCustomWallpaperClass {
	return b.s
}


