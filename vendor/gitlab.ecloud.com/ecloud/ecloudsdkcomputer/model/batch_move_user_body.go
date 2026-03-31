// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchMoveUserBody struct {
    position.Body
    // 移动用户ID列表
	UserIdList []string `json:"userIdList,omitempty"`
    // 组织id
	OrgId *string `json:"orgId,omitempty"`
}

func (s BatchMoveUserBody) String() string {
	return utils.Beautify(s)
}

func (s BatchMoveUserBody) GoString() string {
	return s.String()
}

func (s BatchMoveUserBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchMoveUserBody) SetUserIdList(v []string) *BatchMoveUserBody {
	s.UserIdList = v
	return s
}

func (s *BatchMoveUserBody) SetOrgId(v string) *BatchMoveUserBody {
	s.OrgId = &v
	return s
}


type BatchMoveUserBodyBuilder struct {
	s *BatchMoveUserBody
}

func NewBatchMoveUserBodyBuilder() *BatchMoveUserBodyBuilder {
	s := &BatchMoveUserBody{}
	b := &BatchMoveUserBodyBuilder{s: s}
	return b
}

func (b *BatchMoveUserBodyBuilder) UserIdList(v []string) *BatchMoveUserBodyBuilder {
	b.s.UserIdList = v
	return b
}

func (b *BatchMoveUserBodyBuilder) OrgId(v string) *BatchMoveUserBodyBuilder {
	b.s.OrgId = &v
	return b
}

func (b *BatchMoveUserBodyBuilder) Build() *BatchMoveUserBody {
	return b.s
}


