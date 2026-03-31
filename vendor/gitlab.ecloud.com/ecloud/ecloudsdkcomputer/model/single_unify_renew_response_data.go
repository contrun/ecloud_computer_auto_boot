// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SingleUnifyRenewResponseData struct {

    // 资源id
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单编号
	OrderNo *string `json:"orderNo,omitempty"`
}

func (s SingleUnifyRenewResponseData) String() string {
	return utils.Beautify(s)
}

func (s SingleUnifyRenewResponseData) GoString() string {
	return s.String()
}

func (s SingleUnifyRenewResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SingleUnifyRenewResponseData) SetInstanceId(v string) *SingleUnifyRenewResponseData {
	s.InstanceId = &v
	return s
}

func (s *SingleUnifyRenewResponseData) SetOrderNo(v string) *SingleUnifyRenewResponseData {
	s.OrderNo = &v
	return s
}


type SingleUnifyRenewResponseDataBuilder struct {
	s *SingleUnifyRenewResponseData
}

func NewSingleUnifyRenewResponseDataBuilder() *SingleUnifyRenewResponseDataBuilder {
	s := &SingleUnifyRenewResponseData{}
	b := &SingleUnifyRenewResponseDataBuilder{s: s}
	return b
}

func (b *SingleUnifyRenewResponseDataBuilder) InstanceId(v string) *SingleUnifyRenewResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *SingleUnifyRenewResponseDataBuilder) OrderNo(v string) *SingleUnifyRenewResponseDataBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *SingleUnifyRenewResponseDataBuilder) Build() *SingleUnifyRenewResponseData {
	return b.s
}


