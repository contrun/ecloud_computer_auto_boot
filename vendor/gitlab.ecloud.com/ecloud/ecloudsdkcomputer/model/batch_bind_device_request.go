// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchBindDeviceRequest struct {


	BatchBindDeviceBody *BatchBindDeviceBody `json:"batchBindDeviceBody,omitempty"`
}

func (s BatchBindDeviceRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchBindDeviceRequest) GoString() string {
	return s.String()
}

func (s BatchBindDeviceRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchBindDeviceRequest) SetBatchBindDeviceBody(v *BatchBindDeviceBody) *BatchBindDeviceRequest {
	s.BatchBindDeviceBody = v
	return s
}


type BatchBindDeviceRequestBuilder struct {
	s *BatchBindDeviceRequest
}

func NewBatchBindDeviceRequestBuilder() *BatchBindDeviceRequestBuilder {
	s := &BatchBindDeviceRequest{}
	b := &BatchBindDeviceRequestBuilder{s: s}
	return b
}

func (b *BatchBindDeviceRequestBuilder) BatchBindDeviceBody(v *BatchBindDeviceBody) *BatchBindDeviceRequestBuilder {
	b.s.BatchBindDeviceBody = v
	return b
}

func (b *BatchBindDeviceRequestBuilder) Build() *BatchBindDeviceRequest {
	return b.s
}


