package elements

// The Status element specifies the hold status for a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/status-holdstatustype
import "encoding/xml"

type StatusHoldStatusType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Failed
	StatusHoldStatusTypeFailed string = `Failed`
	// NotOnHold
	StatusHoldStatusTypeNotOnHold string = `NotOnHold`
	// OnHold
	StatusHoldStatusTypeOnHold string = `OnHold`
	// PartialHold
	StatusHoldStatusTypePartialHold string = `PartialHold`
	// Pending
	StatusHoldStatusTypePending string = `Pending`
)

func (S *StatusHoldStatusType) SetForMarshal() {
	S.XMLName.Local = "t:Status"
}

func (S *StatusHoldStatusType) GetSchema() *Schema {
	return &SchemaTypes
}
