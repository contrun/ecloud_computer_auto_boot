// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPoolListResponseBody struct {

    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s GetPoolListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetPoolListResponseBody) GoString() string {
	return s.String()
}

func (s GetPoolListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPoolListResponseBody) SetPoolCode(v string) *GetPoolListResponseBody {
	s.PoolCode = &v
	return s
}

func (s *GetPoolListResponseBody) SetPoolName(v string) *GetPoolListResponseBody {
	s.PoolName = &v
	return s
}


type GetPoolListResponseBodyBuilder struct {
	s *GetPoolListResponseBody
}

func NewGetPoolListResponseBodyBuilder() *GetPoolListResponseBodyBuilder {
	s := &GetPoolListResponseBody{}
	b := &GetPoolListResponseBodyBuilder{s: s}
	return b
}

func (b *GetPoolListResponseBodyBuilder) PoolCode(v string) *GetPoolListResponseBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetPoolListResponseBodyBuilder) PoolName(v string) *GetPoolListResponseBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetPoolListResponseBodyBuilder) Build() *GetPoolListResponseBody {
	return b.s
}


