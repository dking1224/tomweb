package test

import (
	"github.com/dking1224/tomweb/src/web"
	"github.com/gin-gonic/gin"
)

func FindContactById(ctx *gin.Context, contactId uint64) (Contact, error) {
	var contact Contact
	query := &Contact{ContactId: contactId}
	error := web.QueryOne("contact.findById", query, &contact)
	if error != nil {
		web.CError(ctx, error)
	}
	return contact, error
}
