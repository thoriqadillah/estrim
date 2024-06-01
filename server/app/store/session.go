package store

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type sessionStore struct{}

func NewSessionStorage() fiber.Storage {
	return &sessionStore{}
}

// Get gets the value for the given key.
// `nil, nil` is returned when the key does not exist
func (s *sessionStore) Get(key string) ([]byte, error) {
	panic("SESSION GET NOT IMPLEMENTED")
	return nil, nil
}

// Set stores the given value for the given key along
// with an expiration value, 0 means no expiration.
// Empty key or value will be ignored without an error.
func (s *sessionStore) Set(key string, val []byte, exp time.Duration) error {
	panic("SESSION SET NOT IMPLEMENTED")
	return nil
}

// Delete deletes the value for the given key.
// It returns no error if the storage does not contain the key,
func (s *sessionStore) Delete(key string) error {
	panic("SESSION DELETE NOT IMPLEMENTED")
	return nil
}

// Reset resets the storage and delete all keys.
func (s *sessionStore) Reset() error {
	panic("SESSION RESET NOT IMPLEMENTED")
	return nil
}

// Close closes the storage and will stop any running garbage
// collectors and open connections.
func (s *sessionStore) Close() error {
	panic("SESSION CLOSE NOT IMPLEMENTED")
	return nil
}
