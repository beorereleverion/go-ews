package elements

// The ExchangeImpersonation element is used in the SOAP header of a request. When this element is present, the caller is trying to impersonate the account that is contained within the ExchangeImpersonation element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exchangeimpersonation
import "encoding/xml"

type ExchangeImpersonation struct {
	XMLName xml.Name

	// The ConnectingSID element represents an account to impersonate when you are using the ExchangeImpersonation SOAP header.
	ConnectingSID *ConnectingSID `xml:"ConnectingSID"`
}

func (E *ExchangeImpersonation) SetForMarshal() {
	E.XMLName.Local = "t:ExchangeImpersonation"
}

func (E *ExchangeImpersonation) GetSchema() *Schema {
	return &SchemaTypes
}
