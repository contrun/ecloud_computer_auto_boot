// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetIPByUserNameQuery struct {
    position.Query
    // 登录账号
	UserName *string `json:"userName,omitempty"`
}

func (s GetIPByUserNameQuery) String() string {
	return utils.Beautify(s)
}

func (s GetIPByUserNameQuery) GoString() string {
	return s.String()
}

func (s GetIPByUserNameQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetIPByUserNameQuery) SetUserName(v string) *GetIPByUserNameQuery {
	s.UserName = &v
	return s
}


type GetIPByUserNameQueryBuilder struct {
	s *GetIPByUserNameQuery
}

func NewGetIPByUserNameQueryBuilder() *GetIPByUserNameQueryBuilder {
	s := &GetIPByUserNameQuery{}
	b := &GetIPByUserNameQueryBuilder{s: s}
	return b
}

func (b *GetIPByUserNameQueryBuilder) UserName(v string) *GetIPByUserNameQueryBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetIPByUserNameQueryBuilder) Build() *GetIPByUserNameQuery {
	return b.s
}


