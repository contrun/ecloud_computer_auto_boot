// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreatePolicyRequestNetworkPolicysCategoryEnum string

// List of Category
const (
    CreatePolicyRequestNetworkPolicysCategoryEnumCloseMachine CreatePolicyRequestNetworkPolicysCategoryEnum = "CLOSE_MACHINE"
    CreatePolicyRequestNetworkPolicysCategoryEnumBanTerminalType CreatePolicyRequestNetworkPolicysCategoryEnum = "BAN_TERMINAL_TYPE"
    CreatePolicyRequestNetworkPolicysCategoryEnumShutSetting CreatePolicyRequestNetworkPolicysCategoryEnum = "SHUT_SETTING"
    CreatePolicyRequestNetworkPolicysCategoryEnumRestartSetting CreatePolicyRequestNetworkPolicysCategoryEnum = "RESTART_SETTING"
    CreatePolicyRequestNetworkPolicysCategoryEnumReconnectInterval CreatePolicyRequestNetworkPolicysCategoryEnum = "RECONNECT_INTERVAL"
    CreatePolicyRequestNetworkPolicysCategoryEnumReconnectTotal CreatePolicyRequestNetworkPolicysCategoryEnum = "RECONNECT_TOTAL"
    CreatePolicyRequestNetworkPolicysCategoryEnumCloseDeviceSource CreatePolicyRequestNetworkPolicysCategoryEnum = "CLOSE_DEVICE_SOURCE"
    CreatePolicyRequestNetworkPolicysCategoryEnumDisconnect CreatePolicyRequestNetworkPolicysCategoryEnum = "DISCONNECT"
    CreatePolicyRequestNetworkPolicysCategoryEnumStartMode CreatePolicyRequestNetworkPolicysCategoryEnum = "START_MODE"
    CreatePolicyRequestNetworkPolicysCategoryEnumUpdateMachineName CreatePolicyRequestNetworkPolicysCategoryEnum = "UPDATE_MACHINE_NAME"
    CreatePolicyRequestNetworkPolicysCategoryEnumAllowReload CreatePolicyRequestNetworkPolicysCategoryEnum = "ALLOW_RELOAD"
    CreatePolicyRequestNetworkPolicysCategoryEnumBigscreenNetCheck CreatePolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_NET_CHECK"
    CreatePolicyRequestNetworkPolicysCategoryEnumBigscreenAttention CreatePolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_ATTENTION"
    CreatePolicyRequestNetworkPolicysCategoryEnumBigscreenSet CreatePolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_SET"
    CreatePolicyRequestNetworkPolicysCategoryEnumBigscreenMinimize CreatePolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_MINIMIZE"
    CreatePolicyRequestNetworkPolicysCategoryEnumBigscreenWindowed CreatePolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_WINDOWED"
    CreatePolicyRequestNetworkPolicysCategoryEnumBigscreenQuit CreatePolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT"
    CreatePolicyRequestNetworkPolicysCategoryEnumBigscreenQuitBreak CreatePolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT_BREAK"
    CreatePolicyRequestNetworkPolicysCategoryEnumBigscreenQuitRestart CreatePolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT_RESTART"
    CreatePolicyRequestNetworkPolicysCategoryEnumBigscreenQuitShut CreatePolicyRequestNetworkPolicysCategoryEnum = "BIGSCREEN_QUIT_SHUT"
    CreatePolicyRequestNetworkPolicysCategoryEnumMobileManage CreatePolicyRequestNetworkPolicysCategoryEnum = "MOBILE_MANAGE"
    CreatePolicyRequestNetworkPolicysCategoryEnumMobileHelp CreatePolicyRequestNetworkPolicysCategoryEnum = "MOBILE_HELP"
    CreatePolicyRequestNetworkPolicysCategoryEnumMobileQuit CreatePolicyRequestNetworkPolicysCategoryEnum = "MOBILE_QUIT"
    CreatePolicyRequestNetworkPolicysCategoryEnumMobileQuitBreak CreatePolicyRequestNetworkPolicysCategoryEnum = "MOBILE_QUIT_BREAK"
    CreatePolicyRequestNetworkPolicysCategoryEnumMobileQuitRestart CreatePolicyRequestNetworkPolicysCategoryEnum = "MOBILE_QUIT_RESTART"
    CreatePolicyRequestNetworkPolicysCategoryEnumMobileQuitShut CreatePolicyRequestNetworkPolicysCategoryEnum = "MOBILE_QUIT_SHUT"
    CreatePolicyRequestNetworkPolicysCategoryEnumConnectType CreatePolicyRequestNetworkPolicysCategoryEnum = "CONNECT_TYPE"
    CreatePolicyRequestNetworkPolicysCategoryEnumBreakShut CreatePolicyRequestNetworkPolicysCategoryEnum = "BREAK_SHUT"
)

type CreatePolicyRequestNetworkPolicys struct {

    // 网络控制枚举值
	Category *CreatePolicyRequestNetworkPolicysCategoryEnum `json:"category,omitempty"`
    // 网络控制value
	Value *string `json:"value,omitempty"`
}

func (s CreatePolicyRequestNetworkPolicys) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyRequestNetworkPolicys) GoString() string {
	return s.String()
}

func (s CreatePolicyRequestNetworkPolicys) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyRequestNetworkPolicys) SetCategory(v CreatePolicyRequestNetworkPolicysCategoryEnum) *CreatePolicyRequestNetworkPolicys {
	s.Category = &v
	return s
}

func (s *CreatePolicyRequestNetworkPolicys) SetValue(v string) *CreatePolicyRequestNetworkPolicys {
	s.Value = &v
	return s
}


type CreatePolicyRequestNetworkPolicysBuilder struct {
	s *CreatePolicyRequestNetworkPolicys
}

func NewCreatePolicyRequestNetworkPolicysBuilder() *CreatePolicyRequestNetworkPolicysBuilder {
	s := &CreatePolicyRequestNetworkPolicys{}
	b := &CreatePolicyRequestNetworkPolicysBuilder{s: s}
	return b
}

func (b *CreatePolicyRequestNetworkPolicysBuilder) Category(v CreatePolicyRequestNetworkPolicysCategoryEnum) *CreatePolicyRequestNetworkPolicysBuilder {
	b.s.Category = &v
	return b
}

func (b *CreatePolicyRequestNetworkPolicysBuilder) Value(v string) *CreatePolicyRequestNetworkPolicysBuilder {
	b.s.Value = &v
	return b
}

func (b *CreatePolicyRequestNetworkPolicysBuilder) Build() *CreatePolicyRequestNetworkPolicys {
	return b.s
}


