package elements

// The IntendedFreeBusyStatus element represents the intended status for the calendar item that is associated with the meeting request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/intendedfreebusystatus
import "encoding/xml"

type IntendedFreeBusyStatus struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *IntendedFreeBusyStatus) SetForMarshal() {
	I.XMLName.Local = "t:IntendedFreeBusyStatus"
}

func (I *IntendedFreeBusyStatus) GetSchema() *Schema {
	return &SchemaTypes
}
