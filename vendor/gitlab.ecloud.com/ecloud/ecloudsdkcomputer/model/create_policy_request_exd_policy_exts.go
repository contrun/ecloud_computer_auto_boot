// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreatePolicyRequestExdPolicyExts struct {

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

func (s CreatePolicyRequestExdPolicyExts) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestExdPolicyExts) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestExdPolicyExts) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestExdPolicyExts) SetAction(v string) *CreatePolicyRequestExdPolicyExts {
	s.Action = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyExts) SetVidFilter(v string) *CreatePolicyRequestExdPolicyExts {
	s.VidFilter = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyExts) SetPidFilter(v string) *CreatePolicyRequestExdPolicyExts {
	s.PidFilter = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyExts) SetClassFilter(v string) *CreatePolicyRequestExdPolicyExts {
	s.ClassFilter = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyExts) SetExtRemark(v string) *CreatePolicyRequestExdPolicyExts {
	s.ExtRemark = &v
	return s
}


type CreatePolicyRequestExdPolicyExtsBuilder struct {
	s *CreatePolicyRequestExdPolicyExts
}

func NewCreatePolicyRequestExdPolicyExtsBuilder() *CreatePolicyRequestExdPolicyExtsBuilder {
	s := &CreatePolicyRequestExdPolicyExts{}
	b := &CreatePolicyRequestExdPolicyExtsBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestExdPolicyExtsBuilder) Action(v string) *CreatePolicyRequestExdPolicyExtsBuilder {
	b.s.Action = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyExtsBuilder) VidFilter(v string) *CreatePolicyRequestExdPolicyExtsBuilder {
	b.s.VidFilter = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyExtsBuilder) PidFilter(v string) *CreatePolicyRequestExdPolicyExtsBuilder {
	b.s.PidFilter = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyExtsBuilder) ClassFilter(v string) *CreatePolicyRequestExdPolicyExtsBuilder {
	b.s.ClassFilter = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyExtsBuilder) ExtRemark(v string) *CreatePolicyRequestExdPolicyExtsBuilder {
	b.s.ExtRemark = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyExtsBuilder) Build() *CreatePolicyRequestExdPolicyExts {
	return b.s
}


