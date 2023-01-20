package elements

// The IsVisible element indicates whether the retention policy is visible to users.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isvisible
import "encoding/xml"

type IsVisible struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsVisiblefalse bool = false
	// true
	IsVisibletrue bool = true
)
