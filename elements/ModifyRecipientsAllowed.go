package elements

// The ModifyRecipientsAllowed element specifies whether modification of the recipients is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/modifyrecipientsallowed
import "encoding/xml"

type ModifyRecipientsAllowed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ModifyRecipientsAllowedfalse bool = false
	// true
	ModifyRecipientsAllowedtrue bool = true
)

func (M *ModifyRecipientsAllowed) SetForMarshal() {
	M.XMLName.Local = "t:ModifyRecipientsAllowed"
}

func (M *ModifyRecipientsAllowed) GetSchema() *Schema {
	return &SchemaTypes
}
