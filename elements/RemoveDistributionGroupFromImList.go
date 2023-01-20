package elements

// The RemoveDistributionGroupFromImList element represents a request to remove a specific instant messaging distribution list group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removedistributiongroupfromimlist
import "encoding/xml"

type RemoveDistributionGroupFromImList struct {
	XMLName xml.Name

	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"GroupId"`
}

func (R *RemoveDistributionGroupFromImList) SetForMarshal() {
	R.XMLName.Local = "m:RemoveDistributionGroupFromImList"
}

func (R *RemoveDistributionGroupFromImList) GetSchema() *Schema {
	return &SchemaMessages
}
