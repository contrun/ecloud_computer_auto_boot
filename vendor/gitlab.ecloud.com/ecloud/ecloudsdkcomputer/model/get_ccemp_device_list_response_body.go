// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetCcempDeviceListResponseBody struct {

    // 数据
	Data *[]GetCcempDeviceListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetCcempDeviceListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetCcempDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s GetCcempDeviceListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCcempDeviceListResponseBody) SetData(v []GetCcempDeviceListResponseData) *GetCcempDeviceListResponseBody {
	s.Data = &v
	return s
}

func (s *GetCcempDeviceListResponseBody) SetTotalCount(v int32) *GetCcempDeviceListResponseBody {
	s.TotalCount = &v
	return s
}


type GetCcempDeviceListResponseBodyBuilder struct {
	s *GetCcempDeviceListResponseBody
}

func NewGetCcempDeviceListResponseBodyBuilder() *GetCcempDeviceListResponseBodyBuilder {
	s := &GetCcempDeviceListResponseBody{}
	b := &GetCcempDeviceListResponseBodyBuilder{s: s}
	return b
}

func (b *GetCcempDeviceListResponseBodyBuilder) Data(v []GetCcempDeviceListResponseData) *GetCcempDeviceListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetCcempDeviceListResponseBodyBuilder) TotalCount(v int32) *GetCcempDeviceListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetCcempDeviceListResponseBodyBuilder) Build() *GetCcempDeviceListResponseBody {
	return b.s
}


