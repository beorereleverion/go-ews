package elements

// The IsEncrypted element indicates whether incoming messages must be S/MIME encrypted in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isencrypted
import "encoding/xml"

type IsEncrypted struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsEncryptedfalse bool = false
	// true
	IsEncryptedtrue bool = true
)

func (I *IsEncrypted) SetForMarshal() {
	I.XMLName.Local = "m:IsEncrypted"
}

func (I *IsEncrypted) GetSchema() *Schema {
	return &SchemaMessages
}
