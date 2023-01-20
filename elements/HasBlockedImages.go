package elements

// The HasBlockedImages element specifies a Boolean value that indicates whether the item has blocked images.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasblockedimages
import "encoding/xml"

type HasBlockedImages struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	HasBlockedImagesfalse bool = false
	// true
	HasBlockedImagestrue bool = true
)

func (H *HasBlockedImages) SetForMarshal() {
	H.XMLName.Local = "t:HasBlockedImages"
}

func (H *HasBlockedImages) GetSchema() *Schema {
	return &SchemaTypes
}
