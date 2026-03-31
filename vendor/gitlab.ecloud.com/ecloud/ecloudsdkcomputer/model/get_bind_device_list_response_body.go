// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBindDeviceListResponseBody struct {

    // 数据
	Data *[]GetBindDeviceListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetBindDeviceListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetBindDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s GetBindDeviceListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBindDeviceListResponseBody) SetData(v []GetBindDeviceListResponseData) *GetBindDeviceListResponseBody {
	s.Data = &v
	return s
}

func (s *GetBindDeviceListResponseBody) SetTotalCount(v int32) *GetBindDeviceListResponseBody {
	s.TotalCount = &v
	return s
}


type GetBindDeviceListResponseBodyBuilder struct {
	s *GetBindDeviceListResponseBody
}

func NewGetBindDeviceListResponseBodyBuilder() *GetBindDeviceListResponseBodyBuilder {
	s := &GetBindDeviceListResponseBody{}
	b := &GetBindDeviceListResponseBodyBuilder{s: s}
	return b
}

func (b *GetBindDeviceListResponseBodyBuilder) Data(v []GetBindDeviceListResponseData) *GetBindDeviceListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetBindDeviceListResponseBodyBuilder) TotalCount(v int32) *GetBindDeviceListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetBindDeviceListResponseBodyBuilder) Build() *GetBindDeviceListResponseBody {
	return b.s
}


