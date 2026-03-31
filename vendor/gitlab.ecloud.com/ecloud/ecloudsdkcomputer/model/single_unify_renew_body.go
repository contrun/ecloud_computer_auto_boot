// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SingleUnifyRenewBody struct {
    position.Body
    // 续订时长，包月：1-60，包年：1-5
	Duration *int32 `json:"duration,omitempty"`
    // 资源id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s SingleUnifyRenewBody) String() string {
	return utils.Beautify(s)
}

func (s SingleUnifyRenewBody) GoString() string {
	return s.String()
}

func (s SingleUnifyRenewBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SingleUnifyRenewBody) SetDuration(v int32) *SingleUnifyRenewBody {
	s.Duration = &v
	return s
}

func (s *SingleUnifyRenewBody) SetInstanceId(v string) *SingleUnifyRenewBody {
	s.InstanceId = &v
	return s
}


type SingleUnifyRenewBodyBuilder struct {
	s *SingleUnifyRenewBody
}

func NewSingleUnifyRenewBodyBuilder() *SingleUnifyRenewBodyBuilder {
	s := &SingleUnifyRenewBody{}
	b := &SingleUnifyRenewBodyBuilder{s: s}
	return b
}

func (b *SingleUnifyRenewBodyBuilder) Duration(v int32) *SingleUnifyRenewBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *SingleUnifyRenewBodyBuilder) InstanceId(v string) *SingleUnifyRenewBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *SingleUnifyRenewBodyBuilder) Build() *SingleUnifyRenewBody {
	return b.s
}


