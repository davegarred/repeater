// Package persist provides functions to store and retrieve key-value pairs
package persist

// KeyConflict is used during an attempt to store a key already that is already in use
const KeyConflict = persistError("Key conflict on store")

type persistError string

func (e persistError) Error() string {
	return string(e)
}
