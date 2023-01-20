package elements

// The SentCcMe element indicates whether the owner of the mailbox has to be in the CcRecipients property of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sentccme
import "encoding/xml"

type SentCcMe struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	SentCcMefalse bool = false
	// true
	SentCcMetrue bool = true
)

func (S *SentCcMe) SetForMarshal() {
	S.XMLName.Local = "m:SentCcMe"
}

func (S *SentCcMe) GetSchema() *Schema {
	return &SchemaMessages
}
