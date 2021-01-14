package storage

type Scanner interface {
	Scan() bool
	Key() string
	Value() []byte
	Close()
}
