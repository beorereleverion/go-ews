package elements

// The Contains element represents a search expression that determines whether a given property contains the supplied constant string value.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contains
import "encoding/xml"

type Contains struct {
	XMLName xml.Name

	// The Constant element identifies a constant value in a restriction.
	Constant *Constant `xml:"Constant"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
	// Determines whether the search ignores cases and spaces.
	ContainmentComparison *string `xml:"ContainmentComparison,attr"`
	// Identifies the boundaries of a search.
	ContainmentMode *string `xml:"ContainmentMode,attr"`
}

const (
	// The comparison must be exact.
	ContainsExact = `Exact`
	// The comparison ignores casing.
	ContainsIgnoreCase = `IgnoreCase`
	// The comparison ignores non-spacing characters.
	ContainsIgnoreNonSpacingCharacters = `IgnoreNonSpacingCharacters`
	// To be removed.
	ContainsLoose = `Loose`
	// The comparison ignores casing and non-spacing characters.
	ContainsIgnoreCaseAndNonSpacingCharacters = `IgnoreCaseAndNonSpacingCharacters`
	// To be removed.
	ContainsLooseAndIgnoreCase = `LooseAndIgnoreCase`
	// To be removed.
	ContainsLooseAndIgnoreNonSpace = `LooseAndIgnoreNonSpace`
	// To be removed.
	ContainsLooseAndIgnoreCaseAndIgnoreNonSpace = `LooseAndIgnoreCaseAndIgnoreNonSpace`
	// The comparison is between the full string and the constant. The property value and the supplied constant are precisely the same.
	ContainsFullString = `FullString`
	// The comparison is between the string prefix and the constant.
	ContainsPrefixed = `Prefixed`
	// The comparison is between a substring of the string and the constant.
	ContainsSubstring = `Substring`
	// The comparison is between a prefix on individual words in the string and the constant.
	ContainsPrefixOnWords = `PrefixOnWords`
	// The comparison is between an exact phrase in the string and the constant.
	ContainsExactPhrase = `ExactPhrase`
)

func (C *Contains) SetForMarshal() {
	C.XMLName.Local = "t:Contains"
}

func (C *Contains) GetSchema() *Schema {
	return &SchemaTypes
}
