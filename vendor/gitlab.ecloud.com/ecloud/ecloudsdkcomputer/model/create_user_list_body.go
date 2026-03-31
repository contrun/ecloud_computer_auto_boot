// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateUserListBody struct {
    position.Body
    // 激活用户信息列表
	AddUserVoList *[]CreateUserListRequestAddUserVoList `json:"addUserVoList,omitempty"`
}

func (s CreateUserListBody) String() string {
	return utils.Beautify(s)
}

func (s CreateUserListBody) GoString() string {
	return s.String()
}

func (s CreateUserListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateUserListBody) SetAddUserVoList(v []CreateUserListRequestAddUserVoList) *CreateUserListBody {
	s.AddUserVoList = &v
	return s
}


type CreateUserListBodyBuilder struct {
	s *CreateUserListBody
}

func NewCreateUserListBodyBuilder() *CreateUserListBodyBuilder {
	s := &CreateUserListBody{}
	b := &CreateUserListBodyBuilder{s: s}
	return b
}

func (b *CreateUserListBodyBuilder) AddUserVoList(v []CreateUserListRequestAddUserVoList) *CreateUserListBodyBuilder {
	b.s.AddUserVoList = &v
	return b
}

func (b *CreateUserListBodyBuilder) Build() *CreateUserListBody {
	return b.s
}


