package elements

// The OutOfOffice element represents the response message and a duration time for sending the response message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/outofoffice
import "encoding/xml"

type OutOfOffice struct {
	XMLName xml.Name

	// The Duration element specifies the duration that the out of office (OOF) status is enabled if the OofState element is set to Scheduled.
	Duration *DurationUserOofSettings `xml:"Duration"`
	// The OofState element is used to get or set the user's Out of Office (OOF) state.
	OofState *OofState `xml:"OofState"`
	// The ReplyBody element contains an Out of Office (OOF) message and the language used for the message.
	ReplyBody *ReplyBody `xml:"ReplyBody"`
}

func (O *OutOfOffice) SetForMarshal() {
	O.XMLName.Local = "t:OutOfOffice"
}

func (O *OutOfOffice) GetSchema() *Schema {
	return &SchemaTypes
}
