// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetImageBody struct {
    position.Body
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

func (s GetImageBody) String() string {
	return utils.Beautify(s)
}

func (s GetImageBody) GoString() string {
	return s.String()
}

func (s GetImageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetImageBody) SetMachineId(v string) *GetImageBody {
	s.MachineId = &v
	return s
}

func (s *GetImageBody) SetPageSize(v int32) *GetImageBody {
	s.PageSize = &v
	return s
}

func (s *GetImageBody) SetCurrentPage(v int32) *GetImageBody {
	s.CurrentPage = &v
	return s
}


type GetImageBodyBuilder struct {
	s *GetImageBody
}

func NewGetImageBodyBuilder() *GetImageBodyBuilder {
	s := &GetImageBody{}
	b := &GetImageBodyBuilder{s: s}
	return b
}

func (b *GetImageBodyBuilder) MachineId(v string) *GetImageBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetImageBodyBuilder) PageSize(v int32) *GetImageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetImageBodyBuilder) CurrentPage(v int32) *GetImageBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetImageBodyBuilder) Build() *GetImageBody {
	return b.s
}


