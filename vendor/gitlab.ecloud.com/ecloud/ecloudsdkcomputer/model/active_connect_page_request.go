// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ActiveConnectPageRequest struct {


	ActiveConnectPageBody *ActiveConnectPageBody `json:"activeConnectPageBody,omitempty"`
}

func (s ActiveConnectPageRequest) String() string {
	return utils.Beautify(s)
}

func (s ActiveConnectPageRequest) GoString() string {
	return s.String()
}

func (s ActiveConnectPageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActiveConnectPageRequest) SetActiveConnectPageBody(v *ActiveConnectPageBody) *ActiveConnectPageRequest {
	s.ActiveConnectPageBody = v
	return s
}


type ActiveConnectPageRequestBuilder struct {
	s *ActiveConnectPageRequest
}

func NewActiveConnectPageRequestBuilder() *ActiveConnectPageRequestBuilder {
	s := &ActiveConnectPageRequest{}
	b := &ActiveConnectPageRequestBuilder{s: s}
	return b
}

func (b *ActiveConnectPageRequestBuilder) ActiveConnectPageBody(v *ActiveConnectPageBody) *ActiveConnectPageRequestBuilder {
	b.s.ActiveConnectPageBody = v
	return b
}

func (b *ActiveConnectPageRequestBuilder) Build() *ActiveConnectPageRequest {
	return b.s
}


