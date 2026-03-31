// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteNetworkRequest struct {


	DeleteNetworkBody *DeleteNetworkBody `json:"deleteNetworkBody,omitempty"`
}

func (s DeleteNetworkRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteNetworkRequest) GoString() string {
	return s.String()
}

func (s DeleteNetworkRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteNetworkRequest) SetDeleteNetworkBody(v *DeleteNetworkBody) *DeleteNetworkRequest {
	s.DeleteNetworkBody = v
	return s
}


type DeleteNetworkRequestBuilder struct {
	s *DeleteNetworkRequest
}

func NewDeleteNetworkRequestBuilder() *DeleteNetworkRequestBuilder {
	s := &DeleteNetworkRequest{}
	b := &DeleteNetworkRequestBuilder{s: s}
	return b
}

func (b *DeleteNetworkRequestBuilder) DeleteNetworkBody(v *DeleteNetworkBody) *DeleteNetworkRequestBuilder {
	b.s.DeleteNetworkBody = v
	return b
}

func (b *DeleteNetworkRequestBuilder) Build() *DeleteNetworkRequest {
	return b.s
}


