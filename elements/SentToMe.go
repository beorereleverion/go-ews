package elements

// The SentToMe element indicates whether the owner of the mailbox has to be in the ToRecipients property of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senttome
import "encoding/xml"

type SentToMe struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	SentToMefalse bool = false
	// true
	SentToMetrue bool = true
)

func (S *SentToMe) SetForMarshal() {
	S.XMLName.Local = "m:SentToMe"
}

func (S *SentToMe) GetSchema() *Schema {
	return &SchemaMessages
}
