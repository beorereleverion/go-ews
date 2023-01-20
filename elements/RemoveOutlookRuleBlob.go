package elements

// The RemoveOutlookRuleBlob element indicates whether to remove the Microsoft Outlook rule blob.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removeoutlookruleblob
import "encoding/xml"

type RemoveOutlookRuleBlob struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	RemoveOutlookRuleBlobfalse bool = false
	// true
	RemoveOutlookRuleBlobtrue bool = true
)

func (R *RemoveOutlookRuleBlob) SetForMarshal() {
	R.XMLName.Local = "t:RemoveOutlookRuleBlob"
}

func (R *RemoveOutlookRuleBlob) GetSchema() *Schema {
	return &SchemaTypes
}
