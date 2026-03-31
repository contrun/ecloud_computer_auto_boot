// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreatePolicyRequestExdPolicyScreenRecordeClass struct {

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

func (s CreatePolicyRequestExdPolicyScreenRecordeClass) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestExdPolicyScreenRecordeClass) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestExdPolicyScreenRecordeClass) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetScreenRecordEnable(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenRecordEnable = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetIntervalRecordEndTime(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.IntervalRecordEndTime = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetScreenDuration(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenDuration = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetAuditRetention(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.AuditRetention = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetScreenType(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenType = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetScreenFPS(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenFPS = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetScreenTriggerMode(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenTriggerMode = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetIntervalRecordStartTime(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.IntervalRecordStartTime = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetScreenRecordMode(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenRecordMode = &v
	return s
}

func (s *CreatePolicyRequestExdPolicyScreenRecordeClass) SetScreenAuditPolicyUid(v string) *CreatePolicyRequestExdPolicyScreenRecordeClass {
	s.ScreenAuditPolicyUid = &v
	return s
}


type CreatePolicyRequestExdPolicyScreenRecordeClassBuilder struct {
	s *CreatePolicyRequestExdPolicyScreenRecordeClass
}

func NewCreatePolicyRequestExdPolicyScreenRecordeClassBuilder() *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	s := &CreatePolicyRequestExdPolicyScreenRecordeClass{}
	b := &CreatePolicyRequestExdPolicyScreenRecordeClassBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenRecordEnable(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenRecordEnable = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) IntervalRecordEndTime(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.IntervalRecordEndTime = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenDuration(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenDuration = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) AuditRetention(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.AuditRetention = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenType(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenType = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenFPS(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenFPS = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenTriggerMode(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenTriggerMode = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) IntervalRecordStartTime(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.IntervalRecordStartTime = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenRecordMode(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenRecordMode = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) ScreenAuditPolicyUid(v string) *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder {
	b.s.ScreenAuditPolicyUid = &v
	return b
}

func (b *CreatePolicyRequestExdPolicyScreenRecordeClassBuilder) Build() *CreatePolicyRequestExdPolicyScreenRecordeClass {
	return b.s
}


