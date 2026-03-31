// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateComputerPolicyRequestExdPolicyScreenRecordeClass struct {

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

func (s UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetScreenRecordEnable(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenRecordEnable = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetIntervalRecordEndTime(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.IntervalRecordEndTime = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetScreenDuration(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenDuration = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetAuditRetention(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.AuditRetention = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetScreenType(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenType = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetScreenFPS(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenFPS = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetScreenTriggerMode(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenTriggerMode = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetIntervalRecordStartTime(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.IntervalRecordStartTime = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetScreenRecordMode(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenRecordMode = &v
	return s
}

func (s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) SetScreenAuditPolicyUid(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenAuditPolicyUid = &v
	return s
}


type UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder struct {
	s *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass
}

func NewUpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder() *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	s := &UpdateComputerPolicyRequestExdPolicyScreenRecordeClass{}
	b := &UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenRecordEnable(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenRecordEnable = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) IntervalRecordEndTime(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.IntervalRecordEndTime = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenDuration(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenDuration = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) AuditRetention(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.AuditRetention = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenType(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenType = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenFPS(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenFPS = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenTriggerMode(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenTriggerMode = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) IntervalRecordStartTime(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.IntervalRecordStartTime = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenRecordMode(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenRecordMode = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenAuditPolicyUid(v string) *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenAuditPolicyUid = &v
	return b
}

func (b *UpdateComputerPolicyRequestExdPolicyScreenRecordeClassBuilder) Build() *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass {
	return b.s
}


