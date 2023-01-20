package elements

// The RecurringMasterItemId (ItemIdType) element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurringmasteritemid-itemidtype
import "encoding/xml"

type RecurringMasterItemIdItemIdType struct {
	XMLName xml.Name

	// Identifies a specific version of a single occurrence of a recurring master item. Additionally, the recurring master item is also identified because it and the single occurrence will contain the same change key. This attribute is optional.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies a single occurrence of a recurring master item. This attribute is required.
	Id *string `xml:"Id,attr"`
}

func (R *RecurringMasterItemIdItemIdType) SetForMarshal() {
	R.XMLName.Local = "t:RecurringMasterItemId"
}

func (R *RecurringMasterItemIdItemIdType) GetSchema() *Schema {
	return &SchemaTypes
}
