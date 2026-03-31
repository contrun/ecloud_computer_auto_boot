// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LockDeviceByIdRequest struct {


	LockDeviceByIdBody *LockDeviceByIdBody `json:"lockDeviceByIdBody,omitempty"`
}

func (s LockDeviceByIdRequest) String() string {
	return utils.Beautify(s)
}

func (s LockDeviceByIdRequest) GoString() string {
	return s.String()
}

func (s LockDeviceByIdRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LockDeviceByIdRequest) SetLockDeviceByIdBody(v *LockDeviceByIdBody) *LockDeviceByIdRequest {
	s.LockDeviceByIdBody = v
	return s
}


type LockDeviceByIdRequestBuilder struct {
	s *LockDeviceByIdRequest
}

func NewLockDeviceByIdRequestBuilder() *LockDeviceByIdRequestBuilder {
	s := &LockDeviceByIdRequest{}
	b := &LockDeviceByIdRequestBuilder{s: s}
	return b
}

func (b *LockDeviceByIdRequestBuilder) LockDeviceByIdBody(v *LockDeviceByIdBody) *LockDeviceByIdRequestBuilder {
	b.s.LockDeviceByIdBody = v
	return b
}

func (b *LockDeviceByIdRequestBuilder) Build() *LockDeviceByIdRequest {
	return b.s
}


