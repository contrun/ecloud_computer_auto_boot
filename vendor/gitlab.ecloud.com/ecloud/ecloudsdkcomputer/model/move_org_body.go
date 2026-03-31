// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MoveOrgBody struct {
    position.Body
    // 被移动组织id
	SourceId *string `json:"sourceId,omitempty"`
    // 目标组织id
	TargetId *string `json:"targetId,omitempty"`
}

func (s MoveOrgBody) String() string {
	return utils.Beautify(s)
}

func (s MoveOrgBody) GoString() string {
	return s.String()
}

func (s MoveOrgBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MoveOrgBody) SetSourceId(v string) *MoveOrgBody {
	s.SourceId = &v
	return s
}

func (s *MoveOrgBody) SetTargetId(v string) *MoveOrgBody {
	s.TargetId = &v
	return s
}


type MoveOrgBodyBuilder struct {
	s *MoveOrgBody
}

func NewMoveOrgBodyBuilder() *MoveOrgBodyBuilder {
	s := &MoveOrgBody{}
	b := &MoveOrgBodyBuilder{s: s}
	return b
}

func (b *MoveOrgBodyBuilder) SourceId(v string) *MoveOrgBodyBuilder {
	b.s.SourceId = &v
	return b
}

func (b *MoveOrgBodyBuilder) TargetId(v string) *MoveOrgBodyBuilder {
	b.s.TargetId = &v
	return b
}

func (b *MoveOrgBodyBuilder) Build() *MoveOrgBody {
	return b.s
}


