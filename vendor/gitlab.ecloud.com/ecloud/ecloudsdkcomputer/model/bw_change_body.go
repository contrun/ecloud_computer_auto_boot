// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BwChangeBody struct {
    position.Body
    // 实例id,示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
    // 带宽大小
	BwSize *int32 `json:"bwSize,omitempty"`
}

func (s BwChangeBody) String() string {
	return utils.Beautify(s)
}

func (s BwChangeBody) GoString() string {
	return s.String()
}

func (s BwChangeBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BwChangeBody) SetInstanceId(v string) *BwChangeBody {
	s.InstanceId = &v
	return s
}

func (s *BwChangeBody) SetBwSize(v int32) *BwChangeBody {
	s.BwSize = &v
	return s
}


type BwChangeBodyBuilder struct {
	s *BwChangeBody
}

func NewBwChangeBodyBuilder() *BwChangeBodyBuilder {
	s := &BwChangeBody{}
	b := &BwChangeBodyBuilder{s: s}
	return b
}

func (b *BwChangeBodyBuilder) InstanceId(v string) *BwChangeBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *BwChangeBodyBuilder) BwSize(v int32) *BwChangeBodyBuilder {
	b.s.BwSize = &v
	return b
}

func (b *BwChangeBodyBuilder) Build() *BwChangeBody {
	return b.s
}


