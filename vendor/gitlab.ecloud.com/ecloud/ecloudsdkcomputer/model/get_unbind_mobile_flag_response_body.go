// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUnbindMobileFlagResponseBody struct {

    // 用户解绑手机号开关 默认 1 :开 0:关 
	UnbindMobileFlag *int32 `json:"unbindMobileFlag,omitempty"`
}

func (s GetUnbindMobileFlagResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetUnbindMobileFlagResponseBody) GoString() string {
	return s.String()
}

func (s GetUnbindMobileFlagResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUnbindMobileFlagResponseBody) SetUnbindMobileFlag(v int32) *GetUnbindMobileFlagResponseBody {
	s.UnbindMobileFlag = &v
	return s
}


type GetUnbindMobileFlagResponseBodyBuilder struct {
	s *GetUnbindMobileFlagResponseBody
}

func NewGetUnbindMobileFlagResponseBodyBuilder() *GetUnbindMobileFlagResponseBodyBuilder {
	s := &GetUnbindMobileFlagResponseBody{}
	b := &GetUnbindMobileFlagResponseBodyBuilder{s: s}
	return b
}

func (b *GetUnbindMobileFlagResponseBodyBuilder) UnbindMobileFlag(v int32) *GetUnbindMobileFlagResponseBodyBuilder {
	b.s.UnbindMobileFlag = &v
	return b
}

func (b *GetUnbindMobileFlagResponseBodyBuilder) Build() *GetUnbindMobileFlagResponseBody {
	return b.s
}


