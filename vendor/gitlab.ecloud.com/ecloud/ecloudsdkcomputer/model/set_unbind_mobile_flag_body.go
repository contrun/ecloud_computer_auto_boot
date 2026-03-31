// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetUnbindMobileFlagBody struct {
    position.Body
    // 用户解绑手机号开关 默认 1 :开 0:关 
	UnbindMobileFlag *int32 `json:"unbindMobileFlag,omitempty"`
}

func (s SetUnbindMobileFlagBody) String() string {
	return utils.Beautify(s)
}

func (s SetUnbindMobileFlagBody) GoString() string {
	return s.String()
}

func (s SetUnbindMobileFlagBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetUnbindMobileFlagBody) SetUnbindMobileFlag(v int32) *SetUnbindMobileFlagBody {
	s.UnbindMobileFlag = &v
	return s
}


type SetUnbindMobileFlagBodyBuilder struct {
	s *SetUnbindMobileFlagBody
}

func NewSetUnbindMobileFlagBodyBuilder() *SetUnbindMobileFlagBodyBuilder {
	s := &SetUnbindMobileFlagBody{}
	b := &SetUnbindMobileFlagBodyBuilder{s: s}
	return b
}

func (b *SetUnbindMobileFlagBodyBuilder) UnbindMobileFlag(v int32) *SetUnbindMobileFlagBodyBuilder {
	b.s.UnbindMobileFlag = &v
	return b
}

func (b *SetUnbindMobileFlagBodyBuilder) Build() *SetUnbindMobileFlagBody {
	return b.s
}


