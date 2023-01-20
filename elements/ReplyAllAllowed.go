package elements

// The ReplyAllAllowed element specifies whether a reply all is allowed for rights managed data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replyallallowed
import "encoding/xml"

type ReplyAllAllowed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ReplyAllAllowedfalse bool = false
	// true
	ReplyAllAllowedtrue bool = true
)

func (R *ReplyAllAllowed) SetForMarshal() {
	R.XMLName.Local = "t:ReplyAllAllowed"
}

func (R *ReplyAllAllowed) GetSchema() *Schema {
	return &SchemaTypes
}
