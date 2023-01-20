package elements

// The WorkingPeriodArray element contains working period information for the mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workingperiodarray
import "encoding/xml"

type WorkingPeriodArray struct {
	XMLName xml.Name

	// The WorkingPeriod element contains the work week days and hours of the mailbox user.
	WorkingPeriod *WorkingPeriod `xml:"WorkingPeriod"`
}

func (W *WorkingPeriodArray) SetForMarshal() {
	W.XMLName.Local = "t:WorkingPeriodArray"
}

func (W *WorkingPeriodArray) GetSchema() *Schema {
	return &SchemaTypes
}
