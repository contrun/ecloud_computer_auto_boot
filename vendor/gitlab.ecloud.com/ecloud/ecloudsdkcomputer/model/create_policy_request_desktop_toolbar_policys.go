// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreatePolicyRequestDesktopToolbarPolicysCategoryEnum string

// List of Category
const (
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumCloseMachine CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "CLOSE_MACHINE"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBanTerminalType CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BAN_TERMINAL_TYPE"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumShutSetting CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "SHUT_SETTING"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumRestartSetting CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "RESTART_SETTING"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumReconnectInterval CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "RECONNECT_INTERVAL"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumReconnectTotal CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "RECONNECT_TOTAL"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumCloseDeviceSource CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "CLOSE_DEVICE_SOURCE"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumDisconnect CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "DISCONNECT"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumStartMode CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "START_MODE"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumUpdateMachineName CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "UPDATE_MACHINE_NAME"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumAllowReload CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "ALLOW_RELOAD"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBigscreenNetCheck CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_NET_CHECK"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBigscreenAttention CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_ATTENTION"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBigscreenSet CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_SET"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBigscreenMinimize CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_MINIMIZE"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBigscreenWindowed CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_WINDOWED"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBigscreenQuit CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_QUIT"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBigscreenQuitBreak CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_QUIT_BREAK"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBigscreenQuitRestart CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_QUIT_RESTART"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBigscreenQuitShut CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_QUIT_SHUT"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumMobileManage CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "MOBILE_MANAGE"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumMobileHelp CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "MOBILE_HELP"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumMobileQuit CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "MOBILE_QUIT"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumMobileQuitBreak CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "MOBILE_QUIT_BREAK"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumMobileQuitRestart CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "MOBILE_QUIT_RESTART"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumMobileQuitShut CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "MOBILE_QUIT_SHUT"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumConnectType CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "CONNECT_TYPE"
    CreatePolicyRequestDesktopToolbarPolicysCategoryEnumBreakShut CreatePolicyRequestDesktopToolbarPolicysCategoryEnum = "BREAK_SHUT"
)

type CreatePolicyRequestDesktopToolbarPolicys struct {

    // 桌面工具栏枚举值
	Category *CreatePolicyRequestDesktopToolbarPolicysCategoryEnum `json:"category,omitempty"`
    // 桌面工具栏value
	Value *string `json:"value,omitempty"`
}

func (s CreatePolicyRequestDesktopToolbarPolicys) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestDesktopToolbarPolicys) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestDesktopToolbarPolicys) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestDesktopToolbarPolicys) SetCategory(v CreatePolicyRequestDesktopToolbarPolicysCategoryEnum) *CreatePolicyRequestDesktopToolbarPolicys {
	s.Category = &v
	return s
}

func (s *CreatePolicyRequestDesktopToolbarPolicys) SetValue(v string) *CreatePolicyRequestDesktopToolbarPolicys {
	s.Value = &v
	return s
}


type CreatePolicyRequestDesktopToolbarPolicysBuilder struct {
	s *CreatePolicyRequestDesktopToolbarPolicys
}

func NewCreatePolicyRequestDesktopToolbarPolicysBuilder() *CreatePolicyRequestDesktopToolbarPolicysBuilder {
	s := &CreatePolicyRequestDesktopToolbarPolicys{}
	b := &CreatePolicyRequestDesktopToolbarPolicysBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestDesktopToolbarPolicysBuilder) Category(v CreatePolicyRequestDesktopToolbarPolicysCategoryEnum) *CreatePolicyRequestDesktopToolbarPolicysBuilder {
	b.s.Category = &v
	return b
}

func (b *CreatePolicyRequestDesktopToolbarPolicysBuilder) Value(v string) *CreatePolicyRequestDesktopToolbarPolicysBuilder {
	b.s.Value = &v
	return b
}

func (b *CreatePolicyRequestDesktopToolbarPolicysBuilder) Build() *CreatePolicyRequestDesktopToolbarPolicys {
	return b.s
}


