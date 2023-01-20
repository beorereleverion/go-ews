package elements

// The Duration element specifies the duration that the out of office (OOF) status is enabled if the OofState element is set to Scheduled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/duration-useroofsettings
import "encoding/xml"

type DurationUserOofSettings struct {
	XMLName xml.Name

	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"EndTime"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"StartTime"`
}

func (D *DurationUserOofSettings) SetForMarshal() {
	D.XMLName.Local = "t:Duration"
}

func (D *DurationUserOofSettings) GetSchema() *Schema {
	return &SchemaTypes
}
