// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum string

// List of Category
const (
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumCloseMachine GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "CLOSE_MACHINE"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBanTerminalType GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BAN_TERMINAL_TYPE"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumShutSetting GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "SHUT_SETTING"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumRestartSetting GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "RESTART_SETTING"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumReconnectInterval GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "RECONNECT_INTERVAL"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumReconnectTotal GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "RECONNECT_TOTAL"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumCloseDeviceSource GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "CLOSE_DEVICE_SOURCE"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumDisconnect GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "DISCONNECT"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumStartMode GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "START_MODE"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumUpdateMachineName GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "UPDATE_MACHINE_NAME"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumAllowReload GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "ALLOW_RELOAD"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBigscreenNetCheck GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BIGSCREEN_NET_CHECK"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBigscreenAttention GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BIGSCREEN_ATTENTION"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBigscreenSet GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BIGSCREEN_SET"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBigscreenMinimize GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BIGSCREEN_MINIMIZE"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBigscreenWindowed GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BIGSCREEN_WINDOWED"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBigscreenQuit GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBigscreenQuitBreak GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT_BREAK"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBigscreenQuitRestart GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT_RESTART"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBigscreenQuitShut GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BIGSCREEN_QUIT_SHUT"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumMobileManage GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "MOBILE_MANAGE"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumMobileHelp GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "MOBILE_HELP"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumMobileQuit GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumMobileQuitBreak GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT_BREAK"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumMobileQuitRestart GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT_RESTART"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumMobileQuitShut GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "MOBILE_QUIT_SHUT"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumConnectType GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "CONNECT_TYPE"
    GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnumBreakShut GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum = "BREAK_SHUT"
)

type GetPolicyInfoByIdResponseTerminalLoginPolicys struct {

    // 桌面工具栏枚举值
	Category *GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum `json:"category,omitempty"`
    // 桌面工具栏value
	Value *string `json:"value,omitempty"`
}

func (s GetPolicyInfoByIdResponseTerminalLoginPolicys) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseTerminalLoginPolicys) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseTerminalLoginPolicys) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseTerminalLoginPolicys) SetCategory(v GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum) *GetPolicyInfoByIdResponseTerminalLoginPolicys {
	s.Category = &v
	return s
}

func (s *GetPolicyInfoByIdResponseTerminalLoginPolicys) SetValue(v string) *GetPolicyInfoByIdResponseTerminalLoginPolicys {
	s.Value = &v
	return s
}


type GetPolicyInfoByIdResponseTerminalLoginPolicysBuilder struct {
	s *GetPolicyInfoByIdResponseTerminalLoginPolicys
}

func NewGetPolicyInfoByIdResponseTerminalLoginPolicysBuilder() *GetPolicyInfoByIdResponseTerminalLoginPolicysBuilder {
	s := &GetPolicyInfoByIdResponseTerminalLoginPolicys{}
	b := &GetPolicyInfoByIdResponseTerminalLoginPolicysBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseTerminalLoginPolicysBuilder) Category(v GetPolicyInfoByIdResponseTerminalLoginPolicysCategoryEnum) *GetPolicyInfoByIdResponseTerminalLoginPolicysBuilder {
	b.s.Category = &v
	return b
}

func (b *GetPolicyInfoByIdResponseTerminalLoginPolicysBuilder) Value(v string) *GetPolicyInfoByIdResponseTerminalLoginPolicysBuilder {
	b.s.Value = &v
	return b
}

func (b *GetPolicyInfoByIdResponseTerminalLoginPolicysBuilder) Build() *GetPolicyInfoByIdResponseTerminalLoginPolicys {
	return b.s
}


