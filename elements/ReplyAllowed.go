package elements

// The ReplyAllowed element specifies whether a reply is allowed for rights managed data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replyallowed
import "encoding/xml"

type ReplyAllowed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ReplyAllowedfalse bool = false
	// true
	ReplyAllowedtrue bool = true
)

func (R *ReplyAllowed) SetForMarshal() {
	R.XMLName.Local = "t:ReplyAllowed"
}

func (R *ReplyAllowed) GetSchema() *Schema {
	return &SchemaTypes
}
