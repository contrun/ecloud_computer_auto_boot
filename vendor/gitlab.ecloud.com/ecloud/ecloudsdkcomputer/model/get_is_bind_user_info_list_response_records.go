// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetIsBindUserInfoListResponseRecords struct {

    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 用户状态
	AuthorizationState *int32 `json:"authorizationState,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
}

func (s GetIsBindUserInfoListResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s GetIsBindUserInfoListResponseRecords) GoString() string {
	return s.String()
}

func (s GetIsBindUserInfoListResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetIsBindUserInfoListResponseRecords) SetMail(v string) *GetIsBindUserInfoListResponseRecords {
	s.Mail = &v
	return s
}

func (s *GetIsBindUserInfoListResponseRecords) SetMobile(v string) *GetIsBindUserInfoListResponseRecords {
	s.Mobile = &v
	return s
}

func (s *GetIsBindUserInfoListResponseRecords) SetAuthorizationState(v int32) *GetIsBindUserInfoListResponseRecords {
	s.AuthorizationState = &v
	return s
}

func (s *GetIsBindUserInfoListResponseRecords) SetUserName(v string) *GetIsBindUserInfoListResponseRecords {
	s.UserName = &v
	return s
}


type GetIsBindUserInfoListResponseRecordsBuilder struct {
	s *GetIsBindUserInfoListResponseRecords
}

func NewGetIsBindUserInfoListResponseRecordsBuilder() *GetIsBindUserInfoListResponseRecordsBuilder {
	s := &GetIsBindUserInfoListResponseRecords{}
	b := &GetIsBindUserInfoListResponseRecordsBuilder{s: s}
	return b
}

func (b *GetIsBindUserInfoListResponseRecordsBuilder) Mail(v string) *GetIsBindUserInfoListResponseRecordsBuilder {
	b.s.Mail = &v
	return b
}

func (b *GetIsBindUserInfoListResponseRecordsBuilder) Mobile(v string) *GetIsBindUserInfoListResponseRecordsBuilder {
	b.s.Mobile = &v
	return b
}

func (b *GetIsBindUserInfoListResponseRecordsBuilder) AuthorizationState(v int32) *GetIsBindUserInfoListResponseRecordsBuilder {
	b.s.AuthorizationState = &v
	return b
}

func (b *GetIsBindUserInfoListResponseRecordsBuilder) UserName(v string) *GetIsBindUserInfoListResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetIsBindUserInfoListResponseRecordsBuilder) Build() *GetIsBindUserInfoListResponseRecords {
	return b.s
}


