// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConfirmDoSignBody struct {
    position.Body
    // 订单号。示例值：1549224402274672641
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s ConfirmDoSignBody) String() string {
	return utils.Beautify(s)
}

func (s ConfirmDoSignBody) GoString() string {
	return s.String()
}

func (s ConfirmDoSignBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConfirmDoSignBody) SetOrderNo(v string) *ConfirmDoSignBody {
	s.OrderNo = &v
	return s
}


type ConfirmDoSignBodyBuilder struct {
	s *ConfirmDoSignBody
}

func NewConfirmDoSignBodyBuilder() *ConfirmDoSignBodyBuilder {
	s := &ConfirmDoSignBody{}
	b := &ConfirmDoSignBodyBuilder{s: s}
	return b
}

func (b *ConfirmDoSignBodyBuilder) OrderNo(v string) *ConfirmDoSignBodyBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *ConfirmDoSignBodyBuilder) Build() *ConfirmDoSignBody {
	return b.s
}


