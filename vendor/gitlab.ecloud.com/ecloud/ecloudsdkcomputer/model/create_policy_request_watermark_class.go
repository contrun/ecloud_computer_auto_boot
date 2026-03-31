// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreatePolicyRequestWatermarkClass struct {

    // 是否显示时间
	WatermarkTimeEnable *string `json:"watermarkTimeEnable,omitempty"`
    // 水印类型，明水印(0)、盲水印（1） 默认 明水印(0)
	WatermarkType *string `json:"watermarkType,omitempty"`
    // 是否开启水印，关闭(disable)、开启（enable）
	WatermarkEnable *string `json:"watermarkEnable,omitempty"`
    // 字体, 幼圆:YouYuan, 宋体:SimSun, 楷体:KaiTi
	WatermarkFont *string `json:"watermarkFont,omitempty"`
    // 水印透明度,新华三中兴 （H3C、H3CEcloud）可设置值：50、100，自研和中兴（CMSSZTE、ZTE、ZTEEcloud）可设置值：1~100 0就是透明，100就是不透明
	WatermarkTransparency *int32 `json:"watermarkTransparency,omitempty"`
    // 水印倾斜度,枚举0，45
	WatermarkAngle *int32 `json:"watermarkAngle,omitempty"`
    // 是否显示自定义字段
	WatermarkCustomFieldEnable *string `json:"watermarkCustomFieldEnable,omitempty"`
    // 是否显示云电脑ID
	WatermarkMachineIdEnable *string `json:"watermarkMachineIdEnable,omitempty"`
    // 是否显示云电脑mac
	WatermarkMachineMacEnable *string `json:"watermarkMachineMacEnable,omitempty"`
    // 是否显示云电脑IP
	WatermarkMachineIpEnable *string `json:"watermarkMachineIpEnable,omitempty"`
    // 自定义字段
	WatermarkCustomField *string `json:"watermarkCustomField,omitempty"`
    // 水印字体颜色 #RRGGBB ,每个颜色用2个字节标识,范围<#000000-#ffffff>
	WatermarkColor *string `json:"watermarkColor,omitempty"`
    // 水印字体大小，单位px,范围<1-60>
	WatermarkSize *int32 `json:"watermarkSize,omitempty"`
}

func (s CreatePolicyRequestWatermarkClass) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestWatermarkClass) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestWatermarkClass) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkTimeEnable(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkTimeEnable = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkType(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkType = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkEnable(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkEnable = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkFont(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkFont = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkTransparency(v int32) *CreatePolicyRequestWatermarkClass {
	s.WatermarkTransparency = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkAngle(v int32) *CreatePolicyRequestWatermarkClass {
	s.WatermarkAngle = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkCustomFieldEnable(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkCustomFieldEnable = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkMachineIdEnable(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkMachineIdEnable = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkMachineMacEnable(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkMachineMacEnable = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkMachineIpEnable(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkMachineIpEnable = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkCustomField(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkCustomField = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkColor(v string) *CreatePolicyRequestWatermarkClass {
	s.WatermarkColor = &v
	return s
}

func (s *CreatePolicyRequestWatermarkClass) SetWatermarkSize(v int32) *CreatePolicyRequestWatermarkClass {
	s.WatermarkSize = &v
	return s
}


type CreatePolicyRequestWatermarkClassBuilder struct {
	s *CreatePolicyRequestWatermarkClass
}

func NewCreatePolicyRequestWatermarkClassBuilder() *CreatePolicyRequestWatermarkClassBuilder {
	s := &CreatePolicyRequestWatermarkClass{}
	b := &CreatePolicyRequestWatermarkClassBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkTimeEnable(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkTimeEnable = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkType(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkType = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkEnable(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkEnable = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkFont(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkFont = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkTransparency(v int32) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkTransparency = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkAngle(v int32) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkAngle = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkCustomFieldEnable(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkCustomFieldEnable = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkMachineIdEnable(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkMachineIdEnable = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkMachineMacEnable(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkMachineMacEnable = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkMachineIpEnable(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkMachineIpEnable = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkCustomField(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkCustomField = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkColor(v string) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkColor = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) WatermarkSize(v int32) *CreatePolicyRequestWatermarkClassBuilder {
	b.s.WatermarkSize = &v
	return b
}

func (b *CreatePolicyRequestWatermarkClassBuilder) Build() *CreatePolicyRequestWatermarkClass {
	return b.s
}


