// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetInstanceListCemBody struct {
    position.Body
    // ram账号
	RamUserName *string `json:"ramUserName,omitempty"`
    // 资源实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // 每页展示数量
	PageSize *int32 `json:"pageSize,omitempty"`
    // 页码
	Page *int32 `json:"page,omitempty"`
}

func (s GetInstanceListCemBody) String() string {
	return utils.Beautify(s)
}

func (s GetInstanceListCemBody) GoString() string {
	return s.String()
}

func (s GetInstanceListCemBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetInstanceListCemBody) SetRamUserName(v string) *GetInstanceListCemBody {
	s.RamUserName = &v
	return s
}

func (s *GetInstanceListCemBody) SetInstanceId(v string) *GetInstanceListCemBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceListCemBody) SetPageSize(v int32) *GetInstanceListCemBody {
	s.PageSize = &v
	return s
}

func (s *GetInstanceListCemBody) SetPage(v int32) *GetInstanceListCemBody {
	s.Page = &v
	return s
}


type GetInstanceListCemBodyBuilder struct {
	s *GetInstanceListCemBody
}

func NewGetInstanceListCemBodyBuilder() *GetInstanceListCemBodyBuilder {
	s := &GetInstanceListCemBody{}
	b := &GetInstanceListCemBodyBuilder{s: s}
	return b
}

func (b *GetInstanceListCemBodyBuilder) RamUserName(v string) *GetInstanceListCemBodyBuilder {
	b.s.RamUserName = &v
	return b
}

func (b *GetInstanceListCemBodyBuilder) InstanceId(v string) *GetInstanceListCemBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetInstanceListCemBodyBuilder) PageSize(v int32) *GetInstanceListCemBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetInstanceListCemBodyBuilder) Page(v int32) *GetInstanceListCemBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *GetInstanceListCemBodyBuilder) Build() *GetInstanceListCemBody {
	return b.s
}


