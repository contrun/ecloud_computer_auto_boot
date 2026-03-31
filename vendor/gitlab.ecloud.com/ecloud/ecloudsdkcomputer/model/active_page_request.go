// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ActivePageRequest struct {


	ActivePageBody *ActivePageBody `json:"activePageBody,omitempty"`
}

func (s ActivePageRequest) String() string {
	return utils.Beautify(s)
}

func (s ActivePageRequest) GoString() string {
	return s.String()
}

func (s ActivePageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActivePageRequest) SetActivePageBody(v *ActivePageBody) *ActivePageRequest {
	s.ActivePageBody = v
	return s
}


type ActivePageRequestBuilder struct {
	s *ActivePageRequest
}

func NewActivePageRequestBuilder() *ActivePageRequestBuilder {
	s := &ActivePageRequest{}
	b := &ActivePageRequestBuilder{s: s}
	return b
}

func (b *ActivePageRequestBuilder) ActivePageBody(v *ActivePageBody) *ActivePageRequestBuilder {
	b.s.ActivePageBody = v
	return b
}

func (b *ActivePageRequestBuilder) Build() *ActivePageRequest {
	return b.s
}


