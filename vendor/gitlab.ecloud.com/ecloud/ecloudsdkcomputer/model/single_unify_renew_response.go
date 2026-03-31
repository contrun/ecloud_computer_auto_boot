// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SingleUnifyRenewResponse struct {

    // 响应数据
	Data *SingleUnifyRenewResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s SingleUnifyRenewResponse) String() string {
	return utils.Beautify(s)
}

func (s SingleUnifyRenewResponse) GoString() string {
	return s.String()
}

func (s SingleUnifyRenewResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SingleUnifyRenewResponse) SetData(v *SingleUnifyRenewResponseData) *SingleUnifyRenewResponse {
	s.Data = v
	return s
}

func (s *SingleUnifyRenewResponse) SetRequestId(v string) *SingleUnifyRenewResponse {
	s.RequestId = &v
	return s
}

func (s *SingleUnifyRenewResponse) SetRespMsg(v string) *SingleUnifyRenewResponse {
	s.RespMsg = &v
	return s
}

func (s *SingleUnifyRenewResponse) SetRespCode(v string) *SingleUnifyRenewResponse {
	s.RespCode = &v
	return s
}


type SingleUnifyRenewResponseBuilder struct {
	s *SingleUnifyRenewResponse
}

func NewSingleUnifyRenewResponseBuilder() *SingleUnifyRenewResponseBuilder {
	s := &SingleUnifyRenewResponse{}
	b := &SingleUnifyRenewResponseBuilder{s: s}
	return b
}

func (b *SingleUnifyRenewResponseBuilder) Data(v *SingleUnifyRenewResponseData) *SingleUnifyRenewResponseBuilder {
	b.s.Data = v
	return b
}

func (b *SingleUnifyRenewResponseBuilder) RequestId(v string) *SingleUnifyRenewResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SingleUnifyRenewResponseBuilder) RespMsg(v string) *SingleUnifyRenewResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *SingleUnifyRenewResponseBuilder) RespCode(v string) *SingleUnifyRenewResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *SingleUnifyRenewResponseBuilder) Build() *SingleUnifyRenewResponse {
	return b.s
}


