// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LogoutBody struct {
    position.Body
    // accessToken
	AccessToken *string `json:"accessToken,omitempty"`
}

func (s LogoutBody) String() string {
	return utils.Beautify(s)
}

func (s LogoutBody) GoString() string {
	return s.String()
}

func (s LogoutBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LogoutBody) SetAccessToken(v string) *LogoutBody {
	s.AccessToken = &v
	return s
}


type LogoutBodyBuilder struct {
	s *LogoutBody
}

func NewLogoutBodyBuilder() *LogoutBodyBuilder {
	s := &LogoutBody{}
	b := &LogoutBodyBuilder{s: s}
	return b
}

func (b *LogoutBodyBuilder) AccessToken(v string) *LogoutBodyBuilder {
	b.s.AccessToken = &v
	return b
}

func (b *LogoutBodyBuilder) Build() *LogoutBody {
	return b.s
}


