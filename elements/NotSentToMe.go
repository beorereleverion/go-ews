package elements

// The NotSentToMe element indicates whether the owner of the mailbox must not be in the ToRecipients property of the incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/notsenttome
import "encoding/xml"

type NotSentToMe struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	NotSentToMefalse bool = false
	// true
	NotSentToMetrue bool = true
)

func (N *NotSentToMe) SetForMarshal() {
	N.XMLName.Local = "m:NotSentToMe"
}

func (N *NotSentToMe) GetSchema() *Schema {
	return &SchemaMessages
}
