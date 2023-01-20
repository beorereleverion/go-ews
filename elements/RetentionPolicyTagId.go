package elements

// The RetentionPolicyTagId element specifies the identifier of the retention policy tag.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionpolicytagid
import "encoding/xml"

type RetentionPolicyTagId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *RetentionPolicyTagId) SetForMarshal() {
	R.XMLName.Local = "t:RetentionPolicyTagId"
}

func (R *RetentionPolicyTagId) GetSchema() *Schema {
	return &SchemaTypes
}
