// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchBindPolicyBody struct {
    position.Body
    // 订购资源实例id列表
	InstanceIdList []string `json:"instanceIdList,omitempty"`
    // 定时任务uuid
	PolicyUid *string `json:"policyUid,omitempty"`
}

func (s BatchBindPolicyBody) String() string {
	return utils.Beautify(s)
}

func (s BatchBindPolicyBody) GoString() string {
	return s.String()
}

func (s BatchBindPolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchBindPolicyBody) SetInstanceIdList(v []string) *BatchBindPolicyBody {
	s.InstanceIdList = v
	return s
}

func (s *BatchBindPolicyBody) SetPolicyUid(v string) *BatchBindPolicyBody {
	s.PolicyUid = &v
	return s
}


type BatchBindPolicyBodyBuilder struct {
	s *BatchBindPolicyBody
}

func NewBatchBindPolicyBodyBuilder() *BatchBindPolicyBodyBuilder {
	s := &BatchBindPolicyBody{}
	b := &BatchBindPolicyBodyBuilder{s: s}
	return b
}

func (b *BatchBindPolicyBodyBuilder) InstanceIdList(v []string) *BatchBindPolicyBodyBuilder {
	b.s.InstanceIdList = v
	return b
}

func (b *BatchBindPolicyBodyBuilder) PolicyUid(v string) *BatchBindPolicyBodyBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *BatchBindPolicyBodyBuilder) Build() *BatchBindPolicyBody {
	return b.s
}


