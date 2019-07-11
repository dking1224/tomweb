{{define "contact.findById"}}
    select * from contact where contact_id={{.ContactId}}
{{end}}
