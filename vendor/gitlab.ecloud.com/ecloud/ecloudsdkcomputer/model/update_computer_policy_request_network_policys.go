// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateComputerPolicyRequestNetworkPolicysCategoryEnum string

// List of Category
const (
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumCloseMachine UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "CLOSE_MACHINE"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBanTerminalType UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BAN_TERMINAL_TYPE"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumShutSetting UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "SHUT_SETTING"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumRestartSetting UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "RESTART_SETTING"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumReconnectInterval UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "RECONNECT_INTERVAL"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumReconnectTotal UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "RECONNECT_TOTAL"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumCloseDeviceSource UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "CLOSE_DEVICE_SOURCE"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumDisconnect UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "DISCONNECT"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumStartMode UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "START_MODE"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumUpdateMachineName UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "UPDATE_MACHINE_NAME"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumAllowReload UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "ALLOW_RELOAD"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBigscreenNetCheck UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_NET_CHECK"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBigscreenAttention UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_ATTENTION"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBigscreenSet UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_SET"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBigscreenMinimize UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_MINIMIZE"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBigscreenWindowed UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_WINDOWED"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBigscreenQuit UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBigscreenQuitBreak UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT_BREAK"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBigscreenQuitRestart UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT_RESTART"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBigscreenQuitShut UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT_SHUT"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumMobileManage UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "MOBILE_MANAGE"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumMobileHelp UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "MOBILE_HELP"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumMobileQuit UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "MOBILE_QUIT"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumMobileQuitBreak UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "MOBILE_QUIT_BREAK"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumMobileQuitRestart UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "MOBILE_QUIT_RESTART"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumMobileQuitShut UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "MOBILE_QUIT_SHUT"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumConnectType UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "CONNECT_TYPE"
    UpdateComputerPolicyRequestNetworkPolicysCategoryEnumBreakShut UpdateComputerPolicyRequestNetworkPolicysCategoryEnum = "BREAK_SHUT"
)

type UpdateComputerPolicyRequestNetworkPolicys struct {

    // 网络控制枚举值
	Category *UpdateComputerPolicyRequestNetworkPolicysCategoryEnum `json:"category,omitempty"`
    // 网络控制value
	Value *string `json:"value,omitempty"`
}

func (s UpdateComputerPolicyRequestNetworkPolicys) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyRequestNetworkPolicys) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyRequestNetworkPolicys) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyRequestNetworkPolicys) SetCategory(v UpdateComputerPolicyRequestNetworkPolicysCategoryEnum) *UpdateComputerPolicyRequestNetworkPolicys {
	s.Category = &v
	return s
}

func (s *UpdateComputerPolicyRequestNetworkPolicys) SetValue(v string) *UpdateComputerPolicyRequestNetworkPolicys {
	s.Value = &v
	return s
}


type UpdateComputerPolicyRequestNetworkPolicysBuilder struct {
	s *UpdateComputerPolicyRequestNetworkPolicys
}

func NewUpdateComputerPolicyRequestNetworkPolicysBuilder() *UpdateComputerPolicyRequestNetworkPolicysBuilder {
	s := &UpdateComputerPolicyRequestNetworkPolicys{}
	b := &UpdateComputerPolicyRequestNetworkPolicysBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyRequestNetworkPolicysBuilder) Category(v UpdateComputerPolicyRequestNetworkPolicysCategoryEnum) *UpdateComputerPolicyRequestNetworkPolicysBuilder {
	b.s.Category = &v
	return b
}

func (b *UpdateComputerPolicyRequestNetworkPolicysBuilder) Value(v string) *UpdateComputerPolicyRequestNetworkPolicysBuilder {
	b.s.Value = &v
	return b
}

func (b *UpdateComputerPolicyRequestNetworkPolicysBuilder) Build() *UpdateComputerPolicyRequestNetworkPolicys {
	return b.s
}


