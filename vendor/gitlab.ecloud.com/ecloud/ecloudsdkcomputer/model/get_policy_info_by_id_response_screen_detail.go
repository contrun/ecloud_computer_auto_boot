// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdResponseScreenDetail struct {

    // 录屏策略开关 取值 disable 、enable（默认）
	ScreenRecordEnable *string `json:"screenRecordEnable,omitempty"`
    // 间隔录屏结束时间
	IntervalRecordEndTime *string `json:"intervalRecordEndTime,omitempty"`
    // 录屏文件阅览时长：单位分钟, 可选时长:10分钟(默认)、20分钟、30分钟、或60分钟
	ScreenDuration *string `json:"screenDuration,omitempty"`
    // 保留时长：单位天, 默认15天 保留时间的范围为:1~180天
	AuditRetention *string `json:"auditRetention,omitempty"`
    // 录屏类型 0:全程录屏 1: 间隔录屏（默认） 2:用户操作录屏 3:监听会话生命周期录屏
	ScreenType *string `json:"screenType,omitempty"`
    // 录屏图像帧率：单位fps 图像帧率：2 fps、5 fps（推荐）、10 fps和15 fps 
	ScreenFPS *string `json:"screenFPS,omitempty"`
    // 用户操作录屏方式 1:云电脑与本地电脑进行文件双向上传、下载录屏 2:执行命令录屏
	ScreenTriggerMode *string `json:"screenTriggerMode,omitempty"`
    // 间隔录屏开始时间
	IntervalRecordStartTime *string `json:"intervalRecordStartTime,omitempty"`
    // 录制云电脑声音  1:仅录制画面不录制声音（默认） 2:音画同期录制
	ScreenRecordMode *string `json:"screenRecordMode,omitempty"`
    // 录屏审计策略表uid
	ScreenAuditPolicyUid *string `json:"screenAuditPolicyUid,omitempty"`
}

func (s GetPolicyInfoByIdResponseScreenDetail) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseScreenDetail) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseScreenDetail) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetScreenRecordEnable(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.ScreenRecordEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetIntervalRecordEndTime(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.IntervalRecordEndTime = &v
	return s
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetScreenDuration(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.ScreenDuration = &v
	return s
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetAuditRetention(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.AuditRetention = &v
	return s
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetScreenType(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.ScreenType = &v
	return s
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetScreenFPS(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.ScreenFPS = &v
	return s
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetScreenTriggerMode(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.ScreenTriggerMode = &v
	return s
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetIntervalRecordStartTime(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.IntervalRecordStartTime = &v
	return s
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetScreenRecordMode(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.ScreenRecordMode = &v
	return s
}

func (s *GetPolicyInfoByIdResponseScreenDetail) SetScreenAuditPolicyUid(v string) *GetPolicyInfoByIdResponseScreenDetail {
	s.ScreenAuditPolicyUid = &v
	return s
}


type GetPolicyInfoByIdResponseScreenDetailBuilder struct {
	s *GetPolicyInfoByIdResponseScreenDetail
}

func NewGetPolicyInfoByIdResponseScreenDetailBuilder() *GetPolicyInfoByIdResponseScreenDetailBuilder {
	s := &GetPolicyInfoByIdResponseScreenDetail{}
	b := &GetPolicyInfoByIdResponseScreenDetailBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) ScreenRecordEnable(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.ScreenRecordEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) IntervalRecordEndTime(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.IntervalRecordEndTime = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) ScreenDuration(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.ScreenDuration = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) AuditRetention(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.AuditRetention = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) ScreenType(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.ScreenType = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) ScreenFPS(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.ScreenFPS = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) ScreenTriggerMode(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.ScreenTriggerMode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) IntervalRecordStartTime(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.IntervalRecordStartTime = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) ScreenRecordMode(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.ScreenRecordMode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) ScreenAuditPolicyUid(v string) *GetPolicyInfoByIdResponseScreenDetailBuilder {
	b.s.ScreenAuditPolicyUid = &v
	return b
}

func (b *GetPolicyInfoByIdResponseScreenDetailBuilder) Build() *GetPolicyInfoByIdResponseScreenDetail {
	return b.s
}


