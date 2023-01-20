package elements

// The StartDate element represents the start date of a recurring task or calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startdate-recurrence

import "encoding/xml"

type StartDateRecurrence struct {
	XMLName xml.Name
	//this is string, because StartDateReccurence has strange timeformat, and i can't find this format in default go formats
	TEXT *string `xml:",chardata"`
}

func (S *StartDateRecurrence) SetForMarshal() {
	S.XMLName.Local = "t:StartDate"
}

func (S *StartDateRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
