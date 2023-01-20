package elements

// The SentOnlyToMe element indicates whether the owner of the mailbox has to be the only one in the ToRecipients property of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sentonlytome
import "encoding/xml"

type SentOnlyToMe struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	SentOnlyToMefalse bool = false
	// true
	SentOnlyToMetrue bool = true
)

func (S *SentOnlyToMe) SetForMarshal() {
	S.XMLName.Local = "m:SentOnlyToMe"
}

func (S *SentOnlyToMe) GetSchema() *Schema {
	return &SchemaMessages
}
