// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchUnbindPolicyBody struct {
    position.Body
    // 订购资源实例id
	InstanceIdList []string `json:"instanceIdList,omitempty"`
    // 定时任务uuid
	PolicyUid *string `json:"policyUid,omitempty"`
}

func (s BatchUnbindPolicyBody) String() string {
	return utils.Beautify(s)
}

func (s BatchUnbindPolicyBody) GoString() string {
	return s.String()
}

func (s BatchUnbindPolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUnbindPolicyBody) SetInstanceIdList(v []string) *BatchUnbindPolicyBody {
	s.InstanceIdList = v
	return s
}

func (s *BatchUnbindPolicyBody) SetPolicyUid(v string) *BatchUnbindPolicyBody {
	s.PolicyUid = &v
	return s
}


type BatchUnbindPolicyBodyBuilder struct {
	s *BatchUnbindPolicyBody
}

func NewBatchUnbindPolicyBodyBuilder() *BatchUnbindPolicyBodyBuilder {
	s := &BatchUnbindPolicyBody{}
	b := &BatchUnbindPolicyBodyBuilder{s: s}
	return b
}

func (b *BatchUnbindPolicyBodyBuilder) InstanceIdList(v []string) *BatchUnbindPolicyBodyBuilder {
	b.s.InstanceIdList = v
	return b
}

func (b *BatchUnbindPolicyBodyBuilder) PolicyUid(v string) *BatchUnbindPolicyBodyBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *BatchUnbindPolicyBodyBuilder) Build() *BatchUnbindPolicyBody {
	return b.s
}


