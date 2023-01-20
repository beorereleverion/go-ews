package elements

// The OofSettings element contains the Out of Office (OOF) settings.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/oofsettings
import "encoding/xml"

type OofSettings struct {
	XMLName xml.Name

	// The Duration element specifies the duration that the out of office (OOF) status is enabled if the OofState element is set to Scheduled.
	Duration *DurationUserOofSettings `xml:"Duration"`
	// The ExternalAudience element sets or contains a value that determines to whom external Out of Office (OOF) messages are sent.
	ExternalAudience *ExternalAudience `xml:"ExternalAudience"`
	// The ExternalReply element contains the out of office (OOF) response that is sent to addresses outside the recipient's domain or trusted domains.
	ExternalReply *ExternalReply `xml:"ExternalReply"`
	// The InternalReply element contains the out of office (OOF) response sent to other users in the user's domain or trusted domains.
	InternalReply *InternalReply `xml:"InternalReply"`
	// The OofState element is used to get or set the user's Out of Office (OOF) state.
	OofState *OofState `xml:"OofState"`
}

func (O *OofSettings) SetForMarshal() {
	O.XMLName.Local = "m:OofSettings"
}

func (O *OofSettings) GetSchema() *Schema {
	return &SchemaMessages
}
