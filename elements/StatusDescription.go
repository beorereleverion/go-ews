package elements

// The StatusDescription element contains an explanation of the task status.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/statusdescription
import "encoding/xml"

type StatusDescription struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *StatusDescription) SetForMarshal() {
	S.XMLName.Local = "t:StatusDescription"
}

func (S *StatusDescription) GetSchema() *Schema {
	return &SchemaTypes
}
