package elements

// The IsVoicemail element indicates whether incoming messages must be voice mail messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isvoicemail
import "encoding/xml"

type IsVoicemail struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsVoicemailfalse bool = false
	// true
	IsVoicemailtrue bool = true
)

func (I *IsVoicemail) SetForMarshal() {
	I.XMLName.Local = "m:IsVoicemail"
}

func (I *IsVoicemail) GetSchema() *Schema {
	return &SchemaMessages
}
