package elements

// The ProposeNewTime element specifies a response object that indicates that the meeting attendee can propose a new meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposenewtime
import "encoding/xml"

type ProposeNewTime struct {
	XMLName xml.Name

	// The name of the response object.
	ObjectName *string `xml:"ObjectName,attr"`
}

func (P *ProposeNewTime) SetForMarshal() {
	P.XMLName.Local = "t:ProposeNewTime"
}

func (P *ProposeNewTime) GetSchema() *Schema {
	return &SchemaTypes
}
