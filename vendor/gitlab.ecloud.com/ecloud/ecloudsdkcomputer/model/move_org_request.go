// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MoveOrgRequest struct {


	MoveOrgBody *MoveOrgBody `json:"moveOrgBody,omitempty"`
}

func (s MoveOrgRequest) String() string {
	return utils.Beautify(s)
}

func (s MoveOrgRequest) GoString() string {
	return s.String()
}

func (s MoveOrgRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MoveOrgRequest) SetMoveOrgBody(v *MoveOrgBody) *MoveOrgRequest {
	s.MoveOrgBody = v
	return s
}


type MoveOrgRequestBuilder struct {
	s *MoveOrgRequest
}

func NewMoveOrgRequestBuilder() *MoveOrgRequestBuilder {
	s := &MoveOrgRequest{}
	b := &MoveOrgRequestBuilder{s: s}
	return b
}

func (b *MoveOrgRequestBuilder) MoveOrgBody(v *MoveOrgBody) *MoveOrgRequestBuilder {
	b.s.MoveOrgBody = v
	return b
}

func (b *MoveOrgRequestBuilder) Build() *MoveOrgRequest {
	return b.s
}


