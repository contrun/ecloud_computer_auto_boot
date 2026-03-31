// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdResponseExdPolicyExts struct {

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

func (s GetPolicyInfoByIdResponseExdPolicyExts) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseExdPolicyExts) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseExdPolicyExts) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseExdPolicyExts) SetAction(v string) *GetPolicyInfoByIdResponseExdPolicyExts {
	s.Action = &v
	return s
}

func (s *GetPolicyInfoByIdResponseExdPolicyExts) SetVidFilter(v string) *GetPolicyInfoByIdResponseExdPolicyExts {
	s.VidFilter = &v
	return s
}

func (s *GetPolicyInfoByIdResponseExdPolicyExts) SetPidFilter(v string) *GetPolicyInfoByIdResponseExdPolicyExts {
	s.PidFilter = &v
	return s
}

func (s *GetPolicyInfoByIdResponseExdPolicyExts) SetClassFilter(v string) *GetPolicyInfoByIdResponseExdPolicyExts {
	s.ClassFilter = &v
	return s
}

func (s *GetPolicyInfoByIdResponseExdPolicyExts) SetExtRemark(v string) *GetPolicyInfoByIdResponseExdPolicyExts {
	s.ExtRemark = &v
	return s
}


type GetPolicyInfoByIdResponseExdPolicyExtsBuilder struct {
	s *GetPolicyInfoByIdResponseExdPolicyExts
}

func NewGetPolicyInfoByIdResponseExdPolicyExtsBuilder() *GetPolicyInfoByIdResponseExdPolicyExtsBuilder {
	s := &GetPolicyInfoByIdResponseExdPolicyExts{}
	b := &GetPolicyInfoByIdResponseExdPolicyExtsBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseExdPolicyExtsBuilder) Action(v string) *GetPolicyInfoByIdResponseExdPolicyExtsBuilder {
	b.s.Action = &v
	return b
}

func (b *GetPolicyInfoByIdResponseExdPolicyExtsBuilder) VidFilter(v string) *GetPolicyInfoByIdResponseExdPolicyExtsBuilder {
	b.s.VidFilter = &v
	return b
}

func (b *GetPolicyInfoByIdResponseExdPolicyExtsBuilder) PidFilter(v string) *GetPolicyInfoByIdResponseExdPolicyExtsBuilder {
	b.s.PidFilter = &v
	return b
}

func (b *GetPolicyInfoByIdResponseExdPolicyExtsBuilder) ClassFilter(v string) *GetPolicyInfoByIdResponseExdPolicyExtsBuilder {
	b.s.ClassFilter = &v
	return b
}

func (b *GetPolicyInfoByIdResponseExdPolicyExtsBuilder) ExtRemark(v string) *GetPolicyInfoByIdResponseExdPolicyExtsBuilder {
	b.s.ExtRemark = &v
	return b
}

func (b *GetPolicyInfoByIdResponseExdPolicyExtsBuilder) Build() *GetPolicyInfoByIdResponseExdPolicyExts {
	return b.s
}


