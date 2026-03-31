// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateComputerPolicyRequestWatermarkClass struct {

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

func (s UpdateComputerPolicyRequestWatermarkClass) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequestWatermarkClass) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequestWatermarkClass) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkTimeEnable(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkTimeEnable = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkType(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkType = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkEnable(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkEnable = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkFont(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkFont = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkTransparency(v int32) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkTransparency = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkAngle(v int32) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkAngle = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkCustomFieldEnable(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkCustomFieldEnable = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkMachineIdEnable(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkMachineIdEnable = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkMachineMacEnable(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkMachineMacEnable = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkMachineIpEnable(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkMachineIpEnable = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkCustomField(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkCustomField = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkColor(v string) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkColor = &v
	return s
}

func (s *UpdateComputerPolicyRequestWatermarkClass) SetWatermarkSize(v int32) *UpdateComputerPolicyRequestWatermarkClass {
	s.WatermarkSize = &v
	return s
}


type UpdateComputerPolicyRequestWatermarkClassBuilder struct {
	s *UpdateComputerPolicyRequestWatermarkClass
}

func NewUpdateComputerPolicyRequestWatermarkClassBuilder() *UpdateComputerPolicyRequestWatermarkClassBuilder {
	s := &UpdateComputerPolicyRequestWatermarkClass{}
	b := &UpdateComputerPolicyRequestWatermarkClassBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkTimeEnable(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkTimeEnable = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkType(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkType = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkEnable(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkEnable = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkFont(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkFont = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkTransparency(v int32) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkTransparency = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkAngle(v int32) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkAngle = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkCustomFieldEnable(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkCustomFieldEnable = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkMachineIdEnable(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkMachineIdEnable = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkMachineMacEnable(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkMachineMacEnable = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkMachineIpEnable(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkMachineIpEnable = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkCustomField(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkCustomField = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkColor(v string) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkColor = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) WatermarkSize(v int32) *UpdateComputerPolicyRequestWatermarkClassBuilder {
	b.s.WatermarkSize = &v
	return b
}

func (b *UpdateComputerPolicyRequestWatermarkClassBuilder) Build() *UpdateComputerPolicyRequestWatermarkClass {
	return b.s
}


