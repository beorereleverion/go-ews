package elements

// The PageDirection element contains the direction for pagination in the search result. The value is Previous or Next.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pagedirection
type PageDirection string

const (
	// Next
	PageDirectionNext PageDirection = `Next`
	// Previous
	PageDirectionPrevious PageDirection = `Previous`
)
