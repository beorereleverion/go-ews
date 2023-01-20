package elements

// The SentToOrCcMe element indicates whether the owner of the mailbox has to be in either a ToRecipients or CcRecipients property of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senttoorccme
import "encoding/xml"

type SentToOrCcMe struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	SentToOrCcMefalse bool = false
	// true
	SentToOrCcMetrue bool = true
)

func (S *SentToOrCcMe) SetForMarshal() {
	S.XMLName.Local = "m:SentToOrCcMe"
}

func (S *SentToOrCcMe) GetSchema() *Schema {
	return &SchemaMessages
}
