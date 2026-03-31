// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByInstanceIdQuery struct {
    position.Query
    // 订购资源实例id
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s GetPolicyInfoByInstanceIdQuery) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByInstanceIdQuery) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByInstanceIdQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByInstanceIdQuery) SetInstanceId(v string) *GetPolicyInfoByInstanceIdQuery {
	s.InstanceId = &v
	return s
}


type GetPolicyInfoByInstanceIdQueryBuilder struct {
	s *GetPolicyInfoByInstanceIdQuery
}

func NewGetPolicyInfoByInstanceIdQueryBuilder() *GetPolicyInfoByInstanceIdQueryBuilder {
	s := &GetPolicyInfoByInstanceIdQuery{}
	b := &GetPolicyInfoByInstanceIdQueryBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByInstanceIdQueryBuilder) InstanceId(v string) *GetPolicyInfoByInstanceIdQueryBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdQueryBuilder) Build() *GetPolicyInfoByInstanceIdQuery {
	return b.s
}


