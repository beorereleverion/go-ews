package elements

// The RetentionId element specifies the retention tag identifier.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionid
import "encoding/xml"

type RetentionId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *RetentionId) SetForMarshal() {
	R.XMLName.Local = "t:RetentionId"
}

func (R *RetentionId) GetSchema() *Schema {
	return &SchemaTypes
}
