// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetEnforcePasswordPeriodMonthBody struct {
    position.Body
    // 密码强制更新月份，只能是-1,1,6,12，-1标识不限制，
	MonthPeriod *int32 `json:"monthPeriod,omitempty"`
}

func (s SetEnforcePasswordPeriodMonthBody) String() string {
	return utils.Beautify(s)
}

func (s SetEnforcePasswordPeriodMonthBody) GoString() string {
	return s.String()
}

func (s SetEnforcePasswordPeriodMonthBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetEnforcePasswordPeriodMonthBody) SetMonthPeriod(v int32) *SetEnforcePasswordPeriodMonthBody {
	s.MonthPeriod = &v
	return s
}


type SetEnforcePasswordPeriodMonthBodyBuilder struct {
	s *SetEnforcePasswordPeriodMonthBody
}

func NewSetEnforcePasswordPeriodMonthBodyBuilder() *SetEnforcePasswordPeriodMonthBodyBuilder {
	s := &SetEnforcePasswordPeriodMonthBody{}
	b := &SetEnforcePasswordPeriodMonthBodyBuilder{s: s}
	return b
}

func (b *SetEnforcePasswordPeriodMonthBodyBuilder) MonthPeriod(v int32) *SetEnforcePasswordPeriodMonthBodyBuilder {
	b.s.MonthPeriod = &v
	return b
}

func (b *SetEnforcePasswordPeriodMonthBodyBuilder) Build() *SetEnforcePasswordPeriodMonthBody {
	return b.s
}


