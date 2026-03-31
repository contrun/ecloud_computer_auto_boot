// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteRequest struct {


	DeleteBody *DeleteBody `json:"deleteBody,omitempty"`
}

func (s DeleteRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteRequest) GoString() string {
	return s.String()
}

func (s DeleteRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteRequest) SetDeleteBody(v *DeleteBody) *DeleteRequest {
	s.DeleteBody = v
	return s
}


type DeleteRequestBuilder struct {
	s *DeleteRequest
}

func NewDeleteRequestBuilder() *DeleteRequestBuilder {
	s := &DeleteRequest{}
	b := &DeleteRequestBuilder{s: s}
	return b
}

func (b *DeleteRequestBuilder) DeleteBody(v *DeleteBody) *DeleteRequestBuilder {
	b.s.DeleteBody = v
	return b
}

func (b *DeleteRequestBuilder) Build() *DeleteRequest {
	return b.s
}


