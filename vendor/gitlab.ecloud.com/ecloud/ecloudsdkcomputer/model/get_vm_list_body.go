// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVmListBody struct {
    position.Body
    // accessToken
	AccessToken *string `json:"accessToken,omitempty"`
}

func (s GetVmListBody) String() string {
	return utils.Beautify(s)
}

func (s GetVmListBody) GoString() string {
	return s.String()
}

func (s GetVmListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVmListBody) SetAccessToken(v string) *GetVmListBody {
	s.AccessToken = &v
	return s
}


type GetVmListBodyBuilder struct {
	s *GetVmListBody
}

func NewGetVmListBodyBuilder() *GetVmListBodyBuilder {
	s := &GetVmListBody{}
	b := &GetVmListBodyBuilder{s: s}
	return b
}

func (b *GetVmListBodyBuilder) AccessToken(v string) *GetVmListBodyBuilder {
	b.s.AccessToken = &v
	return b
}

func (b *GetVmListBodyBuilder) Build() *GetVmListBody {
	return b.s
}


