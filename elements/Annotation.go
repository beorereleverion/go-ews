package elements

// The Annotation element contains optional notes added by a user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/annotation
import "encoding/xml"

type Annotation struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *Annotation) SetForMarshal() {
	A.XMLName.Local = "t:Annotation"
}

func (A *Annotation) GetSchema() *Schema {
	return &SchemaTypes
}
