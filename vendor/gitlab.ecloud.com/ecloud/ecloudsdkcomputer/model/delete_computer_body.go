// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteComputerBody struct {
    position.Body
    // 退订实例id集合，示例值：[CCA-XXX1,CCA-XXX2]
	InstanceIdList []string `json:"instanceIdList,omitempty"`
}

func (s DeleteComputerBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteComputerBody) GoString() string {
	return s.String()
}

func (s DeleteComputerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteComputerBody) SetInstanceIdList(v []string) *DeleteComputerBody {
	s.InstanceIdList = v
	return s
}


type DeleteComputerBodyBuilder struct {
	s *DeleteComputerBody
}

func NewDeleteComputerBodyBuilder() *DeleteComputerBodyBuilder {
	s := &DeleteComputerBody{}
	b := &DeleteComputerBodyBuilder{s: s}
	return b
}

func (b *DeleteComputerBodyBuilder) InstanceIdList(v []string) *DeleteComputerBodyBuilder {
	b.s.InstanceIdList = v
	return b
}

func (b *DeleteComputerBodyBuilder) Build() *DeleteComputerBody {
	return b.s
}


