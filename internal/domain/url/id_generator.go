package url

// IDGenerator generates unique identifactors.
type IDGenerator interface {
	Generate() (string, error)
}
