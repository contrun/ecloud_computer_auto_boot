// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SdkCipherResponseBody struct {

    // sdk认证信息
	AuthInfo *string `json:"authInfo,omitempty"`
}

func (s SdkCipherResponseBody) String() string {
	return utils.Beautify(s)
}

func (s SdkCipherResponseBody) GoString() string {
	return s.String()
}

func (s SdkCipherResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SdkCipherResponseBody) SetAuthInfo(v string) *SdkCipherResponseBody {
	s.AuthInfo = &v
	return s
}


type SdkCipherResponseBodyBuilder struct {
	s *SdkCipherResponseBody
}

func NewSdkCipherResponseBodyBuilder() *SdkCipherResponseBodyBuilder {
	s := &SdkCipherResponseBody{}
	b := &SdkCipherResponseBodyBuilder{s: s}
	return b
}

func (b *SdkCipherResponseBodyBuilder) AuthInfo(v string) *SdkCipherResponseBodyBuilder {
	b.s.AuthInfo = &v
	return b
}

func (b *SdkCipherResponseBodyBuilder) Build() *SdkCipherResponseBody {
	return b.s
}


