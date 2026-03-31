// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreatePolicyRequestTerminalLoginPolicysCategoryEnum string

// List of Category
const (
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumCloseMachine CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "CLOSE_MACHINE"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBanTerminalType CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BAN_TERMINAL_TYPE"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumShutSetting CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "SHUT_SETTING"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumRestartSetting CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "RESTART_SETTING"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumReconnectInterval CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "RECONNECT_INTERVAL"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumReconnectTotal CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "RECONNECT_TOTAL"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumCloseDeviceSource CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "CLOSE_DEVICE_SOURCE"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumDisconnect CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "DISCONNECT"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumStartMode CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "START_MODE"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumUpdateMachineName CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "UPDATE_MACHINE_NAME"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumAllowReload CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "ALLOW_RELOAD"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBigscreenNetCheck CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_NET_CHECK"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBigscreenAttention CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_ATTENTION"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBigscreenSet CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_SET"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBigscreenMinimize CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_MINIMIZE"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBigscreenWindowed CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_WINDOWED"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBigscreenQuit CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBigscreenQuitBreak CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT_BREAK"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBigscreenQuitRestart CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT_RESTART"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBigscreenQuitShut CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT_SHUT"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumMobileManage CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_MANAGE"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumMobileHelp CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_HELP"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumMobileQuit CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumMobileQuitBreak CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT_BREAK"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumMobileQuitRestart CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT_RESTART"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumMobileQuitShut CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT_SHUT"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumConnectType CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "CONNECT_TYPE"
    CreatePolicyRequestTerminalLoginPolicysCategoryEnumBreakShut CreatePolicyRequestTerminalLoginPolicysCategoryEnum = "BREAK_SHUT"
)

type CreatePolicyRequestTerminalLoginPolicys struct {

    // 终端登录枚举值
	Category *CreatePolicyRequestTerminalLoginPolicysCategoryEnum `json:"category,omitempty"`
    // 终端登录value
	Value *string `json:"value,omitempty"`
}

func (s CreatePolicyRequestTerminalLoginPolicys) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestTerminalLoginPolicys) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestTerminalLoginPolicys) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestTerminalLoginPolicys) SetCategory(v CreatePolicyRequestTerminalLoginPolicysCategoryEnum) *CreatePolicyRequestTerminalLoginPolicys {
	s.Category = &v
	return s
}

func (s *CreatePolicyRequestTerminalLoginPolicys) SetValue(v string) *CreatePolicyRequestTerminalLoginPolicys {
	s.Value = &v
	return s
}


type CreatePolicyRequestTerminalLoginPolicysBuilder struct {
	s *CreatePolicyRequestTerminalLoginPolicys
}

func NewCreatePolicyRequestTerminalLoginPolicysBuilder() *CreatePolicyRequestTerminalLoginPolicysBuilder {
	s := &CreatePolicyRequestTerminalLoginPolicys{}
	b := &CreatePolicyRequestTerminalLoginPolicysBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestTerminalLoginPolicysBuilder) Category(v CreatePolicyRequestTerminalLoginPolicysCategoryEnum) *CreatePolicyRequestTerminalLoginPolicysBuilder {
	b.s.Category = &v
	return b
}

func (b *CreatePolicyRequestTerminalLoginPolicysBuilder) Value(v string) *CreatePolicyRequestTerminalLoginPolicysBuilder {
	b.s.Value = &v
	return b
}

func (b *CreatePolicyRequestTerminalLoginPolicysBuilder) Build() *CreatePolicyRequestTerminalLoginPolicys {
	return b.s
}


