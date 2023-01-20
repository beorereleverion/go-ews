package elements

// The SourceId element specifies the identifier of the attributed contact in a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sourceid
import "encoding/xml"

type SourceId struct {
	XMLName xml.Name

	// The text value of the ChangeKey attribute is the change key of the contact.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The text value of the Id attribute is the identifier of the contact.
	Id *string `xml:"Id,attr"`
}

func (S *SourceId) SetForMarshal() {
	S.XMLName.Local = "t:SourceId"
}

func (S *SourceId) GetSchema() *Schema {
	return &SchemaTypes
}
