// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum string

// List of Category
const (
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumCloseMachine GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "CLOSE_MACHINE"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBanTerminalType GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BAN_TERMINAL_TYPE"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumShutSetting GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "SHUT_SETTING"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumRestartSetting GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "RESTART_SETTING"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumReconnectInterval GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "RECONNECT_INTERVAL"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumReconnectTotal GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "RECONNECT_TOTAL"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumCloseDeviceSource GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "CLOSE_DEVICE_SOURCE"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumDisconnect GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "DISCONNECT"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumStartMode GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "START_MODE"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumUpdateMachineName GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "UPDATE_MACHINE_NAME"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumAllowReload GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "ALLOW_RELOAD"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBigscreenNetCheck GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BIGSCREEN_NET_CHECK"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBigscreenAttention GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BIGSCREEN_ATTENTION"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBigscreenSet GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BIGSCREEN_SET"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBigscreenMinimize GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BIGSCREEN_MINIMIZE"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBigscreenWindowed GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BIGSCREEN_WINDOWED"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBigscreenQuit GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBigscreenQuitBreak GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT_BREAK"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBigscreenQuitRestart GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT_RESTART"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBigscreenQuitShut GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT_SHUT"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumMobileManage GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "MOBILE_MANAGE"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumMobileHelp GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "MOBILE_HELP"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumMobileQuit GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "MOBILE_QUIT"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumMobileQuitBreak GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "MOBILE_QUIT_BREAK"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumMobileQuitRestart GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "MOBILE_QUIT_RESTART"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumMobileQuitShut GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "MOBILE_QUIT_SHUT"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumConnectType GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "CONNECT_TYPE"
    GetPolicyInfoByIdResponseNetworkPolicysCategoryEnumBreakShut GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum = "BREAK_SHUT"
)

type GetPolicyInfoByIdResponseNetworkPolicys struct {

    // 桌面工具栏枚举值
	Category *GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum `json:"category,omitempty"`
    // 桌面工具栏value
	Value *string `json:"value,omitempty"`
}

func (s GetPolicyInfoByIdResponseNetworkPolicys) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseNetworkPolicys) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseNetworkPolicys) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseNetworkPolicys) SetCategory(v GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum) *GetPolicyInfoByIdResponseNetworkPolicys {
	s.Category = &v
	return s
}

func (s *GetPolicyInfoByIdResponseNetworkPolicys) SetValue(v string) *GetPolicyInfoByIdResponseNetworkPolicys {
	s.Value = &v
	return s
}


type GetPolicyInfoByIdResponseNetworkPolicysBuilder struct {
	s *GetPolicyInfoByIdResponseNetworkPolicys
}

func NewGetPolicyInfoByIdResponseNetworkPolicysBuilder() *GetPolicyInfoByIdResponseNetworkPolicysBuilder {
	s := &GetPolicyInfoByIdResponseNetworkPolicys{}
	b := &GetPolicyInfoByIdResponseNetworkPolicysBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseNetworkPolicysBuilder) Category(v GetPolicyInfoByIdResponseNetworkPolicysCategoryEnum) *GetPolicyInfoByIdResponseNetworkPolicysBuilder {
	b.s.Category = &v
	return b
}

func (b *GetPolicyInfoByIdResponseNetworkPolicysBuilder) Value(v string) *GetPolicyInfoByIdResponseNetworkPolicysBuilder {
	b.s.Value = &v
	return b
}

func (b *GetPolicyInfoByIdResponseNetworkPolicysBuilder) Build() *GetPolicyInfoByIdResponseNetworkPolicys {
	return b.s
}


