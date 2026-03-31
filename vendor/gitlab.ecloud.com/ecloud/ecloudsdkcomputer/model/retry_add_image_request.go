// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RetryAddImageRequest struct {


	RetryAddImageBody *RetryAddImageBody `json:"retryAddImageBody,omitempty"`
}

func (s RetryAddImageRequest) String() string {
	return utils.Beautify(s)
}

func (s RetryAddImageRequest) GoString() string {
	return s.String()
}

func (s RetryAddImageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RetryAddImageRequest) SetRetryAddImageBody(v *RetryAddImageBody) *RetryAddImageRequest {
	s.RetryAddImageBody = v
	return s
}


type RetryAddImageRequestBuilder struct {
	s *RetryAddImageRequest
}

func NewRetryAddImageRequestBuilder() *RetryAddImageRequestBuilder {
	s := &RetryAddImageRequest{}
	b := &RetryAddImageRequestBuilder{s: s}
	return b
}

func (b *RetryAddImageRequestBuilder) RetryAddImageBody(v *RetryAddImageBody) *RetryAddImageRequestBuilder {
	b.s.RetryAddImageBody = v
	return b
}

func (b *RetryAddImageRequestBuilder) Build() *RetryAddImageRequest {
	return b.s
}


