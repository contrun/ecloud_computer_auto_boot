// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RetryAddImageBody struct {
    position.Body
    // 镜像模板ID
	TemplateId *string `json:"templateId,omitempty"`
}

func (s RetryAddImageBody) String() string {
	return utils.Beautify(s)
}

func (s RetryAddImageBody) GoString() string {
	return s.String()
}

func (s RetryAddImageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RetryAddImageBody) SetTemplateId(v string) *RetryAddImageBody {
	s.TemplateId = &v
	return s
}


type RetryAddImageBodyBuilder struct {
	s *RetryAddImageBody
}

func NewRetryAddImageBodyBuilder() *RetryAddImageBodyBuilder {
	s := &RetryAddImageBody{}
	b := &RetryAddImageBodyBuilder{s: s}
	return b
}

func (b *RetryAddImageBodyBuilder) TemplateId(v string) *RetryAddImageBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *RetryAddImageBodyBuilder) Build() *RetryAddImageBody {
	return b.s
}


