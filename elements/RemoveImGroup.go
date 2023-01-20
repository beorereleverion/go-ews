package elements

// The RemoveImGroup element represents a request to remove an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removeimgroup
import "encoding/xml"

type RemoveImGroup struct {
	XMLName xml.Name

	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"GroupId"`
}

func (R *RemoveImGroup) SetForMarshal() {
	R.XMLName.Local = "m:RemoveImGroup"
}

func (R *RemoveImGroup) GetSchema() *Schema {
	return &SchemaMessages
}
