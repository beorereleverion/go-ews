package elements

// The PerformReminderActionResponse element specifies the response to a PerformReminderAction request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/performreminderactionresponse
import "encoding/xml"

type PerformReminderActionResponse struct {
	XMLName xml.Name

	// The UpdatedItemIds element specifies the identifiers of updated reminder items.
	UpdatedItemIds *UpdatedItemIds `xml:"UpdatedItemIds"`
}

func (P *PerformReminderActionResponse) SetForMarshal() {
	P.XMLName.Local = "m:PerformReminderActionResponse"
}

func (P *PerformReminderActionResponse) GetSchema() *Schema {
	return &SchemaMessages
}
