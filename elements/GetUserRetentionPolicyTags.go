package elements

// The GetUserRetentionPolicyTags element is the request to get the retention tags associated with the user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserretentionpolicytags
import "encoding/xml"

type GetUserRetentionPolicyTags struct {
	XMLName xml.Name
	TEXT    struct{} `xml:",chardata"`
}

func (G *GetUserRetentionPolicyTags) SetForMarshal() {
	G.XMLName.Local = "m:GetUserRetentionPolicyTags"
}

func (G *GetUserRetentionPolicyTags) GetSchema() *Schema {
	return &SchemaMessages
}
