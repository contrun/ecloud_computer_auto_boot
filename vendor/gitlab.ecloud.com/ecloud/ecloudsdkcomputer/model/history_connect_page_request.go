// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type HistoryConnectPageRequest struct {


	HistoryConnectPageBody *HistoryConnectPageBody `json:"historyConnectPageBody,omitempty"`
}

func (s HistoryConnectPageRequest) String() string {
	return utils.Beautify(s)
}

func (s HistoryConnectPageRequest) GoString() string {
	return s.String()
}

func (s HistoryConnectPageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryConnectPageRequest) SetHistoryConnectPageBody(v *HistoryConnectPageBody) *HistoryConnectPageRequest {
	s.HistoryConnectPageBody = v
	return s
}


type HistoryConnectPageRequestBuilder struct {
	s *HistoryConnectPageRequest
}

func NewHistoryConnectPageRequestBuilder() *HistoryConnectPageRequestBuilder {
	s := &HistoryConnectPageRequest{}
	b := &HistoryConnectPageRequestBuilder{s: s}
	return b
}

func (b *HistoryConnectPageRequestBuilder) HistoryConnectPageBody(v *HistoryConnectPageBody) *HistoryConnectPageRequestBuilder {
	b.s.HistoryConnectPageBody = v
	return b
}

func (b *HistoryConnectPageRequestBuilder) Build() *HistoryConnectPageRequest {
	return b.s
}


