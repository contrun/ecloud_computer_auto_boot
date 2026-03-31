// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchUserReActiveNotifyBody struct {
    position.Body
    // 用户id列表
	UserIdList []string `json:"userIdList,omitempty"`
}

func (s BatchUserReActiveNotifyBody) String() string {
	return utils.Beautify(s)
}

func (s BatchUserReActiveNotifyBody) GoString() string {
	return s.String()
}

func (s BatchUserReActiveNotifyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUserReActiveNotifyBody) SetUserIdList(v []string) *BatchUserReActiveNotifyBody {
	s.UserIdList = v
	return s
}


type BatchUserReActiveNotifyBodyBuilder struct {
	s *BatchUserReActiveNotifyBody
}

func NewBatchUserReActiveNotifyBodyBuilder() *BatchUserReActiveNotifyBodyBuilder {
	s := &BatchUserReActiveNotifyBody{}
	b := &BatchUserReActiveNotifyBodyBuilder{s: s}
	return b
}

func (b *BatchUserReActiveNotifyBodyBuilder) UserIdList(v []string) *BatchUserReActiveNotifyBodyBuilder {
	b.s.UserIdList = v
	return b
}

func (b *BatchUserReActiveNotifyBodyBuilder) Build() *BatchUserReActiveNotifyBody {
	return b.s
}


