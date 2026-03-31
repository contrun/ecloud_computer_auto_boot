// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnLockDeviceByIdRequest struct {


	UnLockDeviceByIdBody *UnLockDeviceByIdBody `json:"unLockDeviceByIdBody,omitempty"`
}

func (s UnLockDeviceByIdRequest) String() string {
	return utils.Beautify(s)
}

func (s UnLockDeviceByIdRequest) GoString() string {
	return s.String()
}

func (s UnLockDeviceByIdRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnLockDeviceByIdRequest) SetUnLockDeviceByIdBody(v *UnLockDeviceByIdBody) *UnLockDeviceByIdRequest {
	s.UnLockDeviceByIdBody = v
	return s
}


type UnLockDeviceByIdRequestBuilder struct {
	s *UnLockDeviceByIdRequest
}

func NewUnLockDeviceByIdRequestBuilder() *UnLockDeviceByIdRequestBuilder {
	s := &UnLockDeviceByIdRequest{}
	b := &UnLockDeviceByIdRequestBuilder{s: s}
	return b
}

func (b *UnLockDeviceByIdRequestBuilder) UnLockDeviceByIdBody(v *UnLockDeviceByIdBody) *UnLockDeviceByIdRequestBuilder {
	b.s.UnLockDeviceByIdBody = v
	return b
}

func (b *UnLockDeviceByIdRequestBuilder) Build() *UnLockDeviceByIdRequest {
	return b.s
}


