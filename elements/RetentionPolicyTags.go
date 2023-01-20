package elements

// The RetentionPolicyTags element contains a list of retention tags returned in the response of the GetUserRetentionPolicyTags WSDL operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionpolicytags
import "encoding/xml"

type RetentionPolicyTags struct {
	XMLName xml.Name

	// The RetentionPolicyTag element specifies the retention policy for a mailbox item.
	RetentionPolicyTag *RetentionPolicyTag `xml:"RetentionPolicyTag"`
}

func (R *RetentionPolicyTags) SetForMarshal() {
	R.XMLName.Local = "m:RetentionPolicyTags"
}

func (R *RetentionPolicyTags) GetSchema() *Schema {
	return &SchemaMessages
}
