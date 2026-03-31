// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserInfoListBody struct {
    position.Body
    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 备注
	Description *string `json:"description,omitempty"`
    // 分页大小
	PageSize *string `json:"pageSize,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 当前页
	PageNum *string `json:"pageNum,omitempty"`
}

func (s GetUserInfoListBody) String() string {
	return utils.Beautify(s)
}

func (s GetUserInfoListBody) GoString() string {
	return s.String()
}

func (s GetUserInfoListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserInfoListBody) SetMail(v string) *GetUserInfoListBody {
	s.Mail = &v
	return s
}

func (s *GetUserInfoListBody) SetDescription(v string) *GetUserInfoListBody {
	s.Description = &v
	return s
}

func (s *GetUserInfoListBody) SetPageSize(v string) *GetUserInfoListBody {
	s.PageSize = &v
	return s
}

func (s *GetUserInfoListBody) SetUserName(v string) *GetUserInfoListBody {
	s.UserName = &v
	return s
}

func (s *GetUserInfoListBody) SetPageNum(v string) *GetUserInfoListBody {
	s.PageNum = &v
	return s
}


type GetUserInfoListBodyBuilder struct {
	s *GetUserInfoListBody
}

func NewGetUserInfoListBodyBuilder() *GetUserInfoListBodyBuilder {
	s := &GetUserInfoListBody{}
	b := &GetUserInfoListBodyBuilder{s: s}
	return b
}

func (b *GetUserInfoListBodyBuilder) Mail(v string) *GetUserInfoListBodyBuilder {
	b.s.Mail = &v
	return b
}

func (b *GetUserInfoListBodyBuilder) Description(v string) *GetUserInfoListBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *GetUserInfoListBodyBuilder) PageSize(v string) *GetUserInfoListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetUserInfoListBodyBuilder) UserName(v string) *GetUserInfoListBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetUserInfoListBodyBuilder) PageNum(v string) *GetUserInfoListBodyBuilder {
	b.s.PageNum = &v
	return b
}

func (b *GetUserInfoListBodyBuilder) Build() *GetUserInfoListBody {
	return b.s
}


