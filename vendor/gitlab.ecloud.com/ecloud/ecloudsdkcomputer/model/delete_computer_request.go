// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteComputerRequest struct {


	DeleteComputerBody *DeleteComputerBody `json:"deleteComputerBody,omitempty"`
}

func (s DeleteComputerRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteComputerRequest) GoString() string {
	return s.String()
}

func (s DeleteComputerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteComputerRequest) SetDeleteComputerBody(v *DeleteComputerBody) *DeleteComputerRequest {
	s.DeleteComputerBody = v
	return s
}


type DeleteComputerRequestBuilder struct {
	s *DeleteComputerRequest
}

func NewDeleteComputerRequestBuilder() *DeleteComputerRequestBuilder {
	s := &DeleteComputerRequest{}
	b := &DeleteComputerRequestBuilder{s: s}
	return b
}

func (b *DeleteComputerRequestBuilder) DeleteComputerBody(v *DeleteComputerBody) *DeleteComputerRequestBuilder {
	b.s.DeleteComputerBody = v
	return b
}

func (b *DeleteComputerRequestBuilder) Build() *DeleteComputerRequest {
	return b.s
}


