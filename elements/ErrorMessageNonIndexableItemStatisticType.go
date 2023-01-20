package elements

// The ErrorMessage element contains the error message that describes why an item is not indexable.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/errormessage-nonindexableitemstatistictype
import "encoding/xml"

type ErrorMessageNonIndexableItemStatisticType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *ErrorMessageNonIndexableItemStatisticType) SetForMarshal() {
	E.XMLName.Local = "m:ErrorMessage"
}

func (E *ErrorMessageNonIndexableItemStatisticType) GetSchema() *Schema {
	return &SchemaMessages
}
