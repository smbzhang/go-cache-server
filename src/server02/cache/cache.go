// Package cache provides ...
package cache

type Cache interface {
	/* TODO: add methods */
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStae() Stat
}
