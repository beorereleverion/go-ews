package elements

// The ShowExternalRecipientCount element indicates whether consumers of the GetMailTips operation have to show mail tips that indicate the number of external recipients to which a message is addressed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/showexternalrecipientcount
import "encoding/xml"

type ShowExternalRecipientCount struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ShowExternalRecipientCountfalse bool = false
	// true
	ShowExternalRecipientCounttrue bool = true
)

func (S *ShowExternalRecipientCount) SetForMarshal() {
	S.XMLName.Local = "t:ShowExternalRecipientCount"
}

func (S *ShowExternalRecipientCount) GetSchema() *Schema {
	return &SchemaTypes
}
