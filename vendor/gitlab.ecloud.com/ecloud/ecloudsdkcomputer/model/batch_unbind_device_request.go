// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchUnbindDeviceRequest struct {


	BatchUnbindDeviceBody *BatchUnbindDeviceBody `json:"batchUnbindDeviceBody,omitempty"`
}

func (s BatchUnbindDeviceRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchUnbindDeviceRequest) GoString() string {
	return s.String()
}

func (s BatchUnbindDeviceRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUnbindDeviceRequest) SetBatchUnbindDeviceBody(v *BatchUnbindDeviceBody) *BatchUnbindDeviceRequest {
	s.BatchUnbindDeviceBody = v
	return s
}


type BatchUnbindDeviceRequestBuilder struct {
	s *BatchUnbindDeviceRequest
}

func NewBatchUnbindDeviceRequestBuilder() *BatchUnbindDeviceRequestBuilder {
	s := &BatchUnbindDeviceRequest{}
	b := &BatchUnbindDeviceRequestBuilder{s: s}
	return b
}

func (b *BatchUnbindDeviceRequestBuilder) BatchUnbindDeviceBody(v *BatchUnbindDeviceBody) *BatchUnbindDeviceRequestBuilder {
	b.s.BatchUnbindDeviceBody = v
	return b
}

func (b *BatchUnbindDeviceRequestBuilder) Build() *BatchUnbindDeviceRequest {
	return b.s
}


