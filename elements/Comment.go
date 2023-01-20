package elements

// The Comment element contains the comment that is associated with a managed folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/comment
import "encoding/xml"

type Comment struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *Comment) SetForMarshal() {
	C.XMLName.Local = "t:Comment"
}

func (C *Comment) GetSchema() *Schema {
	return &SchemaTypes
}
