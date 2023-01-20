package elements

// The SpecificUsers element specifies the email accounts that can access the app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/specificusers
import "encoding/xml"

type SpecificUsers struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (S *SpecificUsers) SetForMarshal() {
	S.XMLName.Local = "t:SpecificUsers"
}

func (S *SpecificUsers) GetSchema() *Schema {
	return &SchemaTypes
}
