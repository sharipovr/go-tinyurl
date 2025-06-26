package url

// Repository defines only what domain needs
type Repository interface {
	Save(u URL) error
	Find(id string) (URL, error)
}
