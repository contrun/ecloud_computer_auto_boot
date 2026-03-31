// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetShowOrderDetailsBody struct {
    position.Body
    // 客户端订单信息展示控制开关 默认 1 :开 0:关 
	ShowOrderDetails *int32 `json:"showOrderDetails,omitempty"`
}

func (s SetShowOrderDetailsBody) String() string {
	return utils.Beautify(s)
}

func (s SetShowOrderDetailsBody) GoString() string {
	return s.String()
}

func (s SetShowOrderDetailsBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetShowOrderDetailsBody) SetShowOrderDetails(v int32) *SetShowOrderDetailsBody {
	s.ShowOrderDetails = &v
	return s
}


type SetShowOrderDetailsBodyBuilder struct {
	s *SetShowOrderDetailsBody
}

func NewSetShowOrderDetailsBodyBuilder() *SetShowOrderDetailsBodyBuilder {
	s := &SetShowOrderDetailsBody{}
	b := &SetShowOrderDetailsBodyBuilder{s: s}
	return b
}

func (b *SetShowOrderDetailsBodyBuilder) ShowOrderDetails(v int32) *SetShowOrderDetailsBodyBuilder {
	b.s.ShowOrderDetails = &v
	return b
}

func (b *SetShowOrderDetailsBodyBuilder) Build() *SetShowOrderDetailsBody {
	return b.s
}


