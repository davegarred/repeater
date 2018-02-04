package persist

const KEY_CONFLICT = Error("Key conflict on store")

type Error string

func (e Error) Error() string {
	return string(e)
}
