package elements

// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subject
import "encoding/xml"

type Subject struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *Subject) SetForMarshal() {
	S.XMLName.Local = "t:Subject"
}

func (S *Subject) GetSchema() *Schema {
	return &SchemaTypes
}
