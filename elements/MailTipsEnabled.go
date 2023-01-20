package elements

// The MailTipsEnabled element indicates whether the mail tips service is available.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailtipsenabled
import "encoding/xml"

type MailTipsEnabled struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	MailTipsEnabledfalse bool = false
	// true
	MailTipsEnabledtrue bool = true
)

func (M *MailTipsEnabled) SetForMarshal() {
	M.XMLName.Local = "t:MailTipsEnabled"
}

func (M *MailTipsEnabled) GetSchema() *Schema {
	return &SchemaTypes
}
