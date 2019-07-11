package test

import (
	"github.com/dking1224/tomweb/src/web"
	"gopkg.in/guregu/null.v3/zero"
)

type Contact struct {
	ContactId        uint64       `db:"contact_id" json:"contact_id"`
	CompanyId        uint64       `db:"company_id" json:"company_id" binding:"required"`
	ContactCode      string       `db:"contact_code" json:"contact_code"`
	ContactName      string       `db:"contact_name" json:"contact_name"`
	ContactShortName string       `db:"contact_short_name" json:"contact_short_name"`
	ContactType      int8         `db:"contact_type" json:"contact_type"`
	ISDelete         int8         `db:"isdelete" json:"is_delete"`
	Category         uint32       `db:"category" json:"category"`
	CreateTime       web.JsonTime `db:"createtime" json:"create_time"`
	ModifyTime       web.JsonTime `db:"modifytime" json:"modify_time"`
	ModifyName       zero.String  `db:"modifyname" json:"modify_name"`
	Country          zero.String  `db:"country" json:"country"`
	Province         zero.String  `db:"province" json:"province"`
	City             zero.String  `db:"city" json:"city"`
	District         zero.String  `db:"district" json:"district"`
	Address          zero.String  `db:"address" json:"address"`
	Phone            zero.String  `db:"phone" json:"phone"`
	Email            zero.String  `db:"email" json:"email"`
}
