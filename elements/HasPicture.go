package elements

// The HasPicture element indicates whether the contact item has a file attachment that represents the contact's picture.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/haspicture
import "encoding/xml"

type HasPicture struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	HasPicturefalse bool = false
	// true
	HasPicturetrue bool = true
)

func (H *HasPicture) SetForMarshal() {
	H.XMLName.Local = "t:HasPicture"
}

func (H *HasPicture) GetSchema() *Schema {
	return &SchemaTypes
}
