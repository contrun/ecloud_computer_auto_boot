// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetShowOrderDetailsResponseBody struct {

    // 客户端订单信息展示控制开关 默认 1 :开 0:关 
	ShowOrderDetails *int32 `json:"showOrderDetails,omitempty"`
}

func (s GetShowOrderDetailsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetShowOrderDetailsResponseBody) GoString() string {
	return s.String()
}

func (s GetShowOrderDetailsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShowOrderDetailsResponseBody) SetShowOrderDetails(v int32) *GetShowOrderDetailsResponseBody {
	s.ShowOrderDetails = &v
	return s
}


type GetShowOrderDetailsResponseBodyBuilder struct {
	s *GetShowOrderDetailsResponseBody
}

func NewGetShowOrderDetailsResponseBodyBuilder() *GetShowOrderDetailsResponseBodyBuilder {
	s := &GetShowOrderDetailsResponseBody{}
	b := &GetShowOrderDetailsResponseBodyBuilder{s: s}
	return b
}

func (b *GetShowOrderDetailsResponseBodyBuilder) ShowOrderDetails(v int32) *GetShowOrderDetailsResponseBodyBuilder {
	b.s.ShowOrderDetails = &v
	return b
}

func (b *GetShowOrderDetailsResponseBodyBuilder) Build() *GetShowOrderDetailsResponseBody {
	return b.s
}


