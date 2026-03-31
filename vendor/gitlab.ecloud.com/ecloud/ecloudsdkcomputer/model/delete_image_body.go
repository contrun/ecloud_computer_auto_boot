// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteImageBody struct {
    position.Body
    // 镜像模板ID
	TemplateId *string `json:"templateId,omitempty"`
}

func (s DeleteImageBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteImageBody) GoString() string {
	return s.String()
}

func (s DeleteImageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteImageBody) SetTemplateId(v string) *DeleteImageBody {
	s.TemplateId = &v
	return s
}


type DeleteImageBodyBuilder struct {
	s *DeleteImageBody
}

func NewDeleteImageBodyBuilder() *DeleteImageBodyBuilder {
	s := &DeleteImageBody{}
	b := &DeleteImageBodyBuilder{s: s}
	return b
}

func (b *DeleteImageBodyBuilder) TemplateId(v string) *DeleteImageBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *DeleteImageBodyBuilder) Build() *DeleteImageBody {
	return b.s
}


