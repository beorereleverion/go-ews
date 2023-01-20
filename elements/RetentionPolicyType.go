package elements

// The RetentionPolicyType element specifies the retention policy type applied to items in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionpolicytype
import "encoding/xml"

type RetentionPolicyType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Archive
	RetentionPolicyTypeArchive string = `Archive`
	// Delete
	RetentionPolicyTypeDelete string = `Delete`
)

func (R *RetentionPolicyType) SetForMarshal() {
	R.XMLName.Local = "t:RetentionPolicyType"
}

func (R *RetentionPolicyType) GetSchema() *Schema {
	return &SchemaTypes
}
