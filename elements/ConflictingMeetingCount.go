package elements

// The ConflictingMeetingCount element represents the number of meetings that conflict with the calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conflictingmeetingcount
import "encoding/xml"

type ConflictingMeetingCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (C *ConflictingMeetingCount) SetForMarshal() {
	C.XMLName.Local = "t:ConflictingMeetingCount"
}

func (C *ConflictingMeetingCount) GetSchema() *Schema {
	return &SchemaTypes
}
