package elements

// The ExternalAudience element sets or contains a value that determines to whom external Out of Office (OOF) messages are sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externalaudience
import "encoding/xml"

type ExternalAudience struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// E-mail senders outside the mailbox user's organization who send messages to the user will receive an external OOF message response.
	ExternalAudienceAll string = `All`
	// E-mail senders outside the mailbox user's organization who send messages to the user will only receive an external OOF message response if the sender is in the user's Exchange store contact list.
	ExternalAudienceKnown string = `Known`
	// E-mail senders outside the mailbox user's organization who send messages to the user will not receive an external OOF message response.
	ExternalAudienceNone string = `None`
)

func (E *ExternalAudience) SetForMarshal() {
	E.XMLName.Local = "t:ExternalAudience"
}

func (E *ExternalAudience) GetSchema() *Schema {
	return &SchemaTypes
}
