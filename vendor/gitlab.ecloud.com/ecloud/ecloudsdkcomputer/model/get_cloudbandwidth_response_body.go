// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetCloudbandwidthResponseBody struct {

    // 公网ip地址
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
    // 弹性公网IP带宽Uid
	BandwidthFipUid *string `json:"bandwidthFipUid,omitempty"`
    // 最大带宽
	BandwidthMax *int32 `json:"bandwidthMax,omitempty"`
    // 公网流量（traffic、bandwidth）
	Pnt *string `json:"pnt,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // ip名称
	IpName *string `json:"ipName,omitempty"`
    // 计费方式
	Billing *string `json:"billing,omitempty"`
    // 状态
	Status *string `json:"status,omitempty"`
}

func (s GetCloudbandwidthResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetCloudbandwidthResponseBody) GoString() string {
	return s.String()
}

func (s GetCloudbandwidthResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCloudbandwidthResponseBody) SetPublicIpAddress(v string) *GetCloudbandwidthResponseBody {
	s.PublicIpAddress = &v
	return s
}

func (s *GetCloudbandwidthResponseBody) SetBandwidthFipUid(v string) *GetCloudbandwidthResponseBody {
	s.BandwidthFipUid = &v
	return s
}

func (s *GetCloudbandwidthResponseBody) SetBandwidthMax(v int32) *GetCloudbandwidthResponseBody {
	s.BandwidthMax = &v
	return s
}

func (s *GetCloudbandwidthResponseBody) SetPnt(v string) *GetCloudbandwidthResponseBody {
	s.Pnt = &v
	return s
}

func (s *GetCloudbandwidthResponseBody) SetCreatedTime(v string) *GetCloudbandwidthResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetCloudbandwidthResponseBody) SetPoolCode(v string) *GetCloudbandwidthResponseBody {
	s.PoolCode = &v
	return s
}

func (s *GetCloudbandwidthResponseBody) SetRegion(v string) *GetCloudbandwidthResponseBody {
	s.Region = &v
	return s
}

func (s *GetCloudbandwidthResponseBody) SetIpName(v string) *GetCloudbandwidthResponseBody {
	s.IpName = &v
	return s
}

func (s *GetCloudbandwidthResponseBody) SetBilling(v string) *GetCloudbandwidthResponseBody {
	s.Billing = &v
	return s
}

func (s *GetCloudbandwidthResponseBody) SetStatus(v string) *GetCloudbandwidthResponseBody {
	s.Status = &v
	return s
}


type GetCloudbandwidthResponseBodyBuilder struct {
	s *GetCloudbandwidthResponseBody
}

func NewGetCloudbandwidthResponseBodyBuilder() *GetCloudbandwidthResponseBodyBuilder {
	s := &GetCloudbandwidthResponseBody{}
	b := &GetCloudbandwidthResponseBodyBuilder{s: s}
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) PublicIpAddress(v string) *GetCloudbandwidthResponseBodyBuilder {
	b.s.PublicIpAddress = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) BandwidthFipUid(v string) *GetCloudbandwidthResponseBodyBuilder {
	b.s.BandwidthFipUid = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) BandwidthMax(v int32) *GetCloudbandwidthResponseBodyBuilder {
	b.s.BandwidthMax = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) Pnt(v string) *GetCloudbandwidthResponseBodyBuilder {
	b.s.Pnt = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) CreatedTime(v string) *GetCloudbandwidthResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) PoolCode(v string) *GetCloudbandwidthResponseBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) Region(v string) *GetCloudbandwidthResponseBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) IpName(v string) *GetCloudbandwidthResponseBodyBuilder {
	b.s.IpName = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) Billing(v string) *GetCloudbandwidthResponseBodyBuilder {
	b.s.Billing = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) Status(v string) *GetCloudbandwidthResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetCloudbandwidthResponseBodyBuilder) Build() *GetCloudbandwidthResponseBody {
	return b.s
}


