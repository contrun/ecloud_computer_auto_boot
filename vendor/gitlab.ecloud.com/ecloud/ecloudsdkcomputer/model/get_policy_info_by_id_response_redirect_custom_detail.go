// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdResponseRedirectCustomDetail struct {

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

func (s GetPolicyInfoByIdResponseRedirectCustomDetail) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseRedirectCustomDetail) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseRedirectCustomDetail) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseRedirectCustomDetail) SetAction(v string) *GetPolicyInfoByIdResponseRedirectCustomDetail {
	s.Action = &v
	return s
}

func (s *GetPolicyInfoByIdResponseRedirectCustomDetail) SetVidFilter(v string) *GetPolicyInfoByIdResponseRedirectCustomDetail {
	s.VidFilter = &v
	return s
}

func (s *GetPolicyInfoByIdResponseRedirectCustomDetail) SetPidFilter(v string) *GetPolicyInfoByIdResponseRedirectCustomDetail {
	s.PidFilter = &v
	return s
}

func (s *GetPolicyInfoByIdResponseRedirectCustomDetail) SetClassFilter(v string) *GetPolicyInfoByIdResponseRedirectCustomDetail {
	s.ClassFilter = &v
	return s
}

func (s *GetPolicyInfoByIdResponseRedirectCustomDetail) SetExtRemark(v string) *GetPolicyInfoByIdResponseRedirectCustomDetail {
	s.ExtRemark = &v
	return s
}


type GetPolicyInfoByIdResponseRedirectCustomDetailBuilder struct {
	s *GetPolicyInfoByIdResponseRedirectCustomDetail
}

func NewGetPolicyInfoByIdResponseRedirectCustomDetailBuilder() *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder {
	s := &GetPolicyInfoByIdResponseRedirectCustomDetail{}
	b := &GetPolicyInfoByIdResponseRedirectCustomDetailBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder) Action(v string) *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder {
	b.s.Action = &v
	return b
}

func (b *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder) VidFilter(v string) *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder {
	b.s.VidFilter = &v
	return b
}

func (b *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder) PidFilter(v string) *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder {
	b.s.PidFilter = &v
	return b
}

func (b *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder) ClassFilter(v string) *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder {
	b.s.ClassFilter = &v
	return b
}

func (b *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder) ExtRemark(v string) *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder {
	b.s.ExtRemark = &v
	return b
}

func (b *GetPolicyInfoByIdResponseRedirectCustomDetailBuilder) Build() *GetPolicyInfoByIdResponseRedirectCustomDetail {
	return b.s
}


