package elements

// The IncludeMimeContent element specifies whether the Multipurpose Internet Mail Extensions (MIME) content of an item or attachment is returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includemimecontent
import "encoding/xml"

type IncludeMimeContent struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IncludeMimeContentfalse bool = false
	// true
	IncludeMimeContenttrue bool = true
)

func (I *IncludeMimeContent) SetForMarshal() {
	I.XMLName.Local = "t:IncludeMimeContent"
}

func (I *IncludeMimeContent) GetSchema() *Schema {
	return &SchemaTypes
}
