package elements

// The MessageTrackingReportId element represents the message by its message ID, the organization where the message was found, the server on which the message was submitted, and an internal ID that uniquely identifies the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagetrackingreportid
import "encoding/xml"

type MessageTrackingReportId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MessageTrackingReportId) SetForMarshal() {
	M.XMLName.Local = "t:MessageTrackingReportId"
}

func (M *MessageTrackingReportId) GetSchema() *Schema {
	return &SchemaTypes
}
