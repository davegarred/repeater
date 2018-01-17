package persist

const KEY_CONFLICT = Error("Key conflict on store")

type Error string

func (e Error) Error() string {
	return string(e)
}

type Store interface {
	Store(string, string) error
	Retrieve(string) (string, error)
	Delete(string) error
}
