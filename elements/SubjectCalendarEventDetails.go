package elements

// The Subject element represents the subject of a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subject-calendareventdetails
import "encoding/xml"

type SubjectCalendarEventDetails struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SubjectCalendarEventDetails) SetForMarshal() {
	S.XMLName.Local = "t:Subject"
}

func (S *SubjectCalendarEventDetails) GetSchema() *Schema {
	return &SchemaTypes
}
