// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateComputerPolicyRequestExdPolicyExts struct {

    // 黑白名单：白名单(enable), 黑名单(disable)
	Action *string `json:"action,omitempty"`
    // vid
	VidFilter *string `json:"vidFilter,omitempty"`
    // pid
	PidFilter *string `json:"pidFilter,omitempty"`
    // 扫描仪(scanner),打印机(printer),存储设备(storage),摄像头(camera),智能卡(smartCard),无线设备(wireless),其他设备(other),全部(all)
	ClassFilter *string `json:"classFilter,omitempty"`
    // 备注
	ExtRemark *string `json:"extRemark,omitempty"`
}

func (s UpdateComputerPolicyRequestExdPolicyExts) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequestExdPolicyExts) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequestExdPolicyExts) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequestExdPolicyExts) SetAction(v string) *UpdateComputerPolicyRequestExdPolicyExts {
	s.Action = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyExts) SetVidFilter(v string) *UpdateComputerPolicyRequestExdPolicyExts {
	s.VidFilter = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyExts) SetPidFilter(v string) *UpdateComputerPolicyRequestExdPolicyExts {
	s.PidFilter = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyExts) SetClassFilter(v string) *UpdateComputerPolicyRequestExdPolicyExts {
	s.ClassFilter = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyExts) SetExtRemark(v string) *UpdateComputerPolicyRequestExdPolicyExts {
	s.ExtRemark = &v
	return s
}


type UpdateComputerPolicyRequestExdPolicyExtsBuilder struct {
	s *UpdateComputerPolicyRequestExdPolicyExts
}

func NewUpdateComputerPolicyRequestExdPolicyExtsBuilder() *UpdateComputerPolicyRequestExdPolicyExtsBuilder {
	s := &UpdateComputerPolicyRequestExdPolicyExts{}
	b := &UpdateComputerPolicyRequestExdPolicyExtsBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyExtsBuilder) Action(v string) *UpdateComputerPolicyRequestExdPolicyExtsBuilder {
	b.s.Action = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyExtsBuilder) VidFilter(v string) *UpdateComputerPolicyRequestExdPolicyExtsBuilder {
	b.s.VidFilter = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyExtsBuilder) PidFilter(v string) *UpdateComputerPolicyRequestExdPolicyExtsBuilder {
	b.s.PidFilter = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyExtsBuilder) ClassFilter(v string) *UpdateComputerPolicyRequestExdPolicyExtsBuilder {
	b.s.ClassFilter = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyExtsBuilder) ExtRemark(v string) *UpdateComputerPolicyRequestExdPolicyExtsBuilder {
	b.s.ExtRemark = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyExtsBuilder) Build() *UpdateComputerPolicyRequestExdPolicyExts {
	return b.s
}


