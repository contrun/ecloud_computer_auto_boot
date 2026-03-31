// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum string

// List of Category
const (
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumCloseMachine UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "CLOSE_MACHINE"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBanTerminalType UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BAN_TERMINAL_TYPE"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumShutSetting UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "SHUT_SETTING"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumRestartSetting UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "RESTART_SETTING"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumReconnectInterval UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "RECONNECT_INTERVAL"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumReconnectTotal UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "RECONNECT_TOTAL"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumCloseDeviceSource UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "CLOSE_DEVICE_SOURCE"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumDisconnect UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "DISCONNECT"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumStartMode UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "START_MODE"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumUpdateMachineName UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "UPDATE_MACHINE_NAME"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumAllowReload UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "ALLOW_RELOAD"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBigscreenNetCheck UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_NET_CHECK"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBigscreenAttention UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_ATTENTION"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBigscreenSet UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_SET"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBigscreenMinimize UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_MINIMIZE"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBigscreenWindowed UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_WINDOWED"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBigscreenQuit UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBigscreenQuitBreak UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT_BREAK"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBigscreenQuitRestart UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT_RESTART"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBigscreenQuitShut UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT_SHUT"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumMobileManage UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_MANAGE"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumMobileHelp UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_HELP"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumMobileQuit UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumMobileQuitBreak UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT_BREAK"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumMobileQuitRestart UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT_RESTART"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumMobileQuitShut UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT_SHUT"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumConnectType UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "CONNECT_TYPE"
    UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnumBreakShut UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum = "BREAK_SHUT"
)

type UpdateComputerPolicyRequestTerminalLoginPolicys struct {

    // 终端登录枚举值
	Category *UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum `json:"category,omitempty"`
    // 终端登录value
	Value *string `json:"value,omitempty"`
}

func (s UpdateComputerPolicyRequestTerminalLoginPolicys) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequestTerminalLoginPolicys) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequestTerminalLoginPolicys) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequestTerminalLoginPolicys) SetCategory(v UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum) *UpdateComputerPolicyRequestTerminalLoginPolicys {
	s.Category = &v
	return s
}

func (s *UpdateComputerPolicyRequestTerminalLoginPolicys) SetValue(v string) *UpdateComputerPolicyRequestTerminalLoginPolicys {
	s.Value = &v
	return s
}


type UpdateComputerPolicyRequestTerminalLoginPolicysBuilder struct {
	s *UpdateComputerPolicyRequestTerminalLoginPolicys
}

func NewUpdateComputerPolicyRequestTerminalLoginPolicysBuilder() *UpdateComputerPolicyRequestTerminalLoginPolicysBuilder {
	s := &UpdateComputerPolicyRequestTerminalLoginPolicys{}
	b := &UpdateComputerPolicyRequestTerminalLoginPolicysBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestTerminalLoginPolicysBuilder) Category(v UpdateComputerPolicyRequestTerminalLoginPolicysCategoryEnum) *UpdateComputerPolicyRequestTerminalLoginPolicysBuilder {
	b.s.Category = &v
	return b
}

func (b *UpdateComputerPolicyRequestTerminalLoginPolicysBuilder) Value(v string) *UpdateComputerPolicyRequestTerminalLoginPolicysBuilder {
	b.s.Value = &v
	return b
}

func (b *UpdateComputerPolicyRequestTerminalLoginPolicysBuilder) Build() *UpdateComputerPolicyRequestTerminalLoginPolicys {
	return b.s
}


