package elements

type Schema struct {
	Name           string
	ShortName      string
	URL            string
	ValidationFile string
}

var (
	SchemaMessages = Schema{
		Name:           "Messages schema",
		ShortName:      "m",
		URL:            "https://schemas.microsoft.com/exchange/services/2006/messages",
		ValidationFile: "Messages.xsd",
	}
	SchemaTypes = Schema{
		Name:           "Types schema",
		ShortName:      "t",
		URL:            "https://schemas.microsoft.com/exchange/services/2006/types",
		ValidationFile: "Types.xsd",
	}
)
