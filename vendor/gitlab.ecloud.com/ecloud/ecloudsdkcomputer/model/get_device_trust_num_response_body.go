// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDeviceTrustNumResponseBody struct {

    // 可信认证设备数量
	DeviceTrustNum *int32 `json:"deviceTrustNum,omitempty"`
}

func (s GetDeviceTrustNumResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetDeviceTrustNumResponseBody) GoString() string {
	return s.String()
}

func (s GetDeviceTrustNumResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDeviceTrustNumResponseBody) SetDeviceTrustNum(v int32) *GetDeviceTrustNumResponseBody {
	s.DeviceTrustNum = &v
	return s
}


type GetDeviceTrustNumResponseBodyBuilder struct {
	s *GetDeviceTrustNumResponseBody
}

func NewGetDeviceTrustNumResponseBodyBuilder() *GetDeviceTrustNumResponseBodyBuilder {
	s := &GetDeviceTrustNumResponseBody{}
	b := &GetDeviceTrustNumResponseBodyBuilder{s: s}
	return b
}

func (b *GetDeviceTrustNumResponseBodyBuilder) DeviceTrustNum(v int32) *GetDeviceTrustNumResponseBodyBuilder {
	b.s.DeviceTrustNum = &v
	return b
}

func (b *GetDeviceTrustNumResponseBodyBuilder) Build() *GetDeviceTrustNumResponseBody {
	return b.s
}


