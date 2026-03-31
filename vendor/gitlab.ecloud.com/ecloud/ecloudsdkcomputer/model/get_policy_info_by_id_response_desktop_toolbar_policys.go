// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum string

// List of Category
const (
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumCloseMachine GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "CLOSE_MACHINE"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBanTerminalType GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BAN_TERMINAL_TYPE"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumShutSetting GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "SHUT_SETTING"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumRestartSetting GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "RESTART_SETTING"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumReconnectInterval GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "RECONNECT_INTERVAL"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumReconnectTotal GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "RECONNECT_TOTAL"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumCloseDeviceSource GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "CLOSE_DEVICE_SOURCE"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumDisconnect GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "DISCONNECT"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumStartMode GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "START_MODE"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumUpdateMachineName GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "UPDATE_MACHINE_NAME"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumAllowReload GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "ALLOW_RELOAD"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBigscreenNetCheck GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_NET_CHECK"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBigscreenAttention GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_ATTENTION"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBigscreenSet GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_SET"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBigscreenMinimize GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_MINIMIZE"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBigscreenWindowed GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_WINDOWED"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBigscreenQuit GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_QUIT"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBigscreenQuitBreak GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_QUIT_BREAK"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBigscreenQuitRestart GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_QUIT_RESTART"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBigscreenQuitShut GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BIGSCREEN_QUIT_SHUT"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumMobileManage GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "MOBILE_MANAGE"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumMobileHelp GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "MOBILE_HELP"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumMobileQuit GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "MOBILE_QUIT"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumMobileQuitBreak GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "MOBILE_QUIT_BREAK"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumMobileQuitRestart GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "MOBILE_QUIT_RESTART"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumMobileQuitShut GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "MOBILE_QUIT_SHUT"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumConnectType GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "CONNECT_TYPE"
    GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnumBreakShut GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum = "BREAK_SHUT"
)

type GetPolicyInfoByIdResponseDesktopToolbarPolicys struct {

    // 桌面工具栏枚举值
	Category *GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum `json:"category,omitempty"`
    // 桌面工具栏value
	Value *string `json:"value,omitempty"`
}

func (s GetPolicyInfoByIdResponseDesktopToolbarPolicys) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseDesktopToolbarPolicys) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseDesktopToolbarPolicys) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseDesktopToolbarPolicys) SetCategory(v GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum) *GetPolicyInfoByIdResponseDesktopToolbarPolicys {
	s.Category = &v
	return s
}

func (s *GetPolicyInfoByIdResponseDesktopToolbarPolicys) SetValue(v string) *GetPolicyInfoByIdResponseDesktopToolbarPolicys {
	s.Value = &v
	return s
}


type GetPolicyInfoByIdResponseDesktopToolbarPolicysBuilder struct {
	s *GetPolicyInfoByIdResponseDesktopToolbarPolicys
}

func NewGetPolicyInfoByIdResponseDesktopToolbarPolicysBuilder() *GetPolicyInfoByIdResponseDesktopToolbarPolicysBuilder {
	s := &GetPolicyInfoByIdResponseDesktopToolbarPolicys{}
	b := &GetPolicyInfoByIdResponseDesktopToolbarPolicysBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseDesktopToolbarPolicysBuilder) Category(v GetPolicyInfoByIdResponseDesktopToolbarPolicysCategoryEnum) *GetPolicyInfoByIdResponseDesktopToolbarPolicysBuilder {
	b.s.Category = &v
	return b
}

func (b *GetPolicyInfoByIdResponseDesktopToolbarPolicysBuilder) Value(v string) *GetPolicyInfoByIdResponseDesktopToolbarPolicysBuilder {
	b.s.Value = &v
	return b
}

func (b *GetPolicyInfoByIdResponseDesktopToolbarPolicysBuilder) Build() *GetPolicyInfoByIdResponseDesktopToolbarPolicys {
	return b.s
}


