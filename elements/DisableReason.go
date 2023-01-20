package elements

// The DisableReason element specifies the reason for disabling an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/disablereason
import "encoding/xml"

type DisableReason struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// To improve mobile client performance.
	DisableReasonMobileClientPerformance string = `MobileClientPerformance`
	// No reason given
	DisableReasonNoReason string = `NoReason`
	// To improve Web app client performance.
	DisableReasonOWAClientPerformance string = `OWAClientPerformance`
	// To improve email client performance.
	DisableReasonOutlookClientPerformance string = `OutlookClientPerformance`
)

func (D *DisableReason) SetForMarshal() {
	D.XMLName.Local = "t:DisableReason"
}

func (D *DisableReason) GetSchema() *Schema {
	return &SchemaTypes
}
