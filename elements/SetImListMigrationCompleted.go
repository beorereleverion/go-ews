package elements

// The SetImListMigrationCompleted element represents a request to indicate whether the Exchange store contains the instant messaging items used by instant messaging clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setimlistmigrationcompleted
type SetImListMigrationCompleted struct {
	// The ImListMigrationCompleted element indicates whether the Exchange store contains the instant messaging items used by instant messaging clients.
	ImListMigrationCompleted *ImListMigrationCompleted `xml:"m:ImListMigrationCompleted"`
}
