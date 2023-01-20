package elements

// The ImItemList element contains a list of instant messaging groups and instant messaging contacts.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imitemlist
import "encoding/xml"

type ImItemList struct {
	XMLName xml.Name

	// The Groups element represents an array of instant messaging (IM) groups.
	Groups *GroupsArrayOfImGroupType `xml:"Groups"`
	// The Personas element specifies an array of personas returned from the GetImItems and GetImItemList operations.
	Personas *Personas `xml:"Personas"`
}

func (I *ImItemList) SetForMarshal() {
	I.XMLName.Local = "m:ImItemList"
}

func (I *ImItemList) GetSchema() *Schema {
	return &SchemaMessages
}
