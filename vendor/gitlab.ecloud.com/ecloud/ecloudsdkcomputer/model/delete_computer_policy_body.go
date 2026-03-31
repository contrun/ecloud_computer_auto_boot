// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteComputerPolicyBody struct {
    position.Body
    // 策略id
	PolicyId *string `json:"policyId,omitempty"`
}

func (s DeleteComputerPolicyBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteComputerPolicyBody) GoString() string {
	return s.String()
}

func (s DeleteComputerPolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteComputerPolicyBody) SetPolicyId(v string) *DeleteComputerPolicyBody {
	s.PolicyId = &v
	return s
}


type DeleteComputerPolicyBodyBuilder struct {
	s *DeleteComputerPolicyBody
}

func NewDeleteComputerPolicyBodyBuilder() *DeleteComputerPolicyBodyBuilder {
	s := &DeleteComputerPolicyBody{}
	b := &DeleteComputerPolicyBodyBuilder{s: s}
	return b
}

func (b *DeleteComputerPolicyBodyBuilder) PolicyId(v string) *DeleteComputerPolicyBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *DeleteComputerPolicyBodyBuilder) Build() *DeleteComputerPolicyBody {
	return b.s
}


