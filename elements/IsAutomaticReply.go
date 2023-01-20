package elements

// The IsAutomaticReply element indicates whether incoming messages must be automatic replies in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isautomaticreply
import "encoding/xml"

type IsAutomaticReply struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsAutomaticReplyfalse bool = false
	// true
	IsAutomaticReplytrue bool = true
)

func (I *IsAutomaticReply) SetForMarshal() {
	I.XMLName.Local = "m:IsAutomaticReply"
}

func (I *IsAutomaticReply) GetSchema() *Schema {
	return &SchemaMessages
}
