// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteUserBody struct {
    position.Body
    // 用户ID列表
	UserIdList []string `json:"userIdList,omitempty"`
}

func (s BatchDeleteUserBody) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteUserBody) GoString() string {
	return s.String()
}

func (s BatchDeleteUserBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteUserBody) SetUserIdList(v []string) *BatchDeleteUserBody {
	s.UserIdList = v
	return s
}


type BatchDeleteUserBodyBuilder struct {
	s *BatchDeleteUserBody
}

func NewBatchDeleteUserBodyBuilder() *BatchDeleteUserBodyBuilder {
	s := &BatchDeleteUserBody{}
	b := &BatchDeleteUserBodyBuilder{s: s}
	return b
}

func (b *BatchDeleteUserBodyBuilder) UserIdList(v []string) *BatchDeleteUserBodyBuilder {
	b.s.UserIdList = v
	return b
}

func (b *BatchDeleteUserBodyBuilder) Build() *BatchDeleteUserBody {
	return b.s
}


