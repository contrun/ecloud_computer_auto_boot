// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetIsBindUserInfoListBody struct {
    position.Body
    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 页面大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 设备主键id
	Id *string `json:"id,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 用户授权状态
	AuthorizationState *int32 `json:"authorizationState,omitempty"`
    // 当前页
	PageNum *int32 `json:"pageNum,omitempty"`
}

func (s GetIsBindUserInfoListBody) String() string {
	return utils.Beautify(s)
}

func (s GetIsBindUserInfoListBody) GoString() string {
	return s.String()
}

func (s GetIsBindUserInfoListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetIsBindUserInfoListBody) SetMail(v string) *GetIsBindUserInfoListBody {
	s.Mail = &v
	return s
}

func (s *GetIsBindUserInfoListBody) SetMobile(v string) *GetIsBindUserInfoListBody {
	s.Mobile = &v
	return s
}

func (s *GetIsBindUserInfoListBody) SetPageSize(v int32) *GetIsBindUserInfoListBody {
	s.PageSize = &v
	return s
}

func (s *GetIsBindUserInfoListBody) SetId(v string) *GetIsBindUserInfoListBody {
	s.Id = &v
	return s
}

func (s *GetIsBindUserInfoListBody) SetUserName(v string) *GetIsBindUserInfoListBody {
	s.UserName = &v
	return s
}

func (s *GetIsBindUserInfoListBody) SetAuthorizationState(v int32) *GetIsBindUserInfoListBody {
	s.AuthorizationState = &v
	return s
}

func (s *GetIsBindUserInfoListBody) SetPageNum(v int32) *GetIsBindUserInfoListBody {
	s.PageNum = &v
	return s
}


type GetIsBindUserInfoListBodyBuilder struct {
	s *GetIsBindUserInfoListBody
}

func NewGetIsBindUserInfoListBodyBuilder() *GetIsBindUserInfoListBodyBuilder {
	s := &GetIsBindUserInfoListBody{}
	b := &GetIsBindUserInfoListBodyBuilder{s: s}
	return b
}

func (b *GetIsBindUserInfoListBodyBuilder) Mail(v string) *GetIsBindUserInfoListBodyBuilder {
	b.s.Mail = &v
	return b
}

func (b *GetIsBindUserInfoListBodyBuilder) Mobile(v string) *GetIsBindUserInfoListBodyBuilder {
	b.s.Mobile = &v
	return b
}

func (b *GetIsBindUserInfoListBodyBuilder) PageSize(v int32) *GetIsBindUserInfoListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetIsBindUserInfoListBodyBuilder) Id(v string) *GetIsBindUserInfoListBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetIsBindUserInfoListBodyBuilder) UserName(v string) *GetIsBindUserInfoListBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetIsBindUserInfoListBodyBuilder) AuthorizationState(v int32) *GetIsBindUserInfoListBodyBuilder {
	b.s.AuthorizationState = &v
	return b
}

func (b *GetIsBindUserInfoListBodyBuilder) PageNum(v int32) *GetIsBindUserInfoListBodyBuilder {
	b.s.PageNum = &v
	return b
}

func (b *GetIsBindUserInfoListBodyBuilder) Build() *GetIsBindUserInfoListBody {
	return b.s
}


