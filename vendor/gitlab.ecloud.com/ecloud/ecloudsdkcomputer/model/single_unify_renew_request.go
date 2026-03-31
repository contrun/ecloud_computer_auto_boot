// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SingleUnifyRenewRequest struct {


	SingleUnifyRenewBody *SingleUnifyRenewBody `json:"singleUnifyRenewBody,omitempty"`
}

func (s SingleUnifyRenewRequest) String() string {
	return utils.Beautify(s)
}

func (s SingleUnifyRenewRequest) GoString() string {
	return s.String()
}

func (s SingleUnifyRenewRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SingleUnifyRenewRequest) SetSingleUnifyRenewBody(v *SingleUnifyRenewBody) *SingleUnifyRenewRequest {
	s.SingleUnifyRenewBody = v
	return s
}


type SingleUnifyRenewRequestBuilder struct {
	s *SingleUnifyRenewRequest
}

func NewSingleUnifyRenewRequestBuilder() *SingleUnifyRenewRequestBuilder {
	s := &SingleUnifyRenewRequest{}
	b := &SingleUnifyRenewRequestBuilder{s: s}
	return b
}

func (b *SingleUnifyRenewRequestBuilder) SingleUnifyRenewBody(v *SingleUnifyRenewBody) *SingleUnifyRenewRequestBuilder {
	b.s.SingleUnifyRenewBody = v
	return b
}

func (b *SingleUnifyRenewRequestBuilder) Build() *SingleUnifyRenewRequest {
	return b.s
}


