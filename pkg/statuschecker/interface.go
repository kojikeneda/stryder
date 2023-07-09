package statuschecker

// StatusChecker is an interface that defines methods for checking the status of a service.
type StatusChecker interface {
	CheckStatus() (string, error)
	CanHandle(url string) bool
}
